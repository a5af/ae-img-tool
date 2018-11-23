// This example demonstrates decoding a JPEG image and examining its pixels.
package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
	//_ "image/gif"
	//_ "image/jpeg"
)

func readHex() image.Image {
	reader, err := os.Open("./input/hex.png")
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

func hex() {
	w := 1600
	h := 800

	//Colors
	palette := []color.Color{
		color.RGBA{0, 0, 255, 255},
		color.RGBA{0, 255, 0, 255},
		//color.RGBA{255, 0, 0, 255},
		//color.RGBA{255, 255, 255, 255},
		color.RGBA{0, 0, 0, 255},
	}

	bounds := image.Rect(0, 0, w, h)
	// outImage := image.NewPaletted(bounds, palette)

	const WIDTH = 120
	const HEIGHT = 100

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
		delays = append(delays, 1)

		//2d iteration
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				if x%WIDTH == 0 && y%HEIGHT == 0 {
					norm := r1.NormFloat64()
					r := math.Min(float64(len(palette)-1), math.Max(math.Floor(norm*std+mean), 0))
					draw.Draw(img, image.Rect(x, y, x+WIDTH, y+HEIGHT), &image.Uniform{palette[int(r)]}, image.ZP, draw.Src)
				}
			}
		}
	}

	// f, _ := os.Create("./output/noise.gif")
	// defer f.Close()
	// gif.EncodeAll(f, &gif.GIF{
	// 	Image: images,
	// 	Delay: delays,
	// })

	outputFile, err := os.Create("./output/out.png")
	defer outputFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(outputFile, images[3])

	log.Println("Done.")
}
