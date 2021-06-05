package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func generateStaticImage(ImgSize int, smoothing string) *image.RGBA {
	rand.Seed(time.Now().UnixNano())

	img := image.NewRGBA(image.Rect(0, 0, ImgSize, ImgSize))

	for x := 0; x < ImgSize; x++ {
		for y := 0; y < ImgSize; y++ {
			randomShade := uint8(rand.Intn(255))

			if smoothing == "enabled" {
				if 255-randomShade > 128 {
					randomShade = 255
				} else {
					randomShade = 0
				}
			}

			img.Set(x, y, color.RGBA{randomShade, randomShade, randomShade, 255})
		}
	}

	return img
}

func saveImg(img image.Image) {
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()

	png.Encode(f, img)
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Wrong number of arguments!")
		os.Exit(2)
	}

	ImgSize, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid dimensions passed as first argument!")
		os.Exit(2)
	}

	var smoothing string

	if os.Args[2] == "y" || os.Args[2] == "Y" {
		smoothing = "enabled"
	} else {
		smoothing = "disabled"
	}

	img := generateStaticImage(ImgSize, smoothing)

	saveImg(img)

	fmt.Printf("Generated random noise image with dimensions %d and smoothing %s", ImgSize, smoothing)
}
