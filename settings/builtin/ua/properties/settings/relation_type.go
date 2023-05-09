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

// RelationType. Allows to specify link to the related entity
type RelationType struct {
	DisplayName            string  `json:"displayName"`               // Display name
	EntitySelectorTemplate string  `json:"entitySelectorTemplate"`    // An entity selector that is used to fetch the selected entities. For example for a process screen where you want to link to the process group this might look like this:  \n`type(PROCESS_GROUP), toRelationships.isInstanceOf($(entityConditions))`\n\nPlease mind that the `$(entityConditions)` is a placeholder for the current ME selector.
	FallbackMessage        *string `json:"fallbackMessage,omitempty"` // Display additional message if no entity is found. By default whole property is hidden.
}

func (me *RelationType) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"display_name": {
			Type:        schema.TypeString,
			Description: "Display name",
			Required:    true,
		},
		"entity_selector_template": {
			Type:        schema.TypeString,
			Description: "An entity selector that is used to fetch the selected entities. For example for a process screen where you want to link to the process group this might look like this:  \n`type(PROCESS_GROUP), toRelationships.isInstanceOf($(entityConditions))`\n\nPlease mind that the `$(entityConditions)` is a placeholder for the current ME selector.",
			Required:    true,
		},
		"fallback_message": {
			Type:        schema.TypeString,
			Description: "Display additional message if no entity is found. By default whole property is hidden.",
			Optional:    true, // nullable
		},
	}
}

func (me *RelationType) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"display_name":             me.DisplayName,
		"entity_selector_template": me.EntitySelectorTemplate,
		"fallback_message":         me.FallbackMessage,
	})
}

func (me *RelationType) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"display_name":             &me.DisplayName,
		"entity_selector_template": &me.EntitySelectorTemplate,
		"fallback_message":         &me.FallbackMessage,
	})
}
