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

package config

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Constraints []*Constraint

func (me *Constraints) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"constraint": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(Constraint).Schema()},
		},
	}
}

func (me Constraints) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("constraint", me)
}

func (me *Constraints) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("constraint", me)
}

type Constraint struct {
	Type  string `json:"type"`  // Constraint type
	Value string `json:"value"` // Constraint value
}

func (me *Constraint) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"type": {
			Type:        schema.TypeString,
			Description: "Constraint type",
			Required:    true,
		},
		"value": {
			Type:        schema.TypeString,
			Description: "Constraint value",
			Required:    true,
		},
	}
}

func (me *Constraint) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"type":  me.Type,
		"value": me.Value,
	})
}

func (me *Constraint) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"type":  &me.Type,
		"value": &me.Value,
	})
}
