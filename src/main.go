// This example demonstrates decoding a JPEG image and examining its pixels.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	//_ "image/gif"
	//_ "image/jpeg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

/*
Matrix generator
1) Draw random string of font
2) Animate via gif
3) Figure out the gif speed issue
*/

func main() {
	log.Println("Inside Main")
	//drawAnim(200, 120)
	writeRandomText()
}

func printBounds(b fixed.Rectangle26_6) {
	fmt.Printf("Min.X:%d Min.Y:%d Max.X:%d Max.Y:%d\n", b.Min.X, b.Min.Y, b.Max.X, b.Max.Y)
}

func printGlyph(g *truetype.GlyphBuf) {
	printBounds(g.Bounds)
	fmt.Print("Points:\n---\n")
	e := 0
	for i, p := range g.Points {
		fmt.Printf("%4d, %4d", p.X, p.Y)
		if p.Flags&0x01 != 0 {
			fmt.Print("  on\n")
		} else {
			fmt.Print("  off\n")
		}
		if i+1 == int(g.Ends[e]) {
			fmt.Print("---\n")
			e++
		}
	}
}

func writeRandomText() {
	var fontfile = flag.String("fontfile", "./input/fonts/AnonymousPro-Regular.ttf",
		"filename of the ttf font")
	flag.Parse()
	fmt.Printf("Loading fontfile %q\n", *fontfile)
	b, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := truetype.Parse(b)
	if err != nil {
		log.Println(err)
		return
	}
	fupe := fixed.Int26_6(f.FUnitsPerEm())
	printBounds(f.Bounds(fupe))
	fmt.Printf("FUnitsPerEm:%d\n\n", fupe)

	c0, c1 := 'A', 'V'

	i0 := f.Index(c0)
	hm := f.HMetric(fupe, i0)
	g := &truetype.GlyphBuf{}
	err = g.Load(f, fupe, i0, font.HintingNone)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("'%c' glyph\n", c0)
	fmt.Printf("AdvanceWidth:%d LeftSideBearing:%d\n", hm.AdvanceWidth, hm.LeftSideBearing)
	printGlyph(g)
	i1 := f.Index(c1)
	fmt.Printf("\n'%c', '%c' Kern:%d\n", c0, c1, f.Kern(fupe, i0, i1))

	fmt.Printf("\nThe numbers above are in FUnits.\n" +
		"The numbers below are in 26.6 fixed point pixels, at 12pt and 72dpi.\n\n")
	a := truetype.NewFace(f, &truetype.Options{
		Size: 12,
		DPI:  72,
	})
	fmt.Printf("%#v\n", a.Metrics())

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

func writeFile(images []*image.Paletted, delays []int) {
	t := time.Now().Unix()
	f, _ := os.Create(fmt.Sprintf("./output/noise_%d.gif", t))
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

	writeFile(images, delays)
}
