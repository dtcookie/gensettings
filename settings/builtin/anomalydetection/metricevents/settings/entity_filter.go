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

type EntityFilter struct {
	Conditions   EntityFilterConditions `json:"conditions"`
	DimensionKey *string                `json:"dimensionKey,omitempty"` // Dimension key of entity type to filter
}

func (me *EntityFilter) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"conditions": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(EntityFilterConditions).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"dimension_key": {
			Type:        schema.TypeString,
			Description: "Dimension key of entity type to filter",
			Optional:    true,
		},
	}
}

func (me *EntityFilter) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"conditions":    me.Conditions,
		"dimension_key": me.DimensionKey,
	})
}

func (me *EntityFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"conditions":    &me.Conditions,
		"dimension_key": &me.DimensionKey,
	})
}
