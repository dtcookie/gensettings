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

package optout

import (
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	DisplayOptOut bool    `json:"displayOptOut"`          // When true, the switch to go back to old screen will be visible.
	FeedbackUrl   *string `json:"feedbackUrl,omitempty"`  // The ID of SeaOtter entry that is an URL to feedback form. Leave empty to use default.
	Scope         string  `json:"-" scope:"scope"`        // The scope of this setting (ua-screen)
	UserVariable  *string `json:"userVariable,omitempty"` // The user variable key for opt-out switch. Leave empty to generate.
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"display_opt_out": {
			Type:        schema.TypeBool,
			Description: "When true, the switch to go back to old screen will be visible.",
			Required:    true,
		},
		"feedback_url": {
			Type:        schema.TypeString,
			Description: "The ID of SeaOtter entry that is an URL to feedback form. Leave empty to use default.",
			Optional:    true, // nullable
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (ua-screen)",
			Required:    true,
		},
		"user_variable": {
			Type:        schema.TypeString,
			Description: "The user variable key for opt-out switch. Leave empty to generate.",
			Optional:    true, // nullable & precondition
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"display_opt_out": me.DisplayOptOut,
		"feedback_url":    me.FeedbackUrl,
		"scope":           me.Scope,
		"user_variable":   me.UserVariable,
	})
}

func (me *Settings) HandlePreconditions() error {
	if me.UserVariable == nil && me.DisplayOptOut {
		return fmt.Errorf("'user_variable' must be specified if 'display_opt_out' is set to '%v'", me.DisplayOptOut)
	}
	return nil
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"display_opt_out": &me.DisplayOptOut,
		"feedback_url":    &me.FeedbackUrl,
		"scope":           &me.Scope,
		"user_variable":   &me.UserVariable,
	})
}
