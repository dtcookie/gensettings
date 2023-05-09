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

package spanservicedetectionconfig

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

type BooleanComparisonType string

var BooleanComparisonTypes = struct {
	DoesNotEqual BooleanComparisonType
	Equals       BooleanComparisonType
}{
	"DOES_NOT_EQUAL",
	"EQUALS",
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

type ExtractionMatcherType string

var ExtractionMatcherTypes = struct {
	After          ExtractionMatcherType
	Before         ExtractionMatcherType
	Between        ExtractionMatcherType
	SplitAndSelect ExtractionMatcherType
}{
	"AFTER",
	"BEFORE",
	"BETWEEN",
	"SPLIT_AND_SELECT",
}

type KnownMethodInput string

var KnownMethodInputs = struct {
	ExtractionValue             KnownMethodInput
	InstrumentationScopeName    KnownMethodInput
	InstrumentationScopeVersion KnownMethodInput
	ResourceAttribute           KnownMethodInput
	SpanAttribute               KnownMethodInput
	SpanKind                    KnownMethodInput
	SpanName                    KnownMethodInput
}{
	"EXTRACTION_VALUE",
	"INSTRUMENTATION_SCOPE_NAME",
	"INSTRUMENTATION_SCOPE_VERSION",
	"RESOURCE_ATTRIBUTE",
	"SPAN_ATTRIBUTE",
	"SPAN_KIND",
	"SPAN_NAME",
}

type KnownServiceInput string

var KnownServiceInputs = struct {
	ExtractionValue             KnownServiceInput
	InstrumentationScopeName    KnownServiceInput
	InstrumentationScopeVersion KnownServiceInput
	PgId                        KnownServiceInput
	ResourceAttribute           KnownServiceInput
	SpanAttribute               KnownServiceInput
	SpanKind                    KnownServiceInput
	SpanName                    KnownServiceInput
}{
	"EXTRACTION_VALUE",
	"INSTRUMENTATION_SCOPE_NAME",
	"INSTRUMENTATION_SCOPE_VERSION",
	"PG_ID",
	"RESOURCE_ATTRIBUTE",
	"SPAN_ATTRIBUTE",
	"SPAN_KIND",
	"SPAN_NAME",
}

type MethodRulesType string

var MethodRulesTypes = struct {
	MethodIdAndName MethodRulesType
	NoService       MethodRulesType
}{
	"METHOD_ID_AND_NAME",
	"NO_SERVICE",
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

type ServiceDetectionType string

var ServiceDetectionTypes = struct {
	NoService        ServiceDetectionType
	ServiceIdAndName ServiceDetectionType
}{
	"NO_SERVICE",
	"SERVICE_ID_AND_NAME",
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
	SpanKindClient   SpanKind
	SpanKindConsumer SpanKind
	SpanKindInternal SpanKind
	SpanKindProducer SpanKind
	SpanKindServer   SpanKind
}{
	"SPAN_KIND_CLIENT",
	"SPAN_KIND_CONSUMER",
	"SPAN_KIND_INTERNAL",
	"SPAN_KIND_PRODUCER",
	"SPAN_KIND_SERVER",
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
}{
	"INSTRUMENTATION_SCOPE_NAME",
	"INSTRUMENTATION_SCOPE_VERSION",
	"RESOURCE_ATTRIBUTE",
	"SPAN_ATTRIBUTE",
	"SPAN_EVENT",
	"SPAN_KIND",
	"SPAN_NAME",
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

type ValueSourceType string

var ValueSourceTypes = struct {
	InstrumentationScopeName    ValueSourceType
	InstrumentationScopeVersion ValueSourceType
	ResourceAttribute           ValueSourceType
	SpanAttribute               ValueSourceType
	SpanName                    ValueSourceType
}{
	"INSTRUMENTATION_SCOPE_NAME",
	"INSTRUMENTATION_SCOPE_VERSION",
	"RESOURCE_ATTRIBUTE",
	"SPAN_ATTRIBUTE",
	"SPAN_NAME",
}
