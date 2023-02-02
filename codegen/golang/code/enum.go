package code

import (
	"github.com/dtcookie/dynatrace/gensettings/codegen/compare"

	"github.com/google/uuid"
)

type Enum struct {
	id      string
	Name    string
	Comment Comment
	Package *Package
	Values  []string
}

func NewEnum(name string, comment Comment, pkg *Package, values []string) *Enum {
	return &Enum{uuid.New().String(), name, comment, pkg, values}
}

func (me *Enum) WithPackage(pkg *Package) Type {
	return &Enum{uuid.New().String(), me.Name, me.Comment, pkg, me.Values}
}

func (me *Enum) RenameTo(name string) bool {
	if _, found := me.Package.Types[name]; found {
		return false
	}
	me.Name = name

	return true
}

func (me *Enum) ID() string {
	return me.id
}

func (me *Enum) Equals(other Type) bool {
	if actual, ok := other.(*Enum); ok {
		return compare.StringSlice(me.Values).Equals(actual.Values)
	}
	return false
}
