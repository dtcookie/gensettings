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
	BucketName string `json:"bucketName"` // Events will be stored in the selected bucket. Analyze bucket contents in the [log & event viewer.](/ui/logs-events?advancedQueryMode=true&query=fetch+bizevents)
	Enabled    bool   `json:"enabled"`    // Enabled
	Matcher    string `json:"matcher"`    // [See our documentation](https://dt-url.net/bp234rv)
	RuleName   string `json:"ruleName"`   // Rule name
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"bucket_name": {
			Type:        schema.TypeString,
			Description: "Events will be stored in the selected bucket. Analyze bucket contents in the [log & event viewer.](/ui/logs-events?advancedQueryMode=true&query=fetch+bizevents)",
			Required:    true,
		},
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
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"bucket_name": me.BucketName,
		"enabled":     me.Enabled,
		"matcher":     me.Matcher,
		"rule_name":   me.RuleName,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"bucket_name": &me.BucketName,
		"enabled":     &me.Enabled,
		"matcher":     &me.Matcher,
		"rule_name":   &me.RuleName,
	})
}
