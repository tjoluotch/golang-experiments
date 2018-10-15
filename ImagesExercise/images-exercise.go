package main

import (
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

type Image struct{
	width int
	height int
}

func (img Image) ColorModel() color.Model{
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
		return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	//  makes a function with x and y values and return x^y as a uint8
	img_func := func(x, y int) uint8 {
		return uint8(x^y)
	}
	v := img_func(x, y)
	return color.RGBA{v, v, 255, 255}
}

func main(){
	m := Image{244,86}
	pic.ShowImage(m)
}
