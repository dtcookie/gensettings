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

package notifications

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type WebHookNotificationHeaders []*WebHookNotificationHeader

func (me *WebHookNotificationHeaders) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"header": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(WebHookNotificationHeader).Schema()},
		},
	}
}

func (me WebHookNotificationHeaders) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("header", me)
}

func (me *WebHookNotificationHeaders) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("header", me)
}

type WebHookNotificationHeader struct {
	Name        string  `json:"name"`                  // The name of the HTTP header.
	Secret      bool    `json:"secret"`                // Secret HTTP header value
	SecretValue *string `json:"secretValue,omitempty"` // The secret value of the HTTP header. May contain an empty value.
	Value       *string `json:"value,omitempty"`       // The value of the HTTP header. May contain an empty value.
}

func (me *WebHookNotificationHeader) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Description: "The name of the HTTP header.",
			Required:    true,
		},
		"secret": {
			Type:        schema.TypeBool,
			Description: "Secret HTTP header value",
			Required:    true,
		},
		"secret_value": {
			Type:        schema.TypeString,
			Description: "The secret value of the HTTP header. May contain an empty value.",
			Optional:    true, // precondition
		},
		"value": {
			Type:        schema.TypeString,
			Description: "The value of the HTTP header. May contain an empty value.",
			Optional:    true, // precondition
		},
	}
}

func (me *WebHookNotificationHeader) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"name":         me.Name,
		"secret":       me.Secret,
		"secret_value": me.SecretValue,
		"value":        me.Value,
	})
}

func (me *WebHookNotificationHeader) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"name":         &me.Name,
		"secret":       &me.Secret,
		"secret_value": &me.SecretValue,
		"value":        &me.Value,
	})
}
