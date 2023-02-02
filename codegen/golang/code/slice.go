package code

import "github.com/google/uuid"

type Slice struct {
	id   string
	Elem Type
}

func NewSlice(elem Type) *Slice {
	return &Slice{uuid.New().String(), elem}
}

func (me *Slice) ID() string {
	return me.id
}

func (me *Slice) Equals(other Type) bool {
	if actual, ok := other.(*Slice); ok {
		return me.Elem.Equals(actual.Elem)
	}
	return false
}
