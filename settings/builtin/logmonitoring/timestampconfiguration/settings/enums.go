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

package timestampconfiguration

type MatcherType string

var MatcherTypes = struct {
	ContainerName          MatcherType
	DtEntityContainerGroup MatcherType
	DtEntityProcessGroup   MatcherType
	K8SContainerName       MatcherType
	K8SDeploymentName      MatcherType
	K8SNamespaceName       MatcherType
	LogSource              MatcherType
	ProcessTechnology      MatcherType
}{
	"Container_name",
	"Dt_entity_container_group",
	"Dt_entity_process_group",
	"K8s_container_name",
	"K8s_deployment_name",
	"K8s_namespace_name",
	"Log_source",
	"Process_technology",
}

type Operator string

var Operators = struct {
	Matches Operator
}{
	"MATCHES",
}
