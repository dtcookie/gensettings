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

package diskrules

type DiskMetric string

var DiskMetrics = struct {
	LowDiskSpace       DiskMetric
	LowInodes          DiskMetric
	ReadTimeExceeding  DiskMetric
	WriteTimeExceeding DiskMetric
}{
	DiskMetric("LOW_DISK_SPACE"),
	DiskMetric("LOW_INODES"),
	DiskMetric("READ_TIME_EXCEEDING"),
	DiskMetric("WRITE_TIME_EXCEEDING"),
}

type DiskNameFilterOperator string

var DiskNameFilterOperators = struct {
	Contains         DiskNameFilterOperator
	DoesNotContain   DiskNameFilterOperator
	DoesNotEqual     DiskNameFilterOperator
	DoesNotStartWith DiskNameFilterOperator
	Equals           DiskNameFilterOperator
	StartsWith       DiskNameFilterOperator
}{
	DiskNameFilterOperator("CONTAINS"),
	DiskNameFilterOperator("DOES_NOT_CONTAIN"),
	DiskNameFilterOperator("DOES_NOT_EQUAL"),
	DiskNameFilterOperator("DOES_NOT_START_WITH"),
	DiskNameFilterOperator("EQUALS"),
	DiskNameFilterOperator("STARTS_WITH"),
}
