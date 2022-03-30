package main

import (
	"fmt"
	"github.com/foobaz/lossypng/lossypng"
	"gopkg.in/gographics/imagick.v3/imagick"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
	"time"
)

const PngQuantization = 9

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	tmpDir := os.TempDir()
	files := [2]string{"imgtest.jpg", "imgtest.png"}

	pw := 1920
	//ph := 800

	for _, fileName := range files {
		fmt.Printf("[%s]:\n---\n", fileName)
		start := time.Now()

		if err := mw.ReadImage(fileName); err != nil {
			panic(err)
		}
		fmt.Println("1:", time.Since(start))
		s2 := time.Now()


		//// Get original logo size
		width := mw.GetImageWidth()
		height := mw.GetImageHeight()

		factor := float32(width) / float32(pw)
		hHeight := uint(float32(height) / factor)

		identify, err := ParseIdentify(mw)

		if err != nil {
			panic(err)
		}

		fmt.Println("2:", time.Since(s2))
		s2 = time.Now()
		//fmt.Println("mimetype:", identify.Image.MimeType)

		// Set the compression quality to 95 (high quality = low compression)
		if err = mw.SetImageCompressionQuality(65); err != nil {
			panic(err)
		}

		fmt.Println("3:", time.Since(s2))
		s2 = time.Now()
		// Resize the image using the Lanczos filter
		// The blur factor is a float, where > 1 is blurry, < 1 is sharp
		err = mw.ResizeImage(uint(pw), hHeight, imagick.FILTER_LANCZOS)
		if err != nil {
			panic(err)
		}

		fmt.Println("4:", time.Since(s2))
		s2 = time.Now()
		if err = mw.SetOption("png:compression-level", "9"); err != nil {
			panic(err)
		}
		fmt.Println("5:", time.Since(s2))
		s2 = time.Now()
		outputFileName := fmt.Sprintf("%so_" + strconv.FormatInt(time.Now().Unix(), 10), tmpDir)
		if err = mw.WriteImage(outputFileName); err != nil {
			panic(err)
		}

		fmt.Println("6:", time.Since(s2))
		s2 = time.Now()
		img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))

		fmt.Println("7:", time.Since(s2))
		s2 = time.Now()
		w, err := os.Create(outputFileName)
		if err != nil {
			panic(err)
		}

		fmt.Println("8:", time.Since(s2))
		s2 = time.Now()
		if identify.Image.MimeType == "image/png" {
			if err != nil {
				panic(err)
			}
			optimized := lossypng.Compress(img, lossypng.RGBAConversion, PngQuantization)

			if err = png.Encode(w, optimized); err != nil {
				panic(err)
			}
		} else if identify.Image.MimeType == "image/jpeg" {
			if err = jpeg.Encode(w, img, nil); err != nil {
				panic(err)
			}
		} else {
			panic("invalid format")
		}
		fmt.Println("9:", time.Since(s2))

		fmt.Println("Time:", time.Since(start))
		fmt.Println("")
	}

	os.RemoveAll("tmp/")
}