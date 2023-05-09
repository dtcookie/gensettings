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

package screensettings

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Layout        *LayoutDefinition `json:"layout"`          // ### Screen layout \n Reference cards that should be shown on the screen (like chart groups or entity lists) or use the auto generated screen layout.
	Scope         string            `json:"-" scope:"scope"` // The scope of this setting (ua-screen)
	StaticContent *StaticContent    `json:"staticContent"`   // ### Static screen content \n Enable/Disable static elements like Properties or Problems.
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"layout": {
			Type:        schema.TypeList,
			Description: "### Screen layout \n Reference cards that should be shown on the screen (like chart groups or entity lists) or use the auto generated screen layout.",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(LayoutDefinition).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (ua-screen)",
			Required:    true,
		},
		"static_content": {
			Type:        schema.TypeList,
			Description: "### Static screen content \n Enable/Disable static elements like Properties or Problems.",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(StaticContent).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"layout":         me.Layout,
		"scope":          me.Scope,
		"static_content": me.StaticContent,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"layout":         &me.Layout,
		"scope":          &me.Scope,
		"static_content": &me.StaticContent,
	})
}
