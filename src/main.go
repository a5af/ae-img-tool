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
)

/*
Matrix generator
1) Draw random string of font
2) Animate via gif
3) Figure out the gif speed issue
*/

var (
	dpi     = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	hinting = flag.String("hinting", "none", "none | full")
	size    = flag.Float64("size", 12, "font size in points")
	spacing = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
)

var text = []string{
	"’Twas brillig, and the slithy toves",
	"Did gyre and gimble in the wabe;",
	"All mimsy were the borogoves,",
	"And the mome raths outgrabe.",
	"",
	"“Beware the Jabberwock, my son!",
	"The jaws that bite, the claws that catch!",
	"Beware the Jubjub bird, and shun",
	"The frumious Bandersnatch!”",
	"",
	"He took his vorpal sword in hand:",
	"Long time the manxome foe he sought—",
	"So rested he by the Tumtum tree,",
	"And stood awhile in thought.",
	"",
	"And as in uffish thought he stood,",
	"The Jabberwock, with eyes of flame,",
	"Came whiffling through the tulgey wood,",
	"And burbled as it came!",
	"",
	"One, two! One, two! and through and through",
	"The vorpal blade went snicker-snack!",
	"He left it dead, and with its head",
	"He went galumphing back.",
	"",
	"“And hast thou slain the Jabberwock?",
	"Come to my arms, my beamish boy!",
	"O frabjous day! Callooh! Callay!”",
	"He chortled in his joy.",
	"",
	"’Twas brillig, and the slithy toves",
	"Did gyre and gimble in the wabe;",
	"All mimsy were the borogoves,",
	"And the mome raths outgrabe.",
}

func main() {
	log.Println("Inside Main")
	//drawAnim(200, 120)
	writeRandomText()
}

func writeRandomText() {
	rgba := image.NewRGBA(image.Rect(0, 0, 640, 480))
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

func drawText(c *freetype.Context) {
	// Draw the text.
	pt := freetype.Pt(10, 10+int(c.PointToFixed(*size)>>6))
	for _, s := range text {
		_, err := c.DrawString(s, pt)
		if err != nil {
			log.Println(err)
			return
		}
		pt.Y += c.PointToFixed(*size * *spacing)
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
