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

type OpsGenieNotification struct {
	ApiKey  string `json:"apiKey"`  // The API key to access OpsGenie.\n\nGo to OpsGenie-Integrations and create a new Dynatrace integration. Copy the newly created API key.
	Domain  string `json:"domain"`  // The region domain of the OpsGenie.\n\nFor example, **api.opsgenie.com** for US or **api.eu.opsgenie.com** for EU.
	Message string `json:"message"` // The content of the message. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ImpactedEntityNames}**: The entity impacted by the problem (or multiple impacted entities).
}

func (me *OpsGenieNotification) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"api_key": {
			Type:        schema.TypeString,
			Description: "The API key to access OpsGenie.\n\nGo to OpsGenie-Integrations and create a new Dynatrace integration. Copy the newly created API key.",
			Required:    true,
			Sensitive:   true,
		},
		"domain": {
			Type:        schema.TypeString,
			Description: "The region domain of the OpsGenie.\n\nFor example, **api.opsgenie.com** for US or **api.eu.opsgenie.com** for EU.",
			Required:    true,
		},
		"message": {
			Type:        schema.TypeString,
			Description: "The content of the message. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ImpactedEntityNames}**: The entity impacted by the problem (or multiple impacted entities).",
			Required:    true,
		},
	}
}

func (me *OpsGenieNotification) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"api_key": me.ApiKey,
		"domain":  me.Domain,
		"message": me.Message,
	})
}

func (me *OpsGenieNotification) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"api_key": &me.ApiKey,
		"domain":  &me.Domain,
		"message": &me.Message,
	})
}
