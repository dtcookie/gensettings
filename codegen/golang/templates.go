package golang

const GO_LIST_ALIAS_TEMPLATE = `
type {{.Name}} []*{{.Elem}}

func (me *{{.Name}}) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"{{.Item}}": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new({{.Elem}}).Schema()},
		},
	}
}

func (me {{.Name}}) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("{{.Item}}", me)
}

func (me *{{.Name}}) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("{{.Item}}", me)
}
`
const GO_SET_ALIAS_TEMPLATE = `
type {{.Name}} []*{{.Elem}}

func (me *{{.Name}}) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"{{.Item}}": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new({{.Elem}}).Schema()},
		},
	}
}

func (me {{.Name}}) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("{{.Item}}", me)
}

func (me *{{.Name}}) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("{{.Item}}", me)
}
`

const GO_STRUCT_TEMPLATE = `
{{ range .Comments }}// {{.}}{{end}}
type {{.Name}} struct { {{with .Properties}}{{ range .}}
	{{.Name}} {{.Type}} ` + "`" + `json:"{{.JSONTag}}"` + "`" + ` {{if .Comment}}// {{.Comment}}{{end}}{{end}}{{end}}
}

func (me *{{.Name}}) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{ {{with .Properties}}{{ range .}}
		"{{.HCLTag}}":  {
			Type: {{.HCLType}},
			Description: "{{if .Comment}}{{.Comment}}{{else}}no documentation available{{end}}",
			{{if .Optional}} Optional: true {{else}} Required: true {{end}}, {{if .OptionalComment}}// {{.OptionalComment}}{{end}}{{if .Default}}
			Default: "{{.Default}}",{{end}}{{if .Elem}}
			Elem: {{.Elem}},{{end}}{{if .MinItems}}
			MinItems: {{.MinItems}},{{end}}{{if .MaxItems}}
			MaxItems: {{.MaxItems}},{{end}}
		},{{end}}{{end}}
	}
}

func (me *{{.Name}}) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{ {{with .Properties}}{{ range .}}
		"{{.HCLTag}}":  me.{{.Name}},{{end}}{{end}}
	})
}

func (me *{{.Name}}) HandlePreconditions() { {{if .Preconditions}}
{{.Preconditions}} {{end}}
}

func (me *{{.Name}}) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{ {{with .Properties}}{{ range .}}
		"{{.HCLTag}}":  &me.{{.Name}},{{end}}{{end}}
	})
}
`

const GO_ENUM_TEMPLATE = `
type {{.Name}} string
{{$name := .Name}}
var {{.Plural}} = struct{
	{{with .Instances}}{{ range .}}
	{{.Name}} {{$name}}{{end}}{{end}}
}{
	{{with .Instances}}{{ range .}}"{{.Literal}}",
	{{end}}{{end}} }
`

const GO_FILE_TEMPLATE = `/**
* @license
* Copyright 2020 Dynatrace LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

package {{.Package}}

{{ range .Imports}}import "{{ . }}"
{{end}}
	
{{ .Code }}
`
