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

package pattern

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type IcuPluralItems []*IcuPluralItem

func (me *IcuPluralItems) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"item": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(IcuPluralItem).Schema()},
		},
	}
}

func (me IcuPluralItems) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("item", me)
}

func (me *IcuPluralItems) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("item", me)
}

type IcuPluralItem struct {
	IntegerProperty int `json:"integerProperty"` // Integer
}

func (me *IcuPluralItem) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"integer_property": {
			Type:        schema.TypeInt,
			Description: "Integer",
			Required:    true,
		},
	}
}

func (me *IcuPluralItem) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"integer_property": me.IntegerProperty,
	})
}

func (me *IcuPluralItem) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"integer_property": &me.IntegerProperty,
	})
}
