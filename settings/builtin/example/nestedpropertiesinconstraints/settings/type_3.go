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

// Inner type 3. Inner type with a property referenced in the uniqueness collection constraint
type Type3 struct {
	InnerTextValue string `json:"innerTextValue"` // Inner text value
}

func (me *Type3) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"inner_text_value": {
			Type:        schema.TypeString,
			Description: "Inner text value",
			Required:    true,
		},
	}
}

func (me *Type3) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"inner_text_value": me.InnerTextValue,
	})
}

func (me *Type3) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"inner_text_value": &me.InnerTextValue,
	})
}
