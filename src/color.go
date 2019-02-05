package main

import (
	"image/color"
	"math/rand"
)

//Decode an image file, return the image.Image
func getColors() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 255, 255},
		color.RGBA{0, 255, 0, 255},
		color.RGBA{255, 0, 0, 255},
		color.RGBA{0, 0, 0, 255},
		color.RGBA{255, 255, 255, 255},
	}
}

func getRGB() []color.Color {
	return []color.Color{
		color.RGBA{0, 0, 255, 255},
		color.RGBA{0, 255, 0, 255},
		color.RGBA{255, 0, 0, 255},
		color.RGBA{255, 255, 255, 255},
		color.RGBA{255, 255, 255, 255},
		color.RGBA{255, 255, 255, 255},
		color.RGBA{255, 255, 255, 255},
	}
}

func getRandomColor() color.Color {
	colors := getRGB()
	l := len(colors)
	randColor := colors[rand.Intn(l)]
	return randColor
}
