package generator

type GenerateType int

const (
	IMAGE GenerateType = 1 + iota
	VIDEO
)

type generator interface {
	Generate() error
}
