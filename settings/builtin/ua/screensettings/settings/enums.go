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

package screensettings

type BreadcrumbType string

var BreadcrumbTypes = struct {
	EntityListRef BreadcrumbType
	EntityRef     BreadcrumbType
	Noop          BreadcrumbType
	StaticLink    BreadcrumbType
}{
	"ENTITY_LIST_REF",
	"ENTITY_REF",
	"NOOP",
	"STATIC_LINK",
}

type CardType string

var CardTypes = struct {
	AttributeTable CardType
	BreakLine      CardType
	CardGroup      CardType
	ChartGroup     CardType
	EntitiesList   CardType
	Events         CardType
	HealthCard     CardType
	Injections     CardType
	Logs           CardType
	Message        CardType
	MetricTable    CardType
	SloList        CardType
	Traces         CardType
}{
	"ATTRIBUTE_TABLE",
	"BREAK_LINE",
	"CARD_GROUP",
	"CHART_GROUP",
	"ENTITIES_LIST",
	"EVENTS",
	"HEALTH_CARD",
	"INJECTIONS",
	"LOGS",
	"MESSAGE",
	"METRIC_TABLE",
	"SLO_LIST",
	"TRACES",
}

type WidthType string

var WidthTypes = struct {
	FullSize WidthType
	HalfSize WidthType
}{
	"FULL_SIZE",
	"HALF_SIZE",
}
