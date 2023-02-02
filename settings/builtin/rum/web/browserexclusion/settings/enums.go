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

package browserexclusion

type Browser string

var Browsers = struct {
	AndroidWebkit    Browser
	BotsAndSpiders   Browser
	Chrome           Browser
	Edge             Browser
	Firefox          Browser
	InternetExplorer Browser
	Opera            Browser
	Safari           Browser
}{
	"ANDROID_WEBKIT",
	"BOTS_AND_SPIDERS",
	"CHROME",
	"EDGE",
	"FIREFOX",
	"INTERNET_EXPLORER",
	"OPERA",
	"SAFARI",
}

type Platform string

var Platforms = struct {
	All     Platform
	Desktop Platform
	Mobile  Platform
}{
	"ALL",
	"DESKTOP",
	"MOBILE",
}

type VersionComparator string

var VersionComparators = struct {
	Equals         VersionComparator
	GreaterOrEqual VersionComparator
	LessOrEqual    VersionComparator
}{
	"EQUALS",
	"GREATER_OR_EQUAL",
	"LESS_OR_EQUAL",
}
