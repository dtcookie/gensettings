package code

import "github.com/google/uuid"

type Pointer struct {
	id   string
	Elem Type
}

func NewPointer(elem Type) *Pointer {
	return &Pointer{uuid.New().String(), elem}
}

func (me *Pointer) ID() string {
	return me.id
}

func (me *Pointer) Equals(other Type) bool {
	if actual, ok := other.(*Pointer); ok {
		return me.Elem.Equals(actual.Elem)
	}
	return false
}
