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

type EmailNotification struct {
	BccRecipients        []string `json:"bccRecipients"`        // BCC
	Body                 string   `json:"body"`                 // The template of the email notifications. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntities}**: Details about the entities impacted by the problem in form of a json array.\n\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem.\n\n**{NamesOfImpactedEntities}**: The names of all entities that are impacted by the problem.\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemDetailsHTML}**: All problem event details including root cause as an HTML-formatted string.\n\n**{ProblemDetailsJSONv2}**: Problem as json object following the structure from the [Dynatrace Problems V2 API](https://dt-url.net/7a03ti2). The optional fields evidenceDetails and impactAnalysis are included, but recentComments is not.\n\n**{ProblemDetailsJSON}**: Problem as json object following the structure from the [Dynatrace Problems V1 API](https://dt-url.net/qn23tk2).\n\n**{ProblemDetailsMarkdown}**: All problem event details including root cause as a Markdown-formatted string.\n\n**{ProblemDetailsText}**: All problem event details including root cause as a text-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.
	CcRecipients         []string `json:"ccRecipients"`         // CC
	NotifyClosedProblems bool     `json:"notifyClosedProblems"` // Send email if problem is closed
	Recipients           []string `json:"recipients"`           // To
	Subject              string   `json:"subject"`              // The subject of the email notifications. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem.\n\n**{NamesOfImpactedEntities}**: The names of all entities that are impacted by the problem.\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.
}

func (me *EmailNotification) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"bcc_recipients": {
			Type:        schema.TypeSet,
			Description: "BCC",
			Required:    true,

			Elem: &schema.Schema{Type: schema.TypeString},
		},
		"body": {
			Type:        schema.TypeString,
			Description: "The template of the email notifications. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntities}**: Details about the entities impacted by the problem in form of a json array.\n\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem.\n\n**{NamesOfImpactedEntities}**: The names of all entities that are impacted by the problem.\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemDetailsHTML}**: All problem event details including root cause as an HTML-formatted string.\n\n**{ProblemDetailsJSONv2}**: Problem as json object following the structure from the [Dynatrace Problems V2 API](https://dt-url.net/7a03ti2). The optional fields evidenceDetails and impactAnalysis are included, but recentComments is not.\n\n**{ProblemDetailsJSON}**: Problem as json object following the structure from the [Dynatrace Problems V1 API](https://dt-url.net/qn23tk2).\n\n**{ProblemDetailsMarkdown}**: All problem event details including root cause as a Markdown-formatted string.\n\n**{ProblemDetailsText}**: All problem event details including root cause as a text-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.",
			Required:    true,
		},
		"cc_recipients": {
			Type:        schema.TypeSet,
			Description: "CC",
			Required:    true,

			Elem: &schema.Schema{Type: schema.TypeString},
		},
		"notify_closed_problems": {
			Type:        schema.TypeBool,
			Description: "Send email if problem is closed",
			Required:    true,
		},
		"recipients": {
			Type:        schema.TypeSet,
			Description: "To",
			Required:    true,

			Elem: &schema.Schema{Type: schema.TypeString},
		},
		"subject": {
			Type:        schema.TypeString,
			Description: "The subject of the email notifications. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem.\n\n**{NamesOfImpactedEntities}**: The names of all entities that are impacted by the problem.\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.",
			Required:    true,
		},
	}
}

func (me *EmailNotification) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"bcc_recipients":         me.BccRecipients,
		"body":                   me.Body,
		"cc_recipients":          me.CcRecipients,
		"notify_closed_problems": me.NotifyClosedProblems,
		"recipients":             me.Recipients,
		"subject":                me.Subject,
	})
}

func (me *EmailNotification) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"bcc_recipients":         &me.BccRecipients,
		"body":                   &me.Body,
		"cc_recipients":          &me.CcRecipients,
		"notify_closed_problems": &me.NotifyClosedProblems,
		"recipients":             &me.Recipients,
		"subject":                &me.Subject,
	})
}
