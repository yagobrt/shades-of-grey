package main

import (
	"image"
	"image/color"
)

// Threshold dithering

var BLACK = color.Gray{0}
var WHITE = color.Gray{255}

// Turn white every pixel with brightness above the threshold, turn black the rest
func thresholdDither(img *image.Gray, threshold uint8) *image.Gray {
	bwImg := image.NewGray(img.Bounds())
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := img.GrayAt(x, y)
			if pixel.Y > threshold {
				bwImg.SetGray(x, y, WHITE)
			} else {
				bwImg.SetGray(x, y, BLACK)
			}
		}
	}
	return bwImg
}

// Calculate the threshold for the image based on the average brightness
func meanThreshold(img *image.Gray) uint8 {
	bounds := img.Bounds()
	sum := 0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := img.GrayAt(x, y)
			sum += int(pixel.Y)
		}
	}

	mean := uint8(sum / (bounds.Dx() * bounds.Dy()))

	return mean
}
