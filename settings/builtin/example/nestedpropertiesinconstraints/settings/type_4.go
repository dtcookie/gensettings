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

package nestedpropertiesinconstraints

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Type4s []*Type4

func (me *Type4s) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"type_4": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(Type4).Schema()},
		},
	}
}

func (me Type4s) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("type_4", me)
}

func (me *Type4s) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("type_4", me)
}

// Inner type 4. Inner type with a property referenced in the uniqueness collection constraint
type Type4 struct {
	ComplexValue *Type3 `json:"complexValue"` // Complex value
	TextValue    string `json:"textValue"`    // Text value
}

func (me *Type4) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"complex_value": {
			Type:        schema.TypeList,
			Description: "Complex value",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(Type3).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"text_value": {
			Type:        schema.TypeString,
			Description: "Text value",
			Required:    true,
		},
	}
}

func (me *Type4) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"complex_value": me.ComplexValue,
		"text_value":    me.TextValue,
	})
}

func (me *Type4) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"complex_value": &me.ComplexValue,
		"text_value":    &me.TextValue,
	})
}
