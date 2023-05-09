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

package attributetable

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ColumnDefinitions []*ColumnDefinition

func (me *ColumnDefinitions) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"column": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(ColumnDefinition).Schema()},
		},
	}
}

func (me ColumnDefinitions) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("column", me)
}

func (me *ColumnDefinitions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("column", me)
}

type ColumnDefinition struct {
	Attribute  *AttributeType `json:"attribute"`            // Attribute options
	Conditions []string       `json:"conditions,omitempty"` // All conditions from the list need to be fulfilled for the column to be visible
}

func (me *ColumnDefinition) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attribute": {
			Type:        schema.TypeList,
			Description: "Attribute options",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(AttributeType).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"conditions": {
			Type:        schema.TypeSet,
			Description: "All conditions from the list need to be fulfilled for the column to be visible",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
	}
}

func (me *ColumnDefinition) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"attribute":  me.Attribute,
		"conditions": me.Conditions,
	})
}

func (me *ColumnDefinition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"attribute":  &me.Attribute,
		"conditions": &me.Conditions,
	})
}
