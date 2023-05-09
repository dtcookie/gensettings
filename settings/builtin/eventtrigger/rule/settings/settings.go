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
	Enabled    bool       `json:"enabled"`             // This setting is enabled (`true`) or disabled (`false`)
	FilterLQL  *string    `json:"filterLQL,omitempty"` // LQL query to filter incoming data to which this rule should apply - DEPRECATED
	Matcher    *string    `json:"matcher,omitempty"`   // DQL query to filter incoming data to which this rule should apply.
	RecordType RecordType `json:"recordType"`          // Possible Values: `BIZ_EVENTS`, `EVENTS`
	RuleName   string     `json:"ruleName"`            // Rule name
	TriggerURL string     `json:"triggerURL"`          // Relative to the Platform Base URL
	UserID     *string    `json:"userId,omitempty"`    // Will be used to verify trigger permissions.
	UserName   *string    `json:"userName,omitempty"`  // Will be used to verify trigger permissions.
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Type:        schema.TypeBool,
			Description: "This setting is enabled (`true`) or disabled (`false`)",
			Required:    true,
		},
		"filter_lql": {
			Type:        schema.TypeString,
			Description: "LQL query to filter incoming data to which this rule should apply - DEPRECATED",
			Optional:    true, // nullable
		},
		"matcher": {
			Type:        schema.TypeString,
			Description: "DQL query to filter incoming data to which this rule should apply.",
			Optional:    true, // nullable
		},
		"record_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `BIZ_EVENTS`, `EVENTS`",
			Required:    true,
		},
		"rule_name": {
			Type:        schema.TypeString,
			Description: "Rule name",
			Required:    true,
		},
		"trigger_url": {
			Type:        schema.TypeString,
			Description: "Relative to the Platform Base URL",
			Required:    true,
		},
		"user_id": {
			Type:        schema.TypeString,
			Description: "Will be used to verify trigger permissions.",
			Optional:    true, // nullable
		},
		"user_name": {
			Type:        schema.TypeString,
			Description: "Will be used to verify trigger permissions.",
			Optional:    true, // nullable
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"enabled":     me.Enabled,
		"filter_lql":  me.FilterLQL,
		"matcher":     me.Matcher,
		"record_type": me.RecordType,
		"rule_name":   me.RuleName,
		"trigger_url": me.TriggerURL,
		"user_id":     me.UserID,
		"user_name":   me.UserName,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":     &me.Enabled,
		"filter_lql":  &me.FilterLQL,
		"matcher":     &me.Matcher,
		"record_type": &me.RecordType,
		"rule_name":   &me.RuleName,
		"trigger_url": &me.TriggerURL,
		"user_id":     &me.UserID,
		"user_name":   &me.UserName,
	})
}
