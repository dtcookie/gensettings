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

package nestedpropertiesinconstraints

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Inner type 1. Inner type with a property referenced in the uniqueness schema constraint
type Type1 struct {
	InnerUniqueName string `json:"innerUniqueName"` // Another name, unique on the schema level in a combination with `Schema unique name`
}

func (me *Type1) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"inner_unique_name": {
			Type:        schema.TypeString,
			Description: "Another name, unique on the schema level in a combination with `Schema unique name`",
			Required:    true,
		},
	}
}

func (me *Type1) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"inner_unique_name": me.InnerUniqueName,
	})
}

func (me *Type1) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"inner_unique_name": &me.InnerUniqueName,
	})
}
