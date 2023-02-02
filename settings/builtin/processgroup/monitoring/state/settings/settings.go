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

package state

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	MonitoringState ProcessGroupMonitoringMode `json:"MonitoringState"`          // Possible Values: `MONITORING_OFF`, `MONITORING_ON`, `DEFAULT`
	ProcessGroupID  string                     `json:"-" scope:"processGroupId"` // The scope of this setting (PROCESS_GROUP)
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"monitoring_state": {
			Type:        schema.TypeString,
			Description: "Possible Values: `MONITORING_OFF`, `MONITORING_ON`, `DEFAULT`",
			Required:    true,
		},
		"process_group_id": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (PROCESS_GROUP)",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"monitoring_state": me.MonitoringState,
		"process_group_id": me.ProcessGroupID,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"monitoring_state": &me.MonitoringState,
		"process_group_id": &me.ProcessGroupID,
	})
}
