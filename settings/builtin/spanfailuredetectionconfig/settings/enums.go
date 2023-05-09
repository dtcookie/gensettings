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

package spanfailuredetectionconfig

type AttrKeyBooleanComparisonType string

var AttrKeyBooleanComparisonTypes = struct {
	DoesNotEqual AttrKeyBooleanComparisonType
	DoesNotExist AttrKeyBooleanComparisonType
	Equals       AttrKeyBooleanComparisonType
	Exists       AttrKeyBooleanComparisonType
}{
	"DOES_NOT_EQUAL",
	"DOES_NOT_EXIST",
	"EQUALS",
	"EXISTS",
}

type AttrKeyStringComparisonType string

var AttrKeyStringComparisonTypes = struct {
	Contains         AttrKeyStringComparisonType
	DoesNotContain   AttrKeyStringComparisonType
	DoesNotEndWith   AttrKeyStringComparisonType
	DoesNotEqual     AttrKeyStringComparisonType
	DoesNotExist     AttrKeyStringComparisonType
	DoesNotStartWith AttrKeyStringComparisonType
	EndsWith         AttrKeyStringComparisonType
	Equals           AttrKeyStringComparisonType
	Exists           AttrKeyStringComparisonType
	StartsWith       AttrKeyStringComparisonType
}{
	"CONTAINS",
	"DOES_NOT_CONTAIN",
	"DOES_NOT_END_WITH",
	"DOES_NOT_EQUAL",
	"DOES_NOT_EXIST",
	"DOES_NOT_START_WITH",
	"ENDS_WITH",
	"EQUALS",
	"EXISTS",
	"STARTS_WITH",
}

type BooleanValue string

var BooleanValues = struct {
	False BooleanValue
	True  BooleanValue
}{
	"FALSE",
	"TRUE",
}

type ComparisonDataType string

var ComparisonDataTypes = struct {
	Boolean ComparisonDataType
	Float   ComparisonDataType
	Integer ComparisonDataType
	String  ComparisonDataType
}{
	"BOOLEAN",
	"FLOAT",
	"INTEGER",
	"STRING",
}

type EnumSelectionsType string

var EnumSelectionsTypes = struct {
	NotAnyOf EnumSelectionsType
	OneOf    EnumSelectionsType
}{
	"NOT_ANY_OF",
	"ONE_OF",
}

type NumberComparisonType string

var NumberComparisonTypes = struct {
	DoesNotEqual          NumberComparisonType
	DoesNotExist          NumberComparisonType
	Equals                NumberComparisonType
	Exists                NumberComparisonType
	GreaterThan           NumberComparisonType
	GreaterThanOrEqual    NumberComparisonType
	LowerThan             NumberComparisonType
	LowerThanOrEqual      NumberComparisonType
	NotGreaterThan        NumberComparisonType
	NotGreaterThanOrEqual NumberComparisonType
	NotLowerThan          NumberComparisonType
	NotLowerThanOrEqual   NumberComparisonType
}{
	"DOES_NOT_EQUAL",
	"DOES_NOT_EXIST",
	"EQUALS",
	"EXISTS",
	"GREATER_THAN",
	"GREATER_THAN_OR_EQUAL",
	"LOWER_THAN",
	"LOWER_THAN_OR_EQUAL",
	"NOT_GREATER_THAN",
	"NOT_GREATER_THAN_OR_EQUAL",
	"NOT_LOWER_THAN",
	"NOT_LOWER_THAN_OR_EQUAL",
}

type SpanEventMatcherSource string

var SpanEventMatcherSources = struct {
	SpanEventAttribute SpanEventMatcherSource
	SpanEventName      SpanEventMatcherSource
}{
	"SPAN_EVENT_ATTRIBUTE",
	"SPAN_EVENT_NAME",
}

type SpanKind string

var SpanKinds = struct {
	Client   SpanKind
	Consumer SpanKind
	Internal SpanKind
	Producer SpanKind
	Server   SpanKind
}{
	"CLIENT",
	"CONSUMER",
	"INTERNAL",
	"PRODUCER",
	"SERVER",
}

type SpanMatcherSource string

var SpanMatcherSources = struct {
	InstrumentationScopeName    SpanMatcherSource
	InstrumentationScopeVersion SpanMatcherSource
	ResourceAttribute           SpanMatcherSource
	SpanAttribute               SpanMatcherSource
	SpanEvent                   SpanMatcherSource
	SpanKind                    SpanMatcherSource
	SpanName                    SpanMatcherSource
	SpanStatusCode              SpanMatcherSource
	SpanStatusMessage           SpanMatcherSource
}{
	"INSTRUMENTATION_SCOPE_NAME",
	"INSTRUMENTATION_SCOPE_VERSION",
	"RESOURCE_ATTRIBUTE",
	"SPAN_ATTRIBUTE",
	"SPAN_EVENT",
	"SPAN_KIND",
	"SPAN_NAME",
	"SPAN_STATUS_CODE",
	"SPAN_STATUS_MESSAGE",
}

type SpanStatusCode string

var SpanStatusCodes = struct {
	Error SpanStatusCode
	Ok    SpanStatusCode
	Unset SpanStatusCode
}{
	"ERROR",
	"OK",
	"UNSET",
}

type SpanVerdictIfRuleMatches string

var SpanVerdictIfRuleMatchess = struct {
	SpanVerdictFailure SpanVerdictIfRuleMatches
	SpanVerdictSuccess SpanVerdictIfRuleMatches
}{
	"SPAN_VERDICT_FAILURE",
	"SPAN_VERDICT_SUCCESS",
}

type StringComparisonType string

var StringComparisonTypes = struct {
	Contains         StringComparisonType
	DoesNotContain   StringComparisonType
	DoesNotEndWith   StringComparisonType
	DoesNotEqual     StringComparisonType
	DoesNotStartWith StringComparisonType
	EndsWith         StringComparisonType
	Equals           StringComparisonType
	StartsWith       StringComparisonType
}{
	"CONTAINS",
	"DOES_NOT_CONTAIN",
	"DOES_NOT_END_WITH",
	"DOES_NOT_EQUAL",
	"DOES_NOT_START_WITH",
	"ENDS_WITH",
	"EQUALS",
	"STARTS_WITH",
}
