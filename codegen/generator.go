package codegen

type Generator interface {
	Write(folder string) error
}
