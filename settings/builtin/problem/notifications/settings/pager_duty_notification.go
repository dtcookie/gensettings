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

type PagerDutyNotification struct {
	Account       string `json:"account"`       // The name of the PagerDuty account.
	ServiceApiKey string `json:"serviceApiKey"` // The API key to access PagerDuty.
	ServiceName   string `json:"serviceName"`   // The name of the service.
}

func (me *PagerDutyNotification) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account": {
			Type:        schema.TypeString,
			Description: "The name of the PagerDuty account.",
			Required:    true,
		},
		"service_api_key": {
			Type:        schema.TypeString,
			Description: "The API key to access PagerDuty.",
			Required:    true,
			Sensitive:   true,
		},
		"service_name": {
			Type:        schema.TypeString,
			Description: "The name of the service.",
			Required:    true,
		},
	}
}

func (me *PagerDutyNotification) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"account":         me.Account,
		"service_api_key": me.ServiceApiKey,
		"service_name":    me.ServiceName,
	})
}

func (me *PagerDutyNotification) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"account":         &me.Account,
		"service_api_key": &me.ServiceApiKey,
		"service_name":    &me.ServiceName,
	})
}
