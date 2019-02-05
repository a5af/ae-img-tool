// This example demonstrates decoding a JPEG image and examining its pixels.
package main

import (
	"flag"
	"log"
	//_ "image/gif"
	//_ "image/jpeg"
)

/*
Matrix generator
1) Draw random string of font
2) Animate via gif
3) Figure out the gif speed issue
*/

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	hinting  = flag.String("hinting", "full", "none | full")
	size     = flag.Float64("size", 25, "font size in points")
	spacing  = flag.Float64("spacing", .55, "line spacing (e.g. 2 means double spaced)")
	fontfile = flag.String("fontfile", "../input/fonts/AnonymousPro-Regular.ttf", "")
)

func main() {
	log.Println("Inside Main")

	//w := 1200
	//h := 523

	//rgba := generateMaskedLogo(w, h)
	//writePngFile(rgba)

	//drawAnim(w, h)

	// g1, _ := decodeGifFile("../input/xyb-fast.gif")
	// g2, _ := decodeGifFile("../input/xyb-slow.gif")

	// log.Print(g1.Delay, g2.Delay)
	drawLines()

}
