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

type UniqueLstItems []*UniqueLstItem

func (me *UniqueLstItems) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list_of_unique_item": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(UniqueLstItem).Schema()},
		},
	}
}

func (me UniqueLstItems) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("list_of_unique_item", me)
}

func (me *UniqueLstItems) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("list_of_unique_item", me)
}

type UniqueLstItem struct {
	FieldX string  `json:"fieldX"`           // The combination of Field X and Y has to be unique within the list
	FieldY *string `json:"fieldY,omitempty"` // The combination of Field X and Y has to be unique within the list
	FieldZ *string `json:"fieldZ,omitempty"` // Can be any value
}

func (me *UniqueLstItem) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"field_x": {
			Type:        schema.TypeString,
			Description: "The combination of Field X and Y has to be unique within the list",
			Required:    true,
		},
		"field_y": {
			Type:        schema.TypeString,
			Description: "The combination of Field X and Y has to be unique within the list",
			Optional:    true, // nullable
		},
		"field_z": {
			Type:        schema.TypeString,
			Description: "Can be any value",
			Optional:    true, // nullable
		},
	}
}

func (me *UniqueLstItem) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"field_x": me.FieldX,
		"field_y": me.FieldY,
		"field_z": me.FieldZ,
	})
}

func (me *UniqueLstItem) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"field_x": &me.FieldX,
		"field_y": &me.FieldY,
		"field_z": &me.FieldZ,
	})
}
