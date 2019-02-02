package main

import (
	"image"
	"image/draw"
)

func generateMaskedLogo(w, h int) *image.Paletted {
	rgba := image.NewPaletted(image.Rect(0, 0, w, h), getColors())
	draw.Draw(rgba, rgba.Bounds(), image.Black, image.ZP, draw.Src)

	drawText(rgba)
	applyMask(rgba)
	return rgba
}
