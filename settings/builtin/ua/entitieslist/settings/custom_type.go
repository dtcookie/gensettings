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

package entitieslist

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// CustomType. Allows to specify custom column
type CustomType struct {
	Data        CustomColumnDatas `json:"data,omitempty"` // Additional options
	DisplayName string            `json:"displayName"`    // Display name
	Key         string            `json:"key"`            // Use to reference desired custom column
}

func (me *CustomType) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"data": {
			Type:        schema.TypeList,
			Description: "Additional options",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(CustomColumnDatas).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"display_name": {
			Type:        schema.TypeString,
			Description: "Display name",
			Required:    true,
		},
		"key": {
			Type:        schema.TypeString,
			Description: "Use to reference desired custom column",
			Required:    true,
		},
	}
}

func (me *CustomType) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"data":         me.Data,
		"display_name": me.DisplayName,
		"key":          me.Key,
	})
}

func (me *CustomType) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"data":         &me.Data,
		"display_name": &me.DisplayName,
		"key":          &me.Key,
	})
}
