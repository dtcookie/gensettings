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
	Key              Attribute `json:"key"`                        // Possible Values: `HOST_PAAS_MEMORY_LIMIT`, `HOST_AIX_SIMULTANEOUS_THREADS`, `HOST_ARCHITECTURE`, `WEB_APPLICATION_TAGS`, `SERVICE_CTG_SERVICE_NAME`, `ENTERPRISE_APPLICATION_NAME`, `HOST_CPU_CORES`, `ENTERPRISE_APPLICATION_DECODER_TYPE`, `AWS_RELATIONAL_DATABASE_SERVICE_TAGS`, `OPENSTACK_ACCOUNT_PROJECT_NAME`, `HOST_AZURE_SKU`, `SERVICE_WEB_CONTEXT_ROOT`, `PROCESS_GROUP_TAGS`, `HOST_TECHNOLOGY`, `EXTERNAL_MONITOR_TAGS`, `NAME_OF_COMPUTE_NODE`, `SERVICE_TECHNOLOGY_EDITION`, `OPENSTACK_REGION_NAME`, `PROCESS_GROUP_LISTEN_PORT`, `ESXI_HOST_TAGS`, `HOST_KUBERNETES_LABELS`, `CUSTOM_APPLICATION_NAME`, `DATA_CENTER_SERVICE_TAGS`, `ESXI_HOST_PRODUCT_VERSION`, `KUBERNETES_CLUSTER_NAME`, `QUEUE_TECHNOLOGY`, `AWS_ACCOUNT_NAME`, `KUBERNETES_NODE_NAME`, `VMWARE_DATACENTER_NAME`, `EC2_INSTANCE_AWS_SECURITY_GROUP`, `GOOGLE_COMPUTE_INSTANCE_PROJECT`, `EXTERNAL_MONITOR_ENGINE_NAME`, `AWS_CLASSIC_LOAD_BALANCER_TAGS`, `SERVICE_IBM_CTG_GATEWAY_URL`, `AZURE_MGMT_GROUP_UUID`, `OPENSTACK_AVAILABILITY_ZONE_NAME`, `EC2_INSTANCE_PRIVATE_HOST_NAME`, `PROCESS_GROUP_AZURE_SITE_NAME`, `HOST_HYPERVISOR_TYPE`, `HTTP_MONITOR_TAGS`, `OPENSTACK_VM_SECURITY_GROUP`, `HOST_PAAS_TYPE`, `WEB_APPLICATION_NAME`, `SERVICE_REMOTE_ENDPOINT`, `CUSTOM_DEVICE_GROUP_NAME`, `SERVICE_TECHNOLOGY_VERSION`, `EXTERNAL_MONITOR_ENGINE_TYPE`, `GOOGLE_COMPUTE_INSTANCE_ID`, `GOOGLE_COMPUTE_INSTANCE_MACHINE_TYPE`, `SERVICE_DATABASE_VENDOR`, `CUSTOM_DEVICE_DNS_ADDRESS`, `CUSTOM_DEVICE_NAME`, `AWS_AUTO_SCALING_GROUP_TAGS`, `DOCKER_FULL_IMAGE_NAME`, `SERVICE_DATABASE_NAME`, `CLOUD_FOUNDRY_ORG_NAME`, `HOST_AZURE_WEB_APPLICATION_SITE_NAMES`, `BROWSER_MONITOR_NAME`, `PROCESS_GROUP_DETECTED_NAME`, `CLOUD_APPLICATION_NAME`, `APPMON_SYSTEM_PROFILE_NAME`, `PROCESS_GROUP_TECHNOLOGY`, `AWS_RELATIONAL_DATABASE_SERVICE_PORT`, `HOST_BITNESS`, `GOOGLE_CLOUD_PLATFORM_ZONE_NAME`, `MOBILE_APPLICATION_TAGS`, `DATA_CENTER_SERVICE_DECODER_TYPE`, `EC2_INSTANCE_ID`, `CLOUD_APPLICATION_NAMESPACE_NAME`, `HOST_OS_VERSION`, `SERVICE_PUBLIC_DOMAIN_NAME`, `EXTERNAL_MONITOR_NAME`, `SERVICE_MESSAGING_LISTENER_CLASS_NAME`, `DATA_CENTER_SERVICE_IP_ADDRESS`, `ENTERPRISE_APPLICATION_IP_ADDRESS`, `HOST_CUSTOM_METADATA`, `GOOGLE_COMPUTE_INSTANCE_NAME`, `GOOGLE_COMPUTE_INSTANCE_PUBLIC_IP_ADDRESSES`, `AWS_RELATIONAL_DATABASE_SERVICE_ENDPOINT`, `AWS_CLASSIC_LOAD_BALANCER_NAME`, `ESXI_HOST_PRODUCT_NAME`, `EC2_INSTANCE_BEANSTALK_ENV_NAME`, `HOST_ONEAGENT_CUSTOM_HOST_NAME`, `SERVICE_WEB_SERVER_ENDPOINT`, `ESXI_HOST_HARDWARE_MODEL`, `BROWSER_MONITOR_TAGS`, `ENTERPRISE_APPLICATION_METADATA`, `HOST_BOSH_DEPLOYMENT_ID`, `KUBERNETES_SERVICE_NAME`, `SERVICE_AKKA_ACTOR_SYSTEM`, `EC2_INSTANCE_NAME`, `SERVICE_REMOTE_SERVICE_NAME`, `CLOUD_FOUNDRY_FOUNDATION_NAME`, `EXTERNAL_MONITOR_ENGINE_DESCRIPTION`, `ESXI_HOST_CLUSTER_NAME`, `HOST_BOSH_STEMCELL_VERSION`, `GOOGLE_COMPUTE_INSTANCE_PROJECT_ID`, `AWS_CLASSIC_LOAD_BALANCER_FRONTEND_PORTS`, `CUSTOM_APPLICATION_TYPE`, `AZURE_ENTITY_NAME`, `SERVICE_DATABASE_HOST_NAME`, `EC2_INSTANCE_PUBLIC_HOST_NAME`, `HOST_AWS_NAME_TAG`, `GEOLOCATION_SITE_NAME`, `AZURE_ENTITY_TAGS`, `SERVICE_ESB_APPLICATION_NAME`, `HOST_BOSH_INSTANCE_NAME`, `ESXI_HOST_NAME`, `HOST_LOGICAL_CPU_CORES`, `AWS_AUTO_SCALING_GROUP_NAME`, `HOST_AZURE_WEB_APPLICATION_HOST_NAMES`, `HOST_AIX_VIRTUAL_CPU_COUNT`, `CUSTOM_APPLICATION_PLATFORM`, `ESXI_HOST_HARDWARE_VENDOR`, `CUSTOM_DEVICE_IP_ADDRESS`, `ENTERPRISE_APPLICATION_TAGS`, `AWS_RELATIONAL_DATABASE_SERVICE_ENGINE`, `AWS_NETWORK_LOAD_BALANCER_NAME`, `ENTERPRISE_APPLICATION_PORT`, `HOST_AZURE_COMPUTE_MODE`, `SERVICE_TAGS`, `DATA_CENTER_SERVICE_METADATA`, `HOST_BOSH_AVAILABILITY_ZONE`, `CLOUD_APPLICATION_LABELS`, `HTTP_MONITOR_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_INSTANCE_CLASS`, `QUEUE_NAME`, `AWS_APPLICATION_LOAD_BALANCER_NAME`, `OPENSTACK_VM_INSTANCE_TYPE`, `AWS_APPLICATION_LOAD_BALANCER_TAGS`, `QUEUE_VENDOR`, `PROCESS_GROUP_PREDEFINED_METADATA`, `PROCESS_GROUP_CUSTOM_METADATA`, `OPENSTACK_PROJECT_NAME`, `AZURE_VM_NAME`, `SERVICE_WEB_SERVICE_NAMESPACE`, `PROCESS_GROUP_AZURE_HOST_NAME`, `SERVICE_TECHNOLOGY`, `EC2_INSTANCE_TAGS`, `EC2_INSTANCE_AMI_ID`, `WEB_APPLICATION_NAME_PATTERN`, `AZURE_SCALE_SET_NAME`, `DATA_CENTER_SERVICE_NAME`, `CUSTOM_DEVICE_METADATA`, `HOST_GROUP_ID`, `SERVICE_TOPOLOGY`, `HOST_DETECTED_NAME`, `AZURE_REGION_NAME`, `OPENSTACK_ACCOUNT_NAME`, `HOST_BOSH_NAME`, `AZURE_TENANT_NAME`, `VMWARE_VM_NAME`, `DATA_CENTER_SERVICE_PORT`, `SERVICE_DATABASE_TOPOLOGY`, `AZURE_SUBSCRIPTION_NAME`, `HOST_TAGS`, `HOST_CLOUD_TYPE`, `PROCESS_GROUP_NAME`, `AWS_AVAILABILITY_ZONE_NAME`, `CUSTOM_DEVICE_TECHNOLOGY`, `SERVICE_WEB_SERVER_NAME`, `SERVICE_PORT`, `DOCKER_CONTAINER_NAME`, `SERVICE_WEB_SERVICE_NAME`, `SERVICE_DETECTED_NAME`, `AZURE_TENANT_UUID`, `AWS_ACCOUNT_ID`, `AWS_NETWORK_LOAD_BALANCER_TAGS`, `CUSTOM_DEVICE_PORT`, `HOST_NAME`, `HOST_IP_ADDRESS`, `PROCESS_GROUP_TECHNOLOGY_VERSION`, `EC2_INSTANCE_AWS_INSTANCE_TYPE`, `WEB_APPLICATION_TYPE`, `CUSTOM_DEVICE_GROUP_TAGS`, `AZURE_MGMT_GROUP_NAME`, `CLOUD_APPLICATION_NAMESPACE_LABELS`, `APPMON_SERVER_NAME`, `HOST_GROUP_NAME`, `CUSTOM_APPLICATION_TAGS`, `CUSTOM_DEVICE_TAGS`, `SERVICE_TYPE`, `HOST_BOSH_INSTANCE_ID`, `SERVICE_NAME`, `OPENSTACK_VM_NAME`, `SERVICE_WEB_APPLICATION_ID`, `HOST_OS_TYPE`, `HOST_AIX_LOGICAL_CPU_COUNT`, `AZURE_SUBSCRIPTION_UUID`, `DOCKER_IMAGE_VERSION`, `MOBILE_APPLICATION_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_DB_NAME`, `PROCESS_GROUP_ID`, `PROCESS_GROUP_TECHNOLOGY_EDITION`, `MOBILE_APPLICATION_PLATFORM`
	Operator         Operator  `json:"operator"`                   // Possible Values: `IS_IP_IN_RANGE`, `LOWER_THAN_OR_EQUAL`, `NOT_REGEX_MATCHES`, `EXISTS`, `TAG_KEY_EQUALS`, `ENDS_WITH`, `GREATER_THAN_OR_EQUAL`, `NOT_EQUALS`, `NOT_CONTAINS`, `LOWER_THAN`, `NOT_ENDS_WITH`, `GREATER_THAN`, `NOT_EXISTS`, `NOT_IS_IP_IN_RANGE`, `NOT_LOWER_THAN`, `NOT_LOWER_THAN_OR_EQUAL`, `REGEX_MATCHES`, `EQUALS`, `CONTAINS`, `NOT_GREATER_THAN`, `BEGINS_WITH`, `NOT_GREATER_THAN_OR_EQUAL`, `NOT_TAG_KEY_EQUALS`, `NOT_BEGINS_WITH`
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
			Description: "Possible Values: `HOST_PAAS_MEMORY_LIMIT`, `HOST_AIX_SIMULTANEOUS_THREADS`, `HOST_ARCHITECTURE`, `WEB_APPLICATION_TAGS`, `SERVICE_CTG_SERVICE_NAME`, `ENTERPRISE_APPLICATION_NAME`, `HOST_CPU_CORES`, `ENTERPRISE_APPLICATION_DECODER_TYPE`, `AWS_RELATIONAL_DATABASE_SERVICE_TAGS`, `OPENSTACK_ACCOUNT_PROJECT_NAME`, `HOST_AZURE_SKU`, `SERVICE_WEB_CONTEXT_ROOT`, `PROCESS_GROUP_TAGS`, `HOST_TECHNOLOGY`, `EXTERNAL_MONITOR_TAGS`, `NAME_OF_COMPUTE_NODE`, `SERVICE_TECHNOLOGY_EDITION`, `OPENSTACK_REGION_NAME`, `PROCESS_GROUP_LISTEN_PORT`, `ESXI_HOST_TAGS`, `HOST_KUBERNETES_LABELS`, `CUSTOM_APPLICATION_NAME`, `DATA_CENTER_SERVICE_TAGS`, `ESXI_HOST_PRODUCT_VERSION`, `KUBERNETES_CLUSTER_NAME`, `QUEUE_TECHNOLOGY`, `AWS_ACCOUNT_NAME`, `KUBERNETES_NODE_NAME`, `VMWARE_DATACENTER_NAME`, `EC2_INSTANCE_AWS_SECURITY_GROUP`, `GOOGLE_COMPUTE_INSTANCE_PROJECT`, `EXTERNAL_MONITOR_ENGINE_NAME`, `AWS_CLASSIC_LOAD_BALANCER_TAGS`, `SERVICE_IBM_CTG_GATEWAY_URL`, `AZURE_MGMT_GROUP_UUID`, `OPENSTACK_AVAILABILITY_ZONE_NAME`, `EC2_INSTANCE_PRIVATE_HOST_NAME`, `PROCESS_GROUP_AZURE_SITE_NAME`, `HOST_HYPERVISOR_TYPE`, `HTTP_MONITOR_TAGS`, `OPENSTACK_VM_SECURITY_GROUP`, `HOST_PAAS_TYPE`, `WEB_APPLICATION_NAME`, `SERVICE_REMOTE_ENDPOINT`, `CUSTOM_DEVICE_GROUP_NAME`, `SERVICE_TECHNOLOGY_VERSION`, `EXTERNAL_MONITOR_ENGINE_TYPE`, `GOOGLE_COMPUTE_INSTANCE_ID`, `GOOGLE_COMPUTE_INSTANCE_MACHINE_TYPE`, `SERVICE_DATABASE_VENDOR`, `CUSTOM_DEVICE_DNS_ADDRESS`, `CUSTOM_DEVICE_NAME`, `AWS_AUTO_SCALING_GROUP_TAGS`, `DOCKER_FULL_IMAGE_NAME`, `SERVICE_DATABASE_NAME`, `CLOUD_FOUNDRY_ORG_NAME`, `HOST_AZURE_WEB_APPLICATION_SITE_NAMES`, `BROWSER_MONITOR_NAME`, `PROCESS_GROUP_DETECTED_NAME`, `CLOUD_APPLICATION_NAME`, `APPMON_SYSTEM_PROFILE_NAME`, `PROCESS_GROUP_TECHNOLOGY`, `AWS_RELATIONAL_DATABASE_SERVICE_PORT`, `HOST_BITNESS`, `GOOGLE_CLOUD_PLATFORM_ZONE_NAME`, `MOBILE_APPLICATION_TAGS`, `DATA_CENTER_SERVICE_DECODER_TYPE`, `EC2_INSTANCE_ID`, `CLOUD_APPLICATION_NAMESPACE_NAME`, `HOST_OS_VERSION`, `SERVICE_PUBLIC_DOMAIN_NAME`, `EXTERNAL_MONITOR_NAME`, `SERVICE_MESSAGING_LISTENER_CLASS_NAME`, `DATA_CENTER_SERVICE_IP_ADDRESS`, `ENTERPRISE_APPLICATION_IP_ADDRESS`, `HOST_CUSTOM_METADATA`, `GOOGLE_COMPUTE_INSTANCE_NAME`, `GOOGLE_COMPUTE_INSTANCE_PUBLIC_IP_ADDRESSES`, `AWS_RELATIONAL_DATABASE_SERVICE_ENDPOINT`, `AWS_CLASSIC_LOAD_BALANCER_NAME`, `ESXI_HOST_PRODUCT_NAME`, `EC2_INSTANCE_BEANSTALK_ENV_NAME`, `HOST_ONEAGENT_CUSTOM_HOST_NAME`, `SERVICE_WEB_SERVER_ENDPOINT`, `ESXI_HOST_HARDWARE_MODEL`, `BROWSER_MONITOR_TAGS`, `ENTERPRISE_APPLICATION_METADATA`, `HOST_BOSH_DEPLOYMENT_ID`, `KUBERNETES_SERVICE_NAME`, `SERVICE_AKKA_ACTOR_SYSTEM`, `EC2_INSTANCE_NAME`, `SERVICE_REMOTE_SERVICE_NAME`, `CLOUD_FOUNDRY_FOUNDATION_NAME`, `EXTERNAL_MONITOR_ENGINE_DESCRIPTION`, `ESXI_HOST_CLUSTER_NAME`, `HOST_BOSH_STEMCELL_VERSION`, `GOOGLE_COMPUTE_INSTANCE_PROJECT_ID`, `AWS_CLASSIC_LOAD_BALANCER_FRONTEND_PORTS`, `CUSTOM_APPLICATION_TYPE`, `AZURE_ENTITY_NAME`, `SERVICE_DATABASE_HOST_NAME`, `EC2_INSTANCE_PUBLIC_HOST_NAME`, `HOST_AWS_NAME_TAG`, `GEOLOCATION_SITE_NAME`, `AZURE_ENTITY_TAGS`, `SERVICE_ESB_APPLICATION_NAME`, `HOST_BOSH_INSTANCE_NAME`, `ESXI_HOST_NAME`, `HOST_LOGICAL_CPU_CORES`, `AWS_AUTO_SCALING_GROUP_NAME`, `HOST_AZURE_WEB_APPLICATION_HOST_NAMES`, `HOST_AIX_VIRTUAL_CPU_COUNT`, `CUSTOM_APPLICATION_PLATFORM`, `ESXI_HOST_HARDWARE_VENDOR`, `CUSTOM_DEVICE_IP_ADDRESS`, `ENTERPRISE_APPLICATION_TAGS`, `AWS_RELATIONAL_DATABASE_SERVICE_ENGINE`, `AWS_NETWORK_LOAD_BALANCER_NAME`, `ENTERPRISE_APPLICATION_PORT`, `HOST_AZURE_COMPUTE_MODE`, `SERVICE_TAGS`, `DATA_CENTER_SERVICE_METADATA`, `HOST_BOSH_AVAILABILITY_ZONE`, `CLOUD_APPLICATION_LABELS`, `HTTP_MONITOR_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_INSTANCE_CLASS`, `QUEUE_NAME`, `AWS_APPLICATION_LOAD_BALANCER_NAME`, `OPENSTACK_VM_INSTANCE_TYPE`, `AWS_APPLICATION_LOAD_BALANCER_TAGS`, `QUEUE_VENDOR`, `PROCESS_GROUP_PREDEFINED_METADATA`, `PROCESS_GROUP_CUSTOM_METADATA`, `OPENSTACK_PROJECT_NAME`, `AZURE_VM_NAME`, `SERVICE_WEB_SERVICE_NAMESPACE`, `PROCESS_GROUP_AZURE_HOST_NAME`, `SERVICE_TECHNOLOGY`, `EC2_INSTANCE_TAGS`, `EC2_INSTANCE_AMI_ID`, `WEB_APPLICATION_NAME_PATTERN`, `AZURE_SCALE_SET_NAME`, `DATA_CENTER_SERVICE_NAME`, `CUSTOM_DEVICE_METADATA`, `HOST_GROUP_ID`, `SERVICE_TOPOLOGY`, `HOST_DETECTED_NAME`, `AZURE_REGION_NAME`, `OPENSTACK_ACCOUNT_NAME`, `HOST_BOSH_NAME`, `AZURE_TENANT_NAME`, `VMWARE_VM_NAME`, `DATA_CENTER_SERVICE_PORT`, `SERVICE_DATABASE_TOPOLOGY`, `AZURE_SUBSCRIPTION_NAME`, `HOST_TAGS`, `HOST_CLOUD_TYPE`, `PROCESS_GROUP_NAME`, `AWS_AVAILABILITY_ZONE_NAME`, `CUSTOM_DEVICE_TECHNOLOGY`, `SERVICE_WEB_SERVER_NAME`, `SERVICE_PORT`, `DOCKER_CONTAINER_NAME`, `SERVICE_WEB_SERVICE_NAME`, `SERVICE_DETECTED_NAME`, `AZURE_TENANT_UUID`, `AWS_ACCOUNT_ID`, `AWS_NETWORK_LOAD_BALANCER_TAGS`, `CUSTOM_DEVICE_PORT`, `HOST_NAME`, `HOST_IP_ADDRESS`, `PROCESS_GROUP_TECHNOLOGY_VERSION`, `EC2_INSTANCE_AWS_INSTANCE_TYPE`, `WEB_APPLICATION_TYPE`, `CUSTOM_DEVICE_GROUP_TAGS`, `AZURE_MGMT_GROUP_NAME`, `CLOUD_APPLICATION_NAMESPACE_LABELS`, `APPMON_SERVER_NAME`, `HOST_GROUP_NAME`, `CUSTOM_APPLICATION_TAGS`, `CUSTOM_DEVICE_TAGS`, `SERVICE_TYPE`, `HOST_BOSH_INSTANCE_ID`, `SERVICE_NAME`, `OPENSTACK_VM_NAME`, `SERVICE_WEB_APPLICATION_ID`, `HOST_OS_TYPE`, `HOST_AIX_LOGICAL_CPU_COUNT`, `AZURE_SUBSCRIPTION_UUID`, `DOCKER_IMAGE_VERSION`, `MOBILE_APPLICATION_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_DB_NAME`, `PROCESS_GROUP_ID`, `PROCESS_GROUP_TECHNOLOGY_EDITION`, `MOBILE_APPLICATION_PLATFORM`",
			Required:    true,
		},
		"operator": {
			Type:        schema.TypeString,
			Description: "Possible Values: `IS_IP_IN_RANGE`, `LOWER_THAN_OR_EQUAL`, `NOT_REGEX_MATCHES`, `EXISTS`, `TAG_KEY_EQUALS`, `ENDS_WITH`, `GREATER_THAN_OR_EQUAL`, `NOT_EQUALS`, `NOT_CONTAINS`, `LOWER_THAN`, `NOT_ENDS_WITH`, `GREATER_THAN`, `NOT_EXISTS`, `NOT_IS_IP_IN_RANGE`, `NOT_LOWER_THAN`, `NOT_LOWER_THAN_OR_EQUAL`, `REGEX_MATCHES`, `EQUALS`, `CONTAINS`, `NOT_GREATER_THAN`, `BEGINS_WITH`, `NOT_GREATER_THAN_OR_EQUAL`, `NOT_TAG_KEY_EQUALS`, `NOT_BEGINS_WITH`",
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
