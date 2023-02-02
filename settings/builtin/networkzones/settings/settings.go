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

package networkzones

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Enabled bool    `json:"enabled"`         // For details, see [Network zones](https://www.dynatrace.com/support/help/shortlink/network-zones). \n\n⚠ Warning: You may experience network imbalance if you suddenly introduce network zones into an environment that has a high number of OneAgents. To avoid this and to ensure smooth adoption of network zones, follow best practices for planning and proper naming, as explained in [Network zones](https://www.dynatrace.com/support/help/shortlink/network-zones).
	Scope   *string `json:"-" scope:"scope"` // The scope of this setting (environment environment-default)
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Type:        schema.TypeBool,
			Description: "For details, see [Network zones](https://www.dynatrace.com/support/help/shortlink/network-zones). \n\n⚠ Warning: You may experience network imbalance if you suddenly introduce network zones into an environment that has a high number of OneAgents. To avoid this and to ensure smooth adoption of network zones, follow best practices for planning and proper naming, as explained in [Network zones](https://www.dynatrace.com/support/help/shortlink/network-zones).",
			Required:    true,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (environment environment-default)",
			Optional:    true,
			Default:     "environment",
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"enabled": me.Enabled,
		"scope":   me.Scope,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled": &me.Enabled,
		"scope":   &me.Scope,
	})
}
