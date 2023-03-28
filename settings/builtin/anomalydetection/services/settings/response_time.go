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

package services

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ResponseTime struct {
	AutoDetection  *ResponseTimeAuto  `json:"autoDetection,omitempty"`
	DetectionMode  *DetectionMode     `json:"detectionMode,omitempty"` // Possible Values: `Auto`, `Fixed`
	Enabled        bool               `json:"enabled"`                 // This setting is enabled (`true`) or disabled (`false`)
	FixedDetection *ResponseTimeFixed `json:"fixedDetection,omitempty"`
}

func (me *ResponseTime) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"auto_detection": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(ResponseTimeAuto).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_mode": {
			Type:        schema.TypeString,
			Description: "Possible Values: `Auto`, `Fixed`",
			Optional:    true, // precondition
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "This setting is enabled (`true`) or disabled (`false`)",
			Required:    true,
		},
		"fixed_detection": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(ResponseTimeFixed).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *ResponseTime) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"auto_detection":  me.AutoDetection,
		"detection_mode":  me.DetectionMode,
		"enabled":         me.Enabled,
		"fixed_detection": me.FixedDetection,
	})
}

func (me *ResponseTime) HandlePreconditions() {
	// ---- AutoDetection *ResponseTimeAuto
	// ---- DetectionMode *DetectionMode
	// ---- FixedDetection *ResponseTimeFixed
}

func (me *ResponseTime) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"auto_detection":  &me.AutoDetection,
		"detection_mode":  &me.DetectionMode,
		"enabled":         &me.Enabled,
		"fixed_detection": &me.FixedDetection,
	})
}
