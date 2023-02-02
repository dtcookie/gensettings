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

package rule

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Enabled              bool                 `json:"enabled"`              // Enabled
	Matcher              string               `json:"matcher"`              // [See our documentation](https://dt-url.net/bp234rv)
	RuleName             string               `json:"ruleName"`             // Rule name
	RuleTesting          *RuleTesting         `json:"RuleTesting"`          // ## Rule testing\n### 1. Paste an event sample
	Script               string               `json:"script"`               // [See our documentation](https://dt-url.net/8k03xm2)
	TransformationFields TransformationFields `json:"transformationFields"` // Transformation fields
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Type:        schema.TypeBool,
			Description: "Enabled",
			Required:    true,
		},
		"matcher": {
			Type:        schema.TypeString,
			Description: "[See our documentation](https://dt-url.net/bp234rv)",
			Required:    true,
		},
		"rule_name": {
			Type:        schema.TypeString,
			Description: "Rule name",
			Required:    true,
		},
		"rule_testing": {
			Type:        schema.TypeList,
			Description: "## Rule testing\n### 1. Paste an event sample",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(RuleTesting).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"script": {
			Type:        schema.TypeString,
			Description: "[See our documentation](https://dt-url.net/8k03xm2)",
			Required:    true,
		},
		"transformation_fields": {
			Type:        schema.TypeList,
			Description: "Transformation fields",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(TransformationFields).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"enabled":               me.Enabled,
		"matcher":               me.Matcher,
		"rule_name":             me.RuleName,
		"rule_testing":          me.RuleTesting,
		"script":                me.Script,
		"transformation_fields": me.TransformationFields,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":               &me.Enabled,
		"matcher":               &me.Matcher,
		"rule_name":             &me.RuleName,
		"rule_testing":          &me.RuleTesting,
		"script":                &me.Script,
		"transformation_fields": &me.TransformationFields,
	})
}
