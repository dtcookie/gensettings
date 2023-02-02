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

package appdetection

type Matcher string

var Matchers = struct {
	DomainContains   Matcher
	DomainEndsWith   Matcher
	DomainEquals     Matcher
	DomainMatches    Matcher
	DomainStartsWith Matcher
	UrlContains      Matcher
	UrlEndsWith      Matcher
	UrlEquals        Matcher
	UrlStartsWith    Matcher
}{
	Matcher("DOMAIN_CONTAINS"),
	Matcher("DOMAIN_ENDS_WITH"),
	Matcher("DOMAIN_EQUALS"),
	Matcher("DOMAIN_MATCHES"),
	Matcher("DOMAIN_STARTS_WITH"),
	Matcher("URL_CONTAINS"),
	Matcher("URL_ENDS_WITH"),
	Matcher("URL_EQUALS"),
	Matcher("URL_STARTS_WITH"),
}
