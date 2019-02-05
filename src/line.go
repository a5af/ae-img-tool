package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

func drawLines() {
	const W = 2048
	const H = 1024
	dc := gg.NewContext(W, H)
	dc.SetRGB(0, 0, 0)
	dc.Clear()

	for i := -W; i < 2*W; i += 4 {
		x1 := float64(i)
		y1 := float64(0)
		x2 := float64(i + W)
		y2 := float64(H)

		r := rand.Float64()
		g := rand.Float64()
		b := rand.Float64()
		a := float64(1)
		w := float64(4)
		dc.SetRGBA(r, g, b, a)
		dc.SetLineWidth(w)
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
	}

	t := time.Now().Unix()

	dc.SavePNG(fmt.Sprintf("../output/out_%d.png", t))
}
