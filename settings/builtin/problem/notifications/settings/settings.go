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
	"fmt"

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

func (me *Settings) HandlePreconditions() error {
	if (me.AnsibleTowerNotification == nil) && (string(me.Type) == "ANSIBLETOWER") {
		return fmt.Errorf("'ansible_tower_notification' must be specified if 'type' is set to '%v'", me.Type)
	}
	if (me.AnsibleTowerNotification != nil) && (string(me.Type) != "ANSIBLETOWER") {
		return fmt.Errorf("'ansible_tower_notification' must not be specified if 'type' is set to '%v'", me.Type)
	}
	if (me.EmailNotification == nil) && (string(me.Type) == "EMAIL") {
		return fmt.Errorf("'email_notification' must be specified if 'type' is set to '%v'", me.Type)
	}
	if (me.EmailNotification != nil) && (string(me.Type) != "EMAIL") {
		return fmt.Errorf("'email_notification' must not be specified if 'type' is set to '%v'", me.Type)
	}
	if (me.JiraNotification == nil) && (string(me.Type) == "JIRA") {
		return fmt.Errorf("'jira_notification' must be specified if 'type' is set to '%v'", me.Type)
	}
	if (me.JiraNotification != nil) && (string(me.Type) != "JIRA") {
		return fmt.Errorf("'jira_notification' must not be specified if 'type' is set to '%v'", me.Type)
	}
	if (me.OpsGenieNotification == nil) && (string(me.Type) == "OPS_GENIE") {
		return fmt.Errorf("'ops_genie_notification' must be specified if 'type' is set to '%v'", me.Type)
	}
	if (me.OpsGenieNotification != nil) && (string(me.Type) != "OPS_GENIE") {
		return fmt.Errorf("'ops_genie_notification' must not be specified if 'type' is set to '%v'", me.Type)
	}
	if (me.PagerDutyNotification == nil) && (string(me.Type) == "PAGER_DUTY") {
		return fmt.Errorf("'pager_duty_notification' must be specified if 'type' is set to '%v'", me.Type)
	}
	if (me.PagerDutyNotification != nil) && (string(me.Type) != "PAGER_DUTY") {
		return fmt.Errorf("'pager_duty_notification' must not be specified if 'type' is set to '%v'", me.Type)
	}
	if (me.ServiceNowNotification == nil) && (string(me.Type) == "SERVICE_NOW") {
		return fmt.Errorf("'service_now_notification' must be specified if 'type' is set to '%v'", me.Type)
	}
	if (me.ServiceNowNotification != nil) && (string(me.Type) != "SERVICE_NOW") {
		return fmt.Errorf("'service_now_notification' must not be specified if 'type' is set to '%v'", me.Type)
	}
	if (me.SlackNotification == nil) && (string(me.Type) == "SLACK") {
		return fmt.Errorf("'slack_notification' must be specified if 'type' is set to '%v'", me.Type)
	}
	if (me.SlackNotification != nil) && (string(me.Type) != "SLACK") {
		return fmt.Errorf("'slack_notification' must not be specified if 'type' is set to '%v'", me.Type)
	}
	if (me.TrelloNotification == nil) && (string(me.Type) == "TRELLO") {
		return fmt.Errorf("'trello_notification' must be specified if 'type' is set to '%v'", me.Type)
	}
	if (me.TrelloNotification != nil) && (string(me.Type) != "TRELLO") {
		return fmt.Errorf("'trello_notification' must not be specified if 'type' is set to '%v'", me.Type)
	}
	if (me.VictorOpsNotification == nil) && (string(me.Type) == "VICTOROPS") {
		return fmt.Errorf("'victor_ops_notification' must be specified if 'type' is set to '%v'", me.Type)
	}
	if (me.VictorOpsNotification != nil) && (string(me.Type) != "VICTOROPS") {
		return fmt.Errorf("'victor_ops_notification' must not be specified if 'type' is set to '%v'", me.Type)
	}
	if (me.WebHookNotification == nil) && (string(me.Type) == "WEBHOOK") {
		return fmt.Errorf("'web_hook_notification' must be specified if 'type' is set to '%v'", me.Type)
	}
	if (me.WebHookNotification != nil) && (string(me.Type) != "WEBHOOK") {
		return fmt.Errorf("'web_hook_notification' must not be specified if 'type' is set to '%v'", me.Type)
	}
	if (me.XMattersNotification == nil) && (string(me.Type) == "XMATTERS") {
		return fmt.Errorf("'x_matters_notification' must be specified if 'type' is set to '%v'", me.Type)
	}
	if (me.XMattersNotification != nil) && (string(me.Type) != "XMATTERS") {
		return fmt.Errorf("'x_matters_notification' must not be specified if 'type' is set to '%v'", me.Type)
	}
	return nil
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
