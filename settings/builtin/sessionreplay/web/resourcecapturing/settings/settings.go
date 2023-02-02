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

package resourcecapturing

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	EnableResourceCapturing                bool     `json:"enableResourceCapturing"`                // When turned on, all CSS resources from all sessions are captured. For details, see [Resource capture](https://dt-url.net/sr-resource-capturing).
	ResourceCaptureUrlExclusionPatternList []string `json:"resourceCaptureUrlExclusionPatternList"` // Add exclusion rules to avoid the capture of resources from certain pages.
	Scope                                  string   `json:"-" scope:"scope"`                        // The scope of this setting (APPLICATION environment)
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enable_resource_capturing": {
			Type:        schema.TypeBool,
			Description: "When turned on, all CSS resources from all sessions are captured. For details, see [Resource capture](https://dt-url.net/sr-resource-capturing).",
			Required:    true,
		},
		"resource_capture_url_exclusion_pattern_list": {
			Type:        schema.TypeSet,
			Description: "Add exclusion rules to avoid the capture of resources from certain pages.",
			Required:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (APPLICATION environment)",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"enable_resource_capturing":                   me.EnableResourceCapturing,
		"resource_capture_url_exclusion_pattern_list": me.ResourceCaptureUrlExclusionPatternList,
		"scope": me.Scope,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enable_resource_capturing":                   &me.EnableResourceCapturing,
		"resource_capture_url_exclusion_pattern_list": &me.ResourceCaptureUrlExclusionPatternList,
		"scope": &me.Scope,
	})
}