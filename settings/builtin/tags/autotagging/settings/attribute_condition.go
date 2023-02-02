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
	Key              Attribute `json:"key"`                        // Possible Values: `OPENSTACK_ACCOUNT_NAME`, `KUBERNETES_SERVICE_NAME`, `GOOGLE_COMPUTE_INSTANCE_PUBLIC_IP_ADDRESSES`, `SERVICE_CTG_SERVICE_NAME`, `CUSTOM_APPLICATION_PLATFORM`, `PROCESS_GROUP_DETECTED_NAME`, `NAME_OF_COMPUTE_NODE`, `PROCESS_GROUP_NAME`, `DATA_CENTER_SERVICE_METADATA`, `DATA_CENTER_SERVICE_TAGS`, `BROWSER_MONITOR_NAME`, `HOST_GROUP_NAME`, `CLOUD_APPLICATION_NAMESPACE_NAME`, `HOST_AZURE_WEB_APPLICATION_HOST_NAMES`, `HOST_CLOUD_TYPE`, `QUEUE_NAME`, `APPMON_SYSTEM_PROFILE_NAME`, `GOOGLE_COMPUTE_INSTANCE_PROJECT`, `HOST_AZURE_SKU`, `HOST_OS_TYPE`, `SERVICE_DATABASE_NAME`, `SERVICE_PUBLIC_DOMAIN_NAME`, `CUSTOM_APPLICATION_TAGS`, `AWS_RELATIONAL_DATABASE_SERVICE_PORT`, `PROCESS_GROUP_ID`, `HOST_BOSH_INSTANCE_ID`, `EXTERNAL_MONITOR_NAME`, `CUSTOM_DEVICE_GROUP_TAGS`, `AZURE_MGMT_GROUP_NAME`, `AWS_ACCOUNT_NAME`, `EXTERNAL_MONITOR_ENGINE_DESCRIPTION`, `HOST_ARCHITECTURE`, `PROCESS_GROUP_AZURE_SITE_NAME`, `DATA_CENTER_SERVICE_NAME`, `SERVICE_DATABASE_HOST_NAME`, `AWS_APPLICATION_LOAD_BALANCER_TAGS`, `DOCKER_CONTAINER_NAME`, `OPENSTACK_PROJECT_NAME`, `CLOUD_FOUNDRY_ORG_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_ENDPOINT`, `VMWARE_VM_NAME`, `AZURE_REGION_NAME`, `AWS_CLASSIC_LOAD_BALANCER_TAGS`, `KUBERNETES_NODE_NAME`, `ENTERPRISE_APPLICATION_METADATA`, `SERVICE_TOPOLOGY`, `AWS_NETWORK_LOAD_BALANCER_TAGS`, `AZURE_ENTITY_TAGS`, `SERVICE_WEB_SERVER_NAME`, `OPENSTACK_VM_SECURITY_GROUP`, `EC2_INSTANCE_BEANSTALK_ENV_NAME`, `CUSTOM_DEVICE_TECHNOLOGY`, `SERVICE_WEB_APPLICATION_ID`, `AWS_RELATIONAL_DATABASE_SERVICE_TAGS`, `PROCESS_GROUP_PREDEFINED_METADATA`, `MOBILE_APPLICATION_NAME`, `APPMON_SERVER_NAME`, `EC2_INSTANCE_AWS_INSTANCE_TYPE`, `AWS_APPLICATION_LOAD_BALANCER_NAME`, `WEB_APPLICATION_TAGS`, `ESXI_HOST_PRODUCT_VERSION`, `ENTERPRISE_APPLICATION_IP_ADDRESS`, `HOST_ONEAGENT_CUSTOM_HOST_NAME`, `SERVICE_MESSAGING_LISTENER_CLASS_NAME`, `SERVICE_WEB_SERVICE_NAMESPACE`, `HOST_HYPERVISOR_TYPE`, `DOCKER_FULL_IMAGE_NAME`, `HOST_CPU_CORES`, `AWS_RELATIONAL_DATABASE_SERVICE_ENGINE`, `HOST_AIX_SIMULTANEOUS_THREADS`, `AZURE_SUBSCRIPTION_UUID`, `CUSTOM_DEVICE_TAGS`, `CUSTOM_DEVICE_GROUP_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_INSTANCE_CLASS`, `DATA_CENTER_SERVICE_DECODER_TYPE`, `WEB_APPLICATION_NAME_PATTERN`, `AZURE_MGMT_GROUP_UUID`, `HOST_NAME`, `OPENSTACK_AVAILABILITY_ZONE_NAME`, `PROCESS_GROUP_TAGS`, `VMWARE_DATACENTER_NAME`, `SERVICE_TECHNOLOGY_VERSION`, `AZURE_TENANT_NAME`, `SERVICE_TYPE`, `DATA_CENTER_SERVICE_IP_ADDRESS`, `SERVICE_REMOTE_ENDPOINT`, `PROCESS_GROUP_TECHNOLOGY_VERSION`, `CUSTOM_APPLICATION_TYPE`, `PROCESS_GROUP_AZURE_HOST_NAME`, `QUEUE_TECHNOLOGY`, `HOST_PAAS_TYPE`, `CLOUD_APPLICATION_LABELS`, `ENTERPRISE_APPLICATION_NAME`, `QUEUE_VENDOR`, `HOST_AIX_VIRTUAL_CPU_COUNT`, `DOCKER_IMAGE_VERSION`, `HOST_PAAS_MEMORY_LIMIT`, `SERVICE_DETECTED_NAME`, `PROCESS_GROUP_TECHNOLOGY`, `ESXI_HOST_TAGS`, `SERVICE_WEB_CONTEXT_ROOT`, `HOST_AWS_NAME_TAG`, `WEB_APPLICATION_NAME`, `ESXI_HOST_CLUSTER_NAME`, `ESXI_HOST_HARDWARE_MODEL`, `BROWSER_MONITOR_TAGS`, `CLOUD_FOUNDRY_FOUNDATION_NAME`, `HOST_AZURE_WEB_APPLICATION_SITE_NAMES`, `SERVICE_AKKA_ACTOR_SYSTEM`, `HOST_CUSTOM_METADATA`, `MOBILE_APPLICATION_TAGS`, `AWS_ACCOUNT_ID`, `AZURE_VM_NAME`, `AWS_AVAILABILITY_ZONE_NAME`, `SERVICE_TECHNOLOGY`, `EC2_INSTANCE_ID`, `CUSTOM_DEVICE_DNS_ADDRESS`, `HOST_OS_VERSION`, `SERVICE_DATABASE_TOPOLOGY`, `HTTP_MONITOR_NAME`, `GOOGLE_COMPUTE_INSTANCE_ID`, `ESXI_HOST_HARDWARE_VENDOR`, `HOST_BOSH_INSTANCE_NAME`, `PROCESS_GROUP_LISTEN_PORT`, `OPENSTACK_VM_NAME`, `CUSTOM_DEVICE_METADATA`, `HOST_KUBERNETES_LABELS`, `CLOUD_APPLICATION_NAMESPACE_LABELS`, `ENTERPRISE_APPLICATION_TAGS`, `OPENSTACK_REGION_NAME`, `HOST_AIX_LOGICAL_CPU_COUNT`, `DATA_CENTER_SERVICE_PORT`, `GOOGLE_COMPUTE_INSTANCE_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_DB_NAME`, `HOST_DETECTED_NAME`, `GOOGLE_COMPUTE_INSTANCE_MACHINE_TYPE`, `AWS_AUTO_SCALING_GROUP_TAGS`, `AZURE_ENTITY_NAME`, `SERVICE_WEB_SERVICE_NAME`, `HOST_TAGS`, `AWS_AUTO_SCALING_GROUP_NAME`, `ESXI_HOST_NAME`, `GEOLOCATION_SITE_NAME`, `SERVICE_TAGS`, `HOST_BOSH_NAME`, `CUSTOM_DEVICE_PORT`, `HOST_AZURE_COMPUTE_MODE`, `CLOUD_APPLICATION_NAME`, `AWS_NETWORK_LOAD_BALANCER_NAME`, `ENTERPRISE_APPLICATION_DECODER_TYPE`, `SERVICE_DATABASE_VENDOR`, `AWS_CLASSIC_LOAD_BALANCER_NAME`, `SERVICE_WEB_SERVER_ENDPOINT`, `HOST_GROUP_ID`, `EC2_INSTANCE_NAME`, `EC2_INSTANCE_PUBLIC_HOST_NAME`, `EC2_INSTANCE_AMI_ID`, `AWS_CLASSIC_LOAD_BALANCER_FRONTEND_PORTS`, `CUSTOM_DEVICE_IP_ADDRESS`, `SERVICE_NAME`, `HOST_BITNESS`, `SERVICE_ESB_APPLICATION_NAME`, `SERVICE_TECHNOLOGY_EDITION`, `EXTERNAL_MONITOR_TAGS`, `CUSTOM_DEVICE_NAME`, `CUSTOM_APPLICATION_NAME`, `HTTP_MONITOR_TAGS`, `HOST_TECHNOLOGY`, `EC2_INSTANCE_TAGS`, `GOOGLE_COMPUTE_INSTANCE_PROJECT_ID`, `OPENSTACK_ACCOUNT_PROJECT_NAME`, `PROCESS_GROUP_TECHNOLOGY_EDITION`, `HOST_BOSH_AVAILABILITY_ZONE`, `EC2_INSTANCE_AWS_SECURITY_GROUP`, `HOST_BOSH_DEPLOYMENT_ID`, `OPENSTACK_VM_INSTANCE_TYPE`, `ENTERPRISE_APPLICATION_PORT`, `SERVICE_REMOTE_SERVICE_NAME`, `HOST_LOGICAL_CPU_CORES`, `HOST_BOSH_STEMCELL_VERSION`, `AZURE_TENANT_UUID`, `EXTERNAL_MONITOR_ENGINE_NAME`, `KUBERNETES_CLUSTER_NAME`, `AZURE_SUBSCRIPTION_NAME`, `SERVICE_IBM_CTG_GATEWAY_URL`, `GOOGLE_CLOUD_PLATFORM_ZONE_NAME`, `PROCESS_GROUP_CUSTOM_METADATA`, `EC2_INSTANCE_PRIVATE_HOST_NAME`, `MOBILE_APPLICATION_PLATFORM`, `ESXI_HOST_PRODUCT_NAME`, `HOST_IP_ADDRESS`, `AWS_RELATIONAL_DATABASE_SERVICE_NAME`, `WEB_APPLICATION_TYPE`, `SERVICE_PORT`, `EXTERNAL_MONITOR_ENGINE_TYPE`, `AZURE_SCALE_SET_NAME`
	Operator         Operator  `json:"operator"`                   // Possible Values: `NOT_LOWER_THAN_OR_EQUAL`, `NOT_CONTAINS`, `NOT_EQUALS`, `NOT_LOWER_THAN`, `GREATER_THAN_OR_EQUAL`, `LOWER_THAN`, `IS_IP_IN_RANGE`, `NOT_GREATER_THAN_OR_EQUAL`, `GREATER_THAN`, `NOT_REGEX_MATCHES`, `TAG_KEY_EQUALS`, `NOT_ENDS_WITH`, `NOT_GREATER_THAN`, `NOT_EXISTS`, `LOWER_THAN_OR_EQUAL`, `EXISTS`, `BEGINS_WITH`, `NOT_IS_IP_IN_RANGE`, `REGEX_MATCHES`, `CONTAINS`, `ENDS_WITH`, `NOT_BEGINS_WITH`, `NOT_TAG_KEY_EQUALS`, `EQUALS`
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
			Description: "Possible Values: `OPENSTACK_ACCOUNT_NAME`, `KUBERNETES_SERVICE_NAME`, `GOOGLE_COMPUTE_INSTANCE_PUBLIC_IP_ADDRESSES`, `SERVICE_CTG_SERVICE_NAME`, `CUSTOM_APPLICATION_PLATFORM`, `PROCESS_GROUP_DETECTED_NAME`, `NAME_OF_COMPUTE_NODE`, `PROCESS_GROUP_NAME`, `DATA_CENTER_SERVICE_METADATA`, `DATA_CENTER_SERVICE_TAGS`, `BROWSER_MONITOR_NAME`, `HOST_GROUP_NAME`, `CLOUD_APPLICATION_NAMESPACE_NAME`, `HOST_AZURE_WEB_APPLICATION_HOST_NAMES`, `HOST_CLOUD_TYPE`, `QUEUE_NAME`, `APPMON_SYSTEM_PROFILE_NAME`, `GOOGLE_COMPUTE_INSTANCE_PROJECT`, `HOST_AZURE_SKU`, `HOST_OS_TYPE`, `SERVICE_DATABASE_NAME`, `SERVICE_PUBLIC_DOMAIN_NAME`, `CUSTOM_APPLICATION_TAGS`, `AWS_RELATIONAL_DATABASE_SERVICE_PORT`, `PROCESS_GROUP_ID`, `HOST_BOSH_INSTANCE_ID`, `EXTERNAL_MONITOR_NAME`, `CUSTOM_DEVICE_GROUP_TAGS`, `AZURE_MGMT_GROUP_NAME`, `AWS_ACCOUNT_NAME`, `EXTERNAL_MONITOR_ENGINE_DESCRIPTION`, `HOST_ARCHITECTURE`, `PROCESS_GROUP_AZURE_SITE_NAME`, `DATA_CENTER_SERVICE_NAME`, `SERVICE_DATABASE_HOST_NAME`, `AWS_APPLICATION_LOAD_BALANCER_TAGS`, `DOCKER_CONTAINER_NAME`, `OPENSTACK_PROJECT_NAME`, `CLOUD_FOUNDRY_ORG_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_ENDPOINT`, `VMWARE_VM_NAME`, `AZURE_REGION_NAME`, `AWS_CLASSIC_LOAD_BALANCER_TAGS`, `KUBERNETES_NODE_NAME`, `ENTERPRISE_APPLICATION_METADATA`, `SERVICE_TOPOLOGY`, `AWS_NETWORK_LOAD_BALANCER_TAGS`, `AZURE_ENTITY_TAGS`, `SERVICE_WEB_SERVER_NAME`, `OPENSTACK_VM_SECURITY_GROUP`, `EC2_INSTANCE_BEANSTALK_ENV_NAME`, `CUSTOM_DEVICE_TECHNOLOGY`, `SERVICE_WEB_APPLICATION_ID`, `AWS_RELATIONAL_DATABASE_SERVICE_TAGS`, `PROCESS_GROUP_PREDEFINED_METADATA`, `MOBILE_APPLICATION_NAME`, `APPMON_SERVER_NAME`, `EC2_INSTANCE_AWS_INSTANCE_TYPE`, `AWS_APPLICATION_LOAD_BALANCER_NAME`, `WEB_APPLICATION_TAGS`, `ESXI_HOST_PRODUCT_VERSION`, `ENTERPRISE_APPLICATION_IP_ADDRESS`, `HOST_ONEAGENT_CUSTOM_HOST_NAME`, `SERVICE_MESSAGING_LISTENER_CLASS_NAME`, `SERVICE_WEB_SERVICE_NAMESPACE`, `HOST_HYPERVISOR_TYPE`, `DOCKER_FULL_IMAGE_NAME`, `HOST_CPU_CORES`, `AWS_RELATIONAL_DATABASE_SERVICE_ENGINE`, `HOST_AIX_SIMULTANEOUS_THREADS`, `AZURE_SUBSCRIPTION_UUID`, `CUSTOM_DEVICE_TAGS`, `CUSTOM_DEVICE_GROUP_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_INSTANCE_CLASS`, `DATA_CENTER_SERVICE_DECODER_TYPE`, `WEB_APPLICATION_NAME_PATTERN`, `AZURE_MGMT_GROUP_UUID`, `HOST_NAME`, `OPENSTACK_AVAILABILITY_ZONE_NAME`, `PROCESS_GROUP_TAGS`, `VMWARE_DATACENTER_NAME`, `SERVICE_TECHNOLOGY_VERSION`, `AZURE_TENANT_NAME`, `SERVICE_TYPE`, `DATA_CENTER_SERVICE_IP_ADDRESS`, `SERVICE_REMOTE_ENDPOINT`, `PROCESS_GROUP_TECHNOLOGY_VERSION`, `CUSTOM_APPLICATION_TYPE`, `PROCESS_GROUP_AZURE_HOST_NAME`, `QUEUE_TECHNOLOGY`, `HOST_PAAS_TYPE`, `CLOUD_APPLICATION_LABELS`, `ENTERPRISE_APPLICATION_NAME`, `QUEUE_VENDOR`, `HOST_AIX_VIRTUAL_CPU_COUNT`, `DOCKER_IMAGE_VERSION`, `HOST_PAAS_MEMORY_LIMIT`, `SERVICE_DETECTED_NAME`, `PROCESS_GROUP_TECHNOLOGY`, `ESXI_HOST_TAGS`, `SERVICE_WEB_CONTEXT_ROOT`, `HOST_AWS_NAME_TAG`, `WEB_APPLICATION_NAME`, `ESXI_HOST_CLUSTER_NAME`, `ESXI_HOST_HARDWARE_MODEL`, `BROWSER_MONITOR_TAGS`, `CLOUD_FOUNDRY_FOUNDATION_NAME`, `HOST_AZURE_WEB_APPLICATION_SITE_NAMES`, `SERVICE_AKKA_ACTOR_SYSTEM`, `HOST_CUSTOM_METADATA`, `MOBILE_APPLICATION_TAGS`, `AWS_ACCOUNT_ID`, `AZURE_VM_NAME`, `AWS_AVAILABILITY_ZONE_NAME`, `SERVICE_TECHNOLOGY`, `EC2_INSTANCE_ID`, `CUSTOM_DEVICE_DNS_ADDRESS`, `HOST_OS_VERSION`, `SERVICE_DATABASE_TOPOLOGY`, `HTTP_MONITOR_NAME`, `GOOGLE_COMPUTE_INSTANCE_ID`, `ESXI_HOST_HARDWARE_VENDOR`, `HOST_BOSH_INSTANCE_NAME`, `PROCESS_GROUP_LISTEN_PORT`, `OPENSTACK_VM_NAME`, `CUSTOM_DEVICE_METADATA`, `HOST_KUBERNETES_LABELS`, `CLOUD_APPLICATION_NAMESPACE_LABELS`, `ENTERPRISE_APPLICATION_TAGS`, `OPENSTACK_REGION_NAME`, `HOST_AIX_LOGICAL_CPU_COUNT`, `DATA_CENTER_SERVICE_PORT`, `GOOGLE_COMPUTE_INSTANCE_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_DB_NAME`, `HOST_DETECTED_NAME`, `GOOGLE_COMPUTE_INSTANCE_MACHINE_TYPE`, `AWS_AUTO_SCALING_GROUP_TAGS`, `AZURE_ENTITY_NAME`, `SERVICE_WEB_SERVICE_NAME`, `HOST_TAGS`, `AWS_AUTO_SCALING_GROUP_NAME`, `ESXI_HOST_NAME`, `GEOLOCATION_SITE_NAME`, `SERVICE_TAGS`, `HOST_BOSH_NAME`, `CUSTOM_DEVICE_PORT`, `HOST_AZURE_COMPUTE_MODE`, `CLOUD_APPLICATION_NAME`, `AWS_NETWORK_LOAD_BALANCER_NAME`, `ENTERPRISE_APPLICATION_DECODER_TYPE`, `SERVICE_DATABASE_VENDOR`, `AWS_CLASSIC_LOAD_BALANCER_NAME`, `SERVICE_WEB_SERVER_ENDPOINT`, `HOST_GROUP_ID`, `EC2_INSTANCE_NAME`, `EC2_INSTANCE_PUBLIC_HOST_NAME`, `EC2_INSTANCE_AMI_ID`, `AWS_CLASSIC_LOAD_BALANCER_FRONTEND_PORTS`, `CUSTOM_DEVICE_IP_ADDRESS`, `SERVICE_NAME`, `HOST_BITNESS`, `SERVICE_ESB_APPLICATION_NAME`, `SERVICE_TECHNOLOGY_EDITION`, `EXTERNAL_MONITOR_TAGS`, `CUSTOM_DEVICE_NAME`, `CUSTOM_APPLICATION_NAME`, `HTTP_MONITOR_TAGS`, `HOST_TECHNOLOGY`, `EC2_INSTANCE_TAGS`, `GOOGLE_COMPUTE_INSTANCE_PROJECT_ID`, `OPENSTACK_ACCOUNT_PROJECT_NAME`, `PROCESS_GROUP_TECHNOLOGY_EDITION`, `HOST_BOSH_AVAILABILITY_ZONE`, `EC2_INSTANCE_AWS_SECURITY_GROUP`, `HOST_BOSH_DEPLOYMENT_ID`, `OPENSTACK_VM_INSTANCE_TYPE`, `ENTERPRISE_APPLICATION_PORT`, `SERVICE_REMOTE_SERVICE_NAME`, `HOST_LOGICAL_CPU_CORES`, `HOST_BOSH_STEMCELL_VERSION`, `AZURE_TENANT_UUID`, `EXTERNAL_MONITOR_ENGINE_NAME`, `KUBERNETES_CLUSTER_NAME`, `AZURE_SUBSCRIPTION_NAME`, `SERVICE_IBM_CTG_GATEWAY_URL`, `GOOGLE_CLOUD_PLATFORM_ZONE_NAME`, `PROCESS_GROUP_CUSTOM_METADATA`, `EC2_INSTANCE_PRIVATE_HOST_NAME`, `MOBILE_APPLICATION_PLATFORM`, `ESXI_HOST_PRODUCT_NAME`, `HOST_IP_ADDRESS`, `AWS_RELATIONAL_DATABASE_SERVICE_NAME`, `WEB_APPLICATION_TYPE`, `SERVICE_PORT`, `EXTERNAL_MONITOR_ENGINE_TYPE`, `AZURE_SCALE_SET_NAME`",
			Required:    true,
		},
		"operator": {
			Type:        schema.TypeString,
			Description: "Possible Values: `NOT_LOWER_THAN_OR_EQUAL`, `NOT_CONTAINS`, `NOT_EQUALS`, `NOT_LOWER_THAN`, `GREATER_THAN_OR_EQUAL`, `LOWER_THAN`, `IS_IP_IN_RANGE`, `NOT_GREATER_THAN_OR_EQUAL`, `GREATER_THAN`, `NOT_REGEX_MATCHES`, `TAG_KEY_EQUALS`, `NOT_ENDS_WITH`, `NOT_GREATER_THAN`, `NOT_EXISTS`, `LOWER_THAN_OR_EQUAL`, `EXISTS`, `BEGINS_WITH`, `NOT_IS_IP_IN_RANGE`, `REGEX_MATCHES`, `CONTAINS`, `ENDS_WITH`, `NOT_BEGINS_WITH`, `NOT_TAG_KEY_EQUALS`, `EQUALS`",
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
