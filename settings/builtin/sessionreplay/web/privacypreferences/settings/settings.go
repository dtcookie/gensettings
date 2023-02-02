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

package privacypreferences

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	EnableOptInMode         bool                 `json:"enableOptInMode"`         // When [Session Replay opt-in mode](https://dt-url.net/sr-opt-in-mode) is turned on, Session Replay is deactivated until explicitly activated via an API call.
	MaskingPresets          *MaskingPresetConfig `json:"maskingPresets"`          // To protect your end users' privacy, select or customize [predefined masking options](https://dt-url.net/sr-masking-preset-options) that suit your content recording and playback requirements.
	Scope                   string               `json:"-" scope:"scope"`         // The scope of this setting (APPLICATION environment)
	UrlExclusionPatternList []string             `json:"urlExclusionPatternList"` // Exclude webpages or views from Session Replay recording by adding [URL exclusion rules](https://dt-url.net/sr-url-exclusion)
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enable_opt_in_mode": {
			Type:        schema.TypeBool,
			Description: "When [Session Replay opt-in mode](https://dt-url.net/sr-opt-in-mode) is turned on, Session Replay is deactivated until explicitly activated via an API call.",
			Required:    true,
		},
		"masking_presets": {
			Type:        schema.TypeList,
			Description: "To protect your end users' privacy, select or customize [predefined masking options](https://dt-url.net/sr-masking-preset-options) that suit your content recording and playback requirements.",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(MaskingPresetConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (APPLICATION environment)",
			Required:    true,
		},
		"url_exclusion_pattern_list": {
			Type:        schema.TypeSet,
			Description: "Exclude webpages or views from Session Replay recording by adding [URL exclusion rules](https://dt-url.net/sr-url-exclusion)",
			Required:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"enable_opt_in_mode":         me.EnableOptInMode,
		"masking_presets":            me.MaskingPresets,
		"scope":                      me.Scope,
		"url_exclusion_pattern_list": me.UrlExclusionPatternList,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enable_opt_in_mode":         &me.EnableOptInMode,
		"masking_presets":            &me.MaskingPresets,
		"scope":                      &me.Scope,
		"url_exclusion_pattern_list": &me.UrlExclusionPatternList,
	})
}
