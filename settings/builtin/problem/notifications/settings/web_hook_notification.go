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

type WebHookNotification struct {
	AcceptAnyCertificate     bool                       `json:"acceptAnyCertificate"`        // Accept any SSL certificate (including self-signed and invalid certificates)
	Headers                  WebHookNotificationHeaders `json:"headers,omitempty"`           // A list of the additional HTTP headers.
	NotifyClosedProblems     bool                       `json:"notifyClosedProblems"`        // Call webhook if problem is closed
	NotifyEventMergesEnabled bool                       `json:"notifyEventMergesEnabled"`    // Call webhook if new events merge into existing problems
	OAuth2Credentials        *OAuth2Credentials         `json:"oAuth2Credentials,omitempty"` // To authenticate your integration, the OAuth 2.0 *Client Credentials* Flow (Grant Type) is used. For details see [Client Credentials Flow](https://dt-url.net/ym22wsm)).\n\nThe obtained Access Token is subsequently provided in the *Authorization* header of the request carrying the notification payload.
	Payload                  string                     `json:"payload"`                     // The content of the notification message. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntities}**: Details about the entities impacted by the problem in form of a json array.\n\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem.\n\n**{NamesOfImpactedEntities}**: The names of all entities that are impacted by the problem.\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemDetailsHTML}**: All problem event details including root cause as an HTML-formatted string.\n\n**{ProblemDetailsJSONv2}**: Problem as json object following the structure from the [Dynatrace Problems V2 API](https://dt-url.net/7a03ti2). The optional fields evidenceDetails and impactAnalysis are included, but recentComments is not.\n\n**{ProblemDetailsJSON}**: Problem as json object following the structure from the [Dynatrace Problems V1 API](https://dt-url.net/qn23tk2).\n\n**{ProblemDetailsMarkdown}**: All problem event details including root cause as a Markdown-formatted string.\n\n**{ProblemDetailsText}**: All problem event details including root cause as a text-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.
	Url                      string                     `json:"url"`                         // The URL of the WebHook endpoint.
	UseOAuth2                *bool                      `json:"useOAuth2,omitempty"`         // Use OAuth 2.0 for authentication
}

func (me *WebHookNotification) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"accept_any_certificate": {
			Type:        schema.TypeBool,
			Description: "Accept any SSL certificate (including self-signed and invalid certificates)",
			Required:    true,
		},
		"headers": {
			Type:        schema.TypeList,
			Description: "A list of the additional HTTP headers.",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(WebHookNotificationHeaders).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"notify_closed_problems": {
			Type:        schema.TypeBool,
			Description: "Call webhook if problem is closed",
			Required:    true,
		},
		"notify_event_merges_enabled": {
			Type:        schema.TypeBool,
			Description: "Call webhook if new events merge into existing problems",
			Required:    true,
		},
		"o_auth_2_credentials": {
			Type:        schema.TypeList,
			Description: "To authenticate your integration, the OAuth 2.0 *Client Credentials* Flow (Grant Type) is used. For details see [Client Credentials Flow](https://dt-url.net/ym22wsm)).\n\nThe obtained Access Token is subsequently provided in the *Authorization* header of the request carrying the notification payload.",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(OAuth2Credentials).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"payload": {
			Type:        schema.TypeString,
			Description: "The content of the notification message. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntities}**: Details about the entities impacted by the problem in form of a json array.\n\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem.\n\n**{NamesOfImpactedEntities}**: The names of all entities that are impacted by the problem.\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemDetailsHTML}**: All problem event details including root cause as an HTML-formatted string.\n\n**{ProblemDetailsJSONv2}**: Problem as json object following the structure from the [Dynatrace Problems V2 API](https://dt-url.net/7a03ti2). The optional fields evidenceDetails and impactAnalysis are included, but recentComments is not.\n\n**{ProblemDetailsJSON}**: Problem as json object following the structure from the [Dynatrace Problems V1 API](https://dt-url.net/qn23tk2).\n\n**{ProblemDetailsMarkdown}**: All problem event details including root cause as a Markdown-formatted string.\n\n**{ProblemDetailsText}**: All problem event details including root cause as a text-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.",
			Required:    true,
		},
		"url": {
			Type:        schema.TypeString,
			Description: "The URL of the WebHook endpoint.",
			Required:    true,
		},
		"use_oauth_2": {
			Type:        schema.TypeBool,
			Description: "Use OAuth 2.0 for authentication",
			Optional:    true, // nullable
		},
	}
}

func (me *WebHookNotification) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"accept_any_certificate":      me.AcceptAnyCertificate,
		"headers":                     me.Headers,
		"notify_closed_problems":      me.NotifyClosedProblems,
		"notify_event_merges_enabled": me.NotifyEventMergesEnabled,
		"o_auth_2_credentials":        me.OAuth2Credentials,
		"payload":                     me.Payload,
		"url":                         me.Url,
		"use_oauth_2":                 me.UseOAuth2,
	})
}

func (me *WebHookNotification) HandlePreconditions() {
	// ---- OAuth2Credentials *OAuth2Credentials -> {"expectedValue":true,"property":"useOAuth2","type":"EQUALS"}
}

func (me *WebHookNotification) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"accept_any_certificate":      &me.AcceptAnyCertificate,
		"headers":                     &me.Headers,
		"notify_closed_problems":      &me.NotifyClosedProblems,
		"notify_event_merges_enabled": &me.NotifyEventMergesEnabled,
		"o_auth_2_credentials":        &me.OAuth2Credentials,
		"payload":                     &me.Payload,
		"url":                         &me.Url,
		"use_oauth_2":                 &me.UseOAuth2,
	})
}
