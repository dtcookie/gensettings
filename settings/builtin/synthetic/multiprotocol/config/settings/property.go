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

package config

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Properties []*Property

func (me *Properties) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"propertie": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(Property).Schema()},
		},
	}
}

func (me Properties) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("propertie", me)
}

func (me *Properties) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("propertie", me)
}

type Property struct {
	Key   string `json:"key"`   // Property key
	Value string `json:"value"` // Property value
}

func (me *Property) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"key": {
			Type:        schema.TypeString,
			Description: "Property key",
			Required:    true,
		},
		"value": {
			Type:        schema.TypeString,
			Description: "Property value",
			Required:    true,
		},
	}
}

func (me *Property) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"key":   me.Key,
		"value": me.Value,
	})
}

func (me *Property) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"key":   &me.Key,
		"value": &me.Value,
	})
}
