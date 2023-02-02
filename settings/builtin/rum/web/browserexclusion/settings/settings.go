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

type Settings struct {
	ApplicationID           string                      `json:"-" scope:"applicationId"` // The scope of this setting (APPLICATION)
	BrowserExclusionInclude bool                        `json:"browserExclusionInclude"` // These are the only browsers that should be monitored
	BrowserExclusionList    BrowserExclusionListObjects `json:"browserExclusionList"`    // Browser exclusion list
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_id": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (APPLICATION)",
			Required:    true,
		},
		"browser_exclusion_include": {
			Type:        schema.TypeBool,
			Description: "These are the only browsers that should be monitored",
			Required:    true,
		},
		"browser_exclusion_list": {
			Type:        schema.TypeList,
			Description: "Browser exclusion list",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(BrowserExclusionListObjects).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"application_id":            me.ApplicationID,
		"browser_exclusion_include": me.BrowserExclusionInclude,
		"browser_exclusion_list":    me.BrowserExclusionList,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"application_id":            &me.ApplicationID,
		"browser_exclusion_include": &me.BrowserExclusionInclude,
		"browser_exclusion_list":    &me.BrowserExclusionList,
	})
}
