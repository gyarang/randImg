package generator

func NewGenerator(g GenerateType, w, h int, p, n string) Generator {
	switch g {
	case IMAGE:
		ig := newImageGenerator(w, h, p, n)
		return ig
	}
	return nil
}
