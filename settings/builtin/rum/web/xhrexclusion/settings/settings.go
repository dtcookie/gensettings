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

package xhrexclusion

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Scope            string `json:"-" scope:"scope"`  // The scope of this setting (APPLICATION)
	XhrExclusionRule string `json:"xhrExclusionRule"` // **Examples:**\n\n   - \\\\/segment1\\\\/segment2\n   - dynatrace\\\\.com\n   - www\\\\.dynatrace\\\\.com\\\\/segment1\\\\/.*[a-zA-Z]\n   - www\\\\.dynatrace\\\\.com:8080\n   - www\\\\.dynatrace\\\\.com:([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])\n   - www\\\\.dynatrace\\\\.com\\\\?param1=value1&param2=.*[a-zA-Z]
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (APPLICATION)",
			Required:    true,
		},
		"xhr_exclusion_rule": {
			Type:        schema.TypeString,
			Description: "**Examples:**\n\n   - \\\\/segment1\\\\/segment2\n   - dynatrace\\\\.com\n   - www\\\\.dynatrace\\\\.com\\\\/segment1\\\\/.*[a-zA-Z]\n   - www\\\\.dynatrace\\\\.com:8080\n   - www\\\\.dynatrace\\\\.com:([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])\n   - www\\\\.dynatrace\\\\.com\\\\?param1=value1&param2=.*[a-zA-Z]",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"scope":              me.Scope,
		"xhr_exclusion_rule": me.XhrExclusionRule,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"scope":              &me.Scope,
		"xhr_exclusion_rule": &me.XhrExclusionRule,
	})
}
