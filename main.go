package main

import (
	"math"
	"golang.org/x/tour/pic"
	"image/color"
	"image"
)

type Image struct {
	data [][]uint8
	sx, sy int
}

func getImage(sx, sy int) *Image {
	img := &Image{}

	img.sx = sx
	img.sy = sy

	img.data = make([][]uint8, sy)

	for i := range img.data {
		img.data[i] = make([]uint8, sx)
	}

	return img
}

func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.sx, img.sy)
}

func (img *Image) At(x, y int) color.Color {
	return color.RGBA{img.data[y][x], img.data[y][x], 255, 255}
}

func (img *Image) genMandelbrot() {
	dx, dy := img.Bounds().Max.X, img.Bounds().Max.Y;

	for i := range img.data {
		for j := range img.data[i] {
			img.data[i][j] = getColor(complex((float64(j)/float64(dx)-.7)*3., (float64(i)/float64(dy)-.5)*3.));
		}
	}
}

func getColor(c complex128) uint8 {
	z := c
	for k := 0; k != 256; k++ {
		z = z*z + c
		if (math.Abs(real(z)) > 2) || (math.Abs(imag(z)) > 2) {
			return uint8(k*8)
		}
	}
	return 0
}

func main() {
	m := getImage(1024, 1024)
	m.genMandelbrot();
	pic.ShowImage(m)
}
