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

package spanentrypoints

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SpanEntrypointRule struct {
	Matchers   SpanMatchers         `json:"matchers"`
	RuleAction SpanEntrypointAction `json:"ruleAction"` // Rule action
	RuleName   string               `json:"ruleName"`   // Rule name
}

func (me *SpanEntrypointRule) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"matchers": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(SpanMatchers).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"rule_action": {
			Type:        schema.TypeString,
			Description: "Rule action",
			Required:    true,
		},
		"rule_name": {
			Type:        schema.TypeString,
			Description: "Rule name",
			Required:    true,
		},
	}
}

func (me *SpanEntrypointRule) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"matchers":    me.Matchers,
		"rule_action": me.RuleAction,
		"rule_name":   me.RuleName,
	})
}

func (me *SpanEntrypointRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"matchers":    &me.Matchers,
		"rule_action": &me.RuleAction,
		"rule_name":   &me.RuleName,
	})
}
