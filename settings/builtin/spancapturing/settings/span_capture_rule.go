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

package spancapturing

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SpanCaptureRule struct {
	Matchers   SpanMatchers      `json:"matchers"`
	RuleAction SpanCaptureAction `json:"ruleAction"` // Possible Values: `CAPTURE`, `IGNORE`
	RuleName   string            `json:"ruleName"`   // Rule name
}

func (me *SpanCaptureRule) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"matchers": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(SpanMatchers).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"rule_action": {
			Type:        schema.TypeString,
			Description: "Possible Values: `CAPTURE`, `IGNORE`",
			Required:    true,
		},
		"rule_name": {
			Type:        schema.TypeString,
			Description: "Rule name",
			Required:    true,
		},
	}
}

func (me *SpanCaptureRule) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"matchers":    me.Matchers,
		"rule_action": me.RuleAction,
		"rule_name":   me.RuleName,
	})
}

func (me *SpanCaptureRule) HandlePreconditions() {
}

func (me *SpanCaptureRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"matchers":    &me.Matchers,
		"rule_action": &me.RuleAction,
		"rule_name":   &me.RuleName,
	})
}
