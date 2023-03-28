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

type BasicAuth struct {
	Password string `json:"password"`
	Username string `json:"username"` // User name
}

func (me *BasicAuth) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"password": {
			Type:        schema.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"username": {
			Type:        schema.TypeString,
			Description: "User name",
			Required:    true,
		},
	}
}

func (me *BasicAuth) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"password": me.Password,
		"username": me.Username,
	})
}

func (me *BasicAuth) HandlePreconditions() {
}

func (me *BasicAuth) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"password": &me.Password,
		"username": &me.Username,
	})
}
