package golang

import (
	"fmt"
	"strings"

	"github.com/dtcookie/dynatrace/gensettings/reflection"
)

func PointerIfOptional(t *reflection.Type, isOptional bool) *reflection.Type {
	if !isOptional {
		return t
	}
	if t.Kind == reflection.SetAliasKind || t.Kind == reflection.SliceAliasKind {
		return t
	}
	return t.Pointer()
}

func TypeName(t *reflection.Type) string {
	switch t.Kind {
	case reflection.BooleanKind:
		return "bool"
	case reflection.StringKind:
		return "string"
	case reflection.IntKind:
		return "int"
	case reflection.FloatKind:
		return "float64"
	case reflection.PointerKind:
		return "*" + TypeName(t.Elem)
	case reflection.EnumKind:
		return t.Name
	case reflection.StructKind:
		return t.Name
	case reflection.ListKind, reflection.SetKind:
		return "[]" + TypeName(t.Elem)
	case reflection.SliceAliasKind, reflection.SetAliasKind:
		return t.Name
	default:
		panic(fmt.Sprintf("unable to determine type name for %v", t))
	}
}

func packageName(schemaID string) string {
	schemaID = strings.ReplaceAll(schemaID, ":", ".")
	schemaID = strings.ReplaceAll(schemaID, "generic.type", "generictype")
	parts := strings.Split(schemaID, ".")
	result := strings.ReplaceAll(parts[len(parts)-1], "-", "")
	if result == "go" {
		result = "golang"
	}
	return result
}
