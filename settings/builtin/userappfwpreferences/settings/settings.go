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

package userappfwpreferences

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Scope          string `json:"-" scope:"scope"` // The scope of this setting (user, userdefaults)
	StayOnLatestUi bool   `json:"stayOnLatestUi"`  // Stay on the latest Dynatrace
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (user, userdefaults)",
			Required:    true,
		},
		"stay_on_latest_ui": {
			Type:        schema.TypeBool,
			Description: "Stay on the latest Dynatrace",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"scope":             me.Scope,
		"stay_on_latest_ui": me.StayOnLatestUi,
	})
}

func (me *Settings) HandlePreconditions() {
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"scope":             &me.Scope,
		"stay_on_latest_ui": &me.StayOnLatestUi,
	})
}
