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

package messagecard

type ButtonColorType string

var ButtonColorTypes = struct {
	Primary   ButtonColorType
	Secondary ButtonColorType
}{
	"PRIMARY",
	"SECONDARY",
}

type CardThemeType string

var CardThemeTypes = struct {
	Cta     CardThemeType
	Main    CardThemeType
	Warning CardThemeType
}{
	"CTA",
	"MAIN",
	"WARNING",
}

type MessageColorType string

var MessageColorTypes = struct {
	Error   MessageColorType
	Info    MessageColorType
	Warning MessageColorType
}{
	"ERROR",
	"INFO",
	"WARNING",
}

type VisualizationType string

var VisualizationTypes = struct {
	Card    VisualizationType
	Message VisualizationType
}{
	"CARD",
	"MESSAGE",
}
