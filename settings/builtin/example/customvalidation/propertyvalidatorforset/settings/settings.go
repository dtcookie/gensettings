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

package propertyvalidatorforset

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	EmailSet []string `json:"emailSet,omitempty"` // A set of email addresses. This list may not contain duplicate entries (case insensitive).
	Scope    *string  `json:"-" scope:"scope"`    // The scope of this setting (environment-default). Omit this property if you want to cover the whole environment.
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"email_set": {
			Type:        schema.TypeSet,
			Description: "A set of email addresses. This list may not contain duplicate entries (case insensitive).",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (environment-default). Omit this property if you want to cover the whole environment.",
			Optional:    true,
			Default:     "environment",
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"email_set": me.EmailSet,
		"scope":     me.Scope,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"email_set": &me.EmailSet,
		"scope":     &me.Scope,
	})
}
