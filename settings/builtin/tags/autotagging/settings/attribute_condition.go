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

package autotagging

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AttributeConditions []*AttributeCondition

func (me *AttributeConditions) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"condition": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(AttributeCondition).Schema()},
		},
	}
}

func (me AttributeConditions) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("condition", me)
}

func (me *AttributeConditions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("condition", me)
}

type AttributeCondition struct {
	CaseSensitive    *bool     `json:"caseSensitive,omitempty"`    // Case sensitive
	DynamicKey       *string   `json:"dynamicKey,omitempty"`       // Dynamic key
	DynamicKeySource *string   `json:"dynamicKeySource,omitempty"` // Key source
	EntityID         *string   `json:"entityId,omitempty"`         // Value
	EnumValue        *string   `json:"enumValue,omitempty"`        // Value
	IntegerValue     *int      `json:"integerValue,omitempty"`     // Value
	Key              Attribute `json:"key"`                        // Property
	Operator         Operator  `json:"operator"`
	StringValue      *string   `json:"stringValue,omitempty"` // Value
	Tag              *string   `json:"tag,omitempty"`         // Format: `[CONTEXT]tagKey:tagValue`
}

func (me *AttributeCondition) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"case_sensitive": {
			Type:        schema.TypeBool,
			Description: "Case sensitive",
			Optional:    true,
		},
		"dynamic_key": {
			Type:        schema.TypeString,
			Description: "Dynamic key",
			Optional:    true,
		},
		"dynamic_key_source": {
			Type:        schema.TypeString,
			Description: "Key source",
			Optional:    true,
		},
		"entity_id": {
			Type:        schema.TypeString,
			Description: "Value",
			Optional:    true,
		},
		"enum_value": {
			Type:        schema.TypeString,
			Description: "Value",
			Optional:    true,
		},
		"integer_value": {
			Type:        schema.TypeInt,
			Description: "Value",
			Optional:    true,
		},
		"key": {
			Type:        schema.TypeString,
			Description: "Property",
			Required:    true,
		},
		"operator": {
			Type:        schema.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"string_value": {
			Type:        schema.TypeString,
			Description: "Value",
			Optional:    true,
		},
		"tag": {
			Type:        schema.TypeString,
			Description: "Format: `[CONTEXT]tagKey:tagValue`",
			Optional:    true,
		},
	}
}

func (me *AttributeCondition) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"case_sensitive":     me.CaseSensitive,
		"dynamic_key":        me.DynamicKey,
		"dynamic_key_source": me.DynamicKeySource,
		"entity_id":          me.EntityID,
		"enum_value":         me.EnumValue,
		"integer_value":      me.IntegerValue,
		"key":                me.Key,
		"operator":           me.Operator,
		"string_value":       me.StringValue,
		"tag":                me.Tag,
	})
}

func (me *AttributeCondition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"case_sensitive":     &me.CaseSensitive,
		"dynamic_key":        &me.DynamicKey,
		"dynamic_key_source": &me.DynamicKeySource,
		"entity_id":          &me.EntityID,
		"enum_value":         &me.EnumValue,
		"integer_value":      &me.IntegerValue,
		"key":                &me.Key,
		"operator":           &me.Operator,
		"string_value":       &me.StringValue,
		"tag":                &me.Tag,
	})
}
