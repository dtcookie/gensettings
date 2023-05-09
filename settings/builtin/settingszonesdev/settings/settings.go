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

package settingszonesdev

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	DisplayName         *string             `json:"displayName,omitempty"` // Settings zone display name
	EntitySelectorRules EntitySelectorRules `json:"entitySelectorRules"`   // Entity selector rules
	ID                  string              `json:"id"`                    // Settings zone id
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"display_name": {
			Type:        schema.TypeString,
			Description: "Settings zone display name",
			Optional:    true, // nullable
		},
		"entity_selector_rules": {
			Type:        schema.TypeList,
			Description: "Entity selector rules",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(EntitySelectorRules).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"id": {
			Type:        schema.TypeString,
			Description: "Settings zone id",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"display_name":          me.DisplayName,
		"entity_selector_rules": me.EntitySelectorRules,
		"id":                    me.ID,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"display_name":          &me.DisplayName,
		"entity_selector_rules": &me.EntitySelectorRules,
		"id":                    &me.ID,
	})
}
