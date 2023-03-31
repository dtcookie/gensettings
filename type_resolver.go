package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/dtcookie/dynatrace/gensettings/reflection"

	"github.com/dtcookie/dynatrace/gensettings/schema"
	"github.com/dtcookie/dynatrace/gensettings/schema/property"
)

type TypeResolver struct {
	Library *reflection.Library
	Schema  *schema.Definition
}

func NewTypeResolver(sch *schema.Definition, jsonSchema []byte) *TypeResolver {
	return &TypeResolver{Library: reflection.NewLibrary(jsonSchema), Schema: sch}
}

func (me *TypeResolver) Resolve(name string) *reflection.Type {
	if strings.HasPrefix(name, "#/enums/") {
		return me.ResolveEnum(strings.TrimPrefix(name, "#/enums/"))
	} else if strings.HasPrefix(name, "#/types/") {
		return me.ResolveType(strings.TrimPrefix(name, "#/types/"))
	}
	panic(fmt.Sprintf("unable to resolve type %s", name))
}

func (me *TypeResolver) ResolveEnum(name string) *reflection.Type {
	enumDefinition, found := me.Schema.Enums[name]
	if !found {
		panic(fmt.Sprintf("schema doesn't contain type %s", name))
	}
	enumType := &reflection.Type{
		ID:         name,
		Kind:       reflection.EnumKind,
		Properties: reflection.Properties{},
	}
	enumType = me.Library.Define(enumType, false)
	for _, item := range enumDefinition.Items {
		property := &reflection.Property{
			Name: item.Value.(string),
			Type: enumType,
		}
		enumType.Properties[property.Name] = property
	}
	return enumType
}

func (me *TypeResolver) ResolveType(name string) *reflection.Type {
	typeDefinition, found := me.Schema.Types[name]
	if !found {
		panic(fmt.Sprintf("schema doesn't contain type %s", name))
	}
	comment := typeDefinition.DisplayName
	if len(typeDefinition.Description) > 0 {
		comment = comment + ". " + typeDefinition.Description
	}
	if len(typeDefinition.Documentation) > 0 {
		comment = comment + ". " + typeDefinition.Documentation
	}
	comment = strings.TrimSpace(comment)
	if comment == typeDefinition.DisplayName {
		comment = ""
	}
	structType := &reflection.Type{
		ID:         name,
		Kind:       reflection.StructKind,
		Properties: reflection.Properties{},
		Comment:    comment,
	}
	structType = me.Library.Define(structType, false)
	for propertyName, propertyDefinition := range typeDefinition.Properties {
		comment := ""
		if len(propertyDefinition.Description) > 0 {
			comment = propertyDefinition.Description
		}
		if len(propertyDefinition.Documentation) > 0 {
			comment = comment + ". " + propertyDefinition.Documentation
		}
		if len(comment) == 0 {
			comment = propertyDefinition.DisplayName
		}
		if strings.EqualFold(comment, propertyName) {
			comment = ""
		}
		comment = strings.TrimSpace(strings.TrimPrefix(strings.ReplaceAll(strings.TrimSpace(comment), "Type '{' for placeholder hints.", ""), "."))
		property := &reflection.Property{
			Name:            propertyName,
			Type:            me.resolvePropertyType(propertyName, propertyDefinition),
			Comment:         comment,
			Optional:        propertyDefinition.Nullable || propertyDefinition.Precondition != nil || (propertyDefinition.MinObjects != nil && *propertyDefinition.MinObjects == 0),
			OptionalComment: OptionalComment(propertyDefinition),
			Precondition:    translatePrecondition(propertyDefinition),
			Constraints:     translateConstraints(propertyDefinition),
		}
		structType.Properties[propertyName] = property
	}
	return structType
}

func translateConstraints(propertyDefinition property.Definition) []map[string]any {
	if len(propertyDefinition.Constraints) == 0 {
		return nil
	}
	constraints := []map[string]any{}
	for _, constraint := range propertyDefinition.Constraints {
		data, _ := json.Marshal(constraint)
		m := map[string]any{}
		json.Unmarshal(data, &m)
		constraints = append(constraints, m)
	}
	return constraints
}

func translatePrecondition(propertyDefinition property.Definition) map[string]any {
	if propertyDefinition.Precondition == nil {
		return nil
	}
	data, _ := json.Marshal(propertyDefinition.Precondition)
	m := map[string]any{}
	json.Unmarshal(data, &m)
	return m
}

func OptionalComment(propertyDefinition property.Definition) string {
	s := ""
	if propertyDefinition.Nullable {
		s = s + " & nullable"
	}
	if propertyDefinition.Precondition != nil {
		s = s + " & precondition"
	}
	if propertyDefinition.MinObjects != nil && *propertyDefinition.MinObjects == 0 {
		s = s + " & minobjects == 0"
	}
	return strings.TrimPrefix(s, " & ")
}

func (me *TypeResolver) resolvePropertyType(propertyName string, propertyDefinition property.Definition) *reflection.Type {
	switch propertyType := propertyDefinition.Type.(type) {
	case string:
		switch propertyType {
		case "boolean":
			return &reflection.TypeBoolean
		case "text":
			return &reflection.TypeString
		case "local_date_time":
			return &reflection.TypeString
		case "zoned_date_time":
			return &reflection.TypeString
		case "local_time":
			return &reflection.TypeString
		case "time_zone":
			return &reflection.TypeString
		case "local_date":
			return &reflection.TypeString
		case "setting":
			return &reflection.TypeString
		case "secret":
			return &reflection.TypeString
		case "integer":
			return &reflection.TypeInt
		case "float":
			return &reflection.TypeFloat
		case "list":
			if propertyDefinition.Items == nil {
				panic(fmt.Sprintf("property %s is of type list, but Items is nil", propertyName))
			}
			itemType := me.resolvePropertyType(propertyName, property.Definition{
				Type: propertyDefinition.Items.Type,
			})
			for _, constraint := range propertyDefinition.Constraints {
				if constraint.CustomValidatorID == "unique-tag-list-validator" {
					return itemType.Set()
				}
			}
			return itemType.List()
		case "set":
			if propertyDefinition.Items == nil {
				panic(fmt.Sprintf("property %s is of type list, but Items is nil", propertyName))
			}
			itemType := me.resolvePropertyType(propertyName, property.Definition{
				Type: propertyDefinition.Items.Type,
			})
			return itemType.Set()
		default:
			panic(fmt.Sprintf("unsupported string property type %v", propertyType))
		}
	case map[string]any:
		if len(propertyType) != 1 {
			panic(fmt.Sprintf("found property type %v with unexpected len %d", propertyType, len(propertyType)))
		}
		if refv, found := propertyType["$ref"]; found {
			ref, ok := refv.(string)
			if !ok {
				panic(fmt.Sprintf("found property type %v with $ref entry of type %T", propertyType, refv))
			}
			if strings.HasPrefix(ref, "#/enums/") {
				return me.Resolve(ref)
			} else if strings.HasPrefix(ref, "#/types/") {
				return me.Resolve(ref)
			} else {
				panic(fmt.Sprintf("found property type %v with $ref entry %v", propertyType, refv))
			}
		} else {
			panic(fmt.Sprintf("found property type %v without expected $ref entry", propertyType))
		}
	default:
		panic(fmt.Sprintf("unsupported property type %v", propertyType))
	}
}
