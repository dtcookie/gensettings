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
	AlertingMode             *AlertingMode `json:"alertingMode,omitempty"`             // Possible Values: `ON_INSTANCE_COUNT_VIOLATION`, `ON_PGI_UNAVAILABILITY`
	Enabled                  bool          `json:"enabled"`                            // This setting is enabled (`true`) or disabled (`false`)
	MinimumInstanceThreshold *int          `json:"minimumInstanceThreshold,omitempty"` // Open a new problem if the number of active process instances in the group is fewer than:
	ProcessGroupID           string        `json:"-" scope:"processGroupId"`           // The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"alerting_mode": {
			Type:        schema.TypeString,
			Description: "Possible Values: `ON_INSTANCE_COUNT_VIOLATION`, `ON_PGI_UNAVAILABILITY`",
			Optional:    true, // precondition
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "This setting is enabled (`true`) or disabled (`false`)",
			Required:    true,
		},
		"minimum_instance_threshold": {
			Type:        schema.TypeInt,
			Description: "Open a new problem if the number of active process instances in the group is fewer than:",
			Optional:    true, // precondition
		},
		"process_group_id": {
			Type:        schema.TypeString,
			Description: "The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"alerting_mode":              me.AlertingMode,
		"enabled":                    me.Enabled,
		"minimum_instance_threshold": me.MinimumInstanceThreshold,
		"process_group_id":           me.ProcessGroupID,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"alerting_mode":              &me.AlertingMode,
		"enabled":                    &me.Enabled,
		"minimum_instance_threshold": &me.MinimumInstanceThreshold,
		"process_group_id":           &me.ProcessGroupID,
	})
}
