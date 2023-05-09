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

package messagecard

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ButtonTypes []*ButtonType

func (me *ButtonTypes) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"button": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(ButtonType).Schema()},
		},
	}
}

func (me ButtonTypes) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("button", me)
}

func (me *ButtonTypes) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("button", me)
}

// Button. Definition of button
type ButtonType struct {
	ActionExpression string          `json:"actionExpression"`     // The definition of an action. Single action expression. The format is: actionName|key=param|key2=param2
	Color            ButtonColorType `json:"color"`                // Possible Values: `PRIMARY`, `SECONDARY`
	Conditions       []string        `json:"conditions,omitempty"` // All conditions from the list need to be fulfilled for the button to be visible
	Text             *string         `json:"text,omitempty"`       // The display name of the button. Deduced from action definition if not set.
}

func (me *ButtonType) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action_expression": {
			Type:        schema.TypeString,
			Description: "The definition of an action. Single action expression. The format is: actionName|key=param|key2=param2",
			Required:    true,
		},
		"color": {
			Type:        schema.TypeString,
			Description: "Possible Values: `PRIMARY`, `SECONDARY`",
			Required:    true,
		},
		"conditions": {
			Type:        schema.TypeSet,
			Description: "All conditions from the list need to be fulfilled for the button to be visible",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"text": {
			Type:        schema.TypeString,
			Description: "The display name of the button. Deduced from action definition if not set.",
			Optional:    true, // nullable
		},
	}
}

func (me *ButtonType) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"action_expression": me.ActionExpression,
		"color":             me.Color,
		"conditions":        me.Conditions,
		"text":              me.Text,
	})
}

func (me *ButtonType) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"action_expression": &me.ActionExpression,
		"color":             &me.Color,
		"conditions":        &me.Conditions,
		"text":              &me.Text,
	})
}
