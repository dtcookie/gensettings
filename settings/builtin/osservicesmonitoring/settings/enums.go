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

package osservicesmonitoring

type LinuxServiceProp string

var LinuxServiceProps = struct {
	Servicename LinuxServiceProp
	Startuptype LinuxServiceProp
}{
	LinuxServiceProp("ServiceName"),
	LinuxServiceProp("StartupType"),
}

type System string

var Systems = struct {
	Linux   System
	Windows System
}{
	System("LINUX"),
	System("WINDOWS"),
}

type WindowsServiceProps string

var WindowsServicePropss = struct {
	Displayname  WindowsServiceProps
	Manufacturer WindowsServiceProps
	Path         WindowsServiceProps
	Servicename  WindowsServiceProps
	Startuptype  WindowsServiceProps
}{
	WindowsServiceProps("DisplayName"),
	WindowsServiceProps("Manufacturer"),
	WindowsServiceProps("Path"),
	WindowsServiceProps("ServiceName"),
	WindowsServiceProps("StartupType"),
}
