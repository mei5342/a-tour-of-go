package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	h int
	w int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	v := uint8((x + y) / 4)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{300, 600}
	pic.ShowImage(m)
}