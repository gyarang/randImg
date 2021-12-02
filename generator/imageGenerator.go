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
	path   string
	name   string
	img    image.Image
}

func newImageGenerator(width int, height int, path, name string) *imageGenerator {
	img := generateRandomImage(width, height)
	i := &imageGenerator{width: width, height: height, path: path, name: name, img: img}
	return i
}

func generateRandomImage(w, h int) *image.RGBA {
	start := image.Point{0, 0}
	end := image.Point{w, h}
	img := image.NewRGBA(image.Rectangle{start, end})

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			pixel := getRandomColor()
			img.Set(x, y, pixel)
		}
	}

	return img
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
