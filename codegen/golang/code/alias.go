package code

import "github.com/google/uuid"

type Alias struct {
	id      string
	Name    string
	Comment Comment
	Package *Package
	Elem    Type
	Methods Methods
}

func NewAlias(name string, comment Comment, pkg *Package, elem Type, methods Methods) *Alias {
	return &Alias{uuid.New().String(), name, comment, pkg, elem, methods}
}

func (me *Alias) WithPackage(pkg *Package) Type {
	return &Alias{uuid.New().String(), me.Name, me.Comment, pkg, me.Elem, me.Methods}
}

func (me *Alias) RenameTo(name string) bool {
	if _, found := me.Package.Types[name]; found {
		return false
	}
	me.Name = name

	return true
}

func (me *Alias) ID() string {
	return me.id
}

func (me *Alias) Equals(other Type) bool {
	if actual, ok := other.(*Alias); ok {
		return me.Elem.Equals(actual.Elem)
	}
	return false
}
