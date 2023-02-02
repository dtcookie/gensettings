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

package monitoring

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	DisplayName string  `json:"displayName"`     // Display name
	Enabled     bool    `json:"enabled"`         // Enabled
	Scope       *string `json:"-" scope:"scope"` // The scope of this setting (HOST HOST_GROUP environment)
	ServiceName string  `json:"serviceName"`     // Service name
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"display_name": {
			Type:        schema.TypeString,
			Description: "Display name",
			Required:    true,
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "Enabled",
			Required:    true,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (HOST HOST_GROUP environment)",
			Optional:    true,
			Default:     "environment",
		},
		"service_name": {
			Type:        schema.TypeString,
			Description: "Service name",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"display_name": me.DisplayName,
		"enabled":      me.Enabled,
		"scope":        me.Scope,
		"service_name": me.ServiceName,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"display_name": &me.DisplayName,
		"enabled":      &me.Enabled,
		"scope":        &me.Scope,
		"service_name": &me.ServiceName,
	})
}
