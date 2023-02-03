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

package managementzones

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ManagementZoneAttributeRule struct {
	AzureToPGPropagation                       *bool                `json:"azureToPGPropagation,omitempty"`      // Apply to process groups connected to matching Azure entities
	AzureToServicePropagation                  *bool                `json:"azureToServicePropagation,omitempty"` // Apply to services provided by matching Azure entities
	Conditions                                 AttributeConditions  `json:"conditions"`
	CustomDeviceGroupToCustomDevicePropagation *bool                `json:"customDeviceGroupToCustomDevicePropagation,omitempty"` // Apply to custom devices in a custom device group
	EntityType                                 ManagementZoneMeType `json:"entityType"`                                           // Possible Values: `HOST_GROUP`, `CUSTOM_APPLICATION`, `QUEUE`, `APPMON_SYSTEM_PROFILE`, `AWS_AUTO_SCALING_GROUP`, `ESXI_HOST`, `MOBILE_APPLICATION`, `AZURE`, `CUSTOM_DEVICE_GROUP`, `AWS_RELATIONAL_DATABASE_SERVICE`, `HOST`, `CLOUD_FOUNDRY_FOUNDATION`, `KUBERNETES_SERVICE`, `SERVICE`, `APPMON_SERVER`, `AWS_ACCOUNT`, `OPENSTACK_ACCOUNT`, `CUSTOM_DEVICE`, `CLOUD_APPLICATION_NAMESPACE`, `ENTERPRISE_APPLICATION`, `AWS_APPLICATION_LOAD_BALANCER`, `AWS_NETWORK_LOAD_BALANCER`, `EXTERNAL_MONITOR`, `CLOUD_APPLICATION`, `BROWSER_MONITOR`, `WEB_APPLICATION`, `HTTP_MONITOR`, `KUBERNETES_CLUSTER`, `AWS_CLASSIC_LOAD_BALANCER`, `PROCESS_GROUP`, `DATA_CENTER_SERVICE`
	HostToPGPropagation                        *bool                `json:"hostToPGPropagation,omitempty"`                        // Apply to processes running on matching hosts
	PGToHostPropagation                        *bool                `json:"pgToHostPropagation,omitempty"`                        // Apply to underlying hosts of matching process groups
	PGToServicePropagation                     *bool                `json:"pgToServicePropagation,omitempty"`                     // Apply to all services provided by the process groups
	ServiceToHostPropagation                   *bool                `json:"serviceToHostPropagation,omitempty"`                   // Apply to underlying hosts of matching services
	ServiceToPGPropagation                     *bool                `json:"serviceToPGPropagation,omitempty"`                     // Apply to underlying process groups of matching services
}

func (me *ManagementZoneAttributeRule) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"azure_to_pgpropagation": {
			Type:        schema.TypeBool,
			Description: "Apply to process groups connected to matching Azure entities",
			Optional:    true,
		},
		"azure_to_service_propagation": {
			Type:        schema.TypeBool,
			Description: "Apply to services provided by matching Azure entities",
			Optional:    true,
		},
		"conditions": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(AttributeConditions).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"custom_device_group_to_custom_device_propagation": {
			Type:        schema.TypeBool,
			Description: "Apply to custom devices in a custom device group",
			Optional:    true,
		},
		"entity_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `HOST_GROUP`, `CUSTOM_APPLICATION`, `QUEUE`, `APPMON_SYSTEM_PROFILE`, `AWS_AUTO_SCALING_GROUP`, `ESXI_HOST`, `MOBILE_APPLICATION`, `AZURE`, `CUSTOM_DEVICE_GROUP`, `AWS_RELATIONAL_DATABASE_SERVICE`, `HOST`, `CLOUD_FOUNDRY_FOUNDATION`, `KUBERNETES_SERVICE`, `SERVICE`, `APPMON_SERVER`, `AWS_ACCOUNT`, `OPENSTACK_ACCOUNT`, `CUSTOM_DEVICE`, `CLOUD_APPLICATION_NAMESPACE`, `ENTERPRISE_APPLICATION`, `AWS_APPLICATION_LOAD_BALANCER`, `AWS_NETWORK_LOAD_BALANCER`, `EXTERNAL_MONITOR`, `CLOUD_APPLICATION`, `BROWSER_MONITOR`, `WEB_APPLICATION`, `HTTP_MONITOR`, `KUBERNETES_CLUSTER`, `AWS_CLASSIC_LOAD_BALANCER`, `PROCESS_GROUP`, `DATA_CENTER_SERVICE`",
			Required:    true,
		},
		"host_to_pgpropagation": {
			Type:        schema.TypeBool,
			Description: "Apply to processes running on matching hosts",
			Optional:    true,
		},
		"pg_to_host_propagation": {
			Type:        schema.TypeBool,
			Description: "Apply to underlying hosts of matching process groups",
			Optional:    true,
		},
		"pg_to_service_propagation": {
			Type:        schema.TypeBool,
			Description: "Apply to all services provided by the process groups",
			Optional:    true,
		},
		"service_to_host_propagation": {
			Type:        schema.TypeBool,
			Description: "Apply to underlying hosts of matching services",
			Optional:    true,
		},
		"service_to_pgpropagation": {
			Type:        schema.TypeBool,
			Description: "Apply to underlying process groups of matching services",
			Optional:    true,
		},
	}
}

func (me *ManagementZoneAttributeRule) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"azure_to_pgpropagation":                           me.AzureToPGPropagation,
		"azure_to_service_propagation":                     me.AzureToServicePropagation,
		"conditions":                                       me.Conditions,
		"custom_device_group_to_custom_device_propagation": me.CustomDeviceGroupToCustomDevicePropagation,
		"entity_type":                                      me.EntityType,
		"host_to_pgpropagation":                            me.HostToPGPropagation,
		"pg_to_host_propagation":                           me.PGToHostPropagation,
		"pg_to_service_propagation":                        me.PGToServicePropagation,
		"service_to_host_propagation":                      me.ServiceToHostPropagation,
		"service_to_pgpropagation":                         me.ServiceToPGPropagation,
	})
}

func (me *ManagementZoneAttributeRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"azure_to_pgpropagation":                           &me.AzureToPGPropagation,
		"azure_to_service_propagation":                     &me.AzureToServicePropagation,
		"conditions":                                       &me.Conditions,
		"custom_device_group_to_custom_device_propagation": &me.CustomDeviceGroupToCustomDevicePropagation,
		"entity_type":                                      &me.EntityType,
		"host_to_pgpropagation":                            &me.HostToPGPropagation,
		"pg_to_host_propagation":                           &me.PGToHostPropagation,
		"pg_to_service_propagation":                        &me.PGToServicePropagation,
		"service_to_host_propagation":                      &me.ServiceToHostPropagation,
		"service_to_pgpropagation":                         &me.ServiceToPGPropagation,
	})
}
