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
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Rule struct {
	Conditions                 Conditions                  `json:"conditions"`                           // Specify all conditions which should be evaluated for this rule. A rule is applied if all of the listed conditions match.
	DefaultMethodRule          *DefaultMethodRule          `json:"defaultMethodRule"`                    // Choose whether the span service detection rule should detect a service call or not in case there are **no matching** `service method detection rules` but the `service detection rule` **matches** the input.
	Extractions                Extractions                 `json:"extractions,omitempty"`                // Specify a list of extractions that are inherently bound to the rule they are defined in.\n\nEach extraction is composed of **at least one** transformation but it can also be constructed from multiple chained transformations.\n\nIf several transformations are specified, they are handled sequentially from top to bottom. Each transformation is applied to the result of the preceding transformation. For example, the second transformation is applied to the result of the first transformation.
	RuleName                   string                      `json:"ruleName"`                             // Rule name
	ServiceDetection           *ServiceDetection           `json:"serviceDetection"`                     // Service detection
	ServiceMethodDetectionRule ServiceMethodDetectionRules `json:"serviceMethodDetectionRule,omitempty"` // Specify a list of rules which are evaluated in order. When the **first** rule matches, the service detection will proceed. Subsequent rules will not be evaluated.
	Technology                 *string                     `json:"technology,omitempty"`
}

func (me *Rule) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"conditions": {
			Type:        schema.TypeList,
			Description: "Specify all conditions which should be evaluated for this rule. A rule is applied if all of the listed conditions match.",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(Conditions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"default_method_rule": {
			Type:        schema.TypeList,
			Description: "Choose whether the span service detection rule should detect a service call or not in case there are **no matching** `service method detection rules` but the `service detection rule` **matches** the input.",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(DefaultMethodRule).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"extractions": {
			Type:        schema.TypeList,
			Description: "Specify a list of extractions that are inherently bound to the rule they are defined in.\n\nEach extraction is composed of **at least one** transformation but it can also be constructed from multiple chained transformations.\n\nIf several transformations are specified, they are handled sequentially from top to bottom. Each transformation is applied to the result of the preceding transformation. For example, the second transformation is applied to the result of the first transformation.",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(Extractions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"rule_name": {
			Type:        schema.TypeString,
			Description: "Rule name",
			Required:    true,
		},
		"service_detection": {
			Type:        schema.TypeList,
			Description: "Service detection",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(ServiceDetection).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"service_method_detection_rule": {
			Type:        schema.TypeList,
			Description: "Specify a list of rules which are evaluated in order. When the **first** rule matches, the service detection will proceed. Subsequent rules will not be evaluated.",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(ServiceMethodDetectionRules).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"technology": {
			Type:        schema.TypeString,
			Description: "no documentation available",
			Optional:    true, // nullable
		},
	}
}

func (me *Rule) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"conditions":                    me.Conditions,
		"default_method_rule":           me.DefaultMethodRule,
		"extractions":                   me.Extractions,
		"rule_name":                     me.RuleName,
		"service_detection":             me.ServiceDetection,
		"service_method_detection_rule": me.ServiceMethodDetectionRule,
		"technology":                    me.Technology,
	})
}

func (me *Rule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"conditions":                    &me.Conditions,
		"default_method_rule":           &me.DefaultMethodRule,
		"extractions":                   &me.Extractions,
		"rule_name":                     &me.RuleName,
		"service_detection":             &me.ServiceDetection,
		"service_method_detection_rule": &me.ServiceMethodDetectionRule,
		"technology":                    &me.Technology,
	})
}
