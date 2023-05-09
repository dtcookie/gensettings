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

package uniquepropertiesconstraints

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	FieldA            string         `json:"fieldA"`                      // The combination of Field A and B has to be unique within the settings
	FieldB            *string        `json:"fieldB,omitempty"`            // The combination of Field A and B has to be unique within the settings
	FieldC            *string        `json:"fieldC,omitempty"`            // Can be any value
	ListOfUniqueItems UniqueLstItems `json:"listOfUniqueItems,omitempty"` // A list of unique items
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"field_a": {
			Type:        schema.TypeString,
			Description: "The combination of Field A and B has to be unique within the settings",
			Required:    true,
		},
		"field_b": {
			Type:        schema.TypeString,
			Description: "The combination of Field A and B has to be unique within the settings",
			Optional:    true, // nullable
		},
		"field_c": {
			Type:        schema.TypeString,
			Description: "Can be any value",
			Optional:    true, // nullable
		},
		"list_of_unique_items": {
			Type:        schema.TypeList,
			Description: "A list of unique items",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(UniqueLstItems).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"field_a":              me.FieldA,
		"field_b":              me.FieldB,
		"field_c":              me.FieldC,
		"list_of_unique_items": me.ListOfUniqueItems,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"field_a":              &me.FieldA,
		"field_b":              &me.FieldB,
		"field_c":              &me.FieldC,
		"list_of_unique_items": &me.ListOfUniqueItems,
	})
}
