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

type ServiceNowNotification struct {
	InstanceName  *string `json:"instanceName,omitempty"` // The ServiceNow instance identifier. It refers to the first part of your own ServiceNow URL. \n\n This field is mutually exclusive with the **url** field. You can only use one of them.
	Message       string  `json:"message"`                // The content of the ServiceNow description. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem.\n\n**{NamesOfImpactedEntities}**: The names of all entities that are impacted by the problem.\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemDetailsHTML}**: All problem event details including root cause as an HTML-formatted string.\n\n**{ProblemDetailsText}**: All problem event details including root cause as a text-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.
	Password      string  `json:"password"`               // The password to the ServiceNow account.
	SendEvents    bool    `json:"sendEvents"`             // Send events into ServiceNow ITOM.
	SendIncidents bool    `json:"sendIncidents"`          // Send incidents into ServiceNow ITSM.
	Url           *string `json:"url,omitempty"`          // The URL of the on-premise ServiceNow installation. \n\n This field is mutually exclusive with the **instanceName** field. You can only use one of them.
	Username      string  `json:"username"`               // The username of the ServiceNow account. \n\n Make sure that your user account has the `web_service_admin` and `x_dynat_ruxit.Integration` roles.
}

func (me *ServiceNowNotification) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"instance_name": {
			Type:        schema.TypeString,
			Description: "The ServiceNow instance identifier. It refers to the first part of your own ServiceNow URL. \n\n This field is mutually exclusive with the **url** field. You can only use one of them.",
			Optional:    true, // precondition
		},
		"message": {
			Type:        schema.TypeString,
			Description: "The content of the ServiceNow description. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem.\n\n**{NamesOfImpactedEntities}**: The names of all entities that are impacted by the problem.\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemDetailsHTML}**: All problem event details including root cause as an HTML-formatted string.\n\n**{ProblemDetailsText}**: All problem event details including root cause as a text-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.",
			Required:    true,
		},
		"password": {
			Type:        schema.TypeString,
			Description: "The password to the ServiceNow account.",
			Required:    true,
		},
		"send_events": {
			Type:        schema.TypeBool,
			Description: "Send events into ServiceNow ITOM.",
			Required:    true,
		},
		"send_incidents": {
			Type:        schema.TypeBool,
			Description: "Send incidents into ServiceNow ITSM.",
			Required:    true,
		},
		"url": {
			Type:        schema.TypeString,
			Description: "The URL of the on-premise ServiceNow installation. \n\n This field is mutually exclusive with the **instanceName** field. You can only use one of them.",
			Optional:    true, // nullable
		},
		"username": {
			Type:        schema.TypeString,
			Description: "The username of the ServiceNow account. \n\n Make sure that your user account has the `web_service_admin` and `x_dynat_ruxit.Integration` roles.",
			Required:    true,
		},
	}
}

func (me *ServiceNowNotification) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"instance_name":  me.InstanceName,
		"message":        me.Message,
		"password":       me.Password,
		"send_events":    me.SendEvents,
		"send_incidents": me.SendIncidents,
		"url":            me.Url,
		"username":       me.Username,
	})
}

func (me *ServiceNowNotification) HandlePreconditions() error {
	// ---- InstanceName *string -> {"preconditions":[{"property":"url","type":"NULL"},{"expectedValue":"","property":"url","type":"EQUALS"}],"type":"OR"}
	return nil
}

func (me *ServiceNowNotification) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"instance_name":  &me.InstanceName,
		"message":        &me.Message,
		"password":       &me.Password,
		"send_events":    &me.SendEvents,
		"send_incidents": &me.SendIncidents,
		"url":            &me.Url,
		"username":       &me.Username,
	})
}
