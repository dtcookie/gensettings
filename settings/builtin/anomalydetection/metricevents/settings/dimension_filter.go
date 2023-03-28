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

package metricevents

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type DimensionFilters []*DimensionFilter

func (me *DimensionFilters) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"dimension_filter": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(DimensionFilter).Schema()},
		},
	}
}

func (me DimensionFilters) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("dimension_filter", me)
}

func (me *DimensionFilters) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("dimension_filter", me)
}

type DimensionFilter struct {
	DimensionKey   string `json:"dimensionKey"`   // Dimension key
	DimensionValue string `json:"dimensionValue"` // Dimension value
}

func (me *DimensionFilter) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"dimension_key": {
			Type:        schema.TypeString,
			Description: "Dimension key",
			Required:    true,
		},
		"dimension_value": {
			Type:        schema.TypeString,
			Description: "Dimension value",
			Required:    true,
		},
	}
}

func (me *DimensionFilter) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"dimension_key":   me.DimensionKey,
		"dimension_value": me.DimensionValue,
	})
}

func (me *DimensionFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"dimension_key":   &me.DimensionKey,
		"dimension_value": &me.DimensionValue,
	})
}
