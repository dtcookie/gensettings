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

package spancontextpropagation

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SpanMatchers []*SpanMatcher

func (me *SpanMatchers) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"matcher": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(SpanMatcher).Schema()},
		},
	}
}

func (me SpanMatchers) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("matcher", me)
}

func (me *SpanMatchers) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("matcher", me)
}

type SpanMatcher struct {
	CaseSensitive *bool             `json:"caseSensitive,omitempty"` // affects value and key
	Source        SpanMatcherSource `json:"source"`
	SourceKey     *string           `json:"sourceKey,omitempty"`     // Key
	SpanKindValue *SpanKind         `json:"spanKindValue,omitempty"` // Value
	Type          SpanMatcherType   `json:"type"`                    // affects value
	Value         *string           `json:"value,omitempty"`         // evaluated at context injection
}

func (me *SpanMatcher) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"case_sensitive": {
			Type:        schema.TypeBool,
			Description: "affects value and key",
			Optional:    true,
		},
		"source": {
			Type:        schema.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"source_key": {
			Type:        schema.TypeString,
			Description: "Key",
			Optional:    true,
		},
		"span_kind_value": {
			Type:        schema.TypeString,
			Description: "Value",
			Optional:    true,
		},
		"type": {
			Type:        schema.TypeString,
			Description: "affects value",
			Required:    true,
		},
		"value": {
			Type:        schema.TypeString,
			Description: "evaluated at context injection",
			Optional:    true,
		},
	}
}

func (me *SpanMatcher) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"case_sensitive":  me.CaseSensitive,
		"source":          me.Source,
		"source_key":      me.SourceKey,
		"span_kind_value": me.SpanKindValue,
		"type":            me.Type,
		"value":           me.Value,
	})
}

func (me *SpanMatcher) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"case_sensitive":  &me.CaseSensitive,
		"source":          &me.Source,
		"source_key":      &me.SourceKey,
		"span_kind_value": &me.SpanKindValue,
		"type":            &me.Type,
		"value":           &me.Value,
	})
}
