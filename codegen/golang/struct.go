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
	IsEnumType      bool
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
	Constraints     []map[string]any
	Sensitive       bool
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

func (me *Struct) findProperty(origName string) *Property {
	name := PropertyName(origName)
	for _, property := range me.Properties {
		if property.Name == name {
			return property
		}
	}
	return nil
}

func (me *Struct) HasPreconditions() bool {
	if len(me.Properties) == 0 {
		return false
	}
	for _, property := range me.Properties {
		if len(property.Precondition) > 0 {
			return true
		}
	}
	return false
}

func (me *Struct) Preconditions() string {
	result := ""
	if len(me.Properties) == 0 {
		return result
	}
	unhandledLines := []string{}
	handledLines := []string{}
	throwingLines := []string{}
	for _, property := range me.Properties {
		if len(property.Precondition) == 0 {
			continue
		}
		m := property.Precondition
		precExpectedValues := getPrecExpectedValues(m)
		precExpectedValue := getPrecExpectedValue(m)
		precPropName := getPrecProperty(m)
		precType := getPrecType(m)

		negate := false
		if precType == "NOT" {
			m = m["precondition"].(map[string]any)
			negate = true
			precExpectedValues = getPrecExpectedValues(m)
			precExpectedValue = getPrecExpectedValue(m)
			precPropName = getPrecProperty(m)
			precType = getPrecType(m)
		}
		var checkProperty *Property
		if len(precPropName) > 0 {
			checkProperty = me.findProperty(precPropName)
		}

		line := ""
		if precExpectedValues != nil && checkProperty != nil && precType == "IN" {
			negStr := ""
			if negate {
				negStr = "!"
			}
			nilCheck := ""
			deref := ""
			if strings.HasPrefix(checkProperty.Type, "*") {
				nilCheck = fmt.Sprintf(" && me.%s != nil", checkProperty.Name)
				deref = "*"
			}
			line = fmt.Sprintf(`if me.%s == nil%s && %sslices.Contains([]string{"%s"}, string(%sme.%s)) {
				`, property.Name, nilCheck, negStr, strings.Join(precExpectedValues, `", "`), deref, checkProperty.Name)
		} else if precExpectedValue != nil && checkProperty != nil && precType == "EQUALS" {
			switch ev := precExpectedValue.(type) {
			case string:
				eqStr := "=="
				if negate {
					eqStr = "!="
				}
				nilCheck := ""
				deref := ""
				if strings.HasPrefix(checkProperty.Type, "*") {
					nilCheck = fmt.Sprintf(" && me.%s != nil", checkProperty.Name)
					deref = "*"
				}
				line = fmt.Sprintf(`if me.%s == nil%s && string(%sme.%s) %s "%s" {
					`, property.Name, nilCheck, deref, checkProperty.Name, eqStr, precExpectedValue)
			case bool:
				if negate {
					ev = !ev
				}
				nilCheck := ""
				deref := ""
				if strings.HasPrefix(checkProperty.Type, "*") {
					nilCheck = fmt.Sprintf(" && me.%s != nil", checkProperty.Name)
					deref = "*"
				}

				if ev {
					line = fmt.Sprintf(`if me.%s == nil%s && %sme.%s {
							`, property.Name, nilCheck, deref, checkProperty.Name)
				} else {
					line = fmt.Sprintf(`if me.%s == nil%s && %s!me.%s {
							`, property.Name, nilCheck, deref, checkProperty.Name)
				}
			}
		}
		if len(line) > 0 && property.Type == "*bool" {
			if HasConstraints(property) {
				pat := fmt.Sprintf("'%s' must be specified if '%s' is set to '%%v'", property.HCLTag, checkProperty.HCLTag)
				throwingLines = append(throwingLines, fmt.Sprintf("%s return fmt.Errorf(\"%s\", me.%s)}", line, pat, checkProperty.Name))
			} else {
				handledLines = append(handledLines, fmt.Sprintf("%sme.%s = %s}", line, property.Name, "opt.NewBool(false)"))
			}
		} else if len(line) > 0 && property.Type == "*int" {
			if HasConstraints(property) {
				pat := fmt.Sprintf("'%s' must be specified if '%s' is set to '%%v'", property.HCLTag, checkProperty.HCLTag)
				throwingLines = append(throwingLines, fmt.Sprintf("%s return fmt.Errorf(\"%s\", me.%s)}", line, pat, checkProperty.Name))
			} else {
				handledLines = append(handledLines, fmt.Sprintf("%sme.%s = %s}", line, property.Name, "opt.NewInt(0)"))
			}
		} else if len(line) > 0 && property.Type == "*string" {
			if HasConstraints(property) {
				pat := fmt.Sprintf("'%s' must be specified if '%s' is set to '%%v'", property.HCLTag, checkProperty.HCLTag)
				throwingLines = append(throwingLines, fmt.Sprintf("%s return fmt.Errorf(\"%s\", me.%s)}", line, pat, checkProperty.Name))
			} else {
				handledLines = append(handledLines, fmt.Sprintf("%sme.%s = %s}", line, property.Name, "opt.NewString(\"\")"))
			}
		} else if len(line) > 0 && property.Type == "*float64" {
			if HasConstraints(property) {
				pat := fmt.Sprintf("'%s' must be specified if '%s' is set to '%%v'", property.HCLTag, checkProperty.HCLTag)
				throwingLines = append(throwingLines, fmt.Sprintf("%s return fmt.Errorf(\"%s\", me.%s)}", line, pat, checkProperty.Name))
			} else {
				handledLines = append(handledLines, fmt.Sprintf("%sme.%s = %s}", line, property.Name, "opt.NewFloat64(0.0)"))
			}

		} else if len(line) > 0 && property.IsEnumType {
			pat := fmt.Sprintf("'%s' must be specified if '%s' is set to '%%v'", property.HCLTag, checkProperty.HCLTag)
			throwingLines = append(throwingLines, fmt.Sprintf("%s return fmt.Errorf(\"%s\", me.%s)}", line, pat, checkProperty.Name))
		} else {
			data, _ := json.Marshal(m)
			unhandledLines = append(unhandledLines, "// ---- "+property.Name+" "+property.Type+" -> "+string(data))
		}
	}
	for _, line := range handledLines {
		result = result + "\n" + line
	}
	for _, line := range throwingLines {
		result = result + "\n" + line
	}
	for _, line := range unhandledLines {
		result = result + "\n" + line
	}
	return strings.TrimSpace(result)
}

func HasConstraints(property *Property) bool {
	if property == nil {
		return false
	}
	if len(property.Constraints) == 0 {
		return false
	}
	for _, constraint := range property.Constraints {
		if isRoadBlockConstraint(constraint) {
			return true
		}
	}
	return false
}

func isRoadBlockConstraint(constraint map[string]any) bool {
	if constraint == nil {
		return false
	}
	minimum := 0
	if v, ok := constraint["minimum"]; ok {
		switch tv := v.(type) {
		case int:
			minimum = tv
		case float64:
			minimum = int(tv)
		default:
			panic(fmt.Sprintf("%T", v))
		}
	}
	minLength := 0
	if v, ok := constraint["minLength"]; ok {
		switch tv := v.(type) {
		case int:
			minLength = tv
		case float64:
			minLength = int(tv)
		default:
			panic(fmt.Sprintf("%T", v))
		}
	}

	if constraintType, ok := constraint["type"]; ok {
		switch constraintType {
		case "RANGE":
			if minimum > 0 {
				return true
			}
		case "LENGTH":
			if minLength > 0 {
				return true
			}
		}
	}
	return false
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
			IsEnumType:      property.Type.UnrefPtr().Kind == reflection.EnumKind,
			Optional:        property.Optional,
			OptionalComment: property.OptionalComment,
			Default:         defVal,
			Precondition:    property.Precondition,
			Constraints:     property.Constraints,
			Sensitive:       property.Sensitive,
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
