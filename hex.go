// This example demonstrates decoding a JPEG image and examining its pixels.
package main

import (
	"image"
	"image/draw"
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
	inputImage, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	bounds := inputImage.Bounds()
	outImage := image.NewRGBA(bounds)
	//blue := color.RGBA{0, 0, 255, 100}

	const WIDTH = 30
	const HEIGHT = 60

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if x%WIDTH == 0 && y%HEIGHT == 0 {
				draw.Draw(outImage, image.Rect(x, y, WIDTH, HEIGHT), inputImage, image.ZP, draw.Src)
			}
		}
	}

	//draw.Draw(outImage, outImage.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	// bounds := m.Bounds()

	outputFile, err := os.Create("./output/out.png")
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(outputFile, outImage)
	defer outputFile.Close()

	log.Println("Done.")
}
