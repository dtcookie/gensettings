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

package properties

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	DisplayOnlyConfigured      *bool               `json:"displayOnlyConfigured,omitempty"`      // By default, all available properties are displayed in the card. Use this flag to display only properties defined in the table below.
	Properties                 PropertyDefinitions `json:"properties,omitempty"`                 // Properties list
	Scope                      string              `json:"-" scope:"scope"`                      // The scope of this setting (ua-screen)
	TagsRelatedEntitySelectors []string            `json:"tagsRelatedEntitySelectors,omitempty"` // Display additionally tags of some related entities. Tags of related entities are read-only. The limit is 10 related entities.
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"display_only_configured": {
			Type:        schema.TypeBool,
			Description: "By default, all available properties are displayed in the card. Use this flag to display only properties defined in the table below.",
			Optional:    true, // nullable
		},
		"properties": {
			Type:        schema.TypeList,
			Description: "Properties list",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(PropertyDefinitions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (ua-screen)",
			Required:    true,
		},
		"tags_related_entity_selectors": {
			Type:        schema.TypeSet,
			Description: "Display additionally tags of some related entities. Tags of related entities are read-only. The limit is 10 related entities.",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"display_only_configured":       me.DisplayOnlyConfigured,
		"properties":                    me.Properties,
		"scope":                         me.Scope,
		"tags_related_entity_selectors": me.TagsRelatedEntitySelectors,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"display_only_configured":       &me.DisplayOnlyConfigured,
		"properties":                    &me.Properties,
		"scope":                         &me.Scope,
		"tags_related_entity_selectors": &me.TagsRelatedEntitySelectors,
	})
}
