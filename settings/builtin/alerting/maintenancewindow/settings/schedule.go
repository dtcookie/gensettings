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

package maintenancewindow

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Schedule struct {
	DailyRecurrence   *DailyRecurrence   `json:"dailyRecurrence,omitempty"`
	MonthlyRecurrence *MonthlyRecurrence `json:"monthlyRecurrence,omitempty"`
	OnceRecurrence    *OnceRecurrence    `json:"onceRecurrence,omitempty"`
	ScheduleType      ScheduleType       `json:"scheduleType"` // Possible Values: `DAILY`, `MONTHLY`, `ONCE`, `WEEKLY`
	WeeklyRecurrence  *WeeklyRecurrence  `json:"weeklyRecurrence,omitempty"`
}

func (me *Schedule) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"daily_recurrence": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(DailyRecurrence).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"monthly_recurrence": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(MonthlyRecurrence).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"once_recurrence": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(OnceRecurrence).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"schedule_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `DAILY`, `MONTHLY`, `ONCE`, `WEEKLY`",
			Required:    true,
		},
		"weekly_recurrence": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(WeeklyRecurrence).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Schedule) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"daily_recurrence":   me.DailyRecurrence,
		"monthly_recurrence": me.MonthlyRecurrence,
		"once_recurrence":    me.OnceRecurrence,
		"schedule_type":      me.ScheduleType,
		"weekly_recurrence":  me.WeeklyRecurrence,
	})
}

func (me *Schedule) HandlePreconditions() {
	// ---- DailyRecurrence *DailyRecurrence
	// ---- MonthlyRecurrence *MonthlyRecurrence
	// ---- OnceRecurrence *OnceRecurrence
	// ---- WeeklyRecurrence *WeeklyRecurrence
}

func (me *Schedule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"daily_recurrence":   &me.DailyRecurrence,
		"monthly_recurrence": &me.MonthlyRecurrence,
		"once_recurrence":    &me.OnceRecurrence,
		"schedule_type":      &me.ScheduleType,
		"weekly_recurrence":  &me.WeeklyRecurrence,
	})
}
