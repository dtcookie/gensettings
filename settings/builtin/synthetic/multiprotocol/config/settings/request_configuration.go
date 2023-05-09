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

package config

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type RequestConfigurations []*RequestConfiguration

func (me *RequestConfigurations) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"request_configuration": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(RequestConfiguration).Schema()},
		},
	}
}

func (me RequestConfigurations) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("request_configuration", me)
}

func (me *RequestConfigurations) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("request_configuration", me)
}

type RequestConfiguration struct {
	Constraints    Constraints `json:"constraints,omitempty"`    // List of request constraints
	PatternMatcher *string     `json:"patternMatcher,omitempty"` // Pattern matcher
	Properties     Properties  `json:"properties,omitempty"`
}

func (me *RequestConfiguration) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"constraints": {
			Type:        schema.TypeList,
			Description: "List of request constraints",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(Constraints).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"pattern_matcher": {
			Type:        schema.TypeString,
			Description: "Pattern matcher",
			Optional:    true, // nullable
		},
		"properties": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(Properties).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *RequestConfiguration) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"constraints":     me.Constraints,
		"pattern_matcher": me.PatternMatcher,
		"properties":      me.Properties,
	})
}

func (me *RequestConfiguration) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"constraints":     &me.Constraints,
		"pattern_matcher": &me.PatternMatcher,
		"properties":      &me.Properties,
	})
}
