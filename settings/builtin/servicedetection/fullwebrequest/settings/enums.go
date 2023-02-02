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

package fullwebrequest

type ContextRootTransformationType string

var ContextRootTransformationTypes = struct {
	Before            ContextRootTransformationType
	RemoveCreditCards ContextRootTransformationType
	RemoveIbans       ContextRootTransformationType
	RemoveIps         ContextRootTransformationType
	RemoveNumbers     ContextRootTransformationType
	ReplaceBetween    ContextRootTransformationType
}{
	ContextRootTransformationType("BEFORE"),
	ContextRootTransformationType("REMOVE_CREDIT_CARDS"),
	ContextRootTransformationType("REMOVE_IBANS"),
	ContextRootTransformationType("REMOVE_IPS"),
	ContextRootTransformationType("REMOVE_NUMBERS"),
	ContextRootTransformationType("REPLACE_BETWEEN"),
}

type ContributionType string

var ContributionTypes = struct {
	Originalvalue  ContributionType
	Transformurl   ContributionType
	Transformvalue ContributionType
}{
	ContributionType("OriginalValue"),
	ContributionType("TransformURL"),
	ContributionType("TransformValue"),
}

type ContributionTypeWithOverride string

var ContributionTypeWithOverrides = struct {
	Originalvalue  ContributionTypeWithOverride
	Overridevalue  ContributionTypeWithOverride
	Transformvalue ContributionTypeWithOverride
}{
	ContributionTypeWithOverride("OriginalValue"),
	ContributionTypeWithOverride("OverrideValue"),
	ContributionTypeWithOverride("TransformValue"),
}

type FrameworkType string

var FrameworkTypes = struct {
	Axis       FrameworkType
	Cxf        FrameworkType
	Hessian    FrameworkType
	JaxWsRi    FrameworkType
	Jboss      FrameworkType
	Jersey     FrameworkType
	Progress   FrameworkType
	Resteasy   FrameworkType
	Restlet    FrameworkType
	Spring     FrameworkType
	Tibco      FrameworkType
	Weblogic   FrameworkType
	Webmethods FrameworkType
	Websphere  FrameworkType
	Wink       FrameworkType
}{
	FrameworkType("AXIS"),
	FrameworkType("CXF"),
	FrameworkType("HESSIAN"),
	FrameworkType("JAX_WS_RI"),
	FrameworkType("JBOSS"),
	FrameworkType("JERSEY"),
	FrameworkType("PROGRESS"),
	FrameworkType("RESTEASY"),
	FrameworkType("RESTLET"),
	FrameworkType("SPRING"),
	FrameworkType("TIBCO"),
	FrameworkType("WEBLOGIC"),
	FrameworkType("WEBMETHODS"),
	FrameworkType("WEBSPHERE"),
	FrameworkType("WINK"),
}

type TransformationType string

var TransformationTypes = struct {
	After             TransformationType
	Before            TransformationType
	Between           TransformationType
	RemoveCreditCards TransformationType
	RemoveIbans       TransformationType
	RemoveIps         TransformationType
	RemoveNumbers     TransformationType
	ReplaceBetween    TransformationType
	SplitSelect       TransformationType
	TakeSegments      TransformationType
}{
	TransformationType("AFTER"),
	TransformationType("BEFORE"),
	TransformationType("BETWEEN"),
	TransformationType("REMOVE_CREDIT_CARDS"),
	TransformationType("REMOVE_IBANS"),
	TransformationType("REMOVE_IPS"),
	TransformationType("REMOVE_NUMBERS"),
	TransformationType("REPLACE_BETWEEN"),
	TransformationType("SPLIT_SELECT"),
	TransformationType("TAKE_SEGMENTS"),
}
