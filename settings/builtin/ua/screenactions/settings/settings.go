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
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"
)

type Settings struct {
	ActionLocation *ActionLocation `json:"actionLocation,omitempty"` // Possible Values: `CHART`, `HEADER`, `TABLE_ENTRY`
	ActionScope    ActionScope     `json:"actionScope"`              // Possible Values: `ATTRIBUTE_TABLE`, `CHART_GROUP`, `ENTITIES_LIST`, `EVENTS`, `GLOBAL_DETAILS`, `GLOBAL_LIST`, `LOGS`, `METRIC_TABLE`, `PROBLEMS`, `PROPERTIES`, `SLO_LIST`
	Actions        Actions         `json:"actions"`                  // Define all actions and their parameters you want to show on this element.
	Key            *string         `json:"key,omitempty"`            // Unique key of the card, which is used to reference desired card configuration. To add this actions in every card of this type, leave it empty.
	MainAction     *bool           `json:"mainAction,omitempty"`     // If true, buttons will be displayed outside of actions menu
	Scope          string          `json:"-" scope:"scope"`          // The scope of this setting (ua-screen)
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action_location": {
			Type:        schema.TypeString,
			Description: "Possible Values: `CHART`, `HEADER`, `TABLE_ENTRY`",
			Optional:    true, // precondition
		},
		"action_scope": {
			Type:        schema.TypeString,
			Description: "Possible Values: `ATTRIBUTE_TABLE`, `CHART_GROUP`, `ENTITIES_LIST`, `EVENTS`, `GLOBAL_DETAILS`, `GLOBAL_LIST`, `LOGS`, `METRIC_TABLE`, `PROBLEMS`, `PROPERTIES`, `SLO_LIST`",
			Required:    true,
		},
		"actions": {
			Type:        schema.TypeList,
			Description: "Define all actions and their parameters you want to show on this element.",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(Actions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"key": {
			Type:        schema.TypeString,
			Description: "Unique key of the card, which is used to reference desired card configuration. To add this actions in every card of this type, leave it empty.",
			Optional:    true, // nullable & precondition
		},
		"main_action": {
			Type:        schema.TypeBool,
			Description: "If true, buttons will be displayed outside of actions menu",
			Optional:    true, // nullable
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (ua-screen)",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"action_location": me.ActionLocation,
		"action_scope":    me.ActionScope,
		"actions":         me.Actions,
		"key":             me.Key,
		"main_action":     me.MainAction,
		"scope":           me.Scope,
	})
}

func (me *Settings) HandlePreconditions() error {
	if me.ActionLocation == nil && slices.Contains([]string{"CHART_GROUP", "ENTITIES_LIST", "METRIC_TABLE", "ATTRIBUTE_TABLE", "EVENTS", "LOGS"}, string(me.ActionScope)) {
		return fmt.Errorf("'action_location' must be specified if 'action_scope' is set to '%v'", me.ActionScope)
	}
	if me.ActionLocation != nil && !slices.Contains([]string{"CHART_GROUP", "ENTITIES_LIST", "METRIC_TABLE", "ATTRIBUTE_TABLE", "EVENTS", "LOGS"}, string(me.ActionScope)) {
		return fmt.Errorf("'action_location' must not be specified if 'action_scope' is set to '%v'", me.ActionScope)
	}
	if me.Key == nil && slices.Contains([]string{"CHART_GROUP", "ENTITIES_LIST", "METRIC_TABLE", "ATTRIBUTE_TABLE", "EVENTS", "LOGS", "SLO_LIST"}, string(me.ActionScope)) {
		return fmt.Errorf("'key' must be specified if 'action_scope' is set to '%v'", me.ActionScope)
	}
	return nil
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"action_location": &me.ActionLocation,
		"action_scope":    &me.ActionScope,
		"actions":         &me.Actions,
		"key":             &me.Key,
		"main_action":     &me.MainAction,
		"scope":           &me.Scope,
	})
}
