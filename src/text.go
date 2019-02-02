package main

import (
	"flag"
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"log"
	"math/rand"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func getContext(rgba draw.Image, colr color.Color, size *float64) (*freetype.Context, error) {

	flag.Parse()

	c := freetype.NewContext()

	// Read the font data.
	fontBytes, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	c.SetDPI(*dpi)
	c.SetFont(f)
	c.SetFontSize(*size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(image.NewUniform(colr))
	switch *hinting {
	default:
		c.SetHinting(font.HintingNone)
	case "full":
		c.SetHinting(font.HintingFull)
	}
	return c, nil
}

func drawVerticalString(rgba draw.Image, s string, pt fixed.Point26_6) *freetype.Context {
	c, _ := getContext(rgba, getRandomColor(), size)
	for _, char := range s {

		c, _ = getContext(rgba, getRandomColor(), size)

		c.DrawString(string(char), pt)
		pt.Y += c.PointToFixed(*size * *spacing)
	}
	return c
}

func drawText(rgba draw.Image) {
	// Draw the text.
	pt := freetype.Pt(0, 0)
	for _, s := range getRandomText() {

		c := drawVerticalString(rgba, s, pt)
		pt.X += c.PointToFixed(*size * *spacing)
	}
}

func getRandomText() []string {
	text := []string{}
	for i := 0; i < 140; i++ {
		text = append(text, readTextFile("../input/d3.v5.min.js")[rand.Intn(500):600])
	}
	return text
}
