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

package infrastructurevmware

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type EsxiHighCpuDetectionConfig struct {
	CustomThresholds *EsxiHighCpuDetectionThresholds `json:"customThresholds,omitempty"` // Alert if **all three** conditions are met in 3 out of 5 samples
	DetectionMode    *DetectionMode                  `json:"detectionMode,omitempty"`    // Possible Values: `Auto`, `Custom`
	Enabled          bool                            `json:"enabled"`                    // This setting is enabled (`true`) or disabled (`false`)
}

func (me *EsxiHighCpuDetectionConfig) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"custom_thresholds": {
			Type:        schema.TypeList,
			Description: "Alert if **all three** conditions are met in 3 out of 5 samples",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(EsxiHighCpuDetectionThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_mode": {
			Type:        schema.TypeString,
			Description: "Possible Values: `Auto`, `Custom`",
			Optional:    true, // precondition
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "This setting is enabled (`true`) or disabled (`false`)",
			Required:    true,
		},
	}
}

func (me *EsxiHighCpuDetectionConfig) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *EsxiHighCpuDetectionConfig) HandlePreconditions() {
	// ---- CustomThresholds *EsxiHighCpuDetectionThresholds
	// ---- DetectionMode *DetectionMode
}

func (me *EsxiHighCpuDetectionConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
