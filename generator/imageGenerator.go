package generator

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
)

type imageGenerator struct {
	width  int
	height int
	path string
	name string
	img    image.Image
}

func newImageGenerator(width int, height int, path, name string) *imageGenerator {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			pixel := getRandomColor()
			img.Set(x, y, pixel)
		}
	}

	i := &imageGenerator{width: width, height: height, path: path, name: name, img: img}

	return i
}

func getRandomColor() color.RGBA {
	r := color.RGBA{uint8(rand.Uint32()), uint8(rand.Uint32()), uint8(rand.Uint32()), 0xff}
	return r
}

func (i imageGenerator) Generate() error {
	out, err := os.Create(i.path + i.name + ".png")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer out.Close()
	if err := png.Encode(out, i.img); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
