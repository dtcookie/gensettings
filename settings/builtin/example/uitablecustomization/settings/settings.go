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

package uitablecustomization

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	CollectionProperty []string `json:"collectionProperty,omitempty"` // Collection Property
	ComplexProperty    *Type1   `json:"complexProperty"`              // ComplexType Property
	IntegerProperty    int      `json:"integerProperty"`              // Integer Property
	StringProperty     *string  `json:"stringProperty,omitempty"`     // String Property
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"collection_property": {
			Type:        schema.TypeList,
			Description: "Collection Property",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"complex_property": {
			Type:        schema.TypeList,
			Description: "ComplexType Property",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(Type1).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"integer_property": {
			Type:        schema.TypeInt,
			Description: "Integer Property",
			Required:    true,
		},
		"string_property": {
			Type:        schema.TypeString,
			Description: "String Property",
			Optional:    true, // nullable
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"collection_property": me.CollectionProperty,
		"complex_property":    me.ComplexProperty,
		"integer_property":    me.IntegerProperty,
		"string_property":     me.StringProperty,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"collection_property": &me.CollectionProperty,
		"complex_property":    &me.ComplexProperty,
		"integer_property":    &me.IntegerProperty,
		"string_property":     &me.StringProperty,
	})
}
