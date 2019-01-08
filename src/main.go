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

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "../../testdata/luxisr.ttf", "filename of the ttf font")
	hinting  = flag.String("hinting", "none", "none | full")
	size     = flag.Float64("size", 12, "font size in points")
	spacing  = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	wonb     = flag.Bool("whiteonblack", false, "white text on a black background")
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

	// Read the font data.
	fontBytes, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	// Initialize the context.
	fg, bg := image.Black, image.White
	ruler := color.RGBA{0xdd, 0xdd, 0xdd, 0xff}
	if *wonb {
		fg, bg = image.White, image.Black
		ruler = color.RGBA{0x22, 0x22, 0x22, 0xff}
	}
	rgba := image.NewRGBA(image.Rect(0, 0, 640, 480))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(*dpi)
	c.SetFont(f)
	c.SetFontSize(*size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	switch *hinting {
	default:
		c.SetHinting(font.HintingNone)
	case "full":
		c.SetHinting(font.HintingFull)
	}

	// Draw the guidelines.
	for i := 0; i < 200; i++ {
		rgba.Set(10, 10+i, ruler)
		rgba.Set(10+i, 10, ruler)
	}

	// Draw the text.
	pt := freetype.Pt(10, 10+int(c.PointToFixed(*size)>>6))
	for _, s := range text {
		_, err = c.DrawString(s, pt)
		if err != nil {
			log.Println(err)
			return
		}
		pt.Y += c.PointToFixed(*size * *spacing)
	}

	// Save that RGBA image to disk.
	outFile, err := os.Create("out.png")
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
