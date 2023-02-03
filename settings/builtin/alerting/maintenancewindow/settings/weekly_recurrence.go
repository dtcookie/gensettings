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

type WeeklyRecurrence struct {
	DayOfWeek       DayOfWeekType    `json:"dayOfWeek"`       // Possible Values: `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`, `SUNDAY`, `MONDAY`, `TUESDAY`
	RecurrenceRange *RecurrenceRange `json:"recurrenceRange"` // Recurrence range
	TimeWindow      *TimeWindow      `json:"timeWindow"`      // Time window
}

func (me *WeeklyRecurrence) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"day_of_week": {
			Type:        schema.TypeString,
			Description: "Possible Values: `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`, `SUNDAY`, `MONDAY`, `TUESDAY`",
			Required:    true,
		},
		"recurrence_range": {
			Type:        schema.TypeList,
			Description: "Recurrence range",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(RecurrenceRange).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"time_window": {
			Type:        schema.TypeList,
			Description: "Time window",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(TimeWindow).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
	}
}

func (me *WeeklyRecurrence) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"day_of_week":      me.DayOfWeek,
		"recurrence_range": me.RecurrenceRange,
		"time_window":      me.TimeWindow,
	})
}

func (me *WeeklyRecurrence) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"day_of_week":      &me.DayOfWeek,
		"recurrence_range": &me.RecurrenceRange,
		"time_window":      &me.TimeWindow,
	})
}
