package generator

type GenerateType int

const (
	IMAGE = 0
	VIDEO = 1
)

type generator interface {
	Generate() error
}

func NewGenerator(g GenerateType, w, h int, p, n string) generator {
	switch g {
	case IMAGE:
		ig := newImageGenerator(w, h, p, n)
		return ig
	}
	return nil
}
