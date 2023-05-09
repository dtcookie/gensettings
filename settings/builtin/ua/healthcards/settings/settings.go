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

package healthcards

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Conditions []string `json:"conditions,omitempty"` // All conditions from the list need to be fulfilled for the Health card to be visible
	Key        string   `json:"key"`                  // Unique key, which is used to map to this health card in the screen layout config
	Scope      string   `json:"-" scope:"scope"`      // The scope of this setting (ua-screen)
	Tiles      Tiles    `json:"tiles,omitempty"`      // Collection of tiles definition in health card
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"conditions": {
			Type:        schema.TypeSet,
			Description: "All conditions from the list need to be fulfilled for the Health card to be visible",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"key": {
			Type:        schema.TypeString,
			Description: "Unique key, which is used to map to this health card in the screen layout config",
			Required:    true,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (ua-screen)",
			Required:    true,
		},
		"tiles": {
			Type:        schema.TypeList,
			Description: "Collection of tiles definition in health card",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(Tiles).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"conditions": me.Conditions,
		"key":        me.Key,
		"scope":      me.Scope,
		"tiles":      me.Tiles,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"conditions": &me.Conditions,
		"key":        &me.Key,
		"scope":      &me.Scope,
		"tiles":      &me.Tiles,
	})
}
