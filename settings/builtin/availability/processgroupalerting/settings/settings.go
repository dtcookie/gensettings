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

package processgroupalerting

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	AlertingMode             AlertingMode `json:"alertingMode"`             // **if any process becomes unavailable:**\nDynatrace will open a new problem if a single process in this group shuts down or crashes. \n\n**if minimum threshold is not met:**\nDynatrace tracks the number of process instances that comprise this process group and treats the group as a cluster. This setting enables you to define a minimum number of process instances that must be available. A problem will be opened if this process group has fewer than the minimum number of required process instances. \n\n Details of the related impact on service requests will be included in the problem summary.\n\n**Note:** If a process is intentionally shutdown or retired while this setting is active, you'll need to manually close the problem.
	Enabled                  bool         `json:"enabled"`                  // Enable process group availability monitoring
	MinimumInstanceThreshold int          `json:"minimumInstanceThreshold"` // Open a new problem if the number of active process instances in the group is fewer than:
	Scope                    string       `json:"-" scope:"scope"`          // The scope of this setting (PROCESS_GROUP)
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"alerting_mode": {
			Type:        schema.TypeString,
			Description: "**if any process becomes unavailable:**\nDynatrace will open a new problem if a single process in this group shuts down or crashes. \n\n**if minimum threshold is not met:**\nDynatrace tracks the number of process instances that comprise this process group and treats the group as a cluster. This setting enables you to define a minimum number of process instances that must be available. A problem will be opened if this process group has fewer than the minimum number of required process instances. \n\n Details of the related impact on service requests will be included in the problem summary.\n\n**Note:** If a process is intentionally shutdown or retired while this setting is active, you'll need to manually close the problem.",
			Required:    true,
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "Enable process group availability monitoring",
			Required:    true,
		},
		"minimum_instance_threshold": {
			Type:        schema.TypeInt,
			Description: "Open a new problem if the number of active process instances in the group is fewer than:",
			Required:    true,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (PROCESS_GROUP)",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"alerting_mode":              me.AlertingMode,
		"enabled":                    me.Enabled,
		"minimum_instance_threshold": me.MinimumInstanceThreshold,
		"scope":                      me.Scope,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"alerting_mode":              &me.AlertingMode,
		"enabled":                    &me.Enabled,
		"minimum_instance_threshold": &me.MinimumInstanceThreshold,
		"scope":                      &me.Scope,
	})
}
