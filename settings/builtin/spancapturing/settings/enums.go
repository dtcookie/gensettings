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

package spancapturing

type SpanCaptureAction string

var SpanCaptureActions = struct {
	Capture SpanCaptureAction
	Ignore  SpanCaptureAction
}{
	"CAPTURE",
	"IGNORE",
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
	Attribute                   SpanMatcherSource
	InstrumentationScopeName    SpanMatcherSource
	InstrumentationScopeVersion SpanMatcherSource
	SpanKind                    SpanMatcherSource
	SpanName                    SpanMatcherSource
}{
	"ATTRIBUTE",
	"INSTRUMENTATION_SCOPE_NAME",
	"INSTRUMENTATION_SCOPE_VERSION",
	"SPAN_KIND",
	"SPAN_NAME",
}

type SpanMatcherType string

var SpanMatcherTypes = struct {
	Contains         SpanMatcherType
	DoesNotContain   SpanMatcherType
	DoesNotEndWith   SpanMatcherType
	DoesNotEqual     SpanMatcherType
	DoesNotStartWith SpanMatcherType
	EndsWith         SpanMatcherType
	Equals           SpanMatcherType
	StartsWith       SpanMatcherType
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
