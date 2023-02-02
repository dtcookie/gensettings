package code

import "github.com/google/uuid"

type Packaged interface {
	WithPackage(*Package) Type
}

type Package struct {
	id       string
	Name     string
	Parent   *Package
	Packages map[string]*Package
	Types    map[string]Type
}

func NewPackage(name string, parent *Package) *Package {
	return &Package{id: uuid.New().String(), Name: name, Parent: parent, Packages: map[string]*Package{}, Types: map[string]Type{}}
}

func (me *Package) ID() string {
	return me.id
}

func (me *Package) Equals(other *Package) bool {
	if other == nil {
		return false
	}
	return me.id == other.id
}

func (me *Package) own(t Type) bool {
	switch actual := t.(type) {
	case *Struct:
		actual.Package = me
		return true
	case *Alias:
		actual.Package = me
		return true
	case *Enum:
		actual.Package = me
		return true
	default:
		return false
	}
}

func (me *Package) Replace(current Type, replacement Type) {

	id := current.ID()
	delete(me.Types, id)
	if me.own(replacement) {
		me.Types[id] = replacement
	}

	col := collector{}
	for _, t := range me.Types {
		col.replace(t, current, replacement)
	}
}
