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

package attributetable

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Attribute              string              `json:"attribute"`                        // Complex attribute which should be displayed as table.
	Columns                ColumnDefinitions   `json:"columns"`                          // Define columns for this complex attribute.
	Conditions             []string            `json:"conditions,omitempty"`             // All conditions from the list need to be fulfilled for the table to be visible
	Description            *string             `json:"description,omitempty"`            // The description of the card.
	DisplayName            *string             `json:"displayName,omitempty"`            // The title of the card.
	EmptyState             *EmptyStateType     `json:"emptyState,omitempty"`             // Specifies message shown when there are no entities found.
	EntitySelectorTemplate *string             `json:"entitySelectorTemplate,omitempty"` // An entity selector that is used to fetch data for related entities. For attribute table it should resolve to single entity.
	InitialSorting         *InitialSortingType `json:"initialSorting,omitempty"`         // Specifies initial sorting configuration.
	Key                    string              `json:"key"`                              // Unique key, which is used to map to this entities list in the screen layout config.
	PageSize               int                 `json:"pageSize"`                         // The number of items displayed on one page.
	Scope                  string              `json:"-" scope:"scope"`                  // The scope of this setting (ua-screen)
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attribute": {
			Type:        schema.TypeString,
			Description: "Complex attribute which should be displayed as table.",
			Required:    true,
		},
		"columns": {
			Type:        schema.TypeList,
			Description: "Define columns for this complex attribute.",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(ColumnDefinitions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"conditions": {
			Type:        schema.TypeSet,
			Description: "All conditions from the list need to be fulfilled for the table to be visible",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"description": {
			Type:        schema.TypeString,
			Description: "The description of the card.",
			Optional:    true, // nullable
		},
		"display_name": {
			Type:        schema.TypeString,
			Description: "The title of the card.",
			Optional:    true, // nullable
		},
		"empty_state": {
			Type:        schema.TypeList,
			Description: "Specifies message shown when there are no entities found.",
			Optional:    true, // nullable
			Elem:        &schema.Resource{Schema: new(EmptyStateType).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"entity_selector_template": {
			Type:        schema.TypeString,
			Description: "An entity selector that is used to fetch data for related entities. For attribute table it should resolve to single entity.",
			Optional:    true, // nullable
		},
		"initial_sorting": {
			Type:        schema.TypeList,
			Description: "Specifies initial sorting configuration.",
			Optional:    true, // nullable
			Elem:        &schema.Resource{Schema: new(InitialSortingType).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"key": {
			Type:        schema.TypeString,
			Description: "Unique key, which is used to map to this entities list in the screen layout config.",
			Required:    true,
		},
		"page_size": {
			Type:        schema.TypeInt,
			Description: "The number of items displayed on one page.",
			Required:    true,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (ua-screen)",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"attribute":                me.Attribute,
		"columns":                  me.Columns,
		"conditions":               me.Conditions,
		"description":              me.Description,
		"display_name":             me.DisplayName,
		"empty_state":              me.EmptyState,
		"entity_selector_template": me.EntitySelectorTemplate,
		"initial_sorting":          me.InitialSorting,
		"key":                      me.Key,
		"page_size":                me.PageSize,
		"scope":                    me.Scope,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"attribute":                &me.Attribute,
		"columns":                  &me.Columns,
		"conditions":               &me.Conditions,
		"description":              &me.Description,
		"display_name":             &me.DisplayName,
		"empty_state":              &me.EmptyState,
		"entity_selector_template": &me.EntitySelectorTemplate,
		"initial_sorting":          &me.InitialSorting,
		"key":                      &me.Key,
		"page_size":                &me.PageSize,
		"scope":                    &me.Scope,
	})
}
