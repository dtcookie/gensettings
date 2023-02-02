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

package integration

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Enabled            bool               `json:"enabled"`            // Enabled
	Issuelabel         string             `json:"issuelabel"`         // Set a label to identify these issues, for example, `release_blocker` or `non-critical`
	Issuequery         string             `json:"issuequery"`         // You can use the following placeholders to automatically insert values from the **Release monitoring** page in your query: `{NAME}`, `{VERSION}`, `{STAGE}`, `{PRODUCT}`.
	Issuetheme         IssueTheme         `json:"issuetheme"`         // Possible Values: `INFO`, `ERROR`, `RESOLVED`
	Issuetrackersystem IssueTrackerSystem `json:"issuetrackersystem"` // Possible Values: `JIRA_ON_PREMISE`, `JIRA_CLOUD`, `JIRA`, `GITHUB`, `GITLAB`, `SERVICENOW`
	Password           *string            `json:"password,omitempty"` // Password
	Token              *string            `json:"token,omitempty"`    // Token
	Url                string             `json:"url"`                // For Jira, use the base URL (for example, https://jira.yourcompany.com); for GitHub, use the repository URL (for example, https://github.com/org/repo); for GitLab, use the specific project API for a single project (for example, https://gitlab.com/api/v4/projects/:projectId), and the specific group API for a multiple projects (for example, https://gitlab.com/api/v4/groups/:groupId); for ServiceNow, use your company instance URL (for example, https://yourinstance.service-now.com/)
	Username           string             `json:"username"`           // Username
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Type:        schema.TypeBool,
			Description: "Enabled",
			Required:    true,
		},
		"issuelabel": {
			Type:        schema.TypeString,
			Description: "Set a label to identify these issues, for example, `release_blocker` or `non-critical`",
			Required:    true,
		},
		"issuequery": {
			Type:        schema.TypeString,
			Description: "You can use the following placeholders to automatically insert values from the **Release monitoring** page in your query: `{NAME}`, `{VERSION}`, `{STAGE}`, `{PRODUCT}`.",
			Required:    true,
		},
		"issuetheme": {
			Type:        schema.TypeString,
			Description: "Possible Values: `INFO`, `ERROR`, `RESOLVED`",
			Required:    true,
		},
		"issuetrackersystem": {
			Type:        schema.TypeString,
			Description: "Possible Values: `JIRA_ON_PREMISE`, `JIRA_CLOUD`, `JIRA`, `GITHUB`, `GITLAB`, `SERVICENOW`",
			Required:    true,
		},
		"password": {
			Type:        schema.TypeString,
			Description: "Password",
			Optional:    true,
		},
		"token": {
			Type:        schema.TypeString,
			Description: "Token",
			Optional:    true,
		},
		"url": {
			Type:        schema.TypeString,
			Description: "For Jira, use the base URL (for example, https://jira.yourcompany.com); for GitHub, use the repository URL (for example, https://github.com/org/repo); for GitLab, use the specific project API for a single project (for example, https://gitlab.com/api/v4/projects/:projectId), and the specific group API for a multiple projects (for example, https://gitlab.com/api/v4/groups/:groupId); for ServiceNow, use your company instance URL (for example, https://yourinstance.service-now.com/)",
			Required:    true,
		},
		"username": {
			Type:        schema.TypeString,
			Description: "Username",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"enabled":            me.Enabled,
		"issuelabel":         me.Issuelabel,
		"issuequery":         me.Issuequery,
		"issuetheme":         me.Issuetheme,
		"issuetrackersystem": me.Issuetrackersystem,
		"password":           me.Password,
		"token":              me.Token,
		"url":                me.Url,
		"username":           me.Username,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":            &me.Enabled,
		"issuelabel":         &me.Issuelabel,
		"issuequery":         &me.Issuequery,
		"issuetheme":         &me.Issuetheme,
		"issuetrackersystem": &me.Issuetrackersystem,
		"password":           &me.Password,
		"token":              &me.Token,
		"url":                &me.Url,
		"username":           &me.Username,
	})
}
