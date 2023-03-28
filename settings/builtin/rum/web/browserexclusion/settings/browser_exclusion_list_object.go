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
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/opt"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type BrowserExclusionListObjects []*BrowserExclusionListObject

func (me *BrowserExclusionListObjects) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"browser_exclusion": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(BrowserExclusionListObject).Schema()},
		},
	}
}

func (me BrowserExclusionListObjects) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("browser_exclusion", me)
}

func (me *BrowserExclusionListObjects) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("browser_exclusion", me)
}

type BrowserExclusionListObject struct {
	BrowserName       Browser            `json:"browserName"`        // Possible Values: `ANDROID_WEBKIT`, `BOTS_AND_SPIDERS`, `CHROME`, `EDGE`, `FIREFOX`, `INTERNET_EXPLORER`, `OPERA`, `SAFARI`
	Platform          *Platform          `json:"platform,omitempty"` // Possible Values: `ALL`, `DESKTOP`, `MOBILE`
	Version           *int               `json:"version,omitempty"`
	VersionComparator *VersionComparator `json:"versionComparator,omitempty"` // Possible Values: `EQUALS`, `GREATER_OR_EQUAL`, `LESS_OR_EQUAL`
}

func (me *BrowserExclusionListObject) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"browser_name": {
			Type:        schema.TypeString,
			Description: "Possible Values: `ANDROID_WEBKIT`, `BOTS_AND_SPIDERS`, `CHROME`, `EDGE`, `FIREFOX`, `INTERNET_EXPLORER`, `OPERA`, `SAFARI`",
			Required:    true,
		},
		"platform": {
			Type:        schema.TypeString,
			Description: "Possible Values: `ALL`, `DESKTOP`, `MOBILE`",
			Optional:    true, // precondition
		},
		"version": {
			Type:        schema.TypeInt,
			Description: "no documentation available",
			Optional:    true, // precondition
		},
		"version_comparator": {
			Type:        schema.TypeString,
			Description: "Possible Values: `EQUALS`, `GREATER_OR_EQUAL`, `LESS_OR_EQUAL`",
			Optional:    true, // precondition
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

func (me *BrowserExclusionListObject) HandlePreconditions() {
	if me.Version == nil && string(me.BrowserName) != "BOTS_AND_SPIDERS" {
		me.Version = opt.NewInt(0)
	}
	// ---- Platform *Platform -> {"expectedValue":"BOTS_AND_SPIDERS","property":"browserName","type":"EQUALS"}
	// ---- VersionComparator *VersionComparator -> {"expectedValue":"BOTS_AND_SPIDERS","property":"browserName","type":"EQUALS"}
}

func (me *BrowserExclusionListObject) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"browser_name":       &me.BrowserName,
		"platform":           &me.Platform,
		"version":            &me.Version,
		"version_comparator": &me.VersionComparator,
	})
}
