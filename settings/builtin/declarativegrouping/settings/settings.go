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

package declarativegrouping

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Detection ProcessDefinitions `json:"detection"`       // Enter a descriptive process group display name and a unique identifier that Dynatrace can use to recognize this process group.
	Enabled   bool               `json:"enabled"`         // Enabled
	Name      string             `json:"name"`            // Monitored technology name
	Scope     *string            `json:"-" scope:"scope"` // The scope of this setting (HOST HOST_GROUP environment)
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"detection": {
			Type:        schema.TypeList,
			Description: "Enter a descriptive process group display name and a unique identifier that Dynatrace can use to recognize this process group.",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(ProcessDefinitions).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "Enabled",
			Required:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Description: "Monitored technology name",
			Required:    true,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (HOST HOST_GROUP environment)",
			Optional:    true,
			Default:     "environment",
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"detection": me.Detection,
		"enabled":   me.Enabled,
		"name":      me.Name,
		"scope":     me.Scope,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"detection": &me.Detection,
		"enabled":   &me.Enabled,
		"name":      &me.Name,
		"scope":     &me.Scope,
	})
}
