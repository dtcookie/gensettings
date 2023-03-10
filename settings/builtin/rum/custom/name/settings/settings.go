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

package name

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	ApplicationName string          `json:"applicationName"` // Update application name
	ApplicationType ApplicationType `json:"applicationType"` // Possible Values: `Desktop`, `Echo`, `Embedded_pc`, `Hololens`, `Iot`, `Ufo`
	Scope           string          `json:"-" scope:"scope"` // The scope of this setting (CUSTOM_APPLICATION)
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_name": {
			Type:        schema.TypeString,
			Description: "Update application name",
			Required:    true,
		},
		"application_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `Desktop`, `Echo`, `Embedded_pc`, `Hololens`, `Iot`, `Ufo`",
			Required:    true,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (CUSTOM_APPLICATION)",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"application_name": me.ApplicationName,
		"application_type": me.ApplicationType,
		"scope":            me.Scope,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"application_name": &me.ApplicationName,
		"application_type": &me.ApplicationType,
		"scope":            &me.Scope,
	})
}
