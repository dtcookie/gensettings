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

package screensettings

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type StaticContent struct {
	Breadcrumbs      Breadcrumbs `json:"breadcrumbs,omitempty"`   // A list of breadcrumbs leading to the screen. Define breadcrumbs list in order from root page to leaf(current page would be last on that list).
	Header           *HeaderType `json:"header,omitempty"`        // Header definition
	ShowAddTag       bool        `json:"showAddTag"`              // If true, the \"Add tag\" button will be displayed
	ShowGlobalFilter bool        `json:"showGlobalFilter"`        // If true, the screen show the filter bar
	ShowOwnership    *bool       `json:"showOwnership,omitempty"` // If true, the screen show the ownerships card
	ShowProblems     bool        `json:"showProblems"`            // If true, the screen show the problem card
	ShowProperties   bool        `json:"showProperties"`          // If true, the screen show the properties card
	ShowTags         bool        `json:"showTags"`                // If true, entity tags will be displayed
}

func (me *StaticContent) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"breadcrumbs": {
			Type:        schema.TypeList,
			Description: "A list of breadcrumbs leading to the screen. Define breadcrumbs list in order from root page to leaf(current page would be last on that list).",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(Breadcrumbs).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"header": {
			Type:        schema.TypeList,
			Description: "Header definition",
			Optional:    true, // nullable
			Elem:        &schema.Resource{Schema: new(HeaderType).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"show_add_tag": {
			Type:        schema.TypeBool,
			Description: "If true, the \"Add tag\" button will be displayed",
			Required:    true,
		},
		"show_global_filter": {
			Type:        schema.TypeBool,
			Description: "If true, the screen show the filter bar",
			Required:    true,
		},
		"show_ownership": {
			Type:        schema.TypeBool,
			Description: "If true, the screen show the ownerships card",
			Optional:    true, // nullable
		},
		"show_problems": {
			Type:        schema.TypeBool,
			Description: "If true, the screen show the problem card",
			Required:    true,
		},
		"show_properties": {
			Type:        schema.TypeBool,
			Description: "If true, the screen show the properties card",
			Required:    true,
		},
		"show_tags": {
			Type:        schema.TypeBool,
			Description: "If true, entity tags will be displayed",
			Required:    true,
		},
	}
}

func (me *StaticContent) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"breadcrumbs":        me.Breadcrumbs,
		"header":             me.Header,
		"show_add_tag":       me.ShowAddTag,
		"show_global_filter": me.ShowGlobalFilter,
		"show_ownership":     me.ShowOwnership,
		"show_problems":      me.ShowProblems,
		"show_properties":    me.ShowProperties,
		"show_tags":          me.ShowTags,
	})
}

func (me *StaticContent) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"breadcrumbs":        &me.Breadcrumbs,
		"header":             &me.Header,
		"show_add_tag":       &me.ShowAddTag,
		"show_global_filter": &me.ShowGlobalFilter,
		"show_ownership":     &me.ShowOwnership,
		"show_problems":      &me.ShowProblems,
		"show_properties":    &me.ShowProperties,
		"show_tags":          &me.ShowTags,
	})
}
