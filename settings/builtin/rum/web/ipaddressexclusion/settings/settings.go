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

package ipaddressexclusion

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	IpAddressExclusionInclude bool                    `json:"ipAddressExclusionInclude"` // These are the only IP addresses that should be monitored
	IpExclusionList           IpAddressExclusionRules `json:"ipExclusionList"`           // **Examples:**\n\n   - 84.112.10.5\n   - fe80::10a1:c6b2:5f68:785d
	Scope                     string                  `json:"-" scope:"scope"`           // The scope of this setting (APPLICATION)
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ip_address_exclusion_include": {
			Type:        schema.TypeBool,
			Description: "These are the only IP addresses that should be monitored",
			Required:    true,
		},
		"ip_exclusion_list": {
			Type:        schema.TypeList,
			Description: "**Examples:**\n\n   - 84.112.10.5\n   - fe80::10a1:c6b2:5f68:785d",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(IpAddressExclusionRules).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (APPLICATION)",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"ip_address_exclusion_include": me.IpAddressExclusionInclude,
		"ip_exclusion_list":            me.IpExclusionList,
		"scope":                        me.Scope,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"ip_address_exclusion_include": &me.IpAddressExclusionInclude,
		"ip_exclusion_list":            &me.IpExclusionList,
		"scope":                        &me.Scope,
	})
}
