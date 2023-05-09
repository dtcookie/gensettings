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

package containervalidator

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// StartEnd. A number range. Start must be before End.
type NumberRange struct {
	End   int `json:"end"`   // End of the range. Must be after Start.
	Start int `json:"start"` // Start of the range.
}

func (me *NumberRange) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"end": {
			Type:        schema.TypeInt,
			Description: "End of the range. Must be after Start.",
			Required:    true,
		},
		"start": {
			Type:        schema.TypeInt,
			Description: "Start of the range.",
			Required:    true,
		},
	}
}

func (me *NumberRange) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"end":   me.End,
		"start": me.Start,
	})
}

func (me *NumberRange) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"end":   &me.End,
		"start": &me.Start,
	})
}
