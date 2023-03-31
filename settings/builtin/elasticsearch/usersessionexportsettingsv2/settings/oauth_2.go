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

type OAuth2 struct {
	AccessTokenUrl  string          `json:"accessTokenUrl"`  // Access token URL
	ClientID        string          `json:"clientId"`        // Client ID
	ClientSecret    string          `json:"clientSecret"`    // Client secret
	GrantType       GrantType       `json:"grantType"`       // Possible Values: `ClientCredentials`
	Scope           *string         `json:"scope,omitempty"` // The scope of access you are requesting
	SendCredentials SendCredentials `json:"sendCredentials"` // Possible Values: `Header`
}

func (me *OAuth2) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access_token_url": {
			Type:        schema.TypeString,
			Description: "Access token URL",
			Required:    true,
		},
		"client_id": {
			Type:        schema.TypeString,
			Description: "Client ID",
			Required:    true,
		},
		"client_secret": {
			Type:        schema.TypeString,
			Description: "Client secret",
			Required:    true,
			Sensitive:   true,
		},
		"grant_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `ClientCredentials`",
			Required:    true,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of access you are requesting",
			Optional:    true, // nullable
		},
		"send_credentials": {
			Type:        schema.TypeString,
			Description: "Possible Values: `Header`",
			Required:    true,
		},
	}
}

func (me *OAuth2) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"access_token_url": me.AccessTokenUrl,
		"client_id":        me.ClientID,
		"client_secret":    me.ClientSecret,
		"grant_type":       me.GrantType,
		"scope":            me.Scope,
		"send_credentials": me.SendCredentials,
	})
}

func (me *OAuth2) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"access_token_url": &me.AccessTokenUrl,
		"client_id":        &me.ClientID,
		"client_secret":    &me.ClientSecret,
		"grant_type":       &me.GrantType,
		"scope":            &me.Scope,
		"send_credentials": &me.SendCredentials,
	})
}
