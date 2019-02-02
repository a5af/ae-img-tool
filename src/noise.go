package main

import (
	"image"
	"image/draw"
	"math/rand"
)

func drawNoise(img *image.Paletted, w, h int) {
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
