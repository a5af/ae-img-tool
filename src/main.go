// This example demonstrates decoding a JPEG image and examining its pixels.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"log"
	"math/rand"
	"os"
	"time"
	//_ "image/gif"
	//_ "image/jpeg"
)

/*
Matrix generator
1) Draw random string of font
2) Animate via gif
3) Figure out the gif speed issue
*/

func main() {
	log.Println("Inside Main")
	drawAnim(200, 120)
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
	const WIDTH_0 = 15
	const HEIGHT_0 = 15

	bounds := image.Rect(0, 0, w, h)
	palette := getColors()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if x%WIDTH_0 == 0 && y%HEIGHT_0 == 0 {
				draw.Draw(img,
					image.Rect(x, y, x+WIDTH_0, y+HEIGHT_0),
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
