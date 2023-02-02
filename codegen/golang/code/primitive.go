package code

import (
	"github.com/google/uuid"
)

var Primitives = struct {
	Int     *Primitive
	Int8    *Primitive
	Int16   *Primitive
	Int32   *Primitive
	Int64   *Primitive
	UInt    *Primitive
	UInt8   *Primitive
	UInt16  *Primitive
	UInt32  *Primitive
	UInt64  *Primitive
	Float32 *Primitive
	Float64 *Primitive
	String  *Primitive
	Bool    *Primitive
}{
	&Primitive{uuid.New().String(), "int"},
	&Primitive{uuid.New().String(), "int8"},
	&Primitive{uuid.New().String(), "int16"},
	&Primitive{uuid.New().String(), "int32"},
	&Primitive{uuid.New().String(), "int64"},
	&Primitive{uuid.New().String(), "uint"},
	&Primitive{uuid.New().String(), "uint8"},
	&Primitive{uuid.New().String(), "uint16"},
	&Primitive{uuid.New().String(), "uint32"},
	&Primitive{uuid.New().String(), "uint64"},
	&Primitive{uuid.New().String(), "float32"},
	&Primitive{uuid.New().String(), "float64"},
	&Primitive{uuid.New().String(), "string"},
	&Primitive{uuid.New().String(), "bool"},
}

type Primitive struct {
	id   string
	name string
}

func (me *Primitive) ID() string {
	return me.id
}

func (me *Primitive) Name() string {
	return me.name
}

func (me *Primitive) Equals(other Type) bool {
	if primitive, ok := other.(*Primitive); ok {
		return primitive.id == me.id
	}
	return false
}
