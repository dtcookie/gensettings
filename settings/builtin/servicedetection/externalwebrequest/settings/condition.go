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

package externalwebrequest

import (
	"fmt"

	"github.com/dtcookie/opt"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"
)

type Conditions []*Condition

func (me *Conditions) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"condition": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(Condition).Schema()},
		},
	}
}

func (me Conditions) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("condition", me)
}

func (me *Conditions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("condition", me)
}

type Condition struct {
	Attribute            string          `json:"attribute"`             // Take the value of this attribute
	CompareOperationType string          `json:"compareOperationType"`  // Apply this operation
	Framework            []FrameworkType `json:"framework,omitempty"`   // Technology
	IgnoreCase           *bool           `json:"ignoreCase,omitempty"`  // Ignore case sensitivity for texts.
	IntValue             *int            `json:"intValue,omitempty"`    // Value
	IntValues            []int           `json:"intValues,omitempty"`   // Values
	IpRangeFrom          *string         `json:"ipRangeFrom,omitempty"` // From
	IpRangeTo            *string         `json:"ipRangeTo,omitempty"`   // To
	TagValues            []string        `json:"tagValues,omitempty"`   // If multiple values are specified, at least one of them must match for the condition to match
	TextValues           []string        `json:"textValues,omitempty"`  // If multiple values are specified, at least one of them must match for the condition to match
}

func (me *Condition) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attribute": {
			Type:        schema.TypeString,
			Description: "Take the value of this attribute",
			Required:    true,
		},
		"compare_operation_type": {
			Type:        schema.TypeString,
			Description: "Apply this operation",
			Required:    true,
		},
		"framework": {
			Type:        schema.TypeSet,
			Description: "Technology",
			Optional:    true, // precondition
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"ignore_case": {
			Type:        schema.TypeBool,
			Description: "Ignore case sensitivity for texts.",
			Optional:    true, // precondition
		},
		"int_value": {
			Type:        schema.TypeInt,
			Description: "Value",
			Optional:    true, // precondition
		},
		"int_values": {
			Type:        schema.TypeSet,
			Description: "Values",
			Optional:    true, // precondition
			Elem:        &schema.Schema{Type: schema.TypeInt},
		},
		"ip_range_from": {
			Type:        schema.TypeString,
			Description: "From",
			Optional:    true, // precondition
		},
		"ip_range_to": {
			Type:        schema.TypeString,
			Description: "To",
			Optional:    true, // precondition
		},
		"tag_values": {
			Type:        schema.TypeSet,
			Description: "If multiple values are specified, at least one of them must match for the condition to match",
			Optional:    true, // precondition
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"text_values": {
			Type:        schema.TypeSet,
			Description: "If multiple values are specified, at least one of them must match for the condition to match",
			Optional:    true, // precondition
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
	}
}

func (me *Condition) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"attribute":              me.Attribute,
		"compare_operation_type": me.CompareOperationType,
		"framework":              me.Framework,
		"ignore_case":            me.IgnoreCase,
		"int_value":              me.IntValue,
		"int_values":             me.IntValues,
		"ip_range_from":          me.IpRangeFrom,
		"ip_range_to":            me.IpRangeTo,
		"tag_values":             me.TagValues,
		"text_values":            me.TextValues,
	})
}

func (me *Condition) HandlePreconditions() error {
	if (me.IgnoreCase == nil) && (slices.Contains([]string{"TagEquals", "TagKeyEquals", "StringEndsWith", "NotStringEndsWith", "StringStartsWith", "NotStringStartsWith", "StringContains", "NotStringContains", "StringEquals", "NotStringEquals"}, string(me.CompareOperationType))) {
		me.IgnoreCase = opt.NewBool(false)
	}
	if (me.IntValue == nil) && (slices.Contains([]string{"IntGreaterThan", "IntLessThan"}, string(me.CompareOperationType))) {
		me.IntValue = opt.NewInt(0)
	}
	if (me.IpRangeFrom == nil) && (slices.Contains([]string{"IpInRange", "NotIpInRange"}, string(me.CompareOperationType))) {
		return fmt.Errorf("'ip_range_from' must be specified if 'compare_operation_type' is set to '%v'", me.CompareOperationType)
	}
	if (me.IpRangeTo == nil) && (slices.Contains([]string{"IpInRange", "NotIpInRange"}, string(me.CompareOperationType))) {
		return fmt.Errorf("'ip_range_to' must be specified if 'compare_operation_type' is set to '%v'", me.CompareOperationType)
	}
	// ---- Framework []FrameworkType -> {"expectedValues":["FrameworkEquals","NotFrameworkEquals"],"property":"compareOperationType","type":"IN"}
	// ---- IntValues []int -> {"expectedValues":["IntEquals","NotIntEquals"],"property":"compareOperationType","type":"IN"}
	// ---- TagValues []string -> {"expectedValue":"TagEquals","property":"compareOperationType","type":"EQUALS"}
	// ---- TextValues []string -> {"expectedValues":["TagKeyEquals","StringEndsWith","NotStringEndsWith","StringStartsWith","NotStringStartsWith","StringContains","NotStringContains","StringEquals","NotStringEquals"],"property":"compareOperationType","type":"IN"}
	return nil
}

func (me *Condition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"attribute":              &me.Attribute,
		"compare_operation_type": &me.CompareOperationType,
		"framework":              &me.Framework,
		"ignore_case":            &me.IgnoreCase,
		"int_value":              &me.IntValue,
		"int_values":             &me.IntValues,
		"ip_range_from":          &me.IpRangeFrom,
		"ip_range_to":            &me.IpRangeTo,
		"tag_values":             &me.TagValues,
		"text_values":            &me.TextValues,
	})
}
