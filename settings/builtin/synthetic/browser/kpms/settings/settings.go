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

package kpms

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	LoadActionKpm LoadKpm `json:"loadActionKpm"`   // Possible Values: `CUMULATIVE_LAYOUT_SHIFT`, `DOM_INTERACTIVE`, `LARGEST_CONTENTFUL_PAINT`, `LOAD_EVENT_END`, `LOAD_EVENT_START`, `RESPONSE_END`, `RESPONSE_START`, `SPEED_INDEX`, `USER_ACTION_DURATION`, `VISUALLY_COMPLETE`
	Scope         string  `json:"-" scope:"scope"` // The scope of this setting (SYNTHETIC_TEST)
	XhrActionKpm  XhrKpm  `json:"xhrActionKpm"`    // Possible Values: `RESPONSE_END`, `RESPONSE_START`, `USER_ACTION_DURATION`, `VISUALLY_COMPLETE`
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"load_action_kpm": {
			Type:        schema.TypeString,
			Description: "Possible Values: `CUMULATIVE_LAYOUT_SHIFT`, `DOM_INTERACTIVE`, `LARGEST_CONTENTFUL_PAINT`, `LOAD_EVENT_END`, `LOAD_EVENT_START`, `RESPONSE_END`, `RESPONSE_START`, `SPEED_INDEX`, `USER_ACTION_DURATION`, `VISUALLY_COMPLETE`",
			Required:    true,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (SYNTHETIC_TEST)",
			Required:    true,
		},
		"xhr_action_kpm": {
			Type:        schema.TypeString,
			Description: "Possible Values: `RESPONSE_END`, `RESPONSE_START`, `USER_ACTION_DURATION`, `VISUALLY_COMPLETE`",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"load_action_kpm": me.LoadActionKpm,
		"scope":           me.Scope,
		"xhr_action_kpm":  me.XhrActionKpm,
	})
}

func (me *Settings) HandlePreconditions() {
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"load_action_kpm": &me.LoadActionKpm,
		"scope":           &me.Scope,
		"xhr_action_kpm":  &me.XhrActionKpm,
	})
}
