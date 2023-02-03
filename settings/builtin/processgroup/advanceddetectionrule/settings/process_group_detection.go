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

package advanceddetectionrule

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ProcessGroupDetection struct {
	ContainedString       string  `json:"containedString"` // (case sensitive)
	Property              string  `json:"property"`
	RestrictToProcessType *string `json:"restrictToProcessType,omitempty"` // Note: Not all types can be detected at startup.
}

func (me *ProcessGroupDetection) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"contained_string": {
			Type:        schema.TypeString,
			Description: "(case sensitive)",
			Required:    true,
		},
		"property": {
			Type:        schema.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"restrict_to_process_type": {
			Type:        schema.TypeString,
			Description: "Note: Not all types can be detected at startup.",
			Optional:    true,
		},
	}
}

func (me *ProcessGroupDetection) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"contained_string":         me.ContainedString,
		"property":                 me.Property,
		"restrict_to_process_type": me.RestrictToProcessType,
	})
}

func (me *ProcessGroupDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"contained_string":         &me.ContainedString,
		"property":                 &me.Property,
		"restrict_to_process_type": &me.RestrictToProcessType,
	})
}
