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

package properties

import (
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type PropertyDefinitions []*PropertyDefinition

func (me *PropertyDefinitions) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"propertie": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(PropertyDefinition).Schema()},
		},
	}
}

func (me PropertyDefinitions) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("propertie", me)
}

func (me *PropertyDefinitions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("propertie", me)
}

type PropertyDefinition struct {
	Attribute         *AttributeType `json:"attribute,omitempty"`         // Attribute options
	Conditions        []string       `json:"conditions,omitempty"`        // All conditions from the list need to be fulfilled for the property to be visible
	ImportantProperty *bool          `json:"importantProperty,omitempty"` // If true, this property will be marked as important in the UI. Those will be displayed in the header below the entity name.
	Relation          *RelationType  `json:"relation,omitempty"`          // Relation options
	Type              PropertyType   `json:"type"`                        // Possible Values: `ATTRIBUTE`, `RELATION`
}

func (me *PropertyDefinition) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attribute": {
			Type:        schema.TypeList,
			Description: "Attribute options",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(AttributeType).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"conditions": {
			Type:        schema.TypeSet,
			Description: "All conditions from the list need to be fulfilled for the property to be visible",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"important_property": {
			Type:        schema.TypeBool,
			Description: "If true, this property will be marked as important in the UI. Those will be displayed in the header below the entity name.",
			Optional:    true, // nullable
		},
		"relation": {
			Type:        schema.TypeList,
			Description: "Relation options",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(RelationType).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `ATTRIBUTE`, `RELATION`",
			Required:    true,
		},
	}
}

func (me *PropertyDefinition) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"attribute":          me.Attribute,
		"conditions":         me.Conditions,
		"important_property": me.ImportantProperty,
		"relation":           me.Relation,
		"type":               me.Type,
	})
}

func (me *PropertyDefinition) HandlePreconditions() error {
	if me.Attribute == nil && (string(me.Type) == "ATTRIBUTE") {
		return fmt.Errorf("'attribute' must be specified if 'type' is set to '%v'", me.Type)
	}
	if me.Attribute != nil && (string(me.Type) != "ATTRIBUTE") {
		return fmt.Errorf("'attribute' must not be specified if 'type' is set to '%v'", me.Type)
	}
	if me.Relation == nil && (string(me.Type) == "RELATION") {
		return fmt.Errorf("'relation' must be specified if 'type' is set to '%v'", me.Type)
	}
	if me.Relation != nil && (string(me.Type) != "RELATION") {
		return fmt.Errorf("'relation' must not be specified if 'type' is set to '%v'", me.Type)
	}
	return nil
}

func (me *PropertyDefinition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"attribute":          &me.Attribute,
		"conditions":         &me.Conditions,
		"important_property": &me.ImportantProperty,
		"relation":           &me.Relation,
		"type":               &me.Type,
	})
}
