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

type NotificationType string

var NotificationTypes = struct {
	Ansibletower NotificationType
	Email        NotificationType
	Jira         NotificationType
	OpsGenie     NotificationType
	PagerDuty    NotificationType
	ServiceNow   NotificationType
	Slack        NotificationType
	Trello       NotificationType
	Victorops    NotificationType
	Webhook      NotificationType
	Xmatters     NotificationType
}{
	"ANSIBLETOWER",
	"EMAIL",
	"JIRA",
	"OPS_GENIE",
	"PAGER_DUTY",
	"SERVICE_NOW",
	"SLACK",
	"TRELLO",
	"VICTOROPS",
	"WEBHOOK",
	"XMATTERS",
}
