package generator

type GenerateType int

const (
	IMAGE GenerateType = iota
	VIDEO
)

type generator interface {
	Generate() error
}
