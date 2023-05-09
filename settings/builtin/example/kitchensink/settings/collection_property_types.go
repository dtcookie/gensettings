/**
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

package kitchensink

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type CollectionPropertyTypes struct {
	List_complex AnimalComplexes  `json:"list_complex,omitempty"` // Collection property type: `list`\n\nJava class: `java.util.List`\n\nDefault constraints: -
	List_simple  []string         `json:"list_simple,omitempty"`  // Collection property type: `list`\n\nJava class: `java.util.List`\n\nDefault constraints: -
	Set_complex  AnimalComplexSet `json:"set_complex,omitempty"`  // Collection property type: `set`\n\nJava class: `java.util.Set`\n\nDefault constraints: -
	Set_simple   []string         `json:"set_simple,omitempty"`   // Collection property type: `set`\n\nJava class: `java.util.Set`\n\nDefault constraints: -
}

func (me *CollectionPropertyTypes) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list_complex": {
			Type:        schema.TypeList,
			Description: "Collection property type: `list`\n\nJava class: `java.util.List`\n\nDefault constraints: -",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(AnimalComplexes).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"list_simple": {
			Type:        schema.TypeList,
			Description: "Collection property type: `list`\n\nJava class: `java.util.List`\n\nDefault constraints: -",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"set_complex": {
			Type:        schema.TypeList,
			Description: "Collection property type: `set`\n\nJava class: `java.util.Set`\n\nDefault constraints: -",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(AnimalComplexSet).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"set_simple": {
			Type:        schema.TypeSet,
			Description: "Collection property type: `set`\n\nJava class: `java.util.Set`\n\nDefault constraints: -",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
	}
}

func (me *CollectionPropertyTypes) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"list_complex": me.List_complex,
		"list_simple":  me.List_simple,
		"set_complex":  me.Set_complex,
		"set_simple":   me.Set_simple,
	})
}

func (me *CollectionPropertyTypes) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"list_complex": &me.List_complex,
		"list_simple":  &me.List_simple,
		"set_complex":  &me.Set_complex,
		"set_simple":   &me.Set_simple,
	})
}
