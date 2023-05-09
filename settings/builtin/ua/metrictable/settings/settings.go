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

type Settings struct {
	Charts                     ChartDefinitions    `json:"charts,omitempty"`                 // Define charts to display timeseries data in the table. Table columns are based on dimensions from metrics results of those charts. Make sure metrics of all charts have the same dimensions in result.
	Conditions                 []string            `json:"conditions,omitempty"`             // All conditions from the list need to be fulfilled for the table to be visible
	Description                *string             `json:"description,omitempty"`            // The description of the card
	DisplayCharts              bool                `json:"displayCharts"`                    // When true, the charts will be displayed above the list. Be aware of performance overhead with vast number of metrics and entities.
	DisplayMicroCharts         *bool               `json:"displayMicroCharts,omitempty"`     // When true, micro charts will be displayed for added metrics. Be aware of performance overhead with vast number of metrics and entities.
	DisplayName                *string             `json:"displayName,omitempty"`            // The title of the card.
	EmptyState                 *EmptyStateType     `json:"emptyState,omitempty"`             // Specifies message shown when there are no entities found.
	EnableDetailsExpandability bool                `json:"enableDetailsExpandability"`       // When true, detailed charts for each entity will be available.
	EntitySelectorTemplate     *string             `json:"entitySelectorTemplate,omitempty"` // An entity selector to narrow down dimensions.
	Filtering                  *FilteringType      `json:"filtering,omitempty"`              // Specifies filtering configuration.
	HideEmptyCharts            *bool               `json:"hideEmptyCharts,omitempty"`        // If true, All empty charts will be hidden.
	InitialSorting             *InitialSortingType `json:"initialSorting,omitempty"`         // Specifies initial sorting configuration.
	Key                        string              `json:"key"`                              // Unique key, which is used to map to this entities list in the screen layout config.
	NumberOfVisibleCharts      int                 `json:"numberOfVisibleCharts"`            // The number of charts to display above the list and in details section.
	PageSize                   int                 `json:"pageSize"`                         // The number of entities displayed on one page.
	Scope                      string              `json:"-" scope:"scope"`                  // The scope of this setting (ua-screen)
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"charts": {
			Type:        schema.TypeList,
			Description: "Define charts to display timeseries data in the table. Table columns are based on dimensions from metrics results of those charts. Make sure metrics of all charts have the same dimensions in result.",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(ChartDefinitions).Schema()},
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
			Description: "The description of the card",
			Optional:    true, // nullable
		},
		"display_charts": {
			Type:        schema.TypeBool,
			Description: "When true, the charts will be displayed above the list. Be aware of performance overhead with vast number of metrics and entities.",
			Required:    true,
		},
		"display_micro_charts": {
			Type:        schema.TypeBool,
			Description: "When true, micro charts will be displayed for added metrics. Be aware of performance overhead with vast number of metrics and entities.",
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
		"enable_details_expandability": {
			Type:        schema.TypeBool,
			Description: "When true, detailed charts for each entity will be available.",
			Required:    true,
		},
		"entity_selector_template": {
			Type:        schema.TypeString,
			Description: "An entity selector to narrow down dimensions.",
			Optional:    true, // nullable
		},
		"filtering": {
			Type:        schema.TypeList,
			Description: "Specifies filtering configuration.",
			Optional:    true, // nullable
			Elem:        &schema.Resource{Schema: new(FilteringType).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"hide_empty_charts": {
			Type:        schema.TypeBool,
			Description: "If true, All empty charts will be hidden.",
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
		"number_of_visible_charts": {
			Type:        schema.TypeInt,
			Description: "The number of charts to display above the list and in details section.",
			Required:    true,
		},
		"page_size": {
			Type:        schema.TypeInt,
			Description: "The number of entities displayed on one page.",
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
		"charts":                       me.Charts,
		"conditions":                   me.Conditions,
		"description":                  me.Description,
		"display_charts":               me.DisplayCharts,
		"display_micro_charts":         me.DisplayMicroCharts,
		"display_name":                 me.DisplayName,
		"empty_state":                  me.EmptyState,
		"enable_details_expandability": me.EnableDetailsExpandability,
		"entity_selector_template":     me.EntitySelectorTemplate,
		"filtering":                    me.Filtering,
		"hide_empty_charts":            me.HideEmptyCharts,
		"initial_sorting":              me.InitialSorting,
		"key":                          me.Key,
		"number_of_visible_charts":     me.NumberOfVisibleCharts,
		"page_size":                    me.PageSize,
		"scope":                        me.Scope,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"charts":                       &me.Charts,
		"conditions":                   &me.Conditions,
		"description":                  &me.Description,
		"display_charts":               &me.DisplayCharts,
		"display_micro_charts":         &me.DisplayMicroCharts,
		"display_name":                 &me.DisplayName,
		"empty_state":                  &me.EmptyState,
		"enable_details_expandability": &me.EnableDetailsExpandability,
		"entity_selector_template":     &me.EntitySelectorTemplate,
		"filtering":                    &me.Filtering,
		"hide_empty_charts":            &me.HideEmptyCharts,
		"initial_sorting":              &me.InitialSorting,
		"key":                          &me.Key,
		"number_of_visible_charts":     &me.NumberOfVisibleCharts,
		"page_size":                    &me.PageSize,
		"scope":                        &me.Scope,
	})
}
