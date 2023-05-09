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

package listscreensettings

import (
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/opt"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"
)

type CardReferences []*CardReference

func (me *CardReferences) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"card": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(CardReference).Schema()},
		},
	}
}

func (me CardReferences) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("card", me)
}

func (me *CardReferences) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("card", me)
}

type CardReference struct {
	Collapsed  *bool      `json:"collapsed,omitempty"`  // If true, card content will be collapsed
	Conditions []string   `json:"conditions,omitempty"` // All conditions from the list need to be fulfilled for the card to be visible
	Key        *string    `json:"key,omitempty"`        // Unique key of the card, which is used to reference desired card configuration
	Type       CardType   `json:"type"`                 // Possible Values: `BREAK_LINE`, `CHART_GROUP`, `ENTITIES_LIST`, `INJECTIONS`, `MESSAGE`
	Width      *WidthType `json:"width,omitempty"`      // Possible Values: `FULL_SIZE`, `HALF_SIZE`
}

func (me *CardReference) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"collapsed": {
			Type:        schema.TypeBool,
			Description: "If true, card content will be collapsed",
			Optional:    true, // nullable & precondition
		},
		"conditions": {
			Type:        schema.TypeSet,
			Description: "All conditions from the list need to be fulfilled for the card to be visible",
			Optional:    true, // precondition & minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"key": {
			Type:        schema.TypeString,
			Description: "Unique key of the card, which is used to reference desired card configuration",
			Optional:    true, // precondition
		},
		"type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `BREAK_LINE`, `CHART_GROUP`, `ENTITIES_LIST`, `INJECTIONS`, `MESSAGE`",
			Required:    true,
		},
		"width": {
			Type:        schema.TypeString,
			Description: "Possible Values: `FULL_SIZE`, `HALF_SIZE`",
			Optional:    true, // nullable & precondition
		},
	}
}

func (me *CardReference) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"collapsed":  me.Collapsed,
		"conditions": me.Conditions,
		"key":        me.Key,
		"type":       me.Type,
		"width":      me.Width,
	})
}

func (me *CardReference) HandlePreconditions() error {
	if me.Collapsed == nil && !slices.Contains([]string{"INJECTIONS", "BREAK_LINE"}, string(me.Type)) {
		me.Collapsed = opt.NewBool(false)
	}
	if me.Key == nil && !slices.Contains([]string{"INJECTIONS", "BREAK_LINE"}, string(me.Type)) {
		return fmt.Errorf("'key' must be specified if 'type' is set to '%v'", me.Type)
	}
	if me.Width == nil && !slices.Contains([]string{"INJECTIONS", "BREAK_LINE"}, string(me.Type)) {
		return fmt.Errorf("'width' must be specified if 'type' is set to '%v'", me.Type)
	}
	if me.Width != nil && slices.Contains([]string{"INJECTIONS", "BREAK_LINE"}, string(me.Type)) {
		return fmt.Errorf("'width' must not be specified if 'type' is set to '%v'", me.Type)
	}
	// ---- Conditions []string -> {"precondition":{"expectedValue":"INJECTIONS","property":"type","type":"EQUALS"},"type":"NOT"}
	return nil
}

func (me *CardReference) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"collapsed":  &me.Collapsed,
		"conditions": &me.Conditions,
		"key":        &me.Key,
		"type":       &me.Type,
		"width":      &me.Width,
	})
}
