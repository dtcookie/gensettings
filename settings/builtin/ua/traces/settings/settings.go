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

package traces

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Description       *string `json:"description,omitempty"`       // The description of the card
	DisplayName       *string `json:"displayName,omitempty"`       // The title of the card.
	EmptyStateMessage *string `json:"emptyStateMessage,omitempty"` // Empty state message
	Key               string  `json:"key"`                         // Unique key, which is used to map to this traces card in the screen layout config.
	Scope             string  `json:"-" scope:"scope"`             // The scope of this setting (ua-screen)
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:        schema.TypeString,
			Description: "The description of the card",
			Optional:    true, // nullable
		},
		"display_name": {
			Type:        schema.TypeString,
			Description: "The title of the card.",
			Optional:    true, // nullable
		},
		"empty_state_message": {
			Type:        schema.TypeString,
			Description: "Empty state message",
			Optional:    true, // nullable
		},
		"key": {
			Type:        schema.TypeString,
			Description: "Unique key, which is used to map to this traces card in the screen layout config.",
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
		"description":         me.Description,
		"display_name":        me.DisplayName,
		"empty_state_message": me.EmptyStateMessage,
		"key":                 me.Key,
		"scope":               me.Scope,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"description":         &me.Description,
		"display_name":        &me.DisplayName,
		"empty_state_message": &me.EmptyStateMessage,
		"key":                 &me.Key,
		"scope":               &me.Scope,
	})
}
