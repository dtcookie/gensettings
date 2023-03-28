package golang

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/dtcookie/dynatrace/gensettings/reflection"

	"github.com/dtcookie/dynatrace/gensettings/codegen/camel"
)

type Property struct {
	Name            string
	Type            string
	Comment         string
	HCLTag          string
	JSONTag         string
	Optional        bool
	OptionalComment string
	MinItems        int
	MaxItems        int
	HCLType         string
	Elem            string
	Default         string
	Precondition    map[string]any
}

func (me *Property) Prettify() {
	me.Name = PrettyProp(me.Name)
}

type Struct struct {
	Name       string
	Plural     string
	Comments   []string
	Properties []*Property
}

/*
	if me.IgnoreCase == nil && slices.Contains([]string{"TagEquals", "TagKeyEquals", "StringEndsWith", "NotStringEndsWith", "StringStartsWith", "NotStringStartsWith", "StringContains", "NotStringContains", "StringEquals", "NotStringEquals"}, string(me.CompareOperationType)) {
		me.IgnoreCase = opt.NewBool(false)
	}

	// IgnoreCase *bool -> {"expectedValues":["TagEquals","TagKeyEquals","StringEndsWith","NotStringEndsWith","StringStartsWith","NotStringStartsWith","StringContains","NotStringContains","StringEquals","NotStringEquals"],"property":"compareOperationType","type":"IN"}
*/

func getPrecExpectedValues(m map[string]any) []string {
	untyped, ok := m["expectedValues"]
	if !ok {
		return nil
	}
	switch typed := untyped.(type) {
	case []any:
		strs := []string{}
		for _, uts := range typed {
			strs = append(strs, uts.(string))
		}
		return strs
	case []string:
		return typed
	}
	return nil
}

func getPrecExpectedValue(m map[string]any) any {
	untyped, ok := m["expectedValue"]
	if !ok {
		return nil
	}
	return untyped
}

func getPrecProperty(m map[string]any) string {
	untyped, ok := m["property"]
	if !ok {
		return ""
	}
	return untyped.(string)
}

func getPrecType(m map[string]any) string {
	untyped, ok := m["type"]
	if !ok {
		return ""
	}
	return untyped.(string)
}

func (me *Struct) Preconditions() string {
	result := ""
	if len(me.Properties) == 0 {
		return result
	}
	unhandledLines := []string{}
	for _, property := range me.Properties {
		if len(property.Precondition) == 0 {
			continue
		}
		m := property.Precondition
		precExpectedValues := getPrecExpectedValues(m)
		precExpectedValue := getPrecExpectedValue(m)
		precPropertyName := getPrecProperty(m)
		precType := getPrecType(m)
		line := ""
		if precExpectedValues != nil && len(precPropertyName) > 0 && precType == "IN" {
			precPropertyName = PropertyName(precPropertyName)
			line = fmt.Sprintf(`if me.%s == nil && slices.Contains([]string{"%s"}, string(me.%s)) {
				me.%s = `, property.Name, strings.Join(precExpectedValues, `", "`), precPropertyName, property.Name)
		} else if precExpectedValue != nil && len(precPropertyName) > 0 && precType == "EQUALS" {
			precPropertyName = PropertyName(precPropertyName)
			switch ev := precExpectedValue.(type) {
			case string:
				line = fmt.Sprintf(`if me.%s == nil && string(me.%s) == "%s" {
					me.%s = `, property.Name, precPropertyName, precExpectedValue, property.Name)
			case bool:
				if ev {
					line = fmt.Sprintf(`if me.%s == nil && me.%s {
							me.%s = `, property.Name, precPropertyName, property.Name)
				} else {
					line = fmt.Sprintf(`if me.%s == nil && !me.%s {
							me.%s = `, property.Name, precPropertyName, property.Name)
				}
			}
		}
		if len(line) > 0 && property.Type == "*bool" {
			result = result + "\n" + fmt.Sprintf("%s%s }", line, "opt.NewBool(false)")
		} else if len(line) > 0 && property.Type == "*int" {
			result = result + "\n" + fmt.Sprintf("%s%s }", line, "opt.NewInt(0)")
		} else if len(line) > 0 && property.Type == "*string" {
			result = result + "\n" + fmt.Sprintf("%s%s }", line, "opt.NewString(\"\")")
		} else {
			unhandledLines = append(unhandledLines, "// ---- "+property.Name+" "+property.Type)
		}
	}
	for _, line := range unhandledLines {
		result = result + "\n" + line
	}
	return strings.TrimSpace(result)
}

func (me *Struct) Bytes() ([]byte, error) {
	tmpl, err := template.New("GO_STRUCT").Parse(GO_STRUCT_TEMPLATE)
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

func (me *Struct) Prettify() {
	me.Name = PrettyProp(me.Name)
	me.Plural = Plural(me.Name)
	if len(me.Properties) > 0 {
		for _, property := range me.Properties {
			property.Prettify()
		}
	}
}

func (me *Struct) FileName() string {
	return camel.Strip(me.Name) + ".go"
}

func (me *Struct) Imports() []string {
	return []string{"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl", "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"}
}

func (me *Struct) Kind() CodeContributorKind {
	return StructContrib
}

func (me *Struct) SortName() string {
	return me.Name
}

func NewStruct(t *reflection.Type) *Struct {
	structDef := Struct{
		Name:     TypeName(t),
		Plural:   Plural(TypeName(t)),
		Comments: Comments(strings.Split(t.Comment, "\n")),
	}

	for propertyName, property := range t.Properties {
		propType := property.Type
		if propType.Kind == reflection.StructKind {
			propType = propType.Pointer()
		}
		// if propType.Kind == reflection.ListKind {
		// 	if propType.Elem.Kind == reflection.StructKind {
		// 		propType = propType.Elem.Pointer().List()
		// 	}
		// }
		// if propType.Kind == reflection.SetKind {
		// 	if propType.Elem.Kind == reflection.StructKind {
		// 		propType = propType.Elem.Pointer().Set()
		// 	}
		// }

		jsonTag := propertyName
		if property.Optional {
			jsonTag = jsonTag + ",omitempty"
		}
		defVal := ""
		if len(property.Scope) > 0 {
			jsonTag = "-\" scope:\"" + property.Scope
			if property.Optional {
				defVal = "environment"
			}
		}
		comment := strings.ReplaceAll(Comment(property.Comment), "\n", "\\n")
		if propType.Kind == reflection.EnumKind {
			comment = "Possible Values:"
			pnames := []string{}
			for _, p := range propType.Properties {
				pnames = append(pnames, p.Name)
			}
			sort.Strings(pnames)
			sep := " "
			for _, pname := range pnames {
				comment = comment + sep + "`" + pname + "`"
				sep = ", "
			}
		}
		if strings.ToLower(property.Name) == "enabled" {
			comment = "This setting is enabled (`true`) or disabled (`false`)"
		}
		structDef.Properties = append(structDef.Properties, HCLKind(propType, &Property{
			Name:            PropertyName(propertyName),
			Comment:         comment,
			HCLTag:          camel.Strip(propertyName),
			JSONTag:         jsonTag,
			Type:            TypeName(PointerIfOptional(propType, property.Optional)),
			Optional:        property.Optional,
			OptionalComment: property.OptionalComment,
			Default:         defVal,
			Precondition:    property.Precondition,
		}))
	}
	sort.Slice(structDef.Properties, func(i, j int) bool {
		return strings.Compare(structDef.Properties[i].Name, structDef.Properties[j].Name) == -1
	})
	return &structDef
}

func toString(m map[string]any) string {
	if m == nil {
		return ""
	}
	data, _ := json.Marshal(m)
	return string(data)
}

func StructTypeName(name string) string {
	return strings.ToUpper(name[0:1]) + name[1:]
}

func PropertyName(name string) string {
	propertyName := strings.ToUpper(name[0:1]) + name[1:]
	propertyName = strings.ReplaceAll(propertyName, " ", "")
	propertyName = strings.ReplaceAll(propertyName, "-", "_")
	propertyName = strings.ReplaceAll(propertyName, "+", "_")
	propertyName = strings.ReplaceAll(propertyName, ".", "_")
	propertyName = strings.ReplaceAll(propertyName, ":", "_")
	return PrettyProp(propertyName)
}

func HCLKind(propType *reflection.Type, property *Property) *Property {
	switch propType.UnrefPtr().Kind {
	case reflection.BooleanKind:
		property.HCLType = "schema.TypeBool"
		// property.Optional = true
	case reflection.FloatKind:
		property.HCLType = "schema.TypeFloat"
	case reflection.EnumKind:
		property.HCLType = "schema.TypeString"
	case reflection.StringKind:
		property.HCLType = "schema.TypeString"
	case reflection.IntKind:
		property.HCLType = "schema.TypeInt"
	case reflection.StructKind:
		property.HCLType = "schema.TypeList"
		property.MaxItems = 1
		property.MinItems = 1
		property.Elem = "&schema.Resource{Schema: new(" + TypeName(propType.Elem.Unref()) + ").Schema()}"
	case reflection.SliceAliasKind, reflection.SetAliasKind:
		property.HCLType = "schema.TypeList"
		property.MinItems = 1
		property.MaxItems = 1
		property.Elem = "&schema.Resource{Schema: new(" + TypeName(propType) + ").Schema()}"
	case reflection.ListKind:
		property.HCLType = "schema.TypeList"
		switch propType.Elem.Kind {
		case reflection.StringKind:
			property.Elem = "&schema.Schema{Type: schema.TypeString}"
		case reflection.EnumKind:
			property.Elem = "&schema.Schema{Type: schema.TypeString}"
		case reflection.IntKind:
			property.Elem = "&schema.Schema{Type: schema.TypeInt}"
		case reflection.FloatKind:
			property.Elem = "&schema.Schema{Type: schema.TypeFloat}"
		case reflection.BooleanKind:
			property.Elem = "&schema.Schema{Type: schema.TypeBool}"
		default:
			property.MinItems = 1
			property.Elem = "&schema.Resource{Schema: new(" + TypeName(propType.Elem.Unref()) + ").Schema()}"
		}
	case reflection.SetKind:
		property.HCLType = "schema.TypeSet"
		switch propType.Elem.Kind {
		case reflection.StringKind:
			property.Elem = "&schema.Schema{Type: schema.TypeString}"
		case reflection.EnumKind:
			property.Elem = "&schema.Schema{Type: schema.TypeString}"
		case reflection.IntKind:
			property.Elem = "&schema.Schema{Type: schema.TypeInt}"
		case reflection.FloatKind:
			property.Elem = "&schema.Schema{Type: schema.TypeFloat}"
		case reflection.BooleanKind:
			property.Elem = "&schema.Schema{Type: schema.TypeBool}"
		default:
			property.MinItems = 1
			property.Elem = "&schema.Resource{Schema: new(" + TypeName(propType.Elem.Unref()) + ").Schema()}"
		}
	}
	return property
}
