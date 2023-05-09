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

package propertyvalidatorforlist

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type NumberRanges []*NumberRange

func (me *NumberRanges) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"number_range": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(NumberRange).Schema()},
		},
	}
}

func (me NumberRanges) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("number_range", me)
}

func (me *NumberRanges) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("number_range", me)
}

// StartEnd. A number range.
type NumberRange struct {
	End   int `json:"end"`   // End of the range.
	Start int `json:"start"` // Start of the range.
}

func (me *NumberRange) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"end": {
			Type:        schema.TypeInt,
			Description: "End of the range.",
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
