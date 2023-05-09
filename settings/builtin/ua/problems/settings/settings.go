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

package problems

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	EntitySelectorTemplates []string `json:"entitySelectorTemplates,omitempty"` // Up to 10 entity selectors, used to fetch data for related entities
	Scope                   string   `json:"-" scope:"scope"`                   // The scope of this setting (ua-screen)
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"entity_selector_templates": {
			Type:        schema.TypeSet,
			Description: "Up to 10 entity selectors, used to fetch data for related entities",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
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
		"entity_selector_templates": me.EntitySelectorTemplates,
		"scope":                     me.Scope,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"entity_selector_templates": &me.EntitySelectorTemplates,
		"scope":                     &me.Scope,
	})
}
