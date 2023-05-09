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

	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"
)

type NestedCardReferences []*NestedCardReference

func (me *NestedCardReferences) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"card": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(NestedCardReference).Schema()},
		},
	}
}

func (me NestedCardReferences) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("card", me)
}

func (me *NestedCardReferences) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("card", me)
}

type NestedCardReference struct {
	Conditions             []string `json:"conditions,omitempty"`             // All conditions from the list need to be fulfilled for the card to be visible
	EntitySelectorTemplate *string  `json:"entitySelectorTemplate,omitempty"` // An entity selector that is used to reference card from other ME type screen. For example for a host screen where you want to show the processes charts this might look like this:  \n`type(PROCESS_GROUP_INSTANCE), fromRelationships.isProcessOf($(entityConditions))`\n\nPlease mind that the `$(entityConditions)` is a placeholder for the relation condition.
	Key                    *string  `json:"key,omitempty"`                    // Unique key of the card, which is used to reference desired card configuration
	Type                   CardType `json:"type"`                             // Possible Values: `ATTRIBUTE_TABLE`, `BREAK_LINE`, `CARD_GROUP`, `CHART_GROUP`, `ENTITIES_LIST`, `EVENTS`, `HEALTH_CARD`, `INJECTIONS`, `LOGS`, `MESSAGE`, `METRIC_TABLE`, `SLO_LIST`, `TRACES`
}

func (me *NestedCardReference) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"conditions": {
			Type:        schema.TypeSet,
			Description: "All conditions from the list need to be fulfilled for the card to be visible",
			Optional:    true, // precondition & minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
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
		"type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `ATTRIBUTE_TABLE`, `BREAK_LINE`, `CARD_GROUP`, `CHART_GROUP`, `ENTITIES_LIST`, `EVENTS`, `HEALTH_CARD`, `INJECTIONS`, `LOGS`, `MESSAGE`, `METRIC_TABLE`, `SLO_LIST`, `TRACES`",
			Required:    true,
		},
	}
}

func (me *NestedCardReference) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"conditions":               me.Conditions,
		"entity_selector_template": me.EntitySelectorTemplate,
		"key":                      me.Key,
		"type":                     me.Type,
	})
}

func (me *NestedCardReference) HandlePreconditions() error {
	if me.EntitySelectorTemplate == nil && !slices.Contains([]string{"INJECTIONS", "BREAK_LINE", "TRACES", "SLO_LIST", "CARD_GROUP"}, string(me.Type)) {
		return fmt.Errorf("'entity_selector_template' must be specified if 'type' is set to '%v'", me.Type)
	}
	if me.Key == nil && !slices.Contains([]string{"INJECTIONS", "BREAK_LINE", "CARD_GROUP"}, string(me.Type)) {
		return fmt.Errorf("'key' must be specified if 'type' is set to '%v'", me.Type)
	}
	// ---- Conditions []string -> {"precondition":{"expectedValue":"INJECTIONS","property":"type","type":"EQUALS"},"type":"NOT"}
	return nil
}

func (me *NestedCardReference) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"conditions":               &me.Conditions,
		"entity_selector_template": &me.EntitySelectorTemplate,
		"key":                      &me.Key,
		"type":                     &me.Type,
	})
}
