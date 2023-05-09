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

package logscard

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Description   *string `json:"description,omitempty"` // The description of the card
	DisplayChart  bool    `json:"displayChart"`          // When true, the chart will be displayed above the list.
	DisplayName   *string `json:"displayName,omitempty"` // The title of the card.
	EnablePaging  bool    `json:"enablePaging"`          // Show or hide the link to the next n log records
	FilterQuery   *string `json:"filterQuery,omitempty"` // Use Dynatrace search query language to prefilter log data.
	Key           string  `json:"key"`                   // Unique key, which is used to map to this log card in the screen layout config
	PageSize      int     `json:"pageSize"`              // Number of log rows per page
	Scope         string  `json:"-" scope:"scope"`       // The scope of this setting (ua-screen)
	ShowFiltering bool    `json:"showFiltering"`         // Enable or disable filtering bar at the bottom of the card
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
		"enable_paging": {
			Type:        schema.TypeBool,
			Description: "Show or hide the link to the next n log records",
			Required:    true,
		},
		"filter_query": {
			Type:        schema.TypeString,
			Description: "Use Dynatrace search query language to prefilter log data.",
			Optional:    true, // nullable
		},
		"key": {
			Type:        schema.TypeString,
			Description: "Unique key, which is used to map to this log card in the screen layout config",
			Required:    true,
		},
		"page_size": {
			Type:        schema.TypeInt,
			Description: "Number of log rows per page",
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
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"description":    me.Description,
		"display_chart":  me.DisplayChart,
		"display_name":   me.DisplayName,
		"enable_paging":  me.EnablePaging,
		"filter_query":   me.FilterQuery,
		"key":            me.Key,
		"page_size":      me.PageSize,
		"scope":          me.Scope,
		"show_filtering": me.ShowFiltering,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"description":    &me.Description,
		"display_chart":  &me.DisplayChart,
		"display_name":   &me.DisplayName,
		"enable_paging":  &me.EnablePaging,
		"filter_query":   &me.FilterQuery,
		"key":            &me.Key,
		"page_size":      &me.PageSize,
		"scope":          &me.Scope,
		"show_filtering": &me.ShowFiltering,
	})
}
