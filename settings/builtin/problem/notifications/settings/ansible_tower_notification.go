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

type AnsibleTowerNotification struct {
	AcceptAnyCertificate bool   `json:"acceptAnyCertificate"` // Accept any SSL certificate (including self-signed and invalid certificates)
	CustomMessage        string `json:"customMessage"`        // This message will be displayed in the Extra Variables **Message** field of your job template. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntities}**: Details about the entities impacted by the problem in form of a json array.\n\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem.\n\n**{NamesOfImpactedEntities}**: The names of all entities that are impacted by the problem.\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemDetailsText}**: All problem event details including root cause as a text-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.
	JobTemplateURL       string `json:"jobTemplateURL"`       // The URL of the target job template.\n\nFor example, https://<Ansible server name>/#/templates/job_template/<JobTemplateID>\n\n**Note:** Be sure to select the **Prompt on Launch** option in the Extra Variables section of your job template configuration.
	Password             string `json:"password"`             // Account password.
	Username             string `json:"username"`             // Account username.
}

func (me *AnsibleTowerNotification) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"accept_any_certificate": {
			Type:        schema.TypeBool,
			Description: "Accept any SSL certificate (including self-signed and invalid certificates)",
			Required:    true,
		},
		"custom_message": {
			Type:        schema.TypeString,
			Description: "This message will be displayed in the Extra Variables **Message** field of your job template. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntities}**: Details about the entities impacted by the problem in form of a json array.\n\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem.\n\n**{NamesOfImpactedEntities}**: The names of all entities that are impacted by the problem.\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemDetailsText}**: All problem event details including root cause as a text-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.",
			Required:    true,
		},
		"job_template_url": {
			Type:        schema.TypeString,
			Description: "The URL of the target job template.\n\nFor example, https://<Ansible server name>/#/templates/job_template/<JobTemplateID>\n\n**Note:** Be sure to select the **Prompt on Launch** option in the Extra Variables section of your job template configuration.",
			Required:    true,
		},
		"password": {
			Type:        schema.TypeString,
			Description: "Account password.",
			Required:    true,
		},
		"username": {
			Type:        schema.TypeString,
			Description: "Account username.",
			Required:    true,
		},
	}
}

func (me *AnsibleTowerNotification) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"accept_any_certificate": me.AcceptAnyCertificate,
		"custom_message":         me.CustomMessage,
		"job_template_url":       me.JobTemplateURL,
		"password":               me.Password,
		"username":               me.Username,
	})
}

func (me *AnsibleTowerNotification) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"accept_any_certificate": &me.AcceptAnyCertificate,
		"custom_message":         &me.CustomMessage,
		"job_template_url":       &me.JobTemplateURL,
		"password":               &me.Password,
		"username":               &me.Username,
	})
}
