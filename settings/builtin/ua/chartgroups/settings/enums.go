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

package chartgroups

type AggregationType string

var AggregationTypes = struct {
	Avg    AggregationType
	Count  AggregationType
	Max    AggregationType
	Median AggregationType
	Min    AggregationType
	Sum    AggregationType
	Value  AggregationType
}{
	"AVG",
	"COUNT",
	"MAX",
	"MEDIAN",
	"MIN",
	"SUM",
	"VALUE",
}

type ChartGroupMode string

var ChartGroupModes = struct {
	Normal ChartGroupMode
	Static ChartGroupMode
}{
	"NORMAL",
	"STATIC",
}

type ChartVisualizationType string

var ChartVisualizationTypes = struct {
	GraphChart  ChartVisualizationType
	PieChart    ChartVisualizationType
	SingleValue ChartVisualizationType
}{
	"GRAPH_CHART",
	"PIE_CHART",
	"SINGLE_VALUE",
}

type FoldTransformation string

var FoldTransformations = struct {
	Auto         FoldTransformation
	LastValue    FoldTransformation
	Max          FoldTransformation
	Median       FoldTransformation
	Min          FoldTransformation
	Percentile10 FoldTransformation
	Percentile75 FoldTransformation
	Percentile90 FoldTransformation
	Sum          FoldTransformation
	Value        FoldTransformation
}{
	"AUTO",
	"LAST_VALUE",
	"MAX",
	"MEDIAN",
	"MIN",
	"PERCENTILE_10",
	"PERCENTILE_75",
	"PERCENTILE_90",
	"SUM",
	"VALUE",
}

type Position string

var Positions = struct {
	Left  Position
	Right Position
}{
	"LEFT",
	"RIGHT",
}

type SeriesType string

var SeriesTypes = struct {
	Area   SeriesType
	Column SeriesType
	Line   SeriesType
}{
	"AREA",
	"COLUMN",
	"LINE",
}

type ThemeColorEnum string

var ThemeColorEnums = struct {
	Blue      ThemeColorEnum
	Default   ThemeColorEnum
	Gray      ThemeColorEnum
	Green     ThemeColorEnum
	Orange    ThemeColorEnum
	Purple    ThemeColorEnum
	Red       ThemeColorEnum
	Royalblue ThemeColorEnum
	Turquoise ThemeColorEnum
	Yellow    ThemeColorEnum
}{
	"BLUE",
	"DEFAULT",
	"GRAY",
	"GREEN",
	"ORANGE",
	"PURPLE",
	"RED",
	"ROYALBLUE",
	"TURQUOISE",
	"YELLOW",
}
