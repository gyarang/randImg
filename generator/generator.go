package generator

type GenerateType int

const (
	IMAGE GenerateType = iota + 1
	VIDEO
)

type generator interface {
	Generate() error
}
