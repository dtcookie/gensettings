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

package autotagging

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AttributeConditions []*AttributeCondition

func (me *AttributeConditions) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"condition": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(AttributeCondition).Schema()},
		},
	}
}

func (me AttributeConditions) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("condition", me)
}

func (me *AttributeConditions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("condition", me)
}

type AttributeCondition struct {
	CaseSensitive    *bool     `json:"caseSensitive,omitempty"`    // Case sensitive
	DynamicKey       *string   `json:"dynamicKey,omitempty"`       // Dynamic key
	DynamicKeySource *string   `json:"dynamicKeySource,omitempty"` // Key source
	EntityID         *string   `json:"entityId,omitempty"`         // Value
	EnumValue        *string   `json:"enumValue,omitempty"`        // Value
	IntegerValue     *int      `json:"integerValue,omitempty"`     // Value
	Key              Attribute `json:"key"`                        // Possible Values: `SERVICE_MESSAGING_LISTENER_CLASS_NAME`, `AZURE_TENANT_NAME`, `PROCESS_GROUP_AZURE_HOST_NAME`, `CUSTOM_DEVICE_TAGS`, `HOST_PAAS_MEMORY_LIMIT`, `CUSTOM_DEVICE_NAME`, `VMWARE_DATACENTER_NAME`, `HOST_BOSH_INSTANCE_ID`, `AWS_RELATIONAL_DATABASE_SERVICE_INSTANCE_CLASS`, `GOOGLE_CLOUD_PLATFORM_ZONE_NAME`, `AZURE_SUBSCRIPTION_UUID`, `EC2_INSTANCE_TAGS`, `HOST_LOGICAL_CPU_CORES`, `ENTERPRISE_APPLICATION_TAGS`, `MOBILE_APPLICATION_PLATFORM`, `AWS_CLASSIC_LOAD_BALANCER_NAME`, `PROCESS_GROUP_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_ENDPOINT`, `EXTERNAL_MONITOR_ENGINE_DESCRIPTION`, `WEB_APPLICATION_TAGS`, `CUSTOM_DEVICE_METADATA`, `HOST_AZURE_COMPUTE_MODE`, `HOST_HYPERVISOR_TYPE`, `BROWSER_MONITOR_NAME`, `AWS_CLASSIC_LOAD_BALANCER_TAGS`, `SERVICE_DATABASE_HOST_NAME`, `ENTERPRISE_APPLICATION_NAME`, `AZURE_SUBSCRIPTION_NAME`, `VMWARE_VM_NAME`, `HOST_GROUP_NAME`, `AWS_APPLICATION_LOAD_BALANCER_TAGS`, `SERVICE_DETECTED_NAME`, `EXTERNAL_MONITOR_TAGS`, `DATA_CENTER_SERVICE_PORT`, `HOST_NAME`, `KUBERNETES_SERVICE_NAME`, `CUSTOM_DEVICE_GROUP_TAGS`, `CLOUD_APPLICATION_NAME`, `PROCESS_GROUP_TECHNOLOGY_VERSION`, `SERVICE_WEB_SERVER_ENDPOINT`, `PROCESS_GROUP_ID`, `MOBILE_APPLICATION_NAME`, `HOST_DETECTED_NAME`, `EC2_INSTANCE_BEANSTALK_ENV_NAME`, `DATA_CENTER_SERVICE_NAME`, `OPENSTACK_VM_NAME`, `AZURE_REGION_NAME`, `ESXI_HOST_HARDWARE_VENDOR`, `GEOLOCATION_SITE_NAME`, `WEB_APPLICATION_NAME`, `HOST_AWS_NAME_TAG`, `HOST_AZURE_WEB_APPLICATION_HOST_NAMES`, `CUSTOM_APPLICATION_TYPE`, `SERVICE_TECHNOLOGY_EDITION`, `SERVICE_ESB_APPLICATION_NAME`, `CUSTOM_APPLICATION_PLATFORM`, `AWS_RELATIONAL_DATABASE_SERVICE_PORT`, `CLOUD_FOUNDRY_ORG_NAME`, `HOST_CPU_CORES`, `HOST_ONEAGENT_CUSTOM_HOST_NAME`, `EXTERNAL_MONITOR_ENGINE_NAME`, `DATA_CENTER_SERVICE_TAGS`, `OPENSTACK_REGION_NAME`, `BROWSER_MONITOR_TAGS`, `PROCESS_GROUP_TECHNOLOGY`, `ESXI_HOST_PRODUCT_VERSION`, `HOST_BOSH_AVAILABILITY_ZONE`, `OPENSTACK_VM_INSTANCE_TYPE`, `CUSTOM_DEVICE_TECHNOLOGY`, `ESXI_HOST_NAME`, `EC2_INSTANCE_AMI_ID`, `MOBILE_APPLICATION_TAGS`, `CUSTOM_DEVICE_GROUP_NAME`, `SERVICE_REMOTE_SERVICE_NAME`, `ENTERPRISE_APPLICATION_PORT`, `EC2_INSTANCE_PUBLIC_HOST_NAME`, `HOST_OS_VERSION`, `DATA_CENTER_SERVICE_DECODER_TYPE`, `WEB_APPLICATION_TYPE`, `HOST_BOSH_DEPLOYMENT_ID`, `AWS_NETWORK_LOAD_BALANCER_TAGS`, `PROCESS_GROUP_TECHNOLOGY_EDITION`, `CUSTOM_APPLICATION_TAGS`, `QUEUE_TECHNOLOGY`, `AWS_APPLICATION_LOAD_BALANCER_NAME`, `OPENSTACK_AVAILABILITY_ZONE_NAME`, `GOOGLE_COMPUTE_INSTANCE_PUBLIC_IP_ADDRESSES`, `DOCKER_FULL_IMAGE_NAME`, `PROCESS_GROUP_LISTEN_PORT`, `ENTERPRISE_APPLICATION_IP_ADDRESS`, `HTTP_MONITOR_NAME`, `AWS_NETWORK_LOAD_BALANCER_NAME`, `CLOUD_FOUNDRY_FOUNDATION_NAME`, `GOOGLE_COMPUTE_INSTANCE_ID`, `SERVICE_WEB_APPLICATION_ID`, `HOST_AIX_LOGICAL_CPU_COUNT`, `APPMON_SYSTEM_PROFILE_NAME`, `CLOUD_APPLICATION_LABELS`, `SERVICE_TOPOLOGY`, `SERVICE_TECHNOLOGY_VERSION`, `QUEUE_VENDOR`, `SERVICE_WEB_SERVICE_NAME`, `EXTERNAL_MONITOR_ENGINE_TYPE`, `SERVICE_REMOTE_ENDPOINT`, `AZURE_MGMT_GROUP_NAME`, `HOST_BITNESS`, `EC2_INSTANCE_AWS_INSTANCE_TYPE`, `PROCESS_GROUP_DETECTED_NAME`, `CUSTOM_APPLICATION_NAME`, `AZURE_MGMT_GROUP_UUID`, `ENTERPRISE_APPLICATION_METADATA`, `OPENSTACK_ACCOUNT_NAME`, `ESXI_HOST_TAGS`, `AWS_RELATIONAL_DATABASE_SERVICE_DB_NAME`, `SERVICE_PORT`, `HOST_ARCHITECTURE`, `SERVICE_DATABASE_NAME`, `CUSTOM_DEVICE_DNS_ADDRESS`, `HOST_CLOUD_TYPE`, `AWS_ACCOUNT_ID`, `PROCESS_GROUP_PREDEFINED_METADATA`, `SERVICE_TAGS`, `AZURE_ENTITY_TAGS`, `SERVICE_WEB_CONTEXT_ROOT`, `HOST_IP_ADDRESS`, `APPMON_SERVER_NAME`, `SERVICE_TYPE`, `HOST_AIX_VIRTUAL_CPU_COUNT`, `AWS_AVAILABILITY_ZONE_NAME`, `SERVICE_AKKA_ACTOR_SYSTEM`, `HOST_OS_TYPE`, `GOOGLE_COMPUTE_INSTANCE_PROJECT`, `ESXI_HOST_CLUSTER_NAME`, `EC2_INSTANCE_ID`, `AZURE_ENTITY_NAME`, `KUBERNETES_NODE_NAME`, `AWS_ACCOUNT_NAME`, `SERVICE_WEB_SERVER_NAME`, `ESXI_HOST_HARDWARE_MODEL`, `EC2_INSTANCE_AWS_SECURITY_GROUP`, `SERVICE_IBM_CTG_GATEWAY_URL`, `PROCESS_GROUP_CUSTOM_METADATA`, `AZURE_SCALE_SET_NAME`, `EC2_INSTANCE_NAME`, `DOCKER_CONTAINER_NAME`, `CLOUD_APPLICATION_NAMESPACE_LABELS`, `WEB_APPLICATION_NAME_PATTERN`, `HOST_PAAS_TYPE`, `HOST_KUBERNETES_LABELS`, `DOCKER_IMAGE_VERSION`, `ESXI_HOST_PRODUCT_NAME`, `NAME_OF_COMPUTE_NODE`, `HOST_GROUP_ID`, `OPENSTACK_PROJECT_NAME`, `SERVICE_CTG_SERVICE_NAME`, `GOOGLE_COMPUTE_INSTANCE_MACHINE_TYPE`, `AZURE_VM_NAME`, `SERVICE_PUBLIC_DOMAIN_NAME`, `HOST_AZURE_SKU`, `HOST_BOSH_NAME`, `PROCESS_GROUP_AZURE_SITE_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_NAME`, `AWS_AUTO_SCALING_GROUP_TAGS`, `CUSTOM_DEVICE_PORT`, `AWS_RELATIONAL_DATABASE_SERVICE_ENGINE`, `PROCESS_GROUP_TAGS`, `ENTERPRISE_APPLICATION_DECODER_TYPE`, `AZURE_TENANT_UUID`, `AWS_CLASSIC_LOAD_BALANCER_FRONTEND_PORTS`, `EC2_INSTANCE_PRIVATE_HOST_NAME`, `SERVICE_DATABASE_VENDOR`, `HOST_BOSH_INSTANCE_NAME`, `HOST_AIX_SIMULTANEOUS_THREADS`, `OPENSTACK_ACCOUNT_PROJECT_NAME`, `HOST_CUSTOM_METADATA`, `KUBERNETES_CLUSTER_NAME`, `DATA_CENTER_SERVICE_METADATA`, `SERVICE_NAME`, `HTTP_MONITOR_TAGS`, `SERVICE_TECHNOLOGY`, `HOST_AZURE_WEB_APPLICATION_SITE_NAMES`, `SERVICE_DATABASE_TOPOLOGY`, `CUSTOM_DEVICE_IP_ADDRESS`, `HOST_BOSH_STEMCELL_VERSION`, `GOOGLE_COMPUTE_INSTANCE_NAME`, `GOOGLE_COMPUTE_INSTANCE_PROJECT_ID`, `HOST_TAGS`, `EXTERNAL_MONITOR_NAME`, `AWS_AUTO_SCALING_GROUP_NAME`, `SERVICE_WEB_SERVICE_NAMESPACE`, `CLOUD_APPLICATION_NAMESPACE_NAME`, `QUEUE_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_TAGS`, `DATA_CENTER_SERVICE_IP_ADDRESS`, `HOST_TECHNOLOGY`, `OPENSTACK_VM_SECURITY_GROUP`
	Operator         Operator  `json:"operator"`                   // Possible Values: `NOT_ENDS_WITH`, `NOT_EQUALS`, `REGEX_MATCHES`, `NOT_CONTAINS`, `NOT_LOWER_THAN_OR_EQUAL`, `TAG_KEY_EQUALS`, `NOT_IS_IP_IN_RANGE`, `BEGINS_WITH`, `EXISTS`, `GREATER_THAN_OR_EQUAL`, `EQUALS`, `NOT_TAG_KEY_EQUALS`, `NOT_EXISTS`, `NOT_GREATER_THAN`, `LOWER_THAN_OR_EQUAL`, `NOT_GREATER_THAN_OR_EQUAL`, `CONTAINS`, `GREATER_THAN`, `LOWER_THAN`, `NOT_LOWER_THAN`, `NOT_BEGINS_WITH`, `ENDS_WITH`, `IS_IP_IN_RANGE`, `NOT_REGEX_MATCHES`
	StringValue      *string   `json:"stringValue,omitempty"`      // Value
	Tag              *string   `json:"tag,omitempty"`              // Format: `[CONTEXT]tagKey:tagValue`
}

func (me *AttributeCondition) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"case_sensitive": {
			Type:        schema.TypeBool,
			Description: "Case sensitive",
			Optional:    true,
		},
		"dynamic_key": {
			Type:        schema.TypeString,
			Description: "Dynamic key",
			Optional:    true,
		},
		"dynamic_key_source": {
			Type:        schema.TypeString,
			Description: "Key source",
			Optional:    true,
		},
		"entity_id": {
			Type:        schema.TypeString,
			Description: "Value",
			Optional:    true,
		},
		"enum_value": {
			Type:        schema.TypeString,
			Description: "Value",
			Optional:    true,
		},
		"integer_value": {
			Type:        schema.TypeInt,
			Description: "Value",
			Optional:    true,
		},
		"key": {
			Type:        schema.TypeString,
			Description: "Possible Values: `SERVICE_MESSAGING_LISTENER_CLASS_NAME`, `AZURE_TENANT_NAME`, `PROCESS_GROUP_AZURE_HOST_NAME`, `CUSTOM_DEVICE_TAGS`, `HOST_PAAS_MEMORY_LIMIT`, `CUSTOM_DEVICE_NAME`, `VMWARE_DATACENTER_NAME`, `HOST_BOSH_INSTANCE_ID`, `AWS_RELATIONAL_DATABASE_SERVICE_INSTANCE_CLASS`, `GOOGLE_CLOUD_PLATFORM_ZONE_NAME`, `AZURE_SUBSCRIPTION_UUID`, `EC2_INSTANCE_TAGS`, `HOST_LOGICAL_CPU_CORES`, `ENTERPRISE_APPLICATION_TAGS`, `MOBILE_APPLICATION_PLATFORM`, `AWS_CLASSIC_LOAD_BALANCER_NAME`, `PROCESS_GROUP_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_ENDPOINT`, `EXTERNAL_MONITOR_ENGINE_DESCRIPTION`, `WEB_APPLICATION_TAGS`, `CUSTOM_DEVICE_METADATA`, `HOST_AZURE_COMPUTE_MODE`, `HOST_HYPERVISOR_TYPE`, `BROWSER_MONITOR_NAME`, `AWS_CLASSIC_LOAD_BALANCER_TAGS`, `SERVICE_DATABASE_HOST_NAME`, `ENTERPRISE_APPLICATION_NAME`, `AZURE_SUBSCRIPTION_NAME`, `VMWARE_VM_NAME`, `HOST_GROUP_NAME`, `AWS_APPLICATION_LOAD_BALANCER_TAGS`, `SERVICE_DETECTED_NAME`, `EXTERNAL_MONITOR_TAGS`, `DATA_CENTER_SERVICE_PORT`, `HOST_NAME`, `KUBERNETES_SERVICE_NAME`, `CUSTOM_DEVICE_GROUP_TAGS`, `CLOUD_APPLICATION_NAME`, `PROCESS_GROUP_TECHNOLOGY_VERSION`, `SERVICE_WEB_SERVER_ENDPOINT`, `PROCESS_GROUP_ID`, `MOBILE_APPLICATION_NAME`, `HOST_DETECTED_NAME`, `EC2_INSTANCE_BEANSTALK_ENV_NAME`, `DATA_CENTER_SERVICE_NAME`, `OPENSTACK_VM_NAME`, `AZURE_REGION_NAME`, `ESXI_HOST_HARDWARE_VENDOR`, `GEOLOCATION_SITE_NAME`, `WEB_APPLICATION_NAME`, `HOST_AWS_NAME_TAG`, `HOST_AZURE_WEB_APPLICATION_HOST_NAMES`, `CUSTOM_APPLICATION_TYPE`, `SERVICE_TECHNOLOGY_EDITION`, `SERVICE_ESB_APPLICATION_NAME`, `CUSTOM_APPLICATION_PLATFORM`, `AWS_RELATIONAL_DATABASE_SERVICE_PORT`, `CLOUD_FOUNDRY_ORG_NAME`, `HOST_CPU_CORES`, `HOST_ONEAGENT_CUSTOM_HOST_NAME`, `EXTERNAL_MONITOR_ENGINE_NAME`, `DATA_CENTER_SERVICE_TAGS`, `OPENSTACK_REGION_NAME`, `BROWSER_MONITOR_TAGS`, `PROCESS_GROUP_TECHNOLOGY`, `ESXI_HOST_PRODUCT_VERSION`, `HOST_BOSH_AVAILABILITY_ZONE`, `OPENSTACK_VM_INSTANCE_TYPE`, `CUSTOM_DEVICE_TECHNOLOGY`, `ESXI_HOST_NAME`, `EC2_INSTANCE_AMI_ID`, `MOBILE_APPLICATION_TAGS`, `CUSTOM_DEVICE_GROUP_NAME`, `SERVICE_REMOTE_SERVICE_NAME`, `ENTERPRISE_APPLICATION_PORT`, `EC2_INSTANCE_PUBLIC_HOST_NAME`, `HOST_OS_VERSION`, `DATA_CENTER_SERVICE_DECODER_TYPE`, `WEB_APPLICATION_TYPE`, `HOST_BOSH_DEPLOYMENT_ID`, `AWS_NETWORK_LOAD_BALANCER_TAGS`, `PROCESS_GROUP_TECHNOLOGY_EDITION`, `CUSTOM_APPLICATION_TAGS`, `QUEUE_TECHNOLOGY`, `AWS_APPLICATION_LOAD_BALANCER_NAME`, `OPENSTACK_AVAILABILITY_ZONE_NAME`, `GOOGLE_COMPUTE_INSTANCE_PUBLIC_IP_ADDRESSES`, `DOCKER_FULL_IMAGE_NAME`, `PROCESS_GROUP_LISTEN_PORT`, `ENTERPRISE_APPLICATION_IP_ADDRESS`, `HTTP_MONITOR_NAME`, `AWS_NETWORK_LOAD_BALANCER_NAME`, `CLOUD_FOUNDRY_FOUNDATION_NAME`, `GOOGLE_COMPUTE_INSTANCE_ID`, `SERVICE_WEB_APPLICATION_ID`, `HOST_AIX_LOGICAL_CPU_COUNT`, `APPMON_SYSTEM_PROFILE_NAME`, `CLOUD_APPLICATION_LABELS`, `SERVICE_TOPOLOGY`, `SERVICE_TECHNOLOGY_VERSION`, `QUEUE_VENDOR`, `SERVICE_WEB_SERVICE_NAME`, `EXTERNAL_MONITOR_ENGINE_TYPE`, `SERVICE_REMOTE_ENDPOINT`, `AZURE_MGMT_GROUP_NAME`, `HOST_BITNESS`, `EC2_INSTANCE_AWS_INSTANCE_TYPE`, `PROCESS_GROUP_DETECTED_NAME`, `CUSTOM_APPLICATION_NAME`, `AZURE_MGMT_GROUP_UUID`, `ENTERPRISE_APPLICATION_METADATA`, `OPENSTACK_ACCOUNT_NAME`, `ESXI_HOST_TAGS`, `AWS_RELATIONAL_DATABASE_SERVICE_DB_NAME`, `SERVICE_PORT`, `HOST_ARCHITECTURE`, `SERVICE_DATABASE_NAME`, `CUSTOM_DEVICE_DNS_ADDRESS`, `HOST_CLOUD_TYPE`, `AWS_ACCOUNT_ID`, `PROCESS_GROUP_PREDEFINED_METADATA`, `SERVICE_TAGS`, `AZURE_ENTITY_TAGS`, `SERVICE_WEB_CONTEXT_ROOT`, `HOST_IP_ADDRESS`, `APPMON_SERVER_NAME`, `SERVICE_TYPE`, `HOST_AIX_VIRTUAL_CPU_COUNT`, `AWS_AVAILABILITY_ZONE_NAME`, `SERVICE_AKKA_ACTOR_SYSTEM`, `HOST_OS_TYPE`, `GOOGLE_COMPUTE_INSTANCE_PROJECT`, `ESXI_HOST_CLUSTER_NAME`, `EC2_INSTANCE_ID`, `AZURE_ENTITY_NAME`, `KUBERNETES_NODE_NAME`, `AWS_ACCOUNT_NAME`, `SERVICE_WEB_SERVER_NAME`, `ESXI_HOST_HARDWARE_MODEL`, `EC2_INSTANCE_AWS_SECURITY_GROUP`, `SERVICE_IBM_CTG_GATEWAY_URL`, `PROCESS_GROUP_CUSTOM_METADATA`, `AZURE_SCALE_SET_NAME`, `EC2_INSTANCE_NAME`, `DOCKER_CONTAINER_NAME`, `CLOUD_APPLICATION_NAMESPACE_LABELS`, `WEB_APPLICATION_NAME_PATTERN`, `HOST_PAAS_TYPE`, `HOST_KUBERNETES_LABELS`, `DOCKER_IMAGE_VERSION`, `ESXI_HOST_PRODUCT_NAME`, `NAME_OF_COMPUTE_NODE`, `HOST_GROUP_ID`, `OPENSTACK_PROJECT_NAME`, `SERVICE_CTG_SERVICE_NAME`, `GOOGLE_COMPUTE_INSTANCE_MACHINE_TYPE`, `AZURE_VM_NAME`, `SERVICE_PUBLIC_DOMAIN_NAME`, `HOST_AZURE_SKU`, `HOST_BOSH_NAME`, `PROCESS_GROUP_AZURE_SITE_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_NAME`, `AWS_AUTO_SCALING_GROUP_TAGS`, `CUSTOM_DEVICE_PORT`, `AWS_RELATIONAL_DATABASE_SERVICE_ENGINE`, `PROCESS_GROUP_TAGS`, `ENTERPRISE_APPLICATION_DECODER_TYPE`, `AZURE_TENANT_UUID`, `AWS_CLASSIC_LOAD_BALANCER_FRONTEND_PORTS`, `EC2_INSTANCE_PRIVATE_HOST_NAME`, `SERVICE_DATABASE_VENDOR`, `HOST_BOSH_INSTANCE_NAME`, `HOST_AIX_SIMULTANEOUS_THREADS`, `OPENSTACK_ACCOUNT_PROJECT_NAME`, `HOST_CUSTOM_METADATA`, `KUBERNETES_CLUSTER_NAME`, `DATA_CENTER_SERVICE_METADATA`, `SERVICE_NAME`, `HTTP_MONITOR_TAGS`, `SERVICE_TECHNOLOGY`, `HOST_AZURE_WEB_APPLICATION_SITE_NAMES`, `SERVICE_DATABASE_TOPOLOGY`, `CUSTOM_DEVICE_IP_ADDRESS`, `HOST_BOSH_STEMCELL_VERSION`, `GOOGLE_COMPUTE_INSTANCE_NAME`, `GOOGLE_COMPUTE_INSTANCE_PROJECT_ID`, `HOST_TAGS`, `EXTERNAL_MONITOR_NAME`, `AWS_AUTO_SCALING_GROUP_NAME`, `SERVICE_WEB_SERVICE_NAMESPACE`, `CLOUD_APPLICATION_NAMESPACE_NAME`, `QUEUE_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_TAGS`, `DATA_CENTER_SERVICE_IP_ADDRESS`, `HOST_TECHNOLOGY`, `OPENSTACK_VM_SECURITY_GROUP`",
			Required:    true,
		},
		"operator": {
			Type:        schema.TypeString,
			Description: "Possible Values: `NOT_ENDS_WITH`, `NOT_EQUALS`, `REGEX_MATCHES`, `NOT_CONTAINS`, `NOT_LOWER_THAN_OR_EQUAL`, `TAG_KEY_EQUALS`, `NOT_IS_IP_IN_RANGE`, `BEGINS_WITH`, `EXISTS`, `GREATER_THAN_OR_EQUAL`, `EQUALS`, `NOT_TAG_KEY_EQUALS`, `NOT_EXISTS`, `NOT_GREATER_THAN`, `LOWER_THAN_OR_EQUAL`, `NOT_GREATER_THAN_OR_EQUAL`, `CONTAINS`, `GREATER_THAN`, `LOWER_THAN`, `NOT_LOWER_THAN`, `NOT_BEGINS_WITH`, `ENDS_WITH`, `IS_IP_IN_RANGE`, `NOT_REGEX_MATCHES`",
			Required:    true,
		},
		"string_value": {
			Type:        schema.TypeString,
			Description: "Value",
			Optional:    true,
		},
		"tag": {
			Type:        schema.TypeString,
			Description: "Format: `[CONTEXT]tagKey:tagValue`",
			Optional:    true,
		},
	}
}

func (me *AttributeCondition) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"case_sensitive":     me.CaseSensitive,
		"dynamic_key":        me.DynamicKey,
		"dynamic_key_source": me.DynamicKeySource,
		"entity_id":          me.EntityID,
		"enum_value":         me.EnumValue,
		"integer_value":      me.IntegerValue,
		"key":                me.Key,
		"operator":           me.Operator,
		"string_value":       me.StringValue,
		"tag":                me.Tag,
	})
}

func (me *AttributeCondition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"case_sensitive":     &me.CaseSensitive,
		"dynamic_key":        &me.DynamicKey,
		"dynamic_key_source": &me.DynamicKeySource,
		"entity_id":          &me.EntityID,
		"enum_value":         &me.EnumValue,
		"integer_value":      &me.IntegerValue,
		"key":                &me.Key,
		"operator":           &me.Operator,
		"string_value":       &me.StringValue,
		"tag":                &me.Tag,
	})
}
