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

package ondemandvalidator

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Scope *string `json:"-" scope:"scope"` // The scope of this setting (environment-default). Omit this property if you want to cover the whole environment.
	Start int     `json:"start"`           // A timestamp. It has to be after the current server time when checking the on-demand validator.
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (environment-default). Omit this property if you want to cover the whole environment.",
			Optional:    true,
			Default:     "environment",
		},
		"start": {
			Type:        schema.TypeInt,
			Description: "A timestamp. It has to be after the current server time when checking the on-demand validator.",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"scope": me.Scope,
		"start": me.Start,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"scope": &me.Scope,
		"start": &me.Start,
	})
}
