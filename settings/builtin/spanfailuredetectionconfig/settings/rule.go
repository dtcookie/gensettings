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

package spanfailuredetectionconfig

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Rule struct {
	Conditions               Conditions               `json:"conditions"`               // Specify all conditions which should be evaluated for this rule. A rule is applied if all of the listed conditions match.
	RuleName                 string                   `json:"ruleName"`                 // Rule name
	SpanVerdictIfRuleMatches SpanVerdictIfRuleMatches `json:"spanVerdictIfRuleMatches"` // Possible Values: `SPAN_VERDICT_FAILURE`, `SPAN_VERDICT_SUCCESS`
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
		"rule_name": {
			Type:        schema.TypeString,
			Description: "Rule name",
			Required:    true,
		},
		"span_verdict_if_rule_matches": {
			Type:        schema.TypeString,
			Description: "Possible Values: `SPAN_VERDICT_FAILURE`, `SPAN_VERDICT_SUCCESS`",
			Required:    true,
		},
	}
}

func (me *Rule) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"conditions":                   me.Conditions,
		"rule_name":                    me.RuleName,
		"span_verdict_if_rule_matches": me.SpanVerdictIfRuleMatches,
	})
}

func (me *Rule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"conditions":                   &me.Conditions,
		"rule_name":                    &me.RuleName,
		"span_verdict_if_rule_matches": &me.SpanVerdictIfRuleMatches,
	})
}
