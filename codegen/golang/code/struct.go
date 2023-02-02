package code

import (
	"github.com/google/uuid"
)

type Struct struct {
	id      string
	Name    string
	Comment Comment
	Package *Package
	Fields  Fields
	Methods Methods
}

func NewStruct(name string, comment Comment, pkg *Package, fields Fields, methods Methods) *Struct {
	return &Struct{uuid.New().String(), name, comment, pkg, fields, methods}
}

func (me *Struct) WithPackage(pkg *Package) Type {
	return &Struct{uuid.New().String(), me.Name, me.Comment, pkg, me.Fields, me.Methods}
}

func (me *Struct) RenameTo(name string) bool {
	if _, found := me.Package.Types[name]; found {
		return false
	}
	me.Name = name

	return true
}

func (me *Struct) ID() string {
	return me.id
}

func (me *Struct) Equals(other Type) bool {
	if actual, ok := other.(*Struct); ok {
		return me.Name == actual.Name
	}
	return false
}
