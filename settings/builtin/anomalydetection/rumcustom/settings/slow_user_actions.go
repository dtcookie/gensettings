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

package rumcustom

import (
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SlowUserActions struct {
	DetectionMode        *DetectionMode        `json:"detectionMode,omitempty"` // Possible Values: `Auto`, `Fixed`
	Enabled              bool                  `json:"enabled"`                 // This setting is enabled (`true`) or disabled (`false`)
	SlowUserActionsAuto  *SlowUserActionsAuto  `json:"slowUserActionsAuto,omitempty"`
	SlowUserActionsFixed *SlowUserActionsFixed `json:"slowUserActionsFixed,omitempty"`
}

func (me *SlowUserActions) Schema() map[string]*schema.Schema {
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
		"slow_user_actions_auto": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(SlowUserActionsAuto).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"slow_user_actions_fixed": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(SlowUserActionsFixed).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *SlowUserActions) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"detection_mode":          me.DetectionMode,
		"enabled":                 me.Enabled,
		"slow_user_actions_auto":  me.SlowUserActionsAuto,
		"slow_user_actions_fixed": me.SlowUserActionsFixed,
	})
}

func (me *SlowUserActions) HandlePreconditions() error {
	if (me.DetectionMode == nil) && (me.Enabled) {
		return fmt.Errorf("'detection_mode' must be specified if 'enabled' is set to '%v'", me.Enabled)
	}
	if (me.DetectionMode != nil) && (!me.Enabled) {
		return fmt.Errorf("'detection_mode' must not be specified if 'enabled' is set to '%v'", me.Enabled)
	}
	if (me.SlowUserActionsAuto == nil) && (me.Enabled && (me.DetectionMode != nil && (string(*me.DetectionMode) == "auto"))) {
		return fmt.Errorf("'slow_user_actions_auto' must be specified if 'enabled' is set to '%v' and 'detection_mode' is set to '%v'", me.Enabled, me.DetectionMode)
	}
	if (me.SlowUserActionsAuto != nil) && (!me.Enabled || (me.DetectionMode == nil || (me.DetectionMode != nil && string(*me.DetectionMode) != "auto"))) {
		return fmt.Errorf("'slow_user_actions_auto' must not be specified if 'enabled' is set to '%v'", me.Enabled)
	}
	if (me.SlowUserActionsFixed == nil) && (me.Enabled && (me.DetectionMode != nil && (string(*me.DetectionMode) == "fixed"))) {
		return fmt.Errorf("'slow_user_actions_fixed' must be specified if 'enabled' is set to '%v' and 'detection_mode' is set to '%v'", me.Enabled, me.DetectionMode)
	}
	if (me.SlowUserActionsFixed != nil) && (!me.Enabled || (me.DetectionMode == nil || (me.DetectionMode != nil && string(*me.DetectionMode) != "fixed"))) {
		return fmt.Errorf("'slow_user_actions_fixed' must not be specified if 'enabled' is set to '%v'", me.Enabled)
	}
	return nil
}

func (me *SlowUserActions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"detection_mode":          &me.DetectionMode,
		"enabled":                 &me.Enabled,
		"slow_user_actions_auto":  &me.SlowUserActionsAuto,
		"slow_user_actions_fixed": &me.SlowUserActionsFixed,
	})
}
