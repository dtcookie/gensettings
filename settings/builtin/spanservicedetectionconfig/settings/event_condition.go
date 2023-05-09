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

package spanservicedetectionconfig

import (
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/opt"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"
)

type EventConditions []*EventCondition

func (me *EventConditions) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"event_condition": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(EventCondition).Schema()},
		},
	}
}

func (me EventConditions) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("event_condition", me)
}

func (me *EventConditions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("event_condition", me)
}

type EventCondition struct {
	AttrKeyBoolComparisonType   *AttrKeyBooleanComparisonType `json:"attrKeyBoolComparisonType,omitempty"`   // Possible Values: `DOES_NOT_EQUAL`, `DOES_NOT_EXIST`, `EQUALS`, `EXISTS`
	AttrKeyStringComparisonType *AttrKeyStringComparisonType  `json:"attrKeyStringComparisonType,omitempty"` // Possible Values: `CONTAINS`, `DOES_NOT_CONTAIN`, `DOES_NOT_END_WITH`, `DOES_NOT_EQUAL`, `DOES_NOT_EXIST`, `DOES_NOT_START_WITH`, `ENDS_WITH`, `EQUALS`, `EXISTS`, `STARTS_WITH`
	BoolValue                   *BooleanValue                 `json:"boolValue,omitempty"`                   // Possible Values: `FALSE`, `TRUE`
	CaseSensitive               *bool                         `json:"caseSensitive,omitempty"`               // Affects value(s)
	DataType                    *ComparisonDataType           `json:"dataType,omitempty"`                    // Possible Values: `BOOLEAN`, `FLOAT`, `INTEGER`, `STRING`
	FloatValue                  *float64                      `json:"floatValue,omitempty"`                  // Value
	FloatValues                 []float64                     `json:"floatValues,omitempty"`                 // If regular comparison type i.e. `Equals` is used, then **any** of the listed values should match for the condition to be evaluated to true\n\nIf negated comparison type i.e. `Does not equal` is used, then **none** of the listed values should match for the condition to be evaluated to true
	IntValue                    *int                          `json:"intValue,omitempty"`                    // Value
	IntValues                   []int                         `json:"intValues,omitempty"`                   // If regular comparison type i.e. `Equals` is used, then **any** of the listed values should match for the condition to be evaluated to true\n\nIf negated comparison type i.e. `Does not equal` is used, then **none** of the listed values should match for the condition to be evaluated to true
	NumComparisonType           *NumberComparisonType         `json:"numComparisonType,omitempty"`           // Possible Values: `DOES_NOT_EQUAL`, `DOES_NOT_EXIST`, `EQUALS`, `EXISTS`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LOWER_THAN`, `LOWER_THAN_OR_EQUAL`, `NOT_GREATER_THAN`, `NOT_GREATER_THAN_OR_EQUAL`, `NOT_LOWER_THAN`, `NOT_LOWER_THAN_OR_EQUAL`
	Source                      SpanEventMatcherSource        `json:"source"`                                // Possible Values: `SPAN_EVENT_ATTRIBUTE`, `SPAN_EVENT_NAME`
	SourceKey                   *string                       `json:"sourceKey,omitempty"`                   // Key
	StrValues                   []string                      `json:"strValues,omitempty"`                   // If regular comparison type i.e. `Equals` is used, then **any** of the listed values should match for the condition to be evaluated to true\n\nIf negated comparison type i.e. `Does not equal` is used, then **none** of the listed values should match for the condition to be evaluated to true
	StringComparisonType        *StringComparisonType         `json:"stringComparisonType,omitempty"`        // Possible Values: `CONTAINS`, `DOES_NOT_CONTAIN`, `DOES_NOT_END_WITH`, `DOES_NOT_EQUAL`, `DOES_NOT_START_WITH`, `ENDS_WITH`, `EQUALS`, `STARTS_WITH`
}

func (me *EventCondition) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attr_key_bool_comparison_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `DOES_NOT_EQUAL`, `DOES_NOT_EXIST`, `EQUALS`, `EXISTS`",
			Optional:    true, // precondition
		},
		"attr_key_string_comparison_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `CONTAINS`, `DOES_NOT_CONTAIN`, `DOES_NOT_END_WITH`, `DOES_NOT_EQUAL`, `DOES_NOT_EXIST`, `DOES_NOT_START_WITH`, `ENDS_WITH`, `EQUALS`, `EXISTS`, `STARTS_WITH`",
			Optional:    true, // precondition
		},
		"bool_value": {
			Type:        schema.TypeString,
			Description: "Possible Values: `FALSE`, `TRUE`",
			Optional:    true, // precondition
		},
		"case_sensitive": {
			Type:        schema.TypeBool,
			Description: "Affects value(s)",
			Optional:    true, // precondition
		},
		"data_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `BOOLEAN`, `FLOAT`, `INTEGER`, `STRING`",
			Optional:    true, // precondition
		},
		"float_value": {
			Type:        schema.TypeFloat,
			Description: "Value",
			Optional:    true, // precondition
		},
		"float_values": {
			Type:        schema.TypeSet,
			Description: "If regular comparison type i.e. `Equals` is used, then **any** of the listed values should match for the condition to be evaluated to true\n\nIf negated comparison type i.e. `Does not equal` is used, then **none** of the listed values should match for the condition to be evaluated to true",
			Optional:    true, // precondition
			Elem:        &schema.Schema{Type: schema.TypeFloat},
		},
		"int_value": {
			Type:        schema.TypeInt,
			Description: "Value",
			Optional:    true, // precondition
		},
		"int_values": {
			Type:        schema.TypeSet,
			Description: "If regular comparison type i.e. `Equals` is used, then **any** of the listed values should match for the condition to be evaluated to true\n\nIf negated comparison type i.e. `Does not equal` is used, then **none** of the listed values should match for the condition to be evaluated to true",
			Optional:    true, // precondition
			Elem:        &schema.Schema{Type: schema.TypeInt},
		},
		"num_comparison_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `DOES_NOT_EQUAL`, `DOES_NOT_EXIST`, `EQUALS`, `EXISTS`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LOWER_THAN`, `LOWER_THAN_OR_EQUAL`, `NOT_GREATER_THAN`, `NOT_GREATER_THAN_OR_EQUAL`, `NOT_LOWER_THAN`, `NOT_LOWER_THAN_OR_EQUAL`",
			Optional:    true, // precondition
		},
		"source": {
			Type:        schema.TypeString,
			Description: "Possible Values: `SPAN_EVENT_ATTRIBUTE`, `SPAN_EVENT_NAME`",
			Required:    true,
		},
		"source_key": {
			Type:        schema.TypeString,
			Description: "Key",
			Optional:    true, // precondition
		},
		"str_values": {
			Type:        schema.TypeSet,
			Description: "If regular comparison type i.e. `Equals` is used, then **any** of the listed values should match for the condition to be evaluated to true\n\nIf negated comparison type i.e. `Does not equal` is used, then **none** of the listed values should match for the condition to be evaluated to true",
			Optional:    true, // precondition
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"string_comparison_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `CONTAINS`, `DOES_NOT_CONTAIN`, `DOES_NOT_END_WITH`, `DOES_NOT_EQUAL`, `DOES_NOT_START_WITH`, `ENDS_WITH`, `EQUALS`, `STARTS_WITH`",
			Optional:    true, // precondition
		},
	}
}

func (me *EventCondition) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"attr_key_bool_comparison_type":   me.AttrKeyBoolComparisonType,
		"attr_key_string_comparison_type": me.AttrKeyStringComparisonType,
		"bool_value":                      me.BoolValue,
		"case_sensitive":                  me.CaseSensitive,
		"data_type":                       me.DataType,
		"float_value":                     me.FloatValue,
		"float_values":                    me.FloatValues,
		"int_value":                       me.IntValue,
		"int_values":                      me.IntValues,
		"num_comparison_type":             me.NumComparisonType,
		"source":                          me.Source,
		"source_key":                      me.SourceKey,
		"str_values":                      me.StrValues,
		"string_comparison_type":          me.StringComparisonType,
	})
}

func (me *EventCondition) HandlePreconditions() error {
	if me.FloatValue == nil && (me.DataType != nil && string(*me.DataType) == "FLOAT") && me.NumComparisonType != nil && !slices.Contains([]string{"EQUALS", "DOES_NOT_EQUAL", "EXISTS", "DOES_NOT_EXIST"}, string(*me.NumComparisonType)) {
		me.FloatValue = opt.NewFloat64(0.0)
	}
	if me.IntValue == nil && (me.DataType != nil && string(*me.DataType) == "INTEGER") && me.NumComparisonType != nil && !slices.Contains([]string{"EQUALS", "DOES_NOT_EQUAL", "EXISTS", "DOES_NOT_EXIST"}, string(*me.NumComparisonType)) {
		me.IntValue = opt.NewInt(0)
	}
	if me.AttrKeyBoolComparisonType == nil && (string(me.Source) == "SPAN_EVENT_ATTRIBUTE") && (me.DataType != nil && string(*me.DataType) == "BOOLEAN") {
		return fmt.Errorf("'attr_key_bool_comparison_type' must be specified if 'source' is set to '%v' and 'data_type' is set to '%v'", me.Source, me.DataType)
	}
	if me.AttrKeyBoolComparisonType != nil && (string(me.Source) != "SPAN_EVENT_ATTRIBUTE") || me.DataType == nil || (me.DataType != nil && string(*me.DataType) != "BOOLEAN") {
		return fmt.Errorf("'attr_key_bool_comparison_type' must not be specified if 'source' is set to '%v'", me.Source)
	}
	if me.AttrKeyStringComparisonType == nil && (string(me.Source) == "SPAN_EVENT_ATTRIBUTE") && (me.DataType != nil && string(*me.DataType) == "STRING") {
		return fmt.Errorf("'attr_key_string_comparison_type' must be specified if 'source' is set to '%v' and 'data_type' is set to '%v'", me.Source, me.DataType)
	}
	if me.AttrKeyStringComparisonType != nil && (string(me.Source) != "SPAN_EVENT_ATTRIBUTE") || me.DataType == nil || (me.DataType != nil && string(*me.DataType) != "STRING") {
		return fmt.Errorf("'attr_key_string_comparison_type' must not be specified if 'source' is set to '%v'", me.Source)
	}
	if me.BoolValue == nil && (me.DataType != nil && string(*me.DataType) == "BOOLEAN") && me.AttrKeyBoolComparisonType != nil && slices.Contains([]string{"EQUALS", "DOES_NOT_EQUAL"}, string(*me.AttrKeyBoolComparisonType)) {
		return fmt.Errorf("'bool_value' must be specified if 'data_type' is set to '%v' and 'attr_key_bool_comparison_type' is set to '%v'", me.DataType, me.AttrKeyBoolComparisonType)
	}
	if me.BoolValue != nil && me.DataType == nil || (me.DataType != nil && string(*me.DataType) != "BOOLEAN") || me.AttrKeyBoolComparisonType != nil && !slices.Contains([]string{"EQUALS", "DOES_NOT_EQUAL"}, string(*me.AttrKeyBoolComparisonType)) {
		return fmt.Errorf("'bool_value' must not be specified if 'data_type' is set to '%v'", me.DataType)
	}
	if me.DataType == nil && (string(me.Source) == "SPAN_EVENT_ATTRIBUTE") {
		return fmt.Errorf("'data_type' must be specified if 'source' is set to '%v'", me.Source)
	}
	if me.DataType != nil && (string(me.Source) != "SPAN_EVENT_ATTRIBUTE") {
		return fmt.Errorf("'data_type' must not be specified if 'source' is set to '%v'", me.Source)
	}
	if me.NumComparisonType == nil && me.DataType != nil && slices.Contains([]string{"INTEGER", "FLOAT"}, string(*me.DataType)) {
		return fmt.Errorf("'num_comparison_type' must be specified if 'data_type' is set to '%v'", me.DataType)
	}
	if me.NumComparisonType != nil && me.DataType != nil && !slices.Contains([]string{"INTEGER", "FLOAT"}, string(*me.DataType)) {
		return fmt.Errorf("'num_comparison_type' must not be specified if 'data_type' is set to '%v'", me.DataType)
	}
	if me.SourceKey == nil && (string(me.Source) == "SPAN_EVENT_ATTRIBUTE") {
		return fmt.Errorf("'source_key' must be specified if 'source' is set to '%v'", me.Source)
	}
	if me.StringComparisonType == nil && slices.Contains([]string{"SPAN_EVENT_NAME"}, string(me.Source)) {
		return fmt.Errorf("'string_comparison_type' must be specified if 'source' is set to '%v'", me.Source)
	}
	if me.StringComparisonType != nil && !slices.Contains([]string{"SPAN_EVENT_NAME"}, string(me.Source)) {
		return fmt.Errorf("'string_comparison_type' must not be specified if 'source' is set to '%v'", me.Source)
	}
	// ---- CaseSensitive *bool -> {"preconditions":[{"precondition":{"property":"stringComparisonType","type":"NULL"},"type":"NOT"},{"preconditions":[{"expectedValue":"SPAN_EVENT_ATTRIBUTE","property":"source","type":"EQUALS"},{"expectedValue":"STRING","property":"dataType","type":"EQUALS"},{"precondition":{"expectedValues":["EXISTS","DOES_NOT_EXIST"],"property":"attrKeyStringComparisonType","type":"IN"},"type":"NOT"}],"type":"AND"}],"type":"OR"}
	// ---- FloatValues []float64 -> {"preconditions":[{"expectedValue":"FLOAT","property":"dataType","type":"EQUALS"},{"expectedValues":["EQUALS","DOES_NOT_EQUAL"],"property":"numComparisonType","type":"IN"}],"type":"AND"}
	// ---- IntValues []int -> {"preconditions":[{"expectedValue":"INTEGER","property":"dataType","type":"EQUALS"},{"expectedValues":["EQUALS","DOES_NOT_EQUAL"],"property":"numComparisonType","type":"IN"}],"type":"AND"}
	// ---- StrValues []string -> {"preconditions":[{"precondition":{"expectedValues":["SPAN_EVENT_ATTRIBUTE"],"property":"source","type":"IN"},"type":"NOT"},{"preconditions":[{"precondition":{"expectedValues":["EXISTS","DOES_NOT_EXIST"],"property":"attrKeyStringComparisonType","type":"IN"},"type":"NOT"},{"expectedValue":"STRING","property":"dataType","type":"EQUALS"}],"type":"AND"}],"type":"OR"}
	return nil
}

func (me *EventCondition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"attr_key_bool_comparison_type":   &me.AttrKeyBoolComparisonType,
		"attr_key_string_comparison_type": &me.AttrKeyStringComparisonType,
		"bool_value":                      &me.BoolValue,
		"case_sensitive":                  &me.CaseSensitive,
		"data_type":                       &me.DataType,
		"float_value":                     &me.FloatValue,
		"float_values":                    &me.FloatValues,
		"int_value":                       &me.IntValue,
		"int_values":                      &me.IntValues,
		"num_comparison_type":             &me.NumComparisonType,
		"source":                          &me.Source,
		"source_key":                      &me.SourceKey,
		"str_values":                      &me.StrValues,
		"string_comparison_type":          &me.StringComparisonType,
	})
}
