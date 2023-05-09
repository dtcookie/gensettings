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

package eventscard

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Description             *string  `json:"description,omitempty"`             // The description of the card
	DisplayChart            bool     `json:"displayChart"`                      // When true, the chart will be displayed above the list.
	DisplayName             *string  `json:"displayName,omitempty"`             // The title of the card.
	EntitySelectorTemplates []string `json:"entitySelectorTemplates,omitempty"` // Up to 10 entity selectors, used to fetch data for related entities
	EventSelectorTemplate   *string  `json:"eventSelectorTemplate,omitempty"`   // An event selector that is used to fetch events for related entities
	Key                     string   `json:"key"`                               // Unique key, which is used to map to this events card in the screen layout config
	PageSize                int      `json:"pageSize"`                          // Number of events per page
	Scope                   string   `json:"-" scope:"scope"`                   // The scope of this setting (ua-screen)
	ShowFiltering           bool     `json:"showFiltering"`                     // Enable or disable filtering bar at the bottom of the card
	ShowPagination          bool     `json:"showPagination"`                    // Enable or disable pagination bar at the bottom of the card
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:        schema.TypeString,
			Description: "The description of the card",
			Optional:    true, // nullable
		},
		"display_chart": {
			Type:        schema.TypeBool,
			Description: "When true, the chart will be displayed above the list.",
			Required:    true,
		},
		"display_name": {
			Type:        schema.TypeString,
			Description: "The title of the card.",
			Optional:    true, // nullable
		},
		"entity_selector_templates": {
			Type:        schema.TypeSet,
			Description: "Up to 10 entity selectors, used to fetch data for related entities",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"event_selector_template": {
			Type:        schema.TypeString,
			Description: "An event selector that is used to fetch events for related entities",
			Optional:    true, // nullable
		},
		"key": {
			Type:        schema.TypeString,
			Description: "Unique key, which is used to map to this events card in the screen layout config",
			Required:    true,
		},
		"page_size": {
			Type:        schema.TypeInt,
			Description: "Number of events per page",
			Required:    true,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (ua-screen)",
			Required:    true,
		},
		"show_filtering": {
			Type:        schema.TypeBool,
			Description: "Enable or disable filtering bar at the bottom of the card",
			Required:    true,
		},
		"show_pagination": {
			Type:        schema.TypeBool,
			Description: "Enable or disable pagination bar at the bottom of the card",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"description":               me.Description,
		"display_chart":             me.DisplayChart,
		"display_name":              me.DisplayName,
		"entity_selector_templates": me.EntitySelectorTemplates,
		"event_selector_template":   me.EventSelectorTemplate,
		"key":                       me.Key,
		"page_size":                 me.PageSize,
		"scope":                     me.Scope,
		"show_filtering":            me.ShowFiltering,
		"show_pagination":           me.ShowPagination,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"description":               &me.Description,
		"display_chart":             &me.DisplayChart,
		"display_name":              &me.DisplayName,
		"entity_selector_templates": &me.EntitySelectorTemplates,
		"event_selector_template":   &me.EventSelectorTemplate,
		"key":                       &me.Key,
		"page_size":                 &me.PageSize,
		"scope":                     &me.Scope,
		"show_filtering":            &me.ShowFiltering,
		"show_pagination":           &me.ShowPagination,
	})
}
