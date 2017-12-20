package main

import (
	"image"
	"image/color"
	"math/rand"
)

var gameboy = []color.Color{
	color.RGBA{14, 55, 15, 255},
	color.RGBA{47, 97, 48, 255},
	color.RGBA{138, 171, 25, 255},
	color.RGBA{154, 187, 27, 255},
}

var retro = []color.RGBA{
	color.RGBA{0x00, 0x04, 0x0f, 0xff},
	color.RGBA{0x03, 0x26, 0x28, 0xff},
	color.RGBA{0x07, 0x3e, 0x1e, 0xff},
	color.RGBA{0x18, 0x55, 0x08, 0xff},
	color.RGBA{0x5f, 0x6e, 0x0f, 0xff},
	color.RGBA{0x84, 0x50, 0x19, 0xff},
	color.RGBA{0x9b, 0x30, 0x22, 0xff},
	color.RGBA{0xb4, 0x92, 0x2f, 0xff},
	color.RGBA{0x94, 0xca, 0x3d, 0xff},
	color.RGBA{0x4f, 0xd5, 0x51, 0xff},
	color.RGBA{0x66, 0xff, 0xb3, 0xff},
	color.RGBA{0x82, 0xc9, 0xe5, 0xff},
	color.RGBA{0x9d, 0xa3, 0xeb, 0xff},
	color.RGBA{0xd7, 0xb5, 0xf3, 0xff},
	color.RGBA{0xfd, 0xd6, 0xf6, 0xff},
	color.RGBA{0xff, 0xf0, 0xf2, 0xff},
}

// mandel for each pixel on the screen do
func mandel(w, h, px, py, maxIterations int, zoom float32) color.RGBA {
	x0 := zoom * (3.5*float32(px)/float32(w) - 2.5)
	y0 := zoom * (2*float32(py)/float32(h) - 1.0)
	x := float32(0)
	y := float32(0)

	i := 0
	for x*x+y*y < 2*2 && i < maxIterations {
		xtemp := x*x - y*y + x0
		y = 2*x*y + y0
		x = xtemp
		i++
	}

	var pixelColor color.RGBA
	if i == maxIterations {
		pixelColor = color.RGBA{0, 0, 0, 255}
	} else {
		// c := uint8(iteration * 10)
		rand.Seed(int64(i))
		// n := uint8(rand.Intn(255))
		// fmt.Printf("%v ", i%len(retro))
		pixelColor = retro[i%len(retro)]
		// pixelColor = color.RGBA{n, n, n, 255}
	}

	return pixelColor
}

// CreateImage create an image
func createImage(width, height, maxIterations int, zoom float32) *image.RGBA {

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, mandel(width, height, x, y, maxIterations, zoom))
		}
	}
	return img
}
