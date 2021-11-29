package generator

import (
	"testing"
)

func TestGetRandomColor(t *testing.T) {
	getRandomColor()
}

func BenchmarkGetRandomColor(b *testing.B) {
	getRandomColor()
}

func BenchmarkNewImageGeneratorFHD(b *testing.B) {
	newImageGenerator(1920, 1080, "./", "asd")
}

func BenchmarkNewImageGenerator100(b *testing.B) {
	newImageGenerator(100, 100, "./", "asd")
}

func BenchmarkGenerateFHD(b *testing.B) {
	image := newImageGenerator(1920, 1080, "./", "testFile")
	b.ResetTimer()
	image.Generate()
}

func BenchmarkGenerate100(b *testing.B) {
	image := newImageGenerator(100, 100, "./", "testFile")
	b.ResetTimer()
	image.Generate()
}

