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

package complexdefaults

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	PropertyTypes            *PropertyTypes `json:"propertyTypes"`                      // All supported property types with their default value
	PropertyTypesDefault     *PropertyTypes `json:"propertyTypesDefault"`               // All supported property types with overridden default
	PropertyTypesDefaultList PropertyTypess `json:"propertyTypesDefaultList,omitempty"` // List of complex type with overridden default
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"property_types": {
			Type:        schema.TypeList,
			Description: "All supported property types with their default value",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(PropertyTypes).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"property_types_default": {
			Type:        schema.TypeList,
			Description: "All supported property types with overridden default",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(PropertyTypes).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"property_types_default_list": {
			Type:        schema.TypeList,
			Description: "List of complex type with overridden default",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(PropertyTypess).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"property_types":              me.PropertyTypes,
		"property_types_default":      me.PropertyTypesDefault,
		"property_types_default_list": me.PropertyTypesDefaultList,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"property_types":              &me.PropertyTypes,
		"property_types_default":      &me.PropertyTypesDefault,
		"property_types_default_list": &me.PropertyTypesDefaultList,
	})
}
