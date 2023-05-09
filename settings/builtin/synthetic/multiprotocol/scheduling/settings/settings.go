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

package scheduling

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Frequency int       `json:"frequency"`           // How often the monitor is executed. Supported values are 1, 2, 5, 10, 15, 30 and 60 minutes
	Locations Locations `json:"locations,omitempty"` // Locations
	Scope     string    `json:"-" scope:"scope"`     // The scope of this setting (MULTIPROTOCOL_MONITOR)
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"frequency": {
			Type:        schema.TypeInt,
			Description: "How often the monitor is executed. Supported values are 1, 2, 5, 10, 15, 30 and 60 minutes",
			Required:    true,
		},
		"locations": {
			Type:        schema.TypeList,
			Description: "Locations",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(Locations).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (MULTIPROTOCOL_MONITOR)",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"frequency": me.Frequency,
		"locations": me.Locations,
		"scope":     me.Scope,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"frequency": &me.Frequency,
		"locations": &me.Locations,
		"scope":     &me.Scope,
	})
}
