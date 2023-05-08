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

package infrastructurehosts

import (
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type NetworkDroppedPacketsDetection struct {
	CustomThresholds *NetworkDroppedPacketsDetectionThresholds `json:"customThresholds,omitempty"` // Alert if the dropped packet percentage is higher than the specified threshold **and** the total packets rate is higher than the defined threshold for the defined amount of samples
	DetectionMode    *DetectionMode                            `json:"detectionMode,omitempty"`    // Possible Values: `Auto`, `Custom`
	Enabled          bool                                      `json:"enabled"`                    // This setting is enabled (`true`) or disabled (`false`)
}

func (me *NetworkDroppedPacketsDetection) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"custom_thresholds": {
			Type:        schema.TypeList,
			Description: "Alert if the dropped packet percentage is higher than the specified threshold **and** the total packets rate is higher than the defined threshold for the defined amount of samples",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(NetworkDroppedPacketsDetectionThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_mode": {
			Type:        schema.TypeString,
			Description: "Possible Values: `Auto`, `Custom`",
			Optional:    true, // precondition
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "This setting is enabled (`true`) or disabled (`false`)",
			Required:    true,
		},
	}
}

func (me *NetworkDroppedPacketsDetection) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *NetworkDroppedPacketsDetection) HandlePreconditions() error {
	if me.CustomThresholds == nil && me.Enabled && (me.DetectionMode != nil && string(*me.DetectionMode) == "custom") {
		return fmt.Errorf("'custom_thresholds' must be specified if 'enabled' is set to '%v' and 'detection_mode' is set to '%v'", me.Enabled, me.DetectionMode)
	}
	if me.CustomThresholds != nil && !me.Enabled || me.DetectionMode == nil || (me.DetectionMode != nil && string(*me.DetectionMode) != "custom") {
		return fmt.Errorf("'custom_thresholds' must not be specified if 'enabled' is set to '%v'", me.Enabled)
	}
	if me.DetectionMode == nil && me.Enabled {
		return fmt.Errorf("'detection_mode' must be specified if 'enabled' is set to '%v'", me.Enabled)
	}
	if me.DetectionMode != nil && !me.Enabled {
		return fmt.Errorf("'detection_mode' must not be specified if 'enabled' is set to '%v'", me.Enabled)
	}
	return nil
}

func (me *NetworkDroppedPacketsDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
