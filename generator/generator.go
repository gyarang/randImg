package generator

type GenerateType int
type generator interface {
	Generate() error
}
