package golang

import (
	"bytes"
	"text/template"

	"github.com/dtcookie/dynatrace/gensettings/reflection"

	"github.com/dtcookie/dynatrace/gensettings/codegen/camel"
)

type SetAlias struct {
	Name string
	Elem string
	Item string
}

func (me *SetAlias) Prettify() {
	me.Name = PrettyProp(me.Name)
}

func (me *SetAlias) Bytes() ([]byte, error) {
	tmpl, err := template.New("GO_SET_ALIAS").Parse(GO_SET_ALIAS_TEMPLATE)
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

func (me *SetAlias) FileName() string {
	return camel.Strip(me.Elem) + ".go"
}

func (me *SetAlias) Imports() []string {
	return []string{"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl", "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"}
}

func (me *SetAlias) Kind() CodeContributorKind {
	return AliasContrib
}

func (me *SetAlias) SortName() string {
	return me.Name
}

func NewSetAlias(t *reflection.Type, item string) *SetAlias {
	return &SetAlias{
		Name: Plural(TypeName(t.Elem.Unref())),
		Elem: TypeName(t.Elem.Unref()),
		Item: camel.Strip(item),
	}
}
