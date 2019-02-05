package main

import (
	"bufio"
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func readTextFile(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	return str
}

func writePngFile(rgba image.Image) {
	// Save that RGBA image to disk.
	t := time.Now().Unix()
	outFile, err := os.Create(fmt.Sprintf("../output/out_%d.png", t))

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, rgba)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("Wrote out.png OK.")
}

func readFile(path string) (image.Image, draw.Image) {
	reader, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	inputImage, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	//convert to draw.Image
	asDraw := image.NewRGBA(inputImage.Bounds())
	draw.Draw(asDraw, inputImage.Bounds(), inputImage, image.ZP, draw.Src)

	return inputImage, asDraw
}

func writeGifFile(images []*image.Paletted, delays []int) {
	t := time.Now().Unix()
	f, _ := os.Create(fmt.Sprintf("../output/out_%d.gif", t))
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})

	log.Println("Done.")
}

func decodeGifFile(path string) (*gif.GIF, error) {
	reader, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	return gif.DecodeAll(reader)
}
