package generator

const (
	IMAGE = 0
	VIDEO = 1
)

func NewGenerator(g GenerateType, w, h int, p, n string) generator {
	switch g {
	case IMAGE:
		ig := newImageGenerator(w, h, p, n)
		return ig
	}
	return nil
}
