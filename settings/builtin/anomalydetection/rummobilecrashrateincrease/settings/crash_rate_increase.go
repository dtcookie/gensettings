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

package rummobilecrashrateincrease

import (
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type CrashRateIncrease struct {
	CrashRateIncreaseAuto  *CrashRateIncreaseAuto  `json:"crashRateIncreaseAuto,omitempty"`  // Alert to crash rate increases when the auto-detected baseline is exceeded and the application has a minimum number of active, non-distinctive users.
	CrashRateIncreaseFixed *CrashRateIncreaseFixed `json:"crashRateIncreaseFixed,omitempty"` // Alert to crash rate increases when the defined threshold is exceeded and the application has a minimum number of active, non-distinctive users.
	DetectionMode          *DetectionMode          `json:"detectionMode,omitempty"`          // Possible Values: `Auto`, `Fixed`
	Enabled                bool                    `json:"enabled"`                          // This setting is enabled (`true`) or disabled (`false`)
}

func (me *CrashRateIncrease) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"crash_rate_increase_auto": {
			Type:        schema.TypeList,
			Description: "Alert to crash rate increases when the auto-detected baseline is exceeded and the application has a minimum number of active, non-distinctive users.",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(CrashRateIncreaseAuto).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"crash_rate_increase_fixed": {
			Type:        schema.TypeList,
			Description: "Alert to crash rate increases when the defined threshold is exceeded and the application has a minimum number of active, non-distinctive users.",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(CrashRateIncreaseFixed).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
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
	}
}

func (me *CrashRateIncrease) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"crash_rate_increase_auto":  me.CrashRateIncreaseAuto,
		"crash_rate_increase_fixed": me.CrashRateIncreaseFixed,
		"detection_mode":            me.DetectionMode,
		"enabled":                   me.Enabled,
	})
}

func (me *CrashRateIncrease) HandlePreconditions() error {
	if me.CrashRateIncreaseAuto == nil && me.Enabled && (me.DetectionMode != nil && string(*me.DetectionMode) == "auto") {
		return fmt.Errorf("'crash_rate_increase_auto' must be specified if 'enabled' is set to '%v' and 'detection_mode' is set to '%v'", me.Enabled, me.DetectionMode)
	}
	if me.CrashRateIncreaseAuto != nil && !me.Enabled || me.DetectionMode == nil || (me.DetectionMode != nil && string(*me.DetectionMode) != "auto") {
		return fmt.Errorf("'crash_rate_increase_auto' must not be specified if 'enabled' is set to '%v'", me.Enabled)
	}
	if me.CrashRateIncreaseFixed == nil && me.Enabled && (me.DetectionMode != nil && string(*me.DetectionMode) == "fixed") {
		return fmt.Errorf("'crash_rate_increase_fixed' must be specified if 'enabled' is set to '%v' and 'detection_mode' is set to '%v'", me.Enabled, me.DetectionMode)
	}
	if me.CrashRateIncreaseFixed != nil && !me.Enabled || me.DetectionMode == nil || (me.DetectionMode != nil && string(*me.DetectionMode) != "fixed") {
		return fmt.Errorf("'crash_rate_increase_fixed' must not be specified if 'enabled' is set to '%v'", me.Enabled)
	}
	if me.DetectionMode == nil && me.Enabled {
		return fmt.Errorf("'detection_mode' must be specified if 'enabled' is set to '%v'", me.Enabled)
	}
	if me.DetectionMode != nil && !me.Enabled {
		return fmt.Errorf("'detection_mode' must not be specified if 'enabled' is set to '%v'", me.Enabled)
	}
	return nil
}

func (me *CrashRateIncrease) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"crash_rate_increase_auto":  &me.CrashRateIncreaseAuto,
		"crash_rate_increase_fixed": &me.CrashRateIncreaseFixed,
		"detection_mode":            &me.DetectionMode,
		"enabled":                   &me.Enabled,
	})
}
