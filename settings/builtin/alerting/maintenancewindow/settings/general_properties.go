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

package maintenancewindow

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type GeneralProperties struct {
	Description                      string          `json:"description"`                      // A short description of the maintenance purpose.
	DisableSyntheticMonitorExecution bool            `json:"disableSyntheticMonitorExecution"` // Disables the execution of the synthetic monitors that are within [the scope of this maintenance window](https://dt-url.net/0e0341m).
	MaintenanceType                  MaintenanceType `json:"maintenanceType"`                  // Possible Values: `PLANNED`, `UNPLANNED`
	Name                             string          `json:"name"`
	Suppression                      SuppressionType `json:"suppression"` // Possible Values: `DETECT_PROBLEMS_AND_ALERT`, `DETECT_PROBLEMS_DONT_ALERT`, `DONT_DETECT_PROBLEMS`
}

func (me *GeneralProperties) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:        schema.TypeString,
			Description: "A short description of the maintenance purpose.",
			Required:    true,
		},
		"disable_synthetic_monitor_execution": {
			Type:        schema.TypeBool,
			Description: "Disables the execution of the synthetic monitors that are within [the scope of this maintenance window](https://dt-url.net/0e0341m).",
			Required:    true,
		},
		"maintenance_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `PLANNED`, `UNPLANNED`",
			Required:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"suppression": {
			Type:        schema.TypeString,
			Description: "Possible Values: `DETECT_PROBLEMS_AND_ALERT`, `DETECT_PROBLEMS_DONT_ALERT`, `DONT_DETECT_PROBLEMS`",
			Required:    true,
		},
	}
}

func (me *GeneralProperties) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"description":                         me.Description,
		"disable_synthetic_monitor_execution": me.DisableSyntheticMonitorExecution,
		"maintenance_type":                    me.MaintenanceType,
		"name":                                me.Name,
		"suppression":                         me.Suppression,
	})
}

func (me *GeneralProperties) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"description":                         &me.Description,
		"disable_synthetic_monitor_execution": &me.DisableSyntheticMonitorExecution,
		"maintenance_type":                    &me.MaintenanceType,
		"name":                                &me.Name,
		"suppression":                         &me.Suppression,
	})
}
