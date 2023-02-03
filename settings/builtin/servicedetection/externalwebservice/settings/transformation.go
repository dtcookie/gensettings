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

package externalwebservice

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Transformations []*Transformation

func (me *Transformations) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"transformation": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(Transformation).Schema()},
		},
	}
}

func (me Transformations) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("transformation", me)
}

func (me *Transformations) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("transformation", me)
}

type Transformation struct {
	IncludeHexNumbers  *bool              `json:"includeHexNumbers,omitempty"` // include hexadecimal numbers
	MinDigitCount      *int               `json:"minDigitCount,omitempty"`     // min digit count
	Prefix             *string            `json:"prefix,omitempty"`
	ReplacementValue   *string            `json:"replacementValue,omitempty"` // replacement
	SegmentCount       *int               `json:"segmentCount,omitempty"`     // How many segments should be taken.
	SelectIndex        *int               `json:"selectIndex,omitempty"`      // select index
	SplitDelimiter     *string            `json:"splitDelimiter,omitempty"`   // split by
	Suffix             *string            `json:"suffix,omitempty"`
	TakeFromEnd        *bool              `json:"takeFromEnd,omitempty"` // take from end
	TransformationType TransformationType `json:"transformationType"`    // Possible Values: `REMOVE_IBANS`, `BETWEEN`, `REMOVE_NUMBERS`, `REMOVE_CREDIT_CARDS`, `TAKE_SEGMENTS`, `BEFORE`, `REMOVE_IPS`, `SPLIT_SELECT`, `AFTER`, `REPLACE_BETWEEN`
}

func (me *Transformation) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"include_hex_numbers": {
			Type:        schema.TypeBool,
			Description: "include hexadecimal numbers",
			Optional:    true,
		},
		"min_digit_count": {
			Type:        schema.TypeInt,
			Description: "min digit count",
			Optional:    true,
		},
		"prefix": {
			Type:        schema.TypeString,
			Description: "no documentation available",
			Optional:    true,
		},
		"replacement_value": {
			Type:        schema.TypeString,
			Description: "replacement",
			Optional:    true,
		},
		"segment_count": {
			Type:        schema.TypeInt,
			Description: "How many segments should be taken.",
			Optional:    true,
		},
		"select_index": {
			Type:        schema.TypeInt,
			Description: "select index",
			Optional:    true,
		},
		"split_delimiter": {
			Type:        schema.TypeString,
			Description: "split by",
			Optional:    true,
		},
		"suffix": {
			Type:        schema.TypeString,
			Description: "no documentation available",
			Optional:    true,
		},
		"take_from_end": {
			Type:        schema.TypeBool,
			Description: "take from end",
			Optional:    true,
		},
		"transformation_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `REMOVE_IBANS`, `BETWEEN`, `REMOVE_NUMBERS`, `REMOVE_CREDIT_CARDS`, `TAKE_SEGMENTS`, `BEFORE`, `REMOVE_IPS`, `SPLIT_SELECT`, `AFTER`, `REPLACE_BETWEEN`",
			Required:    true,
		},
	}
}

func (me *Transformation) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"include_hex_numbers": me.IncludeHexNumbers,
		"min_digit_count":     me.MinDigitCount,
		"prefix":              me.Prefix,
		"replacement_value":   me.ReplacementValue,
		"segment_count":       me.SegmentCount,
		"select_index":        me.SelectIndex,
		"split_delimiter":     me.SplitDelimiter,
		"suffix":              me.Suffix,
		"take_from_end":       me.TakeFromEnd,
		"transformation_type": me.TransformationType,
	})
}

func (me *Transformation) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"include_hex_numbers": &me.IncludeHexNumbers,
		"min_digit_count":     &me.MinDigitCount,
		"prefix":              &me.Prefix,
		"replacement_value":   &me.ReplacementValue,
		"segment_count":       &me.SegmentCount,
		"select_index":        &me.SelectIndex,
		"split_delimiter":     &me.SplitDelimiter,
		"suffix":              &me.Suffix,
		"take_from_end":       &me.TakeFromEnd,
		"transformation_type": &me.TransformationType,
	})
}
