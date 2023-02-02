package code

type Type interface {
	ID() string
	Equals(Type) bool
}
