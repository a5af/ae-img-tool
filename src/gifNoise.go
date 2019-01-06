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

func gifNoise(w, h int) {

	//Out image size
	bounds := image.Rect(0, 0, w, h)

	//Base colors
	palette := []color.Color{
		color.RGBA{0, 0, 255, 255},
		color.RGBA{0, 255, 0, 255},
		color.RGBA{255, 0, 0, 255},
		color.RGBA{0, 0, 0, 255},
		// color.RGBA{128, 0, 128, 255},
		// color.RGBA{128, 128, 0, 255},
		// color.RGBA{0, 128, 128, 255},

		color.RGBA{127, 127, 127, 255},
		color.RGBA{255, 255, 255, 255},
	}

	//Generate more randomly
	// const RANDOM_COLORS = 20
	// for i := 0; i < RANDOM_COLORS; i++ {
	// 	palette = append(palette, color.RGBA{
	// 		uint8(rand.Intn(255)),
	// 		uint8(rand.Intn(255)),
	// 		uint8(rand.Intn(255)),
	// 		255,
	// 	})
	// }

	//Open XYB logo
	// x, _ := os.Open("./input/xyb-black-on-trans.png")
	// defer x.Close()
	// ix, _ := png.Decode(x)

	//pixel size
	const WIDTH_0 = 15
	const HEIGHT_0 = 15
	images := []*image.Paletted{}
	var delays []int
	steps := 300

	//Draw the noise onto the gif
	for step := 0; step < steps; step++ {
		img := image.NewPaletted(image.Rect(0, 0, w, h), palette)
		images = append(images, img)
		delays = append(delays, 0)

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

	//Write the output
	t := time.Now().Unix()
	f, _ := os.Create(fmt.Sprintf("../output/noise_%d.gif", t))
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})

	log.Println("Done.")
}
