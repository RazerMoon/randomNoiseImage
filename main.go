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

func generateStaticImage(ImgSize int, smoothing string, colours string) *image.RGBA {
	rand.Seed(time.Now().UnixNano())

	img := image.NewRGBA(image.Rect(0, 0, ImgSize, ImgSize))

	for x := 0; x < ImgSize; x++ {
		for y := 0; y < ImgSize; y++ {
			randomShade1 := uint8(rand.Intn(255))
			randomShade2 := uint8(rand.Intn(255))
			randomShade3 := uint8(rand.Intn(255))

			if smoothing == "disabled" {
				if 255-randomShade1 > 128 {
					randomShade1 = 255
				} else {
					randomShade1 = 0
				}

				if colours == "enabled" {
					if 255-randomShade2 > 128 {
						randomShade2 = 255
					} else {
						randomShade2 = 0
					}

					if 255-randomShade3 > 128 {
						randomShade3 = 255
					} else {
						randomShade3 = 0
					}
				}
			}

			if colours == "disabled" {
				randomShade2 = randomShade1
				randomShade3 = randomShade1
			}

			img.Set(x, y, color.RGBA{randomShade1, randomShade2, randomShade3, 255})
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

	if len(os.Args) != 4 {
		fmt.Println("Wrong number of arguments!")
		os.Exit(2)
	}

	ImgSize, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid dimensions passed as first argument!")
		os.Exit(2)
	}

	var smoothing string
	var colours string

	if os.Args[2] == "y" || os.Args[2] == "Y" {
		smoothing = "enabled"
	} else {
		smoothing = "disabled"
	}

	if os.Args[3] == "y" || os.Args[3] == "Y" {
		colours = "enabled"
	} else {
		colours = "disabled"
	}

	img := generateStaticImage(ImgSize, smoothing, colours)

	saveImg(img)

	fmt.Printf("Generated random noise image with dimensions %d, smoothing %s and colours %s", ImgSize, smoothing, colours)
}
