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

package kpms

type LoadKpm string

var LoadKpms = struct {
	CumulativeLayoutShift  LoadKpm
	DomInteractive         LoadKpm
	LargestContentfulPaint LoadKpm
	LoadEventEnd           LoadKpm
	LoadEventStart         LoadKpm
	ResponseEnd            LoadKpm
	ResponseStart          LoadKpm
	SpeedIndex             LoadKpm
	UserActionDuration     LoadKpm
	VisuallyComplete       LoadKpm
}{
	LoadKpm("CUMULATIVE_LAYOUT_SHIFT"),
	LoadKpm("DOM_INTERACTIVE"),
	LoadKpm("LARGEST_CONTENTFUL_PAINT"),
	LoadKpm("LOAD_EVENT_END"),
	LoadKpm("LOAD_EVENT_START"),
	LoadKpm("RESPONSE_END"),
	LoadKpm("RESPONSE_START"),
	LoadKpm("SPEED_INDEX"),
	LoadKpm("USER_ACTION_DURATION"),
	LoadKpm("VISUALLY_COMPLETE"),
}

type XhrKpm string

var XhrKpms = struct {
	ResponseEnd        XhrKpm
	ResponseStart      XhrKpm
	UserActionDuration XhrKpm
	VisuallyComplete   XhrKpm
}{
	XhrKpm("RESPONSE_END"),
	XhrKpm("RESPONSE_START"),
	XhrKpm("USER_ACTION_DURATION"),
	XhrKpm("VISUALLY_COMPLETE"),
}
