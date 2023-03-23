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

package usersessionexportsettingsv2

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Authentication struct {
	Active    bool       `json:"active"`              // Activate
	AuthType  *AuthType  `json:"authType,omitempty"`  // Possible Values: `Basic`, `Oauth2`
	BasicAuth *BasicAuth `json:"basicAuth,omitempty"` // Basic authentication
	OAuth2    *OAuth2    `json:"oAuth2,omitempty"`    // OAuth 2.0 (Early Adopter)
}

func (me *Authentication) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active": {
			Type:        schema.TypeBool,
			Description: "Activate",
			Required:    true,
		},
		"auth_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `Basic`, `Oauth2`",
			Optional:    true, // precondition
		},
		"basic_auth": {
			Type:        schema.TypeList,
			Description: "Basic authentication",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(BasicAuth).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"o_auth_2": {
			Type:        schema.TypeList,
			Description: "OAuth 2.0 (Early Adopter)",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(OAuth2).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Authentication) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"active":     me.Active,
		"auth_type":  me.AuthType,
		"basic_auth": me.BasicAuth,
		"o_auth_2":   me.OAuth2,
	})
}

func (me *Authentication) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"active":     &me.Active,
		"auth_type":  &me.AuthType,
		"basic_auth": &me.BasicAuth,
		"o_auth_2":   &me.OAuth2,
	})
}
