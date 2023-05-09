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

package variables

type PatternSearchType string

var PatternSearchTypes = struct {
	First PatternSearchType
	Last  PatternSearchType
}{
	"FIRST",
	"LAST",
}

type ProcessingStepType string

var ProcessingStepTypes = struct {
	ConvertToLowerCase ProcessingStepType
	ConvertToUpperCase ProcessingStepType
	Extract            ProcessingStepType
	ExtractWithRegex   ProcessingStepType
	Replace            ProcessingStepType
	ReplaceIds         ProcessingStepType
	ReplaceText        ProcessingStepType
	ReplaceWithRegex   ProcessingStepType
}{
	"CONVERT_TO_LOWER_CASE",
	"CONVERT_TO_UPPER_CASE",
	"EXTRACT",
	"EXTRACT_WITH_REGEX",
	"REPLACE",
	"REPLACE_IDS",
	"REPLACE_TEXT",
	"REPLACE_WITH_REGEX",
}

type SourceType string

var SourceTypes = struct {
	CookieValue        SourceType
	CssSelector        SourceType
	CustomActionName   SourceType
	ElementIdentifier  SourceType
	JavascriptVariable SourceType
	KeypressEvent      SourceType
	MetaTag            SourceType
	PageTitle          SourceType
	PageUrl            SourceType
	PageUrlAnchor      SourceType
	PageUrlPath        SourceType
	QueryString        SourceType
	RequestAttribute   SourceType
	SourceUrl          SourceType
	SourceUrlAnchor    SourceType
	SourceUrlPath      SourceType
	TopXhrUrl          SourceType
	TopXhrUrlAnchor    SourceType
	TopXhrUrlPath      SourceType
	UserInteraction    SourceType
	XhrUrl             SourceType
	XhrUrlAnchor       SourceType
	XhrUrlPath         SourceType
}{
	"COOKIE_VALUE",
	"CSS_SELECTOR",
	"CUSTOM_ACTION_NAME",
	"ELEMENT_IDENTIFIER",
	"JAVASCRIPT_VARIABLE",
	"KEYPRESS_EVENT",
	"META_TAG",
	"PAGE_TITLE",
	"PAGE_URL",
	"PAGE_URL_ANCHOR",
	"PAGE_URL_PATH",
	"QUERY_STRING",
	"REQUEST_ATTRIBUTE",
	"SOURCE_URL",
	"SOURCE_URL_ANCHOR",
	"SOURCE_URL_PATH",
	"TOP_XHR_URL",
	"TOP_XHR_URL_ANCHOR",
	"TOP_XHR_URL_PATH",
	"USER_INTERACTION",
	"XHR_URL",
	"XHR_URL_ANCHOR",
	"XHR_URL_PATH",
}
