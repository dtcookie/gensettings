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

type ServiceDetection struct {
	ID          ServiceIdContributors `json:"id,omitempty"`          // Service id contributors
	OutputType  ServiceDetectionType  `json:"outputType"`            // Possible Values: `NO_SERVICE`, `SERVICE_ID_AND_NAME`
	ServiceName *string               `json:"serviceName,omitempty"` // `Prefixes`: [SpanAttribute, Extraction, ResourceAttribute] i.e. {SpanAttribute:http.url}, {Extraction:my_variable}, {ResourceAttribute:service.name}\n\n`Well-known placeholders`: {SpanName}, {SpanKind}, {InstrumentationScopeName}, {InstrumentationScopeVersion}\n\nSpan attributes must be [allow-listed](/ui/settings/builtin:span-attribute \"Visit Span attribute page\") in order to be used as a service name  \nResource attributes must be [allow-listed](/ui/settings/builtin:resource-attribute \"Visit Resource attribute page\") in order to be used as a service name
}

func (me *ServiceDetection) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeList,
			Description: "Service id contributors",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(ServiceIdContributors).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"output_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `NO_SERVICE`, `SERVICE_ID_AND_NAME`",
			Required:    true,
		},
		"service_name": {
			Type:        schema.TypeString,
			Description: "`Prefixes`: [SpanAttribute, Extraction, ResourceAttribute] i.e. {SpanAttribute:http.url}, {Extraction:my_variable}, {ResourceAttribute:service.name}\n\n`Well-known placeholders`: {SpanName}, {SpanKind}, {InstrumentationScopeName}, {InstrumentationScopeVersion}\n\nSpan attributes must be [allow-listed](/ui/settings/builtin:span-attribute \"Visit Span attribute page\") in order to be used as a service name  \nResource attributes must be [allow-listed](/ui/settings/builtin:resource-attribute \"Visit Resource attribute page\") in order to be used as a service name",
			Optional:    true, // precondition
		},
	}
}

func (me *ServiceDetection) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"id":           me.ID,
		"output_type":  me.OutputType,
		"service_name": me.ServiceName,
	})
}

func (me *ServiceDetection) HandlePreconditions() error {
	if me.ServiceName == nil && (string(me.OutputType) == "SERVICE_ID_AND_NAME") {
		return fmt.Errorf("'service_name' must be specified if 'output_type' is set to '%v'", me.OutputType)
	}
	// ---- ID ServiceIdContributors -> {"expectedValue":"SERVICE_ID_AND_NAME","property":"outputType","type":"EQUALS"}
	return nil
}

func (me *ServiceDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"id":           &me.ID,
		"output_type":  &me.OutputType,
		"service_name": &me.ServiceName,
	})
}
