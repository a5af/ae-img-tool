package main

import (
	"image"
	"log"
)

func drawAnim(w, h int) {
	images := []*image.Paletted{}
	var delays []int
	steps := 15

	for step := 0; step < steps; step++ {
		log.Print("step", step)
		//img := image.NewPaletted(image.Rect(0, 0, w, h), palette)
		rgba := generateMaskedLogo(w, h)
		images = append(images, rgba)
		delays = append(delays, 0)

		//drawNoise(img, w, h)
	}

	writeGifFile(images, delays)
}
