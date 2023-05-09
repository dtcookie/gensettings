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

package chartgroups

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type YAxiss []*YAxis

func (me *YAxiss) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"y_axe": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(YAxis).Schema()},
		},
	}
}

func (me YAxiss) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("y_axe", me)
}

func (me *YAxiss) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("y_axe", me)
}

type YAxis struct {
	Key      string   `json:"key"`            // Y-axis key to be matched with metric's reference
	Max      *string  `json:"max,omitempty"`  // Maximal axis value
	Min      *string  `json:"min,omitempty"`  // Minimal axis value
	Name     *string  `json:"name,omitempty"` // Axis title
	Position Position `json:"position"`       // Possible Values: `LEFT`, `RIGHT`
	Visible  bool     `json:"visible"`        // When true, axis is visible
}

func (me *YAxis) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"key": {
			Type:        schema.TypeString,
			Description: "Y-axis key to be matched with metric's reference",
			Required:    true,
		},
		"max": {
			Type:        schema.TypeString,
			Description: "Maximal axis value",
			Optional:    true, // nullable
		},
		"min": {
			Type:        schema.TypeString,
			Description: "Minimal axis value",
			Optional:    true, // nullable
		},
		"name": {
			Type:        schema.TypeString,
			Description: "Axis title",
			Optional:    true, // nullable
		},
		"position": {
			Type:        schema.TypeString,
			Description: "Possible Values: `LEFT`, `RIGHT`",
			Required:    true,
		},
		"visible": {
			Type:        schema.TypeBool,
			Description: "When true, axis is visible",
			Required:    true,
		},
	}
}

func (me *YAxis) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"key":      me.Key,
		"max":      me.Max,
		"min":      me.Min,
		"name":     me.Name,
		"position": me.Position,
		"visible":  me.Visible,
	})
}

func (me *YAxis) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"key":      &me.Key,
		"max":      &me.Max,
		"min":      &me.Min,
		"name":     &me.Name,
		"position": &me.Position,
		"visible":  &me.Visible,
	})
}
