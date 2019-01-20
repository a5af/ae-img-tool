// This example demonstrates decoding a JPEG image and examining its pixels.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	//_ "image/gif"
	//_ "image/jpeg"
	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

/*
Matrix generator
1) Draw random string of font
2) Animate via gif
3) Figure out the gif speed issue
*/

var (
	dpi     = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	hinting = flag.String("hinting", "full", "none | full")
	size    = flag.Float64("size", 17, "font size in points")
	spacing = flag.Float64("spacing", .75, "line spacing (e.g. 2 means double spaced)")
)

func readTextFile(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	return str
}

func getRandomText() []string {
	text := []string{}
	for i := 0; i < 140; i++ {
		text = append(text, readTextFile("./input/d3.v5.min.js")[rand.Intn(500):])
	}
	return text
}

func main() {
	log.Println("Inside Main")
	//drawAnim(200, 120)
	writeRandomText(1200, 800)
}

func writeRandomText(w, h int) {
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.Draw(rgba, rgba.Bounds(), image.Black, image.ZP, draw.Src)

	c, _ := getContext(rgba)
	drawText(c)
	writePngFile(rgba)
}

func getContext(rgba draw.Image) (*freetype.Context, error) {
	var fontfile = flag.String("fontfile", "./input/fonts/AnonymousPro-Regular.ttf", "")
	flag.Parse()

	c := freetype.NewContext()

	// Read the font data.
	fontBytes, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	c.SetDPI(*dpi)
	c.SetFont(f)
	c.SetFontSize(*size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(image.White)
	switch *hinting {
	default:
		c.SetHinting(font.HintingNone)
	case "full":
		c.SetHinting(font.HintingFull)
	}
	return c, nil
}

func drawVerticalString(c *freetype.Context, s string, pt fixed.Point26_6) {
	for _, char := range s {
		c.DrawString(string(char), pt)
		pt.Y += c.PointToFixed(*size * *spacing)
	}

}

func drawText(c *freetype.Context) {
	// Draw the text.
	pt := freetype.Pt(0, int(*size))
	for _, s := range getRandomText() {
		drawVerticalString(c, s, pt)

		pt.X += c.PointToFixed(*size * *spacing)
	}
}

func writePngFile(rgba image.Image) {
	// Save that RGBA image to disk.
	t := time.Now().Unix()
	outFile, err := os.Create(fmt.Sprintf("./output/out_%d.png", t))

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, rgba)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("Wrote out.png OK.")
}

//Decode an image file, return the image.Image
func readFile(path string) image.Image {
	reader, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	inputImage, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	return inputImage
}

func writeGifFile(images []*image.Paletted, delays []int) {
	t := time.Now().Unix()
	f, _ := os.Create(fmt.Sprintf("./output/out_%d.gif", t))
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})

	log.Println("Done.")
}

func getColors() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 255, 255},
		color.RGBA{0, 255, 0, 255},
		color.RGBA{255, 0, 0, 255},
		color.RGBA{0, 0, 0, 255},
		color.RGBA{127, 127, 127, 255},
		color.RGBA{255, 255, 255, 255},
	}
}

func drawFrame(img *image.Paletted, w, h int) {
	const width0 = 15
	const height0 = 15

	bounds := image.Rect(0, 0, w, h)
	palette := getColors()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if x%width0 == 0 && y%height0 == 0 {
				draw.Draw(img,
					image.Rect(x, y, x+width0, y+height0),
					&image.Uniform{palette[rand.Intn(len(palette))]},
					image.ZP,
					draw.Src)
			}
		}
	}
}

func drawAnim(w, h int) {
	images := []*image.Paletted{}
	var delays []int
	steps := 300
	palette := getColors()

	for step := 0; step < steps; step++ {
		img := image.NewPaletted(image.Rect(0, 0, w, h), palette)
		images = append(images, img)
		delays = append(delays, 0)

		drawFrame(img, w, h)
	}

	writeGifFile(images, delays)
}
