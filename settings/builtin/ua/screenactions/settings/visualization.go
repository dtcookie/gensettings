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

package screenactions

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Visualization struct {
	Color       *ButtonColorType `json:"color,omitempty"`       // Possible Values: `PRIMARY`, `SECONDARY`
	DisplayName *string          `json:"displayName,omitempty"` // Override/set action button display name. It will be deduced from action definition if not set.
	Icon        *string          `json:"icon,omitempty"`        // Include the [barista icon](https://barista.dynatrace.com/components/icon#all-icons) action button.
	IconOnly    *bool            `json:"iconOnly,omitempty"`    // If selected, only icon will be rendered in the action button. Display name will be moved to the tooltip on hover.
}

func (me *Visualization) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"color": {
			Type:        schema.TypeString,
			Description: "Possible Values: `PRIMARY`, `SECONDARY`",
			Optional:    true, // nullable
		},
		"display_name": {
			Type:        schema.TypeString,
			Description: "Override/set action button display name. It will be deduced from action definition if not set.",
			Optional:    true, // nullable
		},
		"icon": {
			Type:        schema.TypeString,
			Description: "Include the [barista icon](https://barista.dynatrace.com/components/icon#all-icons) action button.",
			Optional:    true, // nullable
		},
		"icon_only": {
			Type:        schema.TypeBool,
			Description: "If selected, only icon will be rendered in the action button. Display name will be moved to the tooltip on hover.",
			Optional:    true, // nullable & precondition
		},
	}
}

func (me *Visualization) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"color":        me.Color,
		"display_name": me.DisplayName,
		"icon":         me.Icon,
		"icon_only":    me.IconOnly,
	})
}

func (me *Visualization) HandlePreconditions() error {
	// ---- IconOnly *bool -> {"precondition":{"property":"icon","type":"NULL"},"type":"NOT"}
	return nil
}

func (me *Visualization) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"color":        &me.Color,
		"display_name": &me.DisplayName,
		"icon":         &me.Icon,
		"icon_only":    &me.IconOnly,
	})
}
