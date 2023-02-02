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

type DayOfWeekType string

var DayOfWeekTypes = struct {
	Friday    DayOfWeekType
	Monday    DayOfWeekType
	Saturday  DayOfWeekType
	Sunday    DayOfWeekType
	Thursday  DayOfWeekType
	Tuesday   DayOfWeekType
	Wednesday DayOfWeekType
}{
	DayOfWeekType("FRIDAY"),
	DayOfWeekType("MONDAY"),
	DayOfWeekType("SATURDAY"),
	DayOfWeekType("SUNDAY"),
	DayOfWeekType("THURSDAY"),
	DayOfWeekType("TUESDAY"),
	DayOfWeekType("WEDNESDAY"),
}

type MaintenanceType string

var MaintenanceTypes = struct {
	Planned   MaintenanceType
	Unplanned MaintenanceType
}{
	MaintenanceType("PLANNED"),
	MaintenanceType("UNPLANNED"),
}

type ScheduleType string

var ScheduleTypes = struct {
	Daily   ScheduleType
	Monthly ScheduleType
	Once    ScheduleType
	Weekly  ScheduleType
}{
	ScheduleType("DAILY"),
	ScheduleType("MONTHLY"),
	ScheduleType("ONCE"),
	ScheduleType("WEEKLY"),
}

type SuppressionType string

var SuppressionTypes = struct {
	DetectProblemsAndAlert  SuppressionType
	DetectProblemsDontAlert SuppressionType
	DontDetectProblems      SuppressionType
}{
	SuppressionType("DETECT_PROBLEMS_AND_ALERT"),
	SuppressionType("DETECT_PROBLEMS_DONT_ALERT"),
	SuppressionType("DONT_DETECT_PROBLEMS"),
}
