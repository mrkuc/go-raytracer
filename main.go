package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	nx := 200
	ny := 100

	img := image.NewRGBA(image.Rect(0, 0, nx, ny))
	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for j := 1; j <= ny; j++ {
		for i := 0; i < nx; i++ {
			r := float64(i) / float64(nx)
			g := float64(j) / float64(ny)
			b := 0.2
			img.Set(i, ny-j, color.RGBA{uint8(r * 255.99), uint8(g * 255.99), uint8(b * 255.99), 255})
		}
	}

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}
