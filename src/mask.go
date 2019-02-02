package main

import (
	"image"
	"image/color"
	"image/draw"
)

func applyMask(dst draw.Image) {

	//mask: logo
	//dst: rgba
	//src: Black

	//read xyb logo file
	_, mask := readFile("../input/xyb-black-on-trans.png")

	invertMask(mask)

	black := color.RGBA{0, 0, 0, 255}
	src := &image.Uniform{black}

	draw.DrawMask(dst,
		dst.Bounds(),
		src, src.Bounds().Min, mask, mask.Bounds().Min, draw.Over)
}

func invertMask(mask draw.Image) {
	b := mask.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			_, _, _, a := mask.At(x, y).RGBA()

			//invert alpha
			if a == 0 {
				a = (1 << 16) - 1
			} else {
				a = 0
			}

			mask.Set(x, y, color.RGBA{255, 0, 0, uint8(a)})
		}
	}
}
