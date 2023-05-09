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

package functioncall

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/opt"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	App           string  `json:"app"`                  // App name
	Function      string  `json:"function"`             // Function
	OauthToken    *string `json:"oauthToken,omitempty"` // Function call won't work with the stored token - has to be entered again each time
	OverrideToken bool    `json:"overrideToken"`        // Use custom OAuth token instead of the logged in user for function call
	PostRequest   bool    `json:"postRequest"`          // Use POST request with setting value as body
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app": {
			Type:        schema.TypeString,
			Description: "App name",
			Required:    true,
		},
		"function": {
			Type:        schema.TypeString,
			Description: "Function",
			Required:    true,
		},
		"oauth_token": {
			Type:        schema.TypeString,
			Description: "Function call won't work with the stored token - has to be entered again each time",
			Optional:    true, // precondition
			Sensitive:   true,
		},
		"override_token": {
			Type:        schema.TypeBool,
			Description: "Use custom OAuth token instead of the logged in user for function call",
			Required:    true,
		},
		"post_request": {
			Type:        schema.TypeBool,
			Description: "Use POST request with setting value as body",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"app":            me.App,
		"function":       me.Function,
		"oauth_token":    me.OauthToken,
		"override_token": me.OverrideToken,
		"post_request":   me.PostRequest,
	})
}

func (me *Settings) HandlePreconditions() error {
	if me.OauthToken == nil && me.OverrideToken {
		me.OauthToken = opt.NewString("")
	}
	return nil
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"app":            &me.App,
		"function":       &me.Function,
		"oauth_token":    &me.OauthToken,
		"override_token": &me.OverrideToken,
		"post_request":   &me.PostRequest,
	})
}
