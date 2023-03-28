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

type VictorOpsNotification struct {
	ApiKey     string `json:"apiKey"`     // The API key for the target VictorOps account.\n\nReceive your VictorOps API key by navigating to: Settings -> Integrations -> Rest Endpoint -> Dynatrace within your VictorOps account.
	Message    string `json:"message"`    // The content of the message. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem.\n\n**{NamesOfImpactedEntities}**: The names of all entities that are impacted by the problem.\n\n**{ProblemDetailsText}**: All problem event details including root cause as a text-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.
	RoutingKey string `json:"routingKey"` // The routing key, defining the group to be notified.
}

func (me *VictorOpsNotification) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"api_key": {
			Type:        schema.TypeString,
			Description: "The API key for the target VictorOps account.\n\nReceive your VictorOps API key by navigating to: Settings -> Integrations -> Rest Endpoint -> Dynatrace within your VictorOps account.",
			Required:    true,
		},
		"message": {
			Type:        schema.TypeString,
			Description: "The content of the message. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem.\n\n**{NamesOfImpactedEntities}**: The names of all entities that are impacted by the problem.\n\n**{ProblemDetailsText}**: All problem event details including root cause as a text-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.",
			Required:    true,
		},
		"routing_key": {
			Type:        schema.TypeString,
			Description: "The routing key, defining the group to be notified.",
			Required:    true,
		},
	}
}

func (me *VictorOpsNotification) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"api_key":     me.ApiKey,
		"message":     me.Message,
		"routing_key": me.RoutingKey,
	})
}

func (me *VictorOpsNotification) HandlePreconditions() {
}

func (me *VictorOpsNotification) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"api_key":     &me.ApiKey,
		"message":     &me.Message,
		"routing_key": &me.RoutingKey,
	})
}
