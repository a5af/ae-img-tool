// This example demonstrates decoding a JPEG image and examining its pixels.
package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	//_ "image/gif"
	//_ "image/jpeg"
)

func hex() {
	reader, err := os.Open("./input/hex.png")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()

	// Split image into grid of 25x60 pixels
	// Iterate through the elements setting a random color
	const GRID_W = 25
	const GRID_H = 60

	out := image.NewRGBA(image.Rect(0, 0, bounds.Max.X, bounds.Max.Y))
	out.Pix[3] = 123

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			//r, g, b, a := m.At(x, y).RGBA()
			if x%GRID_W == 0 && y%GRID_H == 0 {
				//out.Pix[x*y] = 123
				fmt.Print(x, y)
			}
			//out.Pix[x*y] = m.At(x, y)

		}
	}

	outputFile, err := os.Create("./output/out.png")
	if err != nil {
	}
	png.Encode(outputFile, out)
	outputFile.Close()

	fmt.Printf("test.")
}
