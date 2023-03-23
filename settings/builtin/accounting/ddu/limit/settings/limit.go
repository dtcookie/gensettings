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

package limit

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Limit struct {
	LimitEnabled bool       `json:"limitEnabled"`         // Enable limit
	LimitType    *LimitType `json:"limitType,omitempty"`  // Possible Values: `ANNUAL`, `MONTHLY`
	LimitValue   *int       `json:"limitValue,omitempty"` // The upper limit for the defined time frame
}

func (me *Limit) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"limit_enabled": {
			Type:        schema.TypeBool,
			Description: "Enable limit",
			Required:    true,
		},
		"limit_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `ANNUAL`, `MONTHLY`",
			Optional:    true, // precondition
		},
		"limit_value": {
			Type:        schema.TypeInt,
			Description: "The upper limit for the defined time frame",
			Optional:    true, // precondition
		},
	}
}

func (me *Limit) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"limit_enabled": me.LimitEnabled,
		"limit_type":    me.LimitType,
		"limit_value":   me.LimitValue,
	})
}

func (me *Limit) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"limit_enabled": &me.LimitEnabled,
		"limit_type":    &me.LimitType,
		"limit_value":   &me.LimitValue,
	})
}
