package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func createLogo(width, height int) {
	min := image.Point{X: 0, Y: 0}
	max := image.Point{X: width, Y: height}
	img := image.NewRGBA(image.Rectangle{Max: max, Min: min})
	background := color.RGBA{139, 23, 42, 0xff}
	green := color.RGBA{58, 201, 20, 0xff}
	blue := color.RGBA{0, 36, 255, 0xff}
	red := color.RGBA{228, 13, 29, 0xff}

	center := image.Point{X: width / 2, Y: height / 2}
	radius := 140
	rad := 20
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if int(math.Sqrt(float64((center.X-x)*(center.X-x)+(center.Y-y)*(center.Y-y)))) <= radius {
				if int(math.Sqrt(float64((center.X-x)*(center.X-x)-(center.Y-y)*(center.Y-y)))) <= rad {
					if y >= 0 && y <= 100 {
						img.Set(x, y, green)
					} else if y >= 100 && y <= 200 {
						img.Set(x, y, blue)
					} else {
						img.Set(x, y, red)
					}
				} else {
					img.Set(x, y, background)
				}
			}
		}
	}

	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func main() {
	width, height := 300, 300
	createLogo(width, height)
}
