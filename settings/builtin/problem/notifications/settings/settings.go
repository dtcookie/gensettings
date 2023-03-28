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
	AnsibleTowerNotification *AnsibleTowerNotification `json:"ansibleTowerNotification,omitempty"`
	DisplayName              string                    `json:"displayName"` // The name of the notification configuration.
	EmailNotification        *EmailNotification        `json:"emailNotification,omitempty"`
	Enabled                  bool                      `json:"enabled"` // This setting is enabled (`true`) or disabled (`false`)
	JiraNotification         *JiraNotification         `json:"jiraNotification,omitempty"`
	OpsGenieNotification     *OpsGenieNotification     `json:"opsGenieNotification,omitempty"`
	PagerDutyNotification    *PagerDutyNotification    `json:"pagerDutyNotification,omitempty"`
	ServiceNowNotification   *ServiceNowNotification   `json:"serviceNowNotification,omitempty"`
	SlackNotification        *SlackNotification        `json:"slackNotification,omitempty"`
	TrelloNotification       *TrelloNotification       `json:"trelloNotification,omitempty"`
	Type                     NotificationType          `json:"type"` // Possible Values: `ANSIBLETOWER`, `EMAIL`, `JIRA`, `OPS_GENIE`, `PAGER_DUTY`, `SERVICE_NOW`, `SLACK`, `TRELLO`, `VICTOROPS`, `WEBHOOK`, `XMATTERS`
	VictorOpsNotification    *VictorOpsNotification    `json:"victorOpsNotification,omitempty"`
	WebHookNotification      *WebHookNotification      `json:"webHookNotification,omitempty"`
	XMattersNotification     *XMattersNotification     `json:"xMattersNotification,omitempty"`
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
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(AnsibleTowerNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"display_name": {
			Type:        schema.TypeString,
			Description: "The name of the notification configuration.",
			Required:    true,
		},
		"email_notification": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(EmailNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "This setting is enabled (`true`) or disabled (`false`)",
			Required:    true,
		},
		"jira_notification": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(JiraNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"ops_genie_notification": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(OpsGenieNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"pager_duty_notification": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(PagerDutyNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"service_now_notification": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(ServiceNowNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"slack_notification": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(SlackNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"trello_notification": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(TrelloNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `ANSIBLETOWER`, `EMAIL`, `JIRA`, `OPS_GENIE`, `PAGER_DUTY`, `SERVICE_NOW`, `SLACK`, `TRELLO`, `VICTOROPS`, `WEBHOOK`, `XMATTERS`",
			Required:    true,
		},
		"victor_ops_notification": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(VictorOpsNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"web_hook_notification": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(WebHookNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"x_matters_notification": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(XMattersNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
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

func (me *Settings) HandlePreconditions() {
	// ---- AnsibleTowerNotification *AnsibleTowerNotification -> {"expectedValue":"ANSIBLETOWER","property":"type","type":"EQUALS"}
	// ---- EmailNotification *EmailNotification -> {"expectedValue":"EMAIL","property":"type","type":"EQUALS"}
	// ---- JiraNotification *JiraNotification -> {"expectedValue":"JIRA","property":"type","type":"EQUALS"}
	// ---- OpsGenieNotification *OpsGenieNotification -> {"expectedValue":"OPS_GENIE","property":"type","type":"EQUALS"}
	// ---- PagerDutyNotification *PagerDutyNotification -> {"expectedValue":"PAGER_DUTY","property":"type","type":"EQUALS"}
	// ---- ServiceNowNotification *ServiceNowNotification -> {"expectedValue":"SERVICE_NOW","property":"type","type":"EQUALS"}
	// ---- SlackNotification *SlackNotification -> {"expectedValue":"SLACK","property":"type","type":"EQUALS"}
	// ---- TrelloNotification *TrelloNotification -> {"expectedValue":"TRELLO","property":"type","type":"EQUALS"}
	// ---- VictorOpsNotification *VictorOpsNotification -> {"expectedValue":"VICTOROPS","property":"type","type":"EQUALS"}
	// ---- WebHookNotification *WebHookNotification -> {"expectedValue":"WEBHOOK","property":"type","type":"EQUALS"}
	// ---- XMattersNotification *XMattersNotification -> {"expectedValue":"XMATTERS","property":"type","type":"EQUALS"}
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
