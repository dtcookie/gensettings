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

package config

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Active     bool       `json:"active"` // Enabled
	Properties Properties `json:"properties,omitempty"`
	Scope      string     `json:"-" scope:"scope"` // The scope of this setting (MULTIPROTOCOL_MONITOR)
	Steps      Steps      `json:"steps,omitempty"`
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active": {
			Type:        schema.TypeBool,
			Description: "Enabled",
			Required:    true,
		},
		"properties": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(Properties).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (MULTIPROTOCOL_MONITOR)",
			Required:    true,
		},
		"steps": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(Steps).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"active":     me.Active,
		"properties": me.Properties,
		"scope":      me.Scope,
		"steps":      me.Steps,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"active":     &me.Active,
		"properties": &me.Properties,
		"scope":      &me.Scope,
		"steps":      &me.Steps,
	})
}
