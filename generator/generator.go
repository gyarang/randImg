package generator

type GenerateType int

const (
	IMAGE GenerateType = iota
	VIDEO
)

type Generator interface {
	Generate() error
}
