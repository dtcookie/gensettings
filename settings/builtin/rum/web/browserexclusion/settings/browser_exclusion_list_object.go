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

package browserexclusion

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type BrowserExclusionListObjects []*BrowserExclusionListObject

func (me *BrowserExclusionListObjects) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"browserExclusion": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(BrowserExclusionListObject).Schema()},
		},
	}
}

func (me BrowserExclusionListObjects) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("browserExclusion", me)
}

func (me *BrowserExclusionListObjects) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("browserExclusion", me)
}

type BrowserExclusionListObject struct {
	BrowserName       Browser            `json:"browserName"`        // Possible Values: `BOTS_AND_SPIDERS`, `ANDROID_WEBKIT`, `CHROME`, `FIREFOX`, `INTERNET_EXPLORER`, `OPERA`, `SAFARI`, `EDGE`
	Platform          *Platform          `json:"platform,omitempty"` // Possible Values: `DESKTOP`, `ALL`, `MOBILE`
	Version           *int               `json:"version,omitempty"`
	VersionComparator *VersionComparator `json:"versionComparator,omitempty"` // Possible Values: `EQUALS`, `LESS_OR_EQUAL`, `GREATER_OR_EQUAL`
}

func (me *BrowserExclusionListObject) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"browser_name": {
			Type:        schema.TypeString,
			Description: "Possible Values: `BOTS_AND_SPIDERS`, `ANDROID_WEBKIT`, `CHROME`, `FIREFOX`, `INTERNET_EXPLORER`, `OPERA`, `SAFARI`, `EDGE`",
			Required:    true,
		},
		"platform": {
			Type:        schema.TypeString,
			Description: "Possible Values: `DESKTOP`, `ALL`, `MOBILE`",
			Optional:    true,
		},
		"version": {
			Type:        schema.TypeInt,
			Description: "no documentation available",
			Optional:    true,
		},
		"version_comparator": {
			Type:        schema.TypeString,
			Description: "Possible Values: `EQUALS`, `LESS_OR_EQUAL`, `GREATER_OR_EQUAL`",
			Optional:    true,
		},
	}
}

func (me *BrowserExclusionListObject) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"browser_name":       me.BrowserName,
		"platform":           me.Platform,
		"version":            me.Version,
		"version_comparator": me.VersionComparator,
	})
}

func (me *BrowserExclusionListObject) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"browser_name":       &me.BrowserName,
		"platform":           &me.Platform,
		"version":            &me.Version,
		"version_comparator": &me.VersionComparator,
	})
}
