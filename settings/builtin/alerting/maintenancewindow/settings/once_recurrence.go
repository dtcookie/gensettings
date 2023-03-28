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

// Maintenance window. The date time window when the maintenance will take place once.
type OnceRecurrence struct {
	EndTime   string `json:"endTime"`   // End time
	StartTime string `json:"startTime"` // Start time
	TimeZone  string `json:"timeZone"`
}

func (me *OnceRecurrence) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"end_time": {
			Type:        schema.TypeString,
			Description: "End time",
			Required:    true,
		},
		"start_time": {
			Type:        schema.TypeString,
			Description: "Start time",
			Required:    true,
		},
		"time_zone": {
			Type:        schema.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *OnceRecurrence) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"end_time":   me.EndTime,
		"start_time": me.StartTime,
		"time_zone":  me.TimeZone,
	})
}

func (me *OnceRecurrence) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"end_time":   &me.EndTime,
		"start_time": &me.StartTime,
		"time_zone":  &me.TimeZone,
	})
}
