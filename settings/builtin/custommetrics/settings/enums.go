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

package custommetrics

type Operator string

var Operators = struct {
	Equals               Operator
	GreaterThan          Operator
	GreaterThanOrEqualTo Operator
	In                   Operator
	IsNotNull            Operator
	IsNull               Operator
	LessThan             Operator
	LessThanOrEqualTo    Operator
	Like                 Operator
	NotEqual             Operator
	NotLike              Operator
	StartsWith           Operator
}{
	Operator("EQUALS"),
	Operator("GREATER_THAN"),
	Operator("GREATER_THAN_OR_EQUAL_TO"),
	Operator("IN"),
	Operator("IS_NOT_NULL"),
	Operator("IS_NULL"),
	Operator("LESS_THAN"),
	Operator("LESS_THAN_OR_EQUAL_TO"),
	Operator("LIKE"),
	Operator("NOT_EQUAL"),
	Operator("NOT_LIKE"),
	Operator("STARTS_WITH"),
}

type ValueType string

var ValueTypes = struct {
	Counter ValueType
	Field   ValueType
}{
	ValueType("COUNTER"),
	ValueType("FIELD"),
}
