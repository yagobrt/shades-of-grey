package main

import (
	"image"
	"image/color"
)

func greyscale(img image.Image) *image.Gray {
	greyImg := image.NewGray(img.Bounds())
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	for x := range width {
		for y := range height {
			pixel := img.At(x, y)
			greyPix := color.GrayModel.Convert(pixel)
			greyImg.Set(x, y, greyPix)
		}
	}
	return greyImg
}
