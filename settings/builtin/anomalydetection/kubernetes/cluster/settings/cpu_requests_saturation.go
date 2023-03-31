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

package cluster

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type CpuRequestsSaturation struct {
	Configuration *CpuRequestsSaturationConfig `json:"configuration,omitempty"` // Alert if
	Enabled       bool                         `json:"enabled"`                 // This setting is enabled (`true`) or disabled (`false`)
}

func (me *CpuRequestsSaturation) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configuration": {
			Type:        schema.TypeList,
			Description: "Alert if",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(CpuRequestsSaturationConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "This setting is enabled (`true`) or disabled (`false`)",
			Required:    true,
		},
	}
}

func (me *CpuRequestsSaturation) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"configuration": me.Configuration,
		"enabled":       me.Enabled,
	})
}

func (me *CpuRequestsSaturation) HandlePreconditions() error {
	// ---- Configuration *CpuRequestsSaturationConfig -> {"expectedValue":true,"property":"enabled","type":"EQUALS"}
	return nil
}

func (me *CpuRequestsSaturation) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"configuration": &me.Configuration,
		"enabled":       &me.Enabled,
	})
}
