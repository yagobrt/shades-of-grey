package main

import (
	"fmt"
	"image"
	"os"
	"time"

	_ "image/jpeg"
	"image/png"
)

// Return the contents of the file in imagePath
func readImage(imagePath string) (image.Image, error) {
	imageFile, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("opening image: %w", err)
	}
	defer imageFile.Close()

	img, _, err := image.Decode(imageFile)
	if err != nil {
		return nil, fmt.Errorf("decoding image: %w", err)
	}
	return img, nil
}

// Save an image as PNG
func savePNG(imagePath string, img image.Image) error {
	imageFile, err := os.Create(imagePath)
	if err != nil {
		return fmt.Errorf("opening output file: %w", err)
	}
	defer imageFile.Close()
	err = png.Encode(imageFile, img)

	if err != nil {
		return fmt.Errorf("encoding image: %w", err)
	}
	return nil

}

func run() error {
	if len(os.Args) == 1 {
		return fmt.Errorf("usage: %s <filename> [output]", os.Args[0])
	}

	// Read input image
	inputPath := os.Args[1]
	inputImg, err := readImage(inputPath)
	if err != nil {
		return err
	}

	// Convert to greyscale
	outputImg := greyscale(inputImg)

	outputPath := fmt.Sprintf("%v-output.png", time.Now().Format("2006-01-02-150405"))
	if len(os.Args) > 2 {
		outputPath = os.Args[2]
	}

	// Reencode image as png
	err = savePNG(outputPath, outputImg)

	return err

}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
