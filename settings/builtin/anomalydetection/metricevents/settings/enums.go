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

package metricevents

type Aggregation string

var Aggregations = struct {
	Avg          Aggregation
	Count        Aggregation
	Max          Aggregation
	Median       Aggregation
	Min          Aggregation
	Percentile90 Aggregation
	Sum          Aggregation
	Value        Aggregation
}{
	Aggregation("AVG"),
	Aggregation("COUNT"),
	Aggregation("MAX"),
	Aggregation("MEDIAN"),
	Aggregation("MIN"),
	Aggregation("PERCENTILE90"),
	Aggregation("SUM"),
	Aggregation("VALUE"),
}

type AlertCondition string

var AlertConditions = struct {
	Above   AlertCondition
	Below   AlertCondition
	Outside AlertCondition
}{
	AlertCondition("ABOVE"),
	AlertCondition("BELOW"),
	AlertCondition("OUTSIDE"),
}

type EntityFilterOperator string

var EntityFilterOperators = struct {
	ContainsCaseInsensitive EntityFilterOperator
	ContainsCaseSensitive   EntityFilterOperator
	Equals                  EntityFilterOperator
}{
	EntityFilterOperator("CONTAINS_CASE_INSENSITIVE"),
	EntityFilterOperator("CONTAINS_CASE_SENSITIVE"),
	EntityFilterOperator("EQUALS"),
}

type EntityFilterType string

var EntityFilterTypes = struct {
	CustomDeviceGroupName EntityFilterType
	EntityId              EntityFilterType
	HostGroupName         EntityFilterType
	HostName              EntityFilterType
	ManagementZone        EntityFilterType
	Name                  EntityFilterType
	ProcessGroupId        EntityFilterType
	ProcessGroupName      EntityFilterType
	Tag                   EntityFilterType
}{
	EntityFilterType("CUSTOM_DEVICE_GROUP_NAME"),
	EntityFilterType("ENTITY_ID"),
	EntityFilterType("HOST_GROUP_NAME"),
	EntityFilterType("HOST_NAME"),
	EntityFilterType("MANAGEMENT_ZONE"),
	EntityFilterType("NAME"),
	EntityFilterType("PROCESS_GROUP_ID"),
	EntityFilterType("PROCESS_GROUP_NAME"),
	EntityFilterType("TAG"),
}

type EventTypeEnum string

var EventTypeEnums = struct {
	Availability         EventTypeEnum
	CustomAlert          EventTypeEnum
	CustomAnnotation     EventTypeEnum
	CustomConfiguration  EventTypeEnum
	CustomDeployment     EventTypeEnum
	Error                EventTypeEnum
	Info                 EventTypeEnum
	MarkedForTermination EventTypeEnum
	Resource             EventTypeEnum
	Slowdown             EventTypeEnum
}{
	EventTypeEnum("AVAILABILITY"),
	EventTypeEnum("CUSTOM_ALERT"),
	EventTypeEnum("CUSTOM_ANNOTATION"),
	EventTypeEnum("CUSTOM_CONFIGURATION"),
	EventTypeEnum("CUSTOM_DEPLOYMENT"),
	EventTypeEnum("ERROR"),
	EventTypeEnum("INFO"),
	EventTypeEnum("MARKED_FOR_TERMINATION"),
	EventTypeEnum("RESOURCE"),
	EventTypeEnum("SLOWDOWN"),
}

type ModelType string

var ModelTypes = struct {
	AutoAdaptiveThreshold ModelType
	SeasonalBaseline      ModelType
	StaticThreshold       ModelType
}{
	ModelType("AUTO_ADAPTIVE_THRESHOLD"),
	ModelType("SEASONAL_BASELINE"),
	ModelType("STATIC_THRESHOLD"),
}

type Type string

var Types = struct {
	MetricKey      Type
	MetricSelector Type
}{
	Type("METRIC_KEY"),
	Type("METRIC_SELECTOR"),
}
