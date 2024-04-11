package generator

import (
	"image"
	"image/color"
	"image/png"
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
	start := image.Point{}
	end := image.Point{X: w, Y: h}
	img := image.NewRGBA(image.Rectangle{Min: start, Max: end})

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			pixel := getRandomColor()
			img.Set(x, y, pixel)
		}
	}

	return img
}

func getRandomColor() color.RGBA {
	base := rand.Uint32()
	r := uint8(base)
	base = base >> 8
	g := uint8(base)
	base = base >> 8
	b := uint8(base)

	c := color.RGBA{R: r, G: g, B: b}
	return c
}

func (i *imageGenerator) Generate() error {
	out, err := os.Create(i.path + i.name + ".png")
	if err != nil {
		return err
	}
	defer func() {
		err = out.Close()
	}()
	_ = png.Encode(out, i.img)

	return err
}
