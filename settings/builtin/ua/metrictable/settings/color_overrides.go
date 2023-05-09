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

package metrictable

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ColorOverridess []*ColorOverrides

func (me *ColorOverridess) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"color_overrides": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(ColorOverrides).Schema()},
		},
	}
}

func (me ColorOverridess) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("color_overrides", me)
}

func (me *ColorOverridess) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("color_overrides", me)
}

type ColorOverrides struct {
	Color      string `json:"color"`      // RGB code of color
	SeriesName string `json:"seriesName"` // Series name should point to whole metric or single dimension name
}

func (me *ColorOverrides) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"color": {
			Type:        schema.TypeString,
			Description: "RGB code of color",
			Required:    true,
		},
		"series_name": {
			Type:        schema.TypeString,
			Description: "Series name should point to whole metric or single dimension name",
			Required:    true,
		},
	}
}

func (me *ColorOverrides) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"color":       me.Color,
		"series_name": me.SeriesName,
	})
}

func (me *ColorOverrides) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"color":       &me.Color,
		"series_name": &me.SeriesName,
	})
}
