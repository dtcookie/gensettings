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

type HeaderType struct {
	Description *string `json:"description,omitempty"` // The description of the screen
	Icon        *string `json:"icon,omitempty"`        // Icon name
}

func (me *HeaderType) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:        schema.TypeString,
			Description: "The description of the screen",
			Optional:    true, // nullable
		},
		"icon": {
			Type:        schema.TypeString,
			Description: "Icon name",
			Optional:    true, // nullable
		},
	}
}

func (me *HeaderType) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"description": me.Description,
		"icon":        me.Icon,
	})
}

func (me *HeaderType) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"description": &me.Description,
		"icon":        &me.Icon,
	})
}
