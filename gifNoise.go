// This example demonstrates decoding a JPEG image and examining its pixels.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
	//_ "image/gif"
	//_ "image/jpeg"
)

func gifNoise(w, h int) {

	//Colors
	palette := []color.Color{
		color.RGBA{0, 0, 255, 255},
		color.RGBA{0, 255, 0, 255},
		color.RGBA{255, 0, 0, 255},
		color.RGBA{255, 255, 255, 255},
		color.RGBA{0, 0, 0, 255},
	}

	bounds := image.Rect(0, 0, w, h)
	// outImage := image.NewPaletted(bounds, palette)

	const WIDTH_0 = 12
	const HEIGHT_0 = 10

	// const WIDTH_1 = 60
	// const HEIGHT_1 = 50

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	std := 2.0
	mean := 2.5

	images := []*image.Paletted{}
	var delays []int
	steps := 8

	for step := 0; step < steps; step++ {
		img := image.NewPaletted(image.Rect(0, 0, w, h), palette)
		images = append(images, img)
		delays = append(delays, 0)

		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				if x%WIDTH_0 == 0 && y%HEIGHT_0 == 0 {
					norm := r1.NormFloat64()
					r := math.Min(
						float64(len(palette)-1),
						math.Max(math.Floor(norm*std+mean), 0))
					draw.Draw(img,
						image.Rect(x, y, x+WIDTH_0, y+HEIGHT_0),
						&image.Uniform{palette[int(r)]},
						image.ZP,
						draw.Src)
				}
			}
		}

		//2d iteration
		// for y := bounds.Max.Y / 2; y < bounds.Max.Y; y++ {
		// 	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		// 		if x%WIDTH_1 == 0 && y%HEIGHT_1 == 0 {
		// 			norm := r1.NormFloat64()
		// 			r := math.Min(
		// 				float64(len(palette)-1),
		// 				math.Max(math.Floor(norm*std+mean), 0))
		// 			draw.Draw(img,
		// 				image.Rect(x, y, x+WIDTH_1, y+HEIGHT_1),
		// 				&image.Uniform{palette[int(r)]},
		// 				image.ZP,
		// 				draw.Src)
		// 		}
		// 	}
		// }
	}
	t := time.Now().Unix()
	f, _ := os.Create(fmt.Sprintf("./output/noise_%d.gif", t))
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})

	log.Println("Done.")
}
