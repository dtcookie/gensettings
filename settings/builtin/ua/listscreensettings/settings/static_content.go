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

package listscreensettings

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type StaticContent struct {
	Breadcrumbs           Breadcrumbs `json:"breadcrumbs,omitempty"`           // A list of breadcrumbs leading to the screen. Define breadcrumbs list in order from root page to leaf(current page would be last on that list).
	Header                *HeaderType `json:"header,omitempty"`                // Header definition
	HideDefaultBreadcrumb *bool       `json:"hideDefaultBreadcrumb,omitempty"` // If true, the default breadcrumb for the current page is hidden
	ShowGlobalFilter      bool        `json:"showGlobalFilter"`                // If true, the screen show the filter bar
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
		"hide_default_breadcrumb": {
			Type:        schema.TypeBool,
			Description: "If true, the default breadcrumb for the current page is hidden",
			Optional:    true, // nullable
		},
		"show_global_filter": {
			Type:        schema.TypeBool,
			Description: "If true, the screen show the filter bar",
			Required:    true,
		},
	}
}

func (me *StaticContent) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"breadcrumbs":             me.Breadcrumbs,
		"header":                  me.Header,
		"hide_default_breadcrumb": me.HideDefaultBreadcrumb,
		"show_global_filter":      me.ShowGlobalFilter,
	})
}

func (me *StaticContent) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"breadcrumbs":             &me.Breadcrumbs,
		"header":                  &me.Header,
		"hide_default_breadcrumb": &me.HideDefaultBreadcrumb,
		"show_global_filter":      &me.ShowGlobalFilter,
	})
}
