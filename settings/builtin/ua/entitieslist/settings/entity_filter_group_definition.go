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

type EntityFilterGroupDefinitions []*EntityFilterGroupDefinition

func (me *EntityFilterGroupDefinitions) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"entity_filter": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(EntityFilterGroupDefinition).Schema()},
		},
	}
}

func (me EntityFilterGroupDefinitions) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("entity_filter", me)
}

func (me *EntityFilterGroupDefinitions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("entity_filter", me)
}

type EntityFilterGroupDefinition struct {
	DisplayName string                  `json:"displayName"`       // Group display name
	Filters     EntityFilterDefinitions `json:"filters,omitempty"` // Define entity filters.
}

func (me *EntityFilterGroupDefinition) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"display_name": {
			Type:        schema.TypeString,
			Description: "Group display name",
			Required:    true,
		},
		"filters": {
			Type:        schema.TypeList,
			Description: "Define entity filters.",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(EntityFilterDefinitions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *EntityFilterGroupDefinition) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"display_name": me.DisplayName,
		"filters":      me.Filters,
	})
}

func (me *EntityFilterGroupDefinition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"display_name": &me.DisplayName,
		"filters":      &me.Filters,
	})
}
