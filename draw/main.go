package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"time"

	"golang.org/x/image/draw"
)

func main() {
	start := time.Now()
	s2 := start
	input, err := os.Open("imgtest.jpg")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("open src:", time.Since(start))
	s2 = time.Now()
	defer input.Close()

	output, _ := os.Create("out.jpg")
	defer output.Close()
	fmt.Println("create dst:", time.Since(s2))
	s2 = time.Now()

	// Decode the image (from PNG to image.Image):
	src, err := jpeg.Decode(input)
	fmt.Println("jpeg decode:", time.Since(s2))
	s2 = time.Now()

	if err != nil {
		log.Fatal(err)
	}

	// Set the expected size that you want:
	presetWidth := 1920
	height := GetHeightByWidths(src.Bounds().Dx(), src.Bounds().Dy(), presetWidth)
	dst := image.NewRGBA(image.Rect(0, 0, presetWidth, height))
	fmt.Println("create rgba:", time.Since(s2))
	s2 = time.Now()

	// Resize:
	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)
	fmt.Println("scale:", time.Since(s2))
	s2 = time.Now()

	// Encode to `output`:
	jpeg.Encode(output, dst, &jpeg.Options{Quality: 65})
	fmt.Println("encode:", time.Since(s2))
	fmt.Println("Total time:", time.Since(start))
}

func GetHeightByWidths(imageWidth int, imageHeight int, presetWidth int) int {
	factor := float32(imageWidth) / float32(presetWidth)
	return int(float32(imageHeight) / factor)
}
