// Gradient prints a PNG gradient image.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

const res = 256

func main() {
	rect := image.Rect(0, 0, res, res)
	img := image.NewRGBA(rect)

	for i := 0; i < res; i++ {
		for j := 0; j < res; j++ {
			img.Set(i, j, color.RGBA{uint8(i), uint8(j), 127, 255})
		}
	}

	err := png.Encode(os.Stdout, img)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gradient: %v\n", err)
	}
}
