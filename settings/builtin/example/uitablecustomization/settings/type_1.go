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

type Type1 struct {
	InnerStringProperty string `json:"innerStringProperty"` // Nested String Property
}

func (me *Type1) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"inner_string_property": {
			Type:        schema.TypeString,
			Description: "Nested String Property",
			Required:    true,
		},
	}
}

func (me *Type1) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"inner_string_property": me.InnerStringProperty,
	})
}

func (me *Type1) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"inner_string_property": &me.InnerStringProperty,
	})
}
