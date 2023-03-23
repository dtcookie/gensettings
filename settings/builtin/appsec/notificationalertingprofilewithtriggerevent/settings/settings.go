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

package notificationalertingprofilewithtriggerevent

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Enabled              bool           `json:"enabled"`                        // This setting is enabled (`true`) or disabled (`false`)
	EnabledRiskLevels    []RiskLevel    `json:"enabledRiskLevels,omitempty"`    // Risk Levels
	EnabledTriggerEvents []TriggerEvent `json:"enabledTriggerEvents,omitempty"` // Alert for the following events:
	ManagementZone       *string        `json:"managementZone,omitempty"`       // Only if the following management zone is affected (optional)
	Name                 string         `json:"name"`                           // Name
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Type:        schema.TypeBool,
			Description: "This setting is enabled (`true`) or disabled (`false`)",
			Required:    true,
		},
		"enabled_risk_levels": {
			Type:        schema.TypeSet,
			Description: "Risk Levels",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"enabled_trigger_events": {
			Type:        schema.TypeSet,
			Description: "Alert for the following events:",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"management_zone": {
			Type:        schema.TypeString,
			Description: "Only if the following management zone is affected (optional)",
			Optional:    true, // nullable
		},
		"name": {
			Type:        schema.TypeString,
			Description: "Name",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"enabled":                me.Enabled,
		"enabled_risk_levels":    me.EnabledRiskLevels,
		"enabled_trigger_events": me.EnabledTriggerEvents,
		"management_zone":        me.ManagementZone,
		"name":                   me.Name,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":                &me.Enabled,
		"enabled_risk_levels":    &me.EnabledRiskLevels,
		"enabled_trigger_events": &me.EnabledTriggerEvents,
		"management_zone":        &me.ManagementZone,
		"name":                   &me.Name,
	})
}
