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

package screensettings

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
	Cards                  NestedCardReferences `json:"cards,omitempty"`                  // List of references to cards displayed in card group
	Collapsed              *bool                `json:"collapsed,omitempty"`              // If true, card content will be collapsed
	Conditions             []string             `json:"conditions,omitempty"`             // All conditions from the list need to be fulfilled for the card to be visible
	DisplayInSidebar       *bool                `json:"displayInSidebar,omitempty"`       // If true, card will be available in sidebar and button in toolbar will appear to show it
	DisplayName            *string              `json:"displayName,omitempty"`            // Display name of a card group
	EntitySelectorTemplate *string              `json:"entitySelectorTemplate,omitempty"` // An entity selector that is used to reference card from other ME type screen. For example for a host screen where you want to show the processes charts this might look like this:  \n`type(PROCESS_GROUP_INSTANCE), fromRelationships.isProcessOf($(entityConditions))`\n\nPlease mind that the `$(entityConditions)` is a placeholder for the relation condition.
	Key                    *string              `json:"key,omitempty"`                    // Unique key of the card, which is used to reference desired card configuration
	RibbonButton           *string              `json:"ribbonButton,omitempty"`           // Custom component key to render as ribbon button. Every card has its default ribbon button. It can be overriden with that option and custom component on frontend.
	Type                   CardType             `json:"type"`                             // Possible Values: `ATTRIBUTE_TABLE`, `BREAK_LINE`, `CARD_GROUP`, `CHART_GROUP`, `ENTITIES_LIST`, `EVENTS`, `HEALTH_CARD`, `INJECTIONS`, `LOGS`, `MESSAGE`, `METRIC_TABLE`, `SLO_LIST`, `TRACES`
	Width                  *WidthType           `json:"width,omitempty"`                  // Possible Values: `FULL_SIZE`, `HALF_SIZE`
}

func (me *CardReference) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cards": {
			Type:        schema.TypeList,
			Description: "List of references to cards displayed in card group",
			Optional:    true, // precondition & minobjects == 0
			Elem:        &schema.Resource{Schema: new(NestedCardReferences).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
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
		"display_in_sidebar": {
			Type:        schema.TypeBool,
			Description: "If true, card will be available in sidebar and button in toolbar will appear to show it",
			Optional:    true, // nullable & precondition
		},
		"display_name": {
			Type:        schema.TypeString,
			Description: "Display name of a card group",
			Optional:    true, // precondition
		},
		"entity_selector_template": {
			Type:        schema.TypeString,
			Description: "An entity selector that is used to reference card from other ME type screen. For example for a host screen where you want to show the processes charts this might look like this:  \n`type(PROCESS_GROUP_INSTANCE), fromRelationships.isProcessOf($(entityConditions))`\n\nPlease mind that the `$(entityConditions)` is a placeholder for the relation condition.",
			Optional:    true, // nullable & precondition
		},
		"key": {
			Type:        schema.TypeString,
			Description: "Unique key of the card, which is used to reference desired card configuration",
			Optional:    true, // precondition
		},
		"ribbon_button": {
			Type:        schema.TypeString,
			Description: "Custom component key to render as ribbon button. Every card has its default ribbon button. It can be overriden with that option and custom component on frontend.",
			Optional:    true, // nullable & precondition
		},
		"type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `ATTRIBUTE_TABLE`, `BREAK_LINE`, `CARD_GROUP`, `CHART_GROUP`, `ENTITIES_LIST`, `EVENTS`, `HEALTH_CARD`, `INJECTIONS`, `LOGS`, `MESSAGE`, `METRIC_TABLE`, `SLO_LIST`, `TRACES`",
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
		"cards":                    me.Cards,
		"collapsed":                me.Collapsed,
		"conditions":               me.Conditions,
		"display_in_sidebar":       me.DisplayInSidebar,
		"display_name":             me.DisplayName,
		"entity_selector_template": me.EntitySelectorTemplate,
		"key":                      me.Key,
		"ribbon_button":            me.RibbonButton,
		"type":                     me.Type,
		"width":                    me.Width,
	})
}

func (me *CardReference) HandlePreconditions() error {
	if me.Collapsed == nil && !slices.Contains([]string{"INJECTIONS", "BREAK_LINE", "HEALTH_CARD"}, string(me.Type)) {
		me.Collapsed = opt.NewBool(false)
	}
	if me.DisplayInSidebar == nil && !slices.Contains([]string{"INJECTIONS", "BREAK_LINE", "MESSAGE", "HEALTH_CARD", "CARD_GROUP"}, string(me.Type)) {
		me.DisplayInSidebar = opt.NewBool(false)
	}
	if me.DisplayName == nil && (string(me.Type) == "CARD_GROUP") {
		return fmt.Errorf("'display_name' must be specified if 'type' is set to '%v'", me.Type)
	}
	if me.EntitySelectorTemplate == nil && !slices.Contains([]string{"INJECTIONS", "BREAK_LINE", "TRACES", "SLO_LIST", "HEALTH_CARD", "CARD_GROUP"}, string(me.Type)) {
		return fmt.Errorf("'entity_selector_template' must be specified if 'type' is set to '%v'", me.Type)
	}
	if me.Key == nil && !slices.Contains([]string{"INJECTIONS", "BREAK_LINE", "CARD_GROUP"}, string(me.Type)) {
		return fmt.Errorf("'key' must be specified if 'type' is set to '%v'", me.Type)
	}
	if me.RibbonButton == nil && me.DisplayInSidebar != nil && *me.DisplayInSidebar {
		return fmt.Errorf("'ribbon_button' must be specified if 'display_in_sidebar' is set to '%v'", me.DisplayInSidebar)
	}
	if me.Width == nil && !slices.Contains([]string{"INJECTIONS", "BREAK_LINE", "HEALTH_CARD"}, string(me.Type)) {
		return fmt.Errorf("'width' must be specified if 'type' is set to '%v'", me.Type)
	}
	if me.Width != nil && slices.Contains([]string{"INJECTIONS", "BREAK_LINE", "HEALTH_CARD"}, string(me.Type)) {
		return fmt.Errorf("'width' must not be specified if 'type' is set to '%v'", me.Type)
	}
	// ---- Cards NestedCardReferences -> {"expectedValue":"CARD_GROUP","property":"type","type":"EQUALS"}
	// ---- Conditions []string -> {"precondition":{"expectedValue":"INJECTIONS","property":"type","type":"EQUALS"},"type":"NOT"}
	return nil
}

func (me *CardReference) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"cards":                    &me.Cards,
		"collapsed":                &me.Collapsed,
		"conditions":               &me.Conditions,
		"display_in_sidebar":       &me.DisplayInSidebar,
		"display_name":             &me.DisplayName,
		"entity_selector_template": &me.EntitySelectorTemplate,
		"key":                      &me.Key,
		"ribbon_button":            &me.RibbonButton,
		"type":                     &me.Type,
		"width":                    &me.Width,
	})
}
