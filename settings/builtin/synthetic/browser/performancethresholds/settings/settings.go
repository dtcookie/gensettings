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

package performancethresholds

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Enabled    bool             `json:"enabled"`         // Generate a problem and send an alert on performance threshold violations
	Scope      string           `json:"-" scope:"scope"` // The scope of this setting (SYNTHETIC_TEST)
	Thresholds ThresholdEntries `json:"thresholds"`      // Performance thresholds
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Type:        schema.TypeBool,
			Description: "Generate a problem and send an alert on performance threshold violations",
			Required:    true,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (SYNTHETIC_TEST)",
			Required:    true,
		},
		"thresholds": {
			Type:        schema.TypeList,
			Description: "Performance thresholds",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(ThresholdEntries).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"enabled":    me.Enabled,
		"scope":      me.Scope,
		"thresholds": me.Thresholds,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":    &me.Enabled,
		"scope":      &me.Scope,
		"thresholds": &me.Thresholds,
	})
}
