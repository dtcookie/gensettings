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

package generictype

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	CreatedBy   string          `json:"createdBy"`   // The user or extension that created this type.
	DisplayName string          `json:"displayName"` // The human readable type name for this entity type.
	Enabled     bool            `json:"enabled"`     // Enables or disables the type
	Name        string          `json:"name"`        // The entity type name. This type name must be unique and must not be changed after creation.
	Rules       ExtractionRules `json:"rules"`       // Specify a list of rules which are evaluated in order. When **any** rule matches, the entity defined according to that rule will be extracted. Subsequent rules will not be evaluated.
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_by": {
			Type:        schema.TypeString,
			Description: "The user or extension that created this type.",
			Required:    true,
		},
		"display_name": {
			Type:        schema.TypeString,
			Description: "The human readable type name for this entity type.",
			Required:    true,
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "Enables or disables the type",
			Required:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Description: "The entity type name. This type name must be unique and must not be changed after creation.",
			Required:    true,
		},
		"rules": {
			Type:        schema.TypeList,
			Description: "Specify a list of rules which are evaluated in order. When **any** rule matches, the entity defined according to that rule will be extracted. Subsequent rules will not be evaluated.",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(ExtractionRules).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"created_by":   me.CreatedBy,
		"display_name": me.DisplayName,
		"enabled":      me.Enabled,
		"name":         me.Name,
		"rules":        me.Rules,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"created_by":   &me.CreatedBy,
		"display_name": &me.DisplayName,
		"enabled":      &me.Enabled,
		"name":         &me.Name,
		"rules":        &me.Rules,
	})
}
