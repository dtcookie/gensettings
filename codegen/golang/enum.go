package golang

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/dtcookie/dynatrace/gensettings/codegen/camel"
	"github.com/dtcookie/dynatrace/gensettings/reflection"
)

type EnumField struct {
	Name    string
	Literal string
}

type Enum struct {
	Name      string
	Plural    string
	Instances []EnumField
}

func (me *Enum) Prettify() {
	me.Name = PrettyProp(me.Name)
}

func (me *Enum) Kind() CodeContributorKind {
	return EnumContrib
}

func (me *Enum) SortName() string {
	return me.Name
}

func NewEnum(t *reflection.Type) *Enum {
	enumDef := &Enum{
		Name:   EnumTypeName(t.ID),
		Plural: Plural(EnumTypeName(t.ID)),
	}
	for propertyName, property := range t.Properties {
		enumDef.Instances = append(enumDef.Instances, EnumField{
			Name:    EnumInstance(enumDef.Name, propertyName),
			Literal: property.RawName,
		})
	}
	sort.Slice(enumDef.Instances, func(i, j int) bool {
		return strings.Compare(enumDef.Instances[i].Name, enumDef.Instances[j].Name) < 0
	})
	return enumDef
}

func (me *Enum) Bytes() ([]byte, error) {
	tmpl, err := template.New("GO_ENUM").Parse(GO_ENUM_TEMPLATE)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, me)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (me *Enum) FileName() string {
	return "enums.go"
}

func (me *Enum) Imports() []string {
	return nil
}

func EnumTypeName(name string) string {
	return camel.Camel(camel.Strip(name))
}

func EnumInstance(typeName string, name string) string {
	propertyName := strings.ToUpper(name[0:1]) + name[1:]
	propertyName = strings.ReplaceAll(propertyName, " ", "")
	propertyName = strings.ReplaceAll(propertyName, "-", "_minus_")
	propertyName = strings.ReplaceAll(propertyName, "+", "_plus_")
	propertyName = strings.ReplaceAll(propertyName, ".", "_")
	propertyName = strings.ReplaceAll(propertyName, ":", "_")
	if i, err := strconv.Atoi(propertyName); err == nil {
		if i < 0 {
			propertyName = fmt.Sprintf("%s_minus_%d", PrettyProp(typeName), -i)
		} else {
			propertyName = fmt.Sprintf("%s%s", PrettyProp(typeName), strings.TrimSpace(propertyName))
		}
		return propertyName
	}
	return camel.Camel(propertyName)
}
