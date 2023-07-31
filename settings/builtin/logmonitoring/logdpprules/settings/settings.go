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

package logdpprules

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Enabled             bool                 `json:"enabled"`             // This setting is enabled (`true`) or disabled (`false`)
	ProcessorDefinition *ProcessorDefinition `json:"ProcessorDefinition"` // ## Processor definition
	Query               string               `json:"query"`               // Matcher
	RuleName            string               `json:"ruleName"`            // Rule name
	RuleTesting         *RuleTesting         `json:"RuleTesting"`         // ## Rule testing\n### 1. Paste a log / JSON sample
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Type:        schema.TypeBool,
			Description: "This setting is enabled (`true`) or disabled (`false`)",
			Required:    true,
		},
		"processor_definition": {
			Type:        schema.TypeList,
			Description: "## Processor definition",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(ProcessorDefinition).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"query": {
			Type:        schema.TypeString,
			Description: "Matcher",
			Required:    true,
		},
		"rule_name": {
			Type:        schema.TypeString,
			Description: "Rule name",
			Required:    true,
		},
		"rule_testing": {
			Type:        schema.TypeList,
			Description: "## Rule testing\n### 1. Paste a log / JSON sample",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(RuleTesting).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"enabled":              me.Enabled,
		"processor_definition": me.ProcessorDefinition,
		"query":                me.Query,
		"rule_name":            me.RuleName,
		"rule_testing":         me.RuleTesting,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":              &me.Enabled,
		"processor_definition": &me.ProcessorDefinition,
		"query":                &me.Query,
		"rule_name":            &me.RuleName,
		"rule_testing":         &me.RuleTesting,
	})
}
