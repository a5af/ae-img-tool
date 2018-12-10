// This example demonstrates decoding a JPEG image and examining its pixels.
package main

import (
	"image"
	"log"
	"os"
	//_ "image/gif"
	//_ "image/jpeg"
)

func main() {
	log.Println("Inside Main")
	// staticCheckers(1600, 800)
	gifNoise(820, 462)
}

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
