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

package entitieslist

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Thresholds []*Threshold

func (me *Thresholds) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"threshold": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(Threshold).Schema()},
		},
	}
}

func (me Thresholds) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("threshold", me)
}

func (me *Thresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("threshold", me)
}

type Threshold struct {
	Color string  `json:"color"` // RGB code of color
	Value float64 `json:"value"` // Value of the threshold. Area above this threshold will be marked by selected color
}

func (me *Threshold) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"color": {
			Type:        schema.TypeString,
			Description: "RGB code of color",
			Required:    true,
		},
		"value": {
			Type:        schema.TypeFloat,
			Description: "Value of the threshold. Area above this threshold will be marked by selected color",
			Required:    true,
		},
	}
}

func (me *Threshold) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"color": me.Color,
		"value": me.Value,
	})
}

func (me *Threshold) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"color": &me.Color,
		"value": &me.Value,
	})
}
