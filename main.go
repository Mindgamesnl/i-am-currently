package main

import (
	"image"
	"image/draw"
	"image/png"
	_ "io/ioutil"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	// Open the input images.
	tomFile, err := os.Open("tom.png")
	if err != nil {
		log.Fatal(err)
	}
	defer tomFile.Close()

	imgFile, err := os.Open("img.png")
	if err != nil {
		log.Fatal(err)
	}
	defer imgFile.Close()

	// Decode the images.
	tomImg, err := png.Decode(tomFile)
	if err != nil {
		log.Fatal(err)
	}

	imgImg, err := png.Decode(imgFile)
	if err != nil {
		log.Fatal(err)
	}

	// Resize the tom image to match the size of the background image.
	resizedTomImg := resize.Resize(uint(imgImg.Bounds().Dx()), uint(imgImg.Bounds().Dy()), tomImg, resize.Lanczos3)

	// Create a new image with the same dimensions as the background image.
	outImg := image.NewRGBA(imgImg.Bounds())

	// Draw the background image onto the output.
	draw.Draw(outImg, outImg.Bounds(), imgImg, image.ZP, draw.Src)

	// Overlay the tom image onto the output.
	draw.Draw(outImg, outImg.Bounds(), resizedTomImg, image.ZP, draw.Over)

	// Create the output file.
	outFile, err := os.Create("out.png")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// Write the output image to the file.
	err = png.Encode(outFile, outImg)
	if err != nil {
		log.Fatal(err)
	}
}
