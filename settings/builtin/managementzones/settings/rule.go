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

package managementzones

import (
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Rules []*Rule

func (me *Rules) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"rule": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(Rule).Schema()},
		},
	}
}

func (me Rules) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("rule", me)
}

func (me *Rules) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("rule", me)
}

type Rule struct {
	AttributeRule  *ManagementZoneAttributeRule `json:"attributeRule,omitempty"`
	DimensionRule  *DimensionRule               `json:"dimensionRule,omitempty"`
	Enabled        bool                         `json:"enabled"`                  // This setting is enabled (`true`) or disabled (`false`)
	EntitySelector *string                      `json:"entitySelector,omitempty"` // The documentation of the entity selector can be found [here](https://dt-url.net/apientityselector).
	Type           RuleType                     `json:"type"`                     // Possible Values: `DIMENSION`, `ME`, `SELECTOR`
}

func (me *Rule) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attribute_rule": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(ManagementZoneAttributeRule).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"dimension_rule": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(DimensionRule).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "This setting is enabled (`true`) or disabled (`false`)",
			Required:    true,
		},
		"entity_selector": {
			Type:        schema.TypeString,
			Description: "The documentation of the entity selector can be found [here](https://dt-url.net/apientityselector).",
			Optional:    true, // precondition
		},
		"type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `DIMENSION`, `ME`, `SELECTOR`",
			Required:    true,
		},
	}
}

func (me *Rule) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"attribute_rule":  me.AttributeRule,
		"dimension_rule":  me.DimensionRule,
		"enabled":         me.Enabled,
		"entity_selector": me.EntitySelector,
		"type":            me.Type,
	})
}

func (me *Rule) HandlePreconditions() error {
	if me.EntitySelector == nil && string(me.Type) == "SELECTOR" {
		return fmt.Errorf("'entity_selector' must be specified if 'type' is set to '%v'", me.Type)
	}
	// ---- AttributeRule *ManagementZoneAttributeRule -> {"expectedValue":"ME","property":"type","type":"EQUALS"}
	// ---- DimensionRule *DimensionRule -> {"expectedValue":"DIMENSION","property":"type","type":"EQUALS"}
	return nil
}

func (me *Rule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"attribute_rule":  &me.AttributeRule,
		"dimension_rule":  &me.DimensionRule,
		"enabled":         &me.Enabled,
		"entity_selector": &me.EntitySelector,
		"type":            &me.Type,
	})
}
