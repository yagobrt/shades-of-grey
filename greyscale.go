package main

import (
	"image"
	"image/color"
)

func greyscale(img image.Image) *image.Gray {
	greyImg := image.NewGray(img.Bounds())
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y += 1 {
		for x := bounds.Min.X; x < bounds.Max.X; x += 1 {
			pixel := img.At(x, y)
			greyPix := color.GrayModel.Convert(pixel)
			greyImg.Set(x, y, greyPix)
		}
	}
	return greyImg
}
