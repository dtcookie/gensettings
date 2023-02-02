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

package fullwebservice

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type IdContributorsType struct {
	ApplicationID             *ServiceIdContributor `json:"applicationId,omitempty"`       // Let the detected application identifier contribute to the service identifier calculation
	ContextRoot               *ContextIdContributor `json:"contextRoot,omitempty"`         // Let the detected context root contribute to the service identifier calculation.\nThe context root is the first segment of the request URL after the server name. For example, in the www.dynatrace.com/support/help/dynatrace-api/ URL the context root is `support`.
	DetectAsWebRequestService bool                  `json:"detectAsWebRequestService"`     // Detect the matching requests as full web services (false) or web request services (true).\n\nSetting this field to true prevents detecting of matching requests as full web services. A web request service is created instead. If you need to further modify the resulting web request service, you need to create a separate [Full web request rule](builtin:service-detection.full-web-request).
	ServerName                *ServiceIdContributor `json:"serverName,omitempty"`          // Let the detected server name contribute to the service identifier calculation
	WebServiceName            *ServiceIdContributor `json:"webServiceName,omitempty"`      // Let the detected web service name contribute to the service identifier calculation
	WebServiceNamespace       *ServiceIdContributor `json:"webServiceNamespace,omitempty"` // Let the detected web service namespace contribute to the service identifier calculation
}

func (me *IdContributorsType) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_id": {
			Type:        schema.TypeList,
			Description: "Let the detected application identifier contribute to the service identifier calculation",
			Optional:    true,
			Elem:        &schema.Resource{Schema: new(ServiceIdContributor).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"context_root": {
			Type:        schema.TypeList,
			Description: "Let the detected context root contribute to the service identifier calculation.\nThe context root is the first segment of the request URL after the server name. For example, in the www.dynatrace.com/support/help/dynatrace-api/ URL the context root is `support`.",
			Optional:    true,
			Elem:        &schema.Resource{Schema: new(ContextIdContributor).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detect_as_web_request_service": {
			Type:        schema.TypeBool,
			Description: "Detect the matching requests as full web services (false) or web request services (true).\n\nSetting this field to true prevents detecting of matching requests as full web services. A web request service is created instead. If you need to further modify the resulting web request service, you need to create a separate [Full web request rule](builtin:service-detection.full-web-request).",
			Required:    true,
		},
		"server_name": {
			Type:        schema.TypeList,
			Description: "Let the detected server name contribute to the service identifier calculation",
			Optional:    true,
			Elem:        &schema.Resource{Schema: new(ServiceIdContributor).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"web_service_name": {
			Type:        schema.TypeList,
			Description: "Let the detected web service name contribute to the service identifier calculation",
			Optional:    true,
			Elem:        &schema.Resource{Schema: new(ServiceIdContributor).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"web_service_namespace": {
			Type:        schema.TypeList,
			Description: "Let the detected web service namespace contribute to the service identifier calculation",
			Optional:    true,
			Elem:        &schema.Resource{Schema: new(ServiceIdContributor).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *IdContributorsType) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"application_id":                me.ApplicationID,
		"context_root":                  me.ContextRoot,
		"detect_as_web_request_service": me.DetectAsWebRequestService,
		"server_name":                   me.ServerName,
		"web_service_name":              me.WebServiceName,
		"web_service_namespace":         me.WebServiceNamespace,
	})
}

func (me *IdContributorsType) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"application_id":                &me.ApplicationID,
		"context_root":                  &me.ContextRoot,
		"detect_as_web_request_service": &me.DetectAsWebRequestService,
		"server_name":                   &me.ServerName,
		"web_service_name":              &me.WebServiceName,
		"web_service_namespace":         &me.WebServiceNamespace,
	})
}
