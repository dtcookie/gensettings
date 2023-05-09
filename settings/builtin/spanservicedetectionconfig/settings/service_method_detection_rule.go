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

package spanservicedetectionconfig

import (
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ServiceMethodDetectionRules []*ServiceMethodDetectionRule

func (me *ServiceMethodDetectionRules) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"service_method_detection_rule": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(ServiceMethodDetectionRule).Schema()},
		},
	}
}

func (me ServiceMethodDetectionRules) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("service_method_detection_rule", me)
}

func (me *ServiceMethodDetectionRules) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("service_method_detection_rule", me)
}

type ServiceMethodDetectionRule struct {
	Conditions     Conditions           `json:"conditions"`           // Specify conditions which should be evaluated for this rule. A rule is evaluated if all of the specified conditions match.
	ID             MethodIdContributors `json:"id,omitempty"`         // Method id contributors
	MethodName     *string              `json:"methodName,omitempty"` // `Prefixes`: [SpanAttribute, Extraction, ResourceAttribute] i.e. {SpanAttribute:http.url}, {Extraction:my_variable}, , {ResourceAttribute:method.name}\n\n`Well-known placeholders`: {SpanName}, {SpanKind}, {InstrumentationScopeName}, {InstrumentationScopeVersion}\n\nSpan attributes must be [allow-listed](/ui/settings/builtin:span-attribute \"Visit Span attribute page\") in order to be used as a method name  \nResource attributes must be [allow-listed](/ui/settings/builtin:resource-attribute \"Visit Resource attribute page\") in order to be used as a method name
	MethodRuleName string               `json:"methodRuleName"`       // Rule name
	OutputType     MethodRulesType      `json:"outputType"`           // Possible Values: `METHOD_ID_AND_NAME`, `NO_SERVICE`
}

func (me *ServiceMethodDetectionRule) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"conditions": {
			Type:        schema.TypeList,
			Description: "Specify conditions which should be evaluated for this rule. A rule is evaluated if all of the specified conditions match.",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(Conditions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"id": {
			Type:        schema.TypeList,
			Description: "Method id contributors",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(MethodIdContributors).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"method_name": {
			Type:        schema.TypeString,
			Description: "`Prefixes`: [SpanAttribute, Extraction, ResourceAttribute] i.e. {SpanAttribute:http.url}, {Extraction:my_variable}, , {ResourceAttribute:method.name}\n\n`Well-known placeholders`: {SpanName}, {SpanKind}, {InstrumentationScopeName}, {InstrumentationScopeVersion}\n\nSpan attributes must be [allow-listed](/ui/settings/builtin:span-attribute \"Visit Span attribute page\") in order to be used as a method name  \nResource attributes must be [allow-listed](/ui/settings/builtin:resource-attribute \"Visit Resource attribute page\") in order to be used as a method name",
			Optional:    true, // precondition
		},
		"method_rule_name": {
			Type:        schema.TypeString,
			Description: "Rule name",
			Required:    true,
		},
		"output_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `METHOD_ID_AND_NAME`, `NO_SERVICE`",
			Required:    true,
		},
	}
}

func (me *ServiceMethodDetectionRule) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"conditions":       me.Conditions,
		"id":               me.ID,
		"method_name":      me.MethodName,
		"method_rule_name": me.MethodRuleName,
		"output_type":      me.OutputType,
	})
}

func (me *ServiceMethodDetectionRule) HandlePreconditions() error {
	if me.MethodName == nil && (string(me.OutputType) == "METHOD_ID_AND_NAME") {
		return fmt.Errorf("'method_name' must be specified if 'output_type' is set to '%v'", me.OutputType)
	}
	// ---- ID MethodIdContributors -> {"expectedValue":"METHOD_ID_AND_NAME","property":"outputType","type":"EQUALS"}
	return nil
}

func (me *ServiceMethodDetectionRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"conditions":       &me.Conditions,
		"id":               &me.ID,
		"method_name":      &me.MethodName,
		"method_rule_name": &me.MethodRuleName,
		"output_type":      &me.OutputType,
	})
}
