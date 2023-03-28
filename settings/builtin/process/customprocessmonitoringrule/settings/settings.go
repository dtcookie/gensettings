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

package customprocessmonitoringrule

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Condition   *Condition     `json:"condition"`             // Condition
	Enabled     bool           `json:"enabled"`               // This setting is enabled (`true`) or disabled (`false`)
	HostGroupID *string        `json:"-" scope:"hostGroupId"` // The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	Mode        MonitoringMode `json:"mode"`                  // Possible Values: `MONITORING_OFF`, `MONITORING_ON`
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"condition": {
			Type:        schema.TypeList,
			Description: "Condition",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(Condition).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "This setting is enabled (`true`) or disabled (`false`)",
			Required:    true,
		},
		"host_group_id": {
			Type:        schema.TypeString,
			Description: "The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.",
			Optional:    true,
			Default:     "environment",
		},
		"mode": {
			Type:        schema.TypeString,
			Description: "Possible Values: `MONITORING_OFF`, `MONITORING_ON`",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"condition":     me.Condition,
		"enabled":       me.Enabled,
		"host_group_id": me.HostGroupID,
		"mode":          me.Mode,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"condition":     &me.Condition,
		"enabled":       &me.Enabled,
		"host_group_id": &me.HostGroupID,
		"mode":          &me.Mode,
	})
}
