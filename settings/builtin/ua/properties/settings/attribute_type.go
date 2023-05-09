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

// AttributeType. Allows to overwrite attribute display options
type AttributeType struct {
	DisplayName            *string `json:"displayName,omitempty"`            // Leave empty to resolve automatically
	EntitySelectorTemplate *string `json:"entitySelectorTemplate,omitempty"` // An optional entity selector to fetch attribute from related entity. For example for a host screen where you want to display property of VM that it runs on, this might look like this:  \n`type(AWS_AVAILABILITY_ZONE), fromRelationships.isSiteOf($(entityConditions))`\n\nPlease mind that the `$(entityConditions)` is a placeholder for the current ME selector.\n\nPlease remember that attribute key needs to be available in the context of related entity.
	Hidden                 *bool   `json:"hidden,omitempty"`                 // Hide the property
	Key                    string  `json:"key"`                              // Use to reference desired attribute
	Unit                   *string `json:"unit,omitempty"`                   // Use to format numeric properties
}

func (me *AttributeType) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"display_name": {
			Type:        schema.TypeString,
			Description: "Leave empty to resolve automatically",
			Optional:    true, // nullable
		},
		"entity_selector_template": {
			Type:        schema.TypeString,
			Description: "An optional entity selector to fetch attribute from related entity. For example for a host screen where you want to display property of VM that it runs on, this might look like this:  \n`type(AWS_AVAILABILITY_ZONE), fromRelationships.isSiteOf($(entityConditions))`\n\nPlease mind that the `$(entityConditions)` is a placeholder for the current ME selector.\n\nPlease remember that attribute key needs to be available in the context of related entity.",
			Optional:    true, // nullable
		},
		"hidden": {
			Type:        schema.TypeBool,
			Description: "Hide the property",
			Optional:    true, // nullable
		},
		"key": {
			Type:        schema.TypeString,
			Description: "Use to reference desired attribute",
			Required:    true,
		},
		"unit": {
			Type:        schema.TypeString,
			Description: "Use to format numeric properties",
			Optional:    true, // nullable
		},
	}
}

func (me *AttributeType) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"display_name":             me.DisplayName,
		"entity_selector_template": me.EntitySelectorTemplate,
		"hidden":                   me.Hidden,
		"key":                      me.Key,
		"unit":                     me.Unit,
	})
}

func (me *AttributeType) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"display_name":             &me.DisplayName,
		"entity_selector_template": &me.EntitySelectorTemplate,
		"hidden":                   &me.Hidden,
		"key":                      &me.Key,
		"unit":                     &me.Unit,
	})
}
