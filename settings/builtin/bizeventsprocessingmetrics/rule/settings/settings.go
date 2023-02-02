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

package rule

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Dimensions       []string `json:"dimensions"`
	Enabled          bool     `json:"enabled"`          // Enabled
	Key              string   `json:"key"`              // Key
	Matcher          string   `json:"matcher"`          // [See our documentation](https://dt-url.net/bp234rv)
	Measure          Measure  `json:"measure"`          // Measure
	MeasureAttribute string   `json:"measureAttribute"` // Attribute
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"dimensions": {
			Type:        schema.TypeSet,
			Description: "no documentation available",
			Required:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "Enabled",
			Required:    true,
		},
		"key": {
			Type:        schema.TypeString,
			Description: "Key",
			Required:    true,
		},
		"matcher": {
			Type:        schema.TypeString,
			Description: "[See our documentation](https://dt-url.net/bp234rv)",
			Required:    true,
		},
		"measure": {
			Type:        schema.TypeString,
			Description: "Measure",
			Required:    true,
		},
		"measure_attribute": {
			Type:        schema.TypeString,
			Description: "Attribute",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"dimensions":        me.Dimensions,
		"enabled":           me.Enabled,
		"key":               me.Key,
		"matcher":           me.Matcher,
		"measure":           me.Measure,
		"measure_attribute": me.MeasureAttribute,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"dimensions":        &me.Dimensions,
		"enabled":           &me.Enabled,
		"key":               &me.Key,
		"matcher":           &me.Matcher,
		"measure":           &me.Measure,
		"measure_attribute": &me.MeasureAttribute,
	})
}
