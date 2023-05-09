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

package screenactions

type ActionLocation string

var ActionLocations = struct {
	Chart      ActionLocation
	Header     ActionLocation
	TableEntry ActionLocation
}{
	"CHART",
	"HEADER",
	"TABLE_ENTRY",
}

type ActionScope string

var ActionScopes = struct {
	AttributeTable ActionScope
	ChartGroup     ActionScope
	EntitiesList   ActionScope
	Events         ActionScope
	GlobalDetails  ActionScope
	GlobalList     ActionScope
	Logs           ActionScope
	MetricTable    ActionScope
	Problems       ActionScope
	Properties     ActionScope
	SloList        ActionScope
}{
	"ATTRIBUTE_TABLE",
	"CHART_GROUP",
	"ENTITIES_LIST",
	"EVENTS",
	"GLOBAL_DETAILS",
	"GLOBAL_LIST",
	"LOGS",
	"METRIC_TABLE",
	"PROBLEMS",
	"PROPERTIES",
	"SLO_LIST",
}

type ButtonColorType string

var ButtonColorTypes = struct {
	Primary   ButtonColorType
	Secondary ButtonColorType
}{
	"PRIMARY",
	"SECONDARY",
}
