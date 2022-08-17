package generator

import (
	"testing"
)

func TestGetRandomColor(t *testing.T) {
	getRandomColor()
}

func BenchmarkGetRandomColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getRandomColor()
	}
}

func BenchmarkNewImageGeneratorFHD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateRandomImage(1920, 1080)
	}
}

func BenchmarkNewImageGenerator100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateRandomImage(100, 100)
	}
}

func BenchmarkGenerateFHD(b *testing.B) {
	image := newImageGenerator(1920, 1080, "./", "testFile")
	for i := 0; i < b.N; i++ {
		image.Generate()
	}
}

func BenchmarkGenerate100(b *testing.B) {
	image := newImageGenerator(100, 100, "./", "testFile")
	for i := 0; i < b.N; i++ {
		image.Generate()
	}
}
