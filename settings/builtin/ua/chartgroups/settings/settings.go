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
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"
)

type Settings struct {
	AppendDetectedMetrics  *bool            `json:"appendDetectedMetrics,omitempty"`  // If true, Charts for all associated metrics will be generated and appended to the list in chart selector.
	Charts                 ChartDefinitions `json:"charts,omitempty"`                 // Collection of charts definition in this chart group
	ChartsInRow            *int             `json:"chartsInRow,omitempty"`            // Maximum number of charts in a row
	Conditions             []string         `json:"conditions,omitempty"`             // All conditions from the list need to be fulfilled for the ChartGroup card to be visible
	Description            *string          `json:"description,omitempty"`            // The description of the card
	DisplayName            *string          `json:"displayName,omitempty"`            // The title of the card.
	EntitySelectorTemplate *string          `json:"entitySelectorTemplate,omitempty"` // An entity selector that is used to fetch data for related entities
	HideEmptyCharts        *bool            `json:"hideEmptyCharts,omitempty"`        // If true, All empty charts will be hidden.
	Key                    string           `json:"key"`                              // Unique key, which is used to map to this chart group in the screen layout config
	Mode                   ChartGroupMode   `json:"mode"`                             // Possible Values: `NORMAL`, `STATIC`
	NumberOfVisibleCharts  int              `json:"numberOfVisibleCharts"`            // Number of charts to display
	Scope                  string           `json:"-" scope:"scope"`                  // The scope of this setting (ua-screen)
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"append_detected_metrics": {
			Type:        schema.TypeBool,
			Description: "If true, Charts for all associated metrics will be generated and appended to the list in chart selector.",
			Optional:    true, // nullable
		},
		"charts": {
			Type:        schema.TypeList,
			Description: "Collection of charts definition in this chart group",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(ChartDefinitions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"charts_in_row": {
			Type:        schema.TypeInt,
			Description: "Maximum number of charts in a row",
			Optional:    true, // nullable
		},
		"conditions": {
			Type:        schema.TypeSet,
			Description: "All conditions from the list need to be fulfilled for the ChartGroup card to be visible",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"description": {
			Type:        schema.TypeString,
			Description: "The description of the card",
			Optional:    true, // nullable & precondition
		},
		"display_name": {
			Type:        schema.TypeString,
			Description: "The title of the card.",
			Optional:    true, // nullable
		},
		"entity_selector_template": {
			Type:        schema.TypeString,
			Description: "An entity selector that is used to fetch data for related entities",
			Optional:    true, // nullable
		},
		"hide_empty_charts": {
			Type:        schema.TypeBool,
			Description: "If true, All empty charts will be hidden.",
			Optional:    true, // nullable
		},
		"key": {
			Type:        schema.TypeString,
			Description: "Unique key, which is used to map to this chart group in the screen layout config",
			Required:    true,
		},
		"mode": {
			Type:        schema.TypeString,
			Description: "Possible Values: `NORMAL`, `STATIC`",
			Required:    true,
		},
		"number_of_visible_charts": {
			Type:        schema.TypeInt,
			Description: "Number of charts to display",
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
		"append_detected_metrics":  me.AppendDetectedMetrics,
		"charts":                   me.Charts,
		"charts_in_row":            me.ChartsInRow,
		"conditions":               me.Conditions,
		"description":              me.Description,
		"display_name":             me.DisplayName,
		"entity_selector_template": me.EntitySelectorTemplate,
		"hide_empty_charts":        me.HideEmptyCharts,
		"key":                      me.Key,
		"mode":                     me.Mode,
		"number_of_visible_charts": me.NumberOfVisibleCharts,
		"scope":                    me.Scope,
	})
}

func (me *Settings) HandlePreconditions() error {
	if me.Description == nil && slices.Contains([]string{"NORMAL", "STATIC"}, string(me.Mode)) {
		return fmt.Errorf("'description' must be specified if 'mode' is set to '%v'", me.Mode)
	}
	return nil
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"append_detected_metrics":  &me.AppendDetectedMetrics,
		"charts":                   &me.Charts,
		"charts_in_row":            &me.ChartsInRow,
		"conditions":               &me.Conditions,
		"description":              &me.Description,
		"display_name":             &me.DisplayName,
		"entity_selector_template": &me.EntitySelectorTemplate,
		"hide_empty_charts":        &me.HideEmptyCharts,
		"key":                      &me.Key,
		"mode":                     &me.Mode,
		"number_of_visible_charts": &me.NumberOfVisibleCharts,
		"scope":                    &me.Scope,
	})
}
