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

package rummobile

import (
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ErrorRateIncrease struct {
	DetectionMode          *DetectionMode          `json:"detectionMode,omitempty"`          // Possible Values: `Auto`, `Fixed`
	Enabled                bool                    `json:"enabled"`                          // This setting is enabled (`true`) or disabled (`false`)
	ErrorRateIncreaseAuto  *ErrorRateIncreaseAuto  `json:"errorRateIncreaseAuto,omitempty"`  // Alert if the percentage of user actions affected by reported errors exceeds **both** the absolute threshold and the relative threshold
	ErrorRateIncreaseFixed *ErrorRateIncreaseFixed `json:"errorRateIncreaseFixed,omitempty"` // Alert if the custom reported error rate threshold is exceeded during any 5-minute period
}

func (me *ErrorRateIncrease) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"detection_mode": {
			Type:        schema.TypeString,
			Description: "Possible Values: `Auto`, `Fixed`",
			Optional:    true, // precondition
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "This setting is enabled (`true`) or disabled (`false`)",
			Required:    true,
		},
		"error_rate_increase_auto": {
			Type:        schema.TypeList,
			Description: "Alert if the percentage of user actions affected by reported errors exceeds **both** the absolute threshold and the relative threshold",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(ErrorRateIncreaseAuto).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"error_rate_increase_fixed": {
			Type:        schema.TypeList,
			Description: "Alert if the custom reported error rate threshold is exceeded during any 5-minute period",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(ErrorRateIncreaseFixed).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *ErrorRateIncrease) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"detection_mode":            me.DetectionMode,
		"enabled":                   me.Enabled,
		"error_rate_increase_auto":  me.ErrorRateIncreaseAuto,
		"error_rate_increase_fixed": me.ErrorRateIncreaseFixed,
	})
}

func (me *ErrorRateIncrease) HandlePreconditions() error {
	if (me.DetectionMode == nil) && (me.Enabled) {
		return fmt.Errorf("'detection_mode' must be specified if 'enabled' is set to '%v'", me.Enabled)
	}
	if (me.DetectionMode != nil) && (!me.Enabled) {
		return fmt.Errorf("'detection_mode' must not be specified if 'enabled' is set to '%v'", me.Enabled)
	}
	if (me.ErrorRateIncreaseAuto == nil) && (me.Enabled && (me.DetectionMode != nil && (string(*me.DetectionMode) == "auto"))) {
		return fmt.Errorf("'error_rate_increase_auto' must be specified if 'enabled' is set to '%v' and 'detection_mode' is set to '%v'", me.Enabled, me.DetectionMode)
	}
	if (me.ErrorRateIncreaseAuto != nil) && (!me.Enabled || (me.DetectionMode == nil || (me.DetectionMode != nil && string(*me.DetectionMode) != "auto"))) {
		return fmt.Errorf("'error_rate_increase_auto' must not be specified if 'enabled' is set to '%v'", me.Enabled)
	}
	if (me.ErrorRateIncreaseFixed == nil) && (me.Enabled && (me.DetectionMode != nil && (string(*me.DetectionMode) == "fixed"))) {
		return fmt.Errorf("'error_rate_increase_fixed' must be specified if 'enabled' is set to '%v' and 'detection_mode' is set to '%v'", me.Enabled, me.DetectionMode)
	}
	if (me.ErrorRateIncreaseFixed != nil) && (!me.Enabled || (me.DetectionMode == nil || (me.DetectionMode != nil && string(*me.DetectionMode) != "fixed"))) {
		return fmt.Errorf("'error_rate_increase_fixed' must not be specified if 'enabled' is set to '%v'", me.Enabled)
	}
	return nil
}

func (me *ErrorRateIncrease) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"detection_mode":            &me.DetectionMode,
		"enabled":                   &me.Enabled,
		"error_rate_increase_auto":  &me.ErrorRateIncreaseAuto,
		"error_rate_increase_fixed": &me.ErrorRateIncreaseFixed,
	})
}
