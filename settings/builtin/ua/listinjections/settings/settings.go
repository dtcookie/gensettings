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

package listinjections

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Collapsed  *bool      `json:"collapsed,omitempty"`  // If true, card content will be collapsed
	Conditions []string   `json:"conditions,omitempty"` // All conditions from the list need to be fulfilled for the card to be visible
	Key        string     `json:"key"`                  // Unique key of the card, which is used to reference desired card configuration
	Scope      string     `json:"-" scope:"scope"`      // The scope of this setting (ua-screen)
	Type       CardType   `json:"type"`                 // Possible Values: `CHART_GROUP`, `ENTITIES_LIST`, `MESSAGE`
	Width      *WidthType `json:"width,omitempty"`      // A value that determines how much vertical space will that card take
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"collapsed": {
			Type:        schema.TypeBool,
			Description: "If true, card content will be collapsed",
			Optional:    true, // nullable
		},
		"conditions": {
			Type:        schema.TypeSet,
			Description: "All conditions from the list need to be fulfilled for the card to be visible",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"key": {
			Type:        schema.TypeString,
			Description: "Unique key of the card, which is used to reference desired card configuration",
			Required:    true,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (ua-screen)",
			Required:    true,
		},
		"type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `CHART_GROUP`, `ENTITIES_LIST`, `MESSAGE`",
			Required:    true,
		},
		"width": {
			Type:        schema.TypeString,
			Description: "A value that determines how much vertical space will that card take",
			Optional:    true, // nullable
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"collapsed":  me.Collapsed,
		"conditions": me.Conditions,
		"key":        me.Key,
		"scope":      me.Scope,
		"type":       me.Type,
		"width":      me.Width,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"collapsed":  &me.Collapsed,
		"conditions": &me.Conditions,
		"key":        &me.Key,
		"scope":      &me.Scope,
		"type":       &me.Type,
		"width":      &me.Width,
	})
}
