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

type Settings struct {
	AlertingProfile          string                    `json:"alertingProfile"` // Select an [alerting profile](/ui/settings/builtin:alerting.profile) to control the delivery of problem notifications related to this integration.
	AnsibleTowerNotification *AnsibleTowerNotification `json:"ansibleTowerNotification"`
	DisplayName              string                    `json:"displayName"` // The name of the notification configuration.
	EmailNotification        *EmailNotification        `json:"emailNotification"`
	Enabled                  bool                      `json:"enabled"` // Enabled
	JiraNotification         *JiraNotification         `json:"jiraNotification"`
	OpsGenieNotification     *OpsGenieNotification     `json:"opsGenieNotification"`
	PagerDutyNotification    *PagerDutyNotification    `json:"pagerDutyNotification"`
	ServiceNowNotification   *ServiceNowNotification   `json:"serviceNowNotification"`
	SlackNotification        *SlackNotification        `json:"slackNotification"`
	TrelloNotification       *TrelloNotification       `json:"trelloNotification"`
	Type                     NotificationType          `json:"type"` // Possible Values: `TRELLO`, `ANSIBLETOWER`, `VICTOROPS`, `PAGER_DUTY`, `XMATTERS`, `JIRA`, `WEBHOOK`, `EMAIL`, `SERVICE_NOW`, `SLACK`, `OPS_GENIE`
	VictorOpsNotification    *VictorOpsNotification    `json:"victorOpsNotification"`
	WebHookNotification      *WebHookNotification      `json:"webHookNotification"`
	XMattersNotification     *XMattersNotification     `json:"xMattersNotification"`
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"alerting_profile": {
			Type:        schema.TypeString,
			Description: "Select an [alerting profile](/ui/settings/builtin:alerting.profile) to control the delivery of problem notifications related to this integration.",
			Required:    true,
		},
		"ansible_tower_notification": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(AnsibleTowerNotification).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"display_name": {
			Type:        schema.TypeString,
			Description: "The name of the notification configuration.",
			Required:    true,
		},
		"email_notification": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(EmailNotification).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "Enabled",
			Required:    true,
		},
		"jira_notification": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(JiraNotification).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"ops_genie_notification": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(OpsGenieNotification).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"pager_duty_notification": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(PagerDutyNotification).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"service_now_notification": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(ServiceNowNotification).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"slack_notification": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(SlackNotification).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"trello_notification": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(TrelloNotification).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `TRELLO`, `ANSIBLETOWER`, `VICTOROPS`, `PAGER_DUTY`, `XMATTERS`, `JIRA`, `WEBHOOK`, `EMAIL`, `SERVICE_NOW`, `SLACK`, `OPS_GENIE`",
			Required:    true,
		},
		"victor_ops_notification": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(VictorOpsNotification).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"web_hook_notification": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(WebHookNotification).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"x_matters_notification": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(XMattersNotification).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"alerting_profile":           me.AlertingProfile,
		"ansible_tower_notification": me.AnsibleTowerNotification,
		"display_name":               me.DisplayName,
		"email_notification":         me.EmailNotification,
		"enabled":                    me.Enabled,
		"jira_notification":          me.JiraNotification,
		"ops_genie_notification":     me.OpsGenieNotification,
		"pager_duty_notification":    me.PagerDutyNotification,
		"service_now_notification":   me.ServiceNowNotification,
		"slack_notification":         me.SlackNotification,
		"trello_notification":        me.TrelloNotification,
		"type":                       me.Type,
		"victor_ops_notification":    me.VictorOpsNotification,
		"web_hook_notification":      me.WebHookNotification,
		"x_matters_notification":     me.XMattersNotification,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"alerting_profile":           &me.AlertingProfile,
		"ansible_tower_notification": &me.AnsibleTowerNotification,
		"display_name":               &me.DisplayName,
		"email_notification":         &me.EmailNotification,
		"enabled":                    &me.Enabled,
		"jira_notification":          &me.JiraNotification,
		"ops_genie_notification":     &me.OpsGenieNotification,
		"pager_duty_notification":    &me.PagerDutyNotification,
		"service_now_notification":   &me.ServiceNowNotification,
		"slack_notification":         &me.SlackNotification,
		"trello_notification":        &me.TrelloNotification,
		"type":                       &me.Type,
		"victor_ops_notification":    &me.VictorOpsNotification,
		"web_hook_notification":      &me.WebHookNotification,
		"x_matters_notification":     &me.XMattersNotification,
	})
}
