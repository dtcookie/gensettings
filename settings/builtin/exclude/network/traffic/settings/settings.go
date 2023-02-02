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

package traffic

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	ExcludeIp  IpAddressForms `json:"excludeIp"`       // Providing a host IP address, you will exclude network traffic only in calculating connectivity (other metrics will still be calculated).
	ExcludeNic NicForms       `json:"excludeNic"`      // Selecting a network interface, you will exclude all network traffic on that interface from being monitored. You can select from the list below what to not monitor, or input it manually using the \"other one\" option.
	Scope      string         `json:"-" scope:"scope"` // The scope of this setting (HOST)
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"exclude_ip": {
			Type:        schema.TypeList,
			Description: "Providing a host IP address, you will exclude network traffic only in calculating connectivity (other metrics will still be calculated).",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(IpAddressForms).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"exclude_nic": {
			Type:        schema.TypeList,
			Description: "Selecting a network interface, you will exclude all network traffic on that interface from being monitored. You can select from the list below what to not monitor, or input it manually using the \"other one\" option.",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(NicForms).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (HOST)",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"exclude_ip":  me.ExcludeIp,
		"exclude_nic": me.ExcludeNic,
		"scope":       me.Scope,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"exclude_ip":  &me.ExcludeIp,
		"exclude_nic": &me.ExcludeNic,
		"scope":       &me.Scope,
	})
}
