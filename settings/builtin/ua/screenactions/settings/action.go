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

package screenactions

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Actions []*Action

func (me *Actions) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(Action).Schema()},
		},
	}
}

func (me Actions) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("action", me)
}

func (me *Actions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("action", me)
}

type Action struct {
	ActionExpression string         `json:"actionExpression"`        // The definition of an action. Single action expression. The format is: actionName|key=param|key2=param2
	Conditions       []string       `json:"conditions,omitempty"`    // All conditions from the list need to be fulfilled for the action to be visible
	Visualization    *Visualization `json:"visualization,omitempty"` // If applicable, you can define visualization properties of action button
}

func (me *Action) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action_expression": {
			Type:        schema.TypeString,
			Description: "The definition of an action. Single action expression. The format is: actionName|key=param|key2=param2",
			Required:    true,
		},
		"conditions": {
			Type:        schema.TypeSet,
			Description: "All conditions from the list need to be fulfilled for the action to be visible",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"visualization": {
			Type:        schema.TypeList,
			Description: "If applicable, you can define visualization properties of action button",
			Optional:    true, // nullable
			Elem:        &schema.Resource{Schema: new(Visualization).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Action) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"action_expression": me.ActionExpression,
		"conditions":        me.Conditions,
		"visualization":     me.Visualization,
	})
}

func (me *Action) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"action_expression": &me.ActionExpression,
		"conditions":        &me.Conditions,
		"visualization":     &me.Visualization,
	})
}
