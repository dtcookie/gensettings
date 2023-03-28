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

package usersessionexportsettingsv2

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ExportBehavior struct {
	CustomConfiguration *string `json:"customConfiguration,omitempty"` // Here you can set additional properties for this export configuration. Changing these values might only be necessary in some rare cases. Please contact support before changing this field.
	DisableNotification bool    `json:"disableNotification"`           // Disable notification
	ManagementZone      *string `json:"managementZone,omitempty"`      // Restrict exported sessions to management zone
}

func (me *ExportBehavior) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"custom_configuration": {
			Type:        schema.TypeString,
			Description: "Here you can set additional properties for this export configuration. Changing these values might only be necessary in some rare cases. Please contact support before changing this field.",
			Optional:    true, // nullable
		},
		"disable_notification": {
			Type:        schema.TypeBool,
			Description: "Disable notification",
			Required:    true,
		},
		"management_zone": {
			Type:        schema.TypeString,
			Description: "Restrict exported sessions to management zone",
			Optional:    true, // nullable
		},
	}
}

func (me *ExportBehavior) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"custom_configuration": me.CustomConfiguration,
		"disable_notification": me.DisableNotification,
		"management_zone":      me.ManagementZone,
	})
}

func (me *ExportBehavior) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_configuration": &me.CustomConfiguration,
		"disable_notification": &me.DisableNotification,
		"management_zone":      &me.ManagementZone,
	})
}
