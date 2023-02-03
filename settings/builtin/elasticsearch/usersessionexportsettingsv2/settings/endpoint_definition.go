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

type EndpointDefinition struct {
	ContentType             string  `json:"contentType"`             // Content type
	EnableUserSessionExport bool    `json:"enableUserSessionExport"` // Enable user session export
	EndpointUrl             *string `json:"endpointUrl,omitempty"`   // Endpoint URL
	UsePost                 bool    `json:"usePost"`                 // Use POST method (instead of PUT)
}

func (me *EndpointDefinition) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"content_type": {
			Type:        schema.TypeString,
			Description: "Content type",
			Required:    true,
		},
		"enable_user_session_export": {
			Type:        schema.TypeBool,
			Description: "Enable user session export",
			Required:    true,
		},
		"endpoint_url": {
			Type:        schema.TypeString,
			Description: "Endpoint URL",
			Optional:    true,
		},
		"use_post": {
			Type:        schema.TypeBool,
			Description: "Use POST method (instead of PUT)",
			Required:    true,
		},
	}
}

func (me *EndpointDefinition) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"content_type":               me.ContentType,
		"enable_user_session_export": me.EnableUserSessionExport,
		"endpoint_url":               me.EndpointUrl,
		"use_post":                   me.UsePost,
	})
}

func (me *EndpointDefinition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"content_type":               &me.ContentType,
		"enable_user_session_export": &me.EnableUserSessionExport,
		"endpoint_url":               &me.EndpointUrl,
		"use_post":                   &me.UsePost,
	})
}
