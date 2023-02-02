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

package spanentrypoints

type SpanEntrypointAction string

var SpanEntrypointActions = struct {
	CreateEntrypoint     SpanEntrypointAction
	DontCreateEntrypoint SpanEntrypointAction
}{
	SpanEntrypointAction("CREATE_ENTRYPOINT"),
	SpanEntrypointAction("DONT_CREATE_ENTRYPOINT"),
}

type SpanKind string

var SpanKinds = struct {
	Client   SpanKind
	Consumer SpanKind
	Internal SpanKind
	Producer SpanKind
	Server   SpanKind
}{
	SpanKind("CLIENT"),
	SpanKind("CONSUMER"),
	SpanKind("INTERNAL"),
	SpanKind("PRODUCER"),
	SpanKind("SERVER"),
}

type SpanMatcherSource string

var SpanMatcherSources = struct {
	Attribute                   SpanMatcherSource
	InstrumentationScopeName    SpanMatcherSource
	InstrumentationScopeVersion SpanMatcherSource
	SpanKind                    SpanMatcherSource
	SpanName                    SpanMatcherSource
}{
	SpanMatcherSource("ATTRIBUTE"),
	SpanMatcherSource("INSTRUMENTATION_SCOPE_NAME"),
	SpanMatcherSource("INSTRUMENTATION_SCOPE_VERSION"),
	SpanMatcherSource("SPAN_KIND"),
	SpanMatcherSource("SPAN_NAME"),
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
	SpanMatcherType("CONTAINS"),
	SpanMatcherType("DOES_NOT_CONTAIN"),
	SpanMatcherType("DOES_NOT_END_WITH"),
	SpanMatcherType("DOES_NOT_EQUAL"),
	SpanMatcherType("DOES_NOT_START_WITH"),
	SpanMatcherType("ENDS_WITH"),
	SpanMatcherType("EQUALS"),
	SpanMatcherType("STARTS_WITH"),
}
