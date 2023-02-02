package code

import "fmt"

type collector map[string]string

func (me collector) replace(t Type, current Type, replacement Type) {
	id := current.ID()
	if _, found := me[id]; found {
		return
	}
	me[id] = id
	switch actual := t.(type) {
	case *Alias:
		if actual.Elem.ID() == id {
			actual.Elem = replacement
		} else {
			me.replace(actual.Elem, current, replacement)
		}
	case *Pointer:
		if actual.Elem.ID() == id {
			actual.Elem = replacement
		} else {
			me.replace(actual.Elem, current, replacement)
		}
	case *Slice:
		if actual.Elem.ID() == id {
			actual.Elem = replacement
		} else {
			me.replace(actual.Elem, current, replacement)
		}
	case *Primitive:
		// ignore
	case *Struct:
		for _, field := range actual.Fields {
			fieldType := field.Type
			if fieldType.ID() == id {
				field.Type = replacement
			} else {
				me.replace(fieldType, current, replacement)
			}
		}
		for _, method := range actual.Methods {
			for _, argument := range method.Arguments {
				argumentType := argument.Type
				if argumentType.ID() == id {
					argument.Type = replacement
				} else {
					me.replace(argumentType, current, replacement)
				}
			}
			for idx, returnType := range method.Returns {
				if returnType.ID() == id {
					method.Returns[idx] = replacement
				} else {
					me.replace(returnType, current, replacement)
				}
			}
		}
	}
}

func (me collector) imports(t Type) []*Package {
	packages := map[string]*Package{}
	me.collectImports(t, packages)
	result := []*Package{}
	for _, pkg := range packages {
		result = append(result, pkg)
	}
	return result
}

func (me collector) collectImports(t Type, packages map[string]*Package) {
	typeID := t.ID()
	if _, found := me[typeID]; found {
		return
	}
	me[typeID] = typeID
	switch actual := t.(type) {
	case *Alias:
		packages[actual.Package.id] = actual.Package
		me.collectImports(actual.Elem, packages)
	case *Pointer:
		me.collectImports(actual.Elem, packages)
	case *Slice:
		me.collectImports(actual.Elem, packages)
	case *Struct:
		packages[actual.Package.id] = actual.Package
	case *Enum:
		packages[actual.Package.id] = actual.Package
	case *Primitive:
	default:
		panic(fmt.Sprintf("import - unsupported type %T", t))
	}
}
