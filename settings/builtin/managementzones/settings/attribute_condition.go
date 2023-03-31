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
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/opt"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"
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
	Key              Attribute `json:"key"`                        // Possible Values: `APPMON_SERVER_NAME`, `APPMON_SYSTEM_PROFILE_NAME`, `AWS_ACCOUNT_ID`, `AWS_ACCOUNT_NAME`, `AWS_APPLICATION_LOAD_BALANCER_NAME`, `AWS_APPLICATION_LOAD_BALANCER_TAGS`, `AWS_AUTO_SCALING_GROUP_NAME`, `AWS_AUTO_SCALING_GROUP_TAGS`, `AWS_AVAILABILITY_ZONE_NAME`, `AWS_CLASSIC_LOAD_BALANCER_FRONTEND_PORTS`, `AWS_CLASSIC_LOAD_BALANCER_NAME`, `AWS_CLASSIC_LOAD_BALANCER_TAGS`, `AWS_NETWORK_LOAD_BALANCER_NAME`, `AWS_NETWORK_LOAD_BALANCER_TAGS`, `AWS_RELATIONAL_DATABASE_SERVICE_DB_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_ENDPOINT`, `AWS_RELATIONAL_DATABASE_SERVICE_ENGINE`, `AWS_RELATIONAL_DATABASE_SERVICE_INSTANCE_CLASS`, `AWS_RELATIONAL_DATABASE_SERVICE_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_PORT`, `AWS_RELATIONAL_DATABASE_SERVICE_TAGS`, `AZURE_ENTITY_NAME`, `AZURE_ENTITY_TAGS`, `AZURE_MGMT_GROUP_NAME`, `AZURE_MGMT_GROUP_UUID`, `AZURE_REGION_NAME`, `AZURE_SCALE_SET_NAME`, `AZURE_SUBSCRIPTION_NAME`, `AZURE_SUBSCRIPTION_UUID`, `AZURE_TENANT_NAME`, `AZURE_TENANT_UUID`, `AZURE_VM_NAME`, `BROWSER_MONITOR_NAME`, `BROWSER_MONITOR_TAGS`, `CLOUD_APPLICATION_LABELS`, `CLOUD_APPLICATION_NAME`, `CLOUD_APPLICATION_NAMESPACE_LABELS`, `CLOUD_APPLICATION_NAMESPACE_NAME`, `CLOUD_FOUNDRY_FOUNDATION_NAME`, `CLOUD_FOUNDRY_ORG_NAME`, `CUSTOM_APPLICATION_NAME`, `CUSTOM_APPLICATION_PLATFORM`, `CUSTOM_APPLICATION_TAGS`, `CUSTOM_APPLICATION_TYPE`, `CUSTOM_DEVICE_DNS_ADDRESS`, `CUSTOM_DEVICE_GROUP_NAME`, `CUSTOM_DEVICE_GROUP_TAGS`, `CUSTOM_DEVICE_IP_ADDRESS`, `CUSTOM_DEVICE_METADATA`, `CUSTOM_DEVICE_NAME`, `CUSTOM_DEVICE_PORT`, `CUSTOM_DEVICE_TAGS`, `CUSTOM_DEVICE_TECHNOLOGY`, `DATA_CENTER_SERVICE_DECODER_TYPE`, `DATA_CENTER_SERVICE_IP_ADDRESS`, `DATA_CENTER_SERVICE_METADATA`, `DATA_CENTER_SERVICE_NAME`, `DATA_CENTER_SERVICE_PORT`, `DATA_CENTER_SERVICE_TAGS`, `DOCKER_CONTAINER_NAME`, `DOCKER_FULL_IMAGE_NAME`, `DOCKER_IMAGE_VERSION`, `EC2_INSTANCE_AMI_ID`, `EC2_INSTANCE_AWS_INSTANCE_TYPE`, `EC2_INSTANCE_AWS_SECURITY_GROUP`, `EC2_INSTANCE_BEANSTALK_ENV_NAME`, `EC2_INSTANCE_ID`, `EC2_INSTANCE_NAME`, `EC2_INSTANCE_PRIVATE_HOST_NAME`, `EC2_INSTANCE_PUBLIC_HOST_NAME`, `EC2_INSTANCE_TAGS`, `ENTERPRISE_APPLICATION_DECODER_TYPE`, `ENTERPRISE_APPLICATION_IP_ADDRESS`, `ENTERPRISE_APPLICATION_METADATA`, `ENTERPRISE_APPLICATION_NAME`, `ENTERPRISE_APPLICATION_PORT`, `ENTERPRISE_APPLICATION_TAGS`, `ESXI_HOST_CLUSTER_NAME`, `ESXI_HOST_HARDWARE_MODEL`, `ESXI_HOST_HARDWARE_VENDOR`, `ESXI_HOST_NAME`, `ESXI_HOST_PRODUCT_NAME`, `ESXI_HOST_PRODUCT_VERSION`, `ESXI_HOST_TAGS`, `EXTERNAL_MONITOR_ENGINE_DESCRIPTION`, `EXTERNAL_MONITOR_ENGINE_NAME`, `EXTERNAL_MONITOR_ENGINE_TYPE`, `EXTERNAL_MONITOR_NAME`, `EXTERNAL_MONITOR_TAGS`, `GEOLOCATION_SITE_NAME`, `GOOGLE_CLOUD_PLATFORM_ZONE_NAME`, `GOOGLE_COMPUTE_INSTANCE_ID`, `GOOGLE_COMPUTE_INSTANCE_MACHINE_TYPE`, `GOOGLE_COMPUTE_INSTANCE_NAME`, `GOOGLE_COMPUTE_INSTANCE_PROJECT`, `GOOGLE_COMPUTE_INSTANCE_PROJECT_ID`, `GOOGLE_COMPUTE_INSTANCE_PUBLIC_IP_ADDRESSES`, `HOST_AIX_LOGICAL_CPU_COUNT`, `HOST_AIX_SIMULTANEOUS_THREADS`, `HOST_AIX_VIRTUAL_CPU_COUNT`, `HOST_ARCHITECTURE`, `HOST_AWS_NAME_TAG`, `HOST_AZURE_COMPUTE_MODE`, `HOST_AZURE_SKU`, `HOST_AZURE_WEB_APPLICATION_HOST_NAMES`, `HOST_AZURE_WEB_APPLICATION_SITE_NAMES`, `HOST_BITNESS`, `HOST_BOSH_AVAILABILITY_ZONE`, `HOST_BOSH_DEPLOYMENT_ID`, `HOST_BOSH_INSTANCE_ID`, `HOST_BOSH_INSTANCE_NAME`, `HOST_BOSH_NAME`, `HOST_BOSH_STEMCELL_VERSION`, `HOST_CLOUD_TYPE`, `HOST_CPU_CORES`, `HOST_CUSTOM_METADATA`, `HOST_DETECTED_NAME`, `HOST_GROUP_ID`, `HOST_GROUP_NAME`, `HOST_HYPERVISOR_TYPE`, `HOST_IP_ADDRESS`, `HOST_KUBERNETES_LABELS`, `HOST_LOGICAL_CPU_CORES`, `HOST_NAME`, `HOST_ONEAGENT_CUSTOM_HOST_NAME`, `HOST_OS_TYPE`, `HOST_OS_VERSION`, `HOST_PAAS_MEMORY_LIMIT`, `HOST_PAAS_TYPE`, `HOST_TAGS`, `HOST_TECHNOLOGY`, `HTTP_MONITOR_NAME`, `HTTP_MONITOR_TAGS`, `KUBERNETES_CLUSTER_NAME`, `KUBERNETES_NODE_NAME`, `KUBERNETES_SERVICE_NAME`, `MOBILE_APPLICATION_NAME`, `MOBILE_APPLICATION_PLATFORM`, `MOBILE_APPLICATION_TAGS`, `NAME_OF_COMPUTE_NODE`, `OPENSTACK_ACCOUNT_NAME`, `OPENSTACK_ACCOUNT_PROJECT_NAME`, `OPENSTACK_AVAILABILITY_ZONE_NAME`, `OPENSTACK_PROJECT_NAME`, `OPENSTACK_REGION_NAME`, `OPENSTACK_VM_INSTANCE_TYPE`, `OPENSTACK_VM_NAME`, `OPENSTACK_VM_SECURITY_GROUP`, `PROCESS_GROUP_AZURE_HOST_NAME`, `PROCESS_GROUP_AZURE_SITE_NAME`, `PROCESS_GROUP_CUSTOM_METADATA`, `PROCESS_GROUP_DETECTED_NAME`, `PROCESS_GROUP_ID`, `PROCESS_GROUP_LISTEN_PORT`, `PROCESS_GROUP_NAME`, `PROCESS_GROUP_PREDEFINED_METADATA`, `PROCESS_GROUP_TAGS`, `PROCESS_GROUP_TECHNOLOGY`, `PROCESS_GROUP_TECHNOLOGY_EDITION`, `PROCESS_GROUP_TECHNOLOGY_VERSION`, `QUEUE_NAME`, `QUEUE_TECHNOLOGY`, `QUEUE_VENDOR`, `SERVICE_AKKA_ACTOR_SYSTEM`, `SERVICE_CTG_SERVICE_NAME`, `SERVICE_DATABASE_HOST_NAME`, `SERVICE_DATABASE_NAME`, `SERVICE_DATABASE_TOPOLOGY`, `SERVICE_DATABASE_VENDOR`, `SERVICE_DETECTED_NAME`, `SERVICE_ESB_APPLICATION_NAME`, `SERVICE_IBM_CTG_GATEWAY_URL`, `SERVICE_MESSAGING_LISTENER_CLASS_NAME`, `SERVICE_NAME`, `SERVICE_PORT`, `SERVICE_PUBLIC_DOMAIN_NAME`, `SERVICE_REMOTE_ENDPOINT`, `SERVICE_REMOTE_SERVICE_NAME`, `SERVICE_TAGS`, `SERVICE_TECHNOLOGY`, `SERVICE_TECHNOLOGY_EDITION`, `SERVICE_TECHNOLOGY_VERSION`, `SERVICE_TOPOLOGY`, `SERVICE_TYPE`, `SERVICE_WEB_APPLICATION_ID`, `SERVICE_WEB_CONTEXT_ROOT`, `SERVICE_WEB_SERVER_ENDPOINT`, `SERVICE_WEB_SERVER_NAME`, `SERVICE_WEB_SERVICE_NAME`, `SERVICE_WEB_SERVICE_NAMESPACE`, `VMWARE_DATACENTER_NAME`, `VMWARE_VM_NAME`, `WEB_APPLICATION_NAME`, `WEB_APPLICATION_NAME_PATTERN`, `WEB_APPLICATION_TAGS`, `WEB_APPLICATION_TYPE`
	Operator         Operator  `json:"operator"`                   // Possible Values: `BEGINS_WITH`, `CONTAINS`, `ENDS_WITH`, `EQUALS`, `EXISTS`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `IS_IP_IN_RANGE`, `LOWER_THAN`, `LOWER_THAN_OR_EQUAL`, `NOT_BEGINS_WITH`, `NOT_CONTAINS`, `NOT_ENDS_WITH`, `NOT_EQUALS`, `NOT_EXISTS`, `NOT_GREATER_THAN`, `NOT_GREATER_THAN_OR_EQUAL`, `NOT_IS_IP_IN_RANGE`, `NOT_LOWER_THAN`, `NOT_LOWER_THAN_OR_EQUAL`, `NOT_REGEX_MATCHES`, `NOT_TAG_KEY_EQUALS`, `REGEX_MATCHES`, `TAG_KEY_EQUALS`
	StringValue      *string   `json:"stringValue,omitempty"`      // Value
	Tag              *string   `json:"tag,omitempty"`              // Format: `[CONTEXT]tagKey:tagValue`
}

func (me *AttributeCondition) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"case_sensitive": {
			Type:        schema.TypeBool,
			Description: "Case sensitive",
			Optional:    true, // precondition
		},
		"dynamic_key": {
			Type:        schema.TypeString,
			Description: "Dynamic key",
			Optional:    true, // precondition
		},
		"dynamic_key_source": {
			Type:        schema.TypeString,
			Description: "Key source",
			Optional:    true, // precondition
		},
		"entity_id": {
			Type:        schema.TypeString,
			Description: "Value",
			Optional:    true, // precondition
		},
		"enum_value": {
			Type:        schema.TypeString,
			Description: "Value",
			Optional:    true, // precondition
		},
		"integer_value": {
			Type:        schema.TypeInt,
			Description: "Value",
			Optional:    true, // precondition
		},
		"key": {
			Type:        schema.TypeString,
			Description: "Possible Values: `APPMON_SERVER_NAME`, `APPMON_SYSTEM_PROFILE_NAME`, `AWS_ACCOUNT_ID`, `AWS_ACCOUNT_NAME`, `AWS_APPLICATION_LOAD_BALANCER_NAME`, `AWS_APPLICATION_LOAD_BALANCER_TAGS`, `AWS_AUTO_SCALING_GROUP_NAME`, `AWS_AUTO_SCALING_GROUP_TAGS`, `AWS_AVAILABILITY_ZONE_NAME`, `AWS_CLASSIC_LOAD_BALANCER_FRONTEND_PORTS`, `AWS_CLASSIC_LOAD_BALANCER_NAME`, `AWS_CLASSIC_LOAD_BALANCER_TAGS`, `AWS_NETWORK_LOAD_BALANCER_NAME`, `AWS_NETWORK_LOAD_BALANCER_TAGS`, `AWS_RELATIONAL_DATABASE_SERVICE_DB_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_ENDPOINT`, `AWS_RELATIONAL_DATABASE_SERVICE_ENGINE`, `AWS_RELATIONAL_DATABASE_SERVICE_INSTANCE_CLASS`, `AWS_RELATIONAL_DATABASE_SERVICE_NAME`, `AWS_RELATIONAL_DATABASE_SERVICE_PORT`, `AWS_RELATIONAL_DATABASE_SERVICE_TAGS`, `AZURE_ENTITY_NAME`, `AZURE_ENTITY_TAGS`, `AZURE_MGMT_GROUP_NAME`, `AZURE_MGMT_GROUP_UUID`, `AZURE_REGION_NAME`, `AZURE_SCALE_SET_NAME`, `AZURE_SUBSCRIPTION_NAME`, `AZURE_SUBSCRIPTION_UUID`, `AZURE_TENANT_NAME`, `AZURE_TENANT_UUID`, `AZURE_VM_NAME`, `BROWSER_MONITOR_NAME`, `BROWSER_MONITOR_TAGS`, `CLOUD_APPLICATION_LABELS`, `CLOUD_APPLICATION_NAME`, `CLOUD_APPLICATION_NAMESPACE_LABELS`, `CLOUD_APPLICATION_NAMESPACE_NAME`, `CLOUD_FOUNDRY_FOUNDATION_NAME`, `CLOUD_FOUNDRY_ORG_NAME`, `CUSTOM_APPLICATION_NAME`, `CUSTOM_APPLICATION_PLATFORM`, `CUSTOM_APPLICATION_TAGS`, `CUSTOM_APPLICATION_TYPE`, `CUSTOM_DEVICE_DNS_ADDRESS`, `CUSTOM_DEVICE_GROUP_NAME`, `CUSTOM_DEVICE_GROUP_TAGS`, `CUSTOM_DEVICE_IP_ADDRESS`, `CUSTOM_DEVICE_METADATA`, `CUSTOM_DEVICE_NAME`, `CUSTOM_DEVICE_PORT`, `CUSTOM_DEVICE_TAGS`, `CUSTOM_DEVICE_TECHNOLOGY`, `DATA_CENTER_SERVICE_DECODER_TYPE`, `DATA_CENTER_SERVICE_IP_ADDRESS`, `DATA_CENTER_SERVICE_METADATA`, `DATA_CENTER_SERVICE_NAME`, `DATA_CENTER_SERVICE_PORT`, `DATA_CENTER_SERVICE_TAGS`, `DOCKER_CONTAINER_NAME`, `DOCKER_FULL_IMAGE_NAME`, `DOCKER_IMAGE_VERSION`, `EC2_INSTANCE_AMI_ID`, `EC2_INSTANCE_AWS_INSTANCE_TYPE`, `EC2_INSTANCE_AWS_SECURITY_GROUP`, `EC2_INSTANCE_BEANSTALK_ENV_NAME`, `EC2_INSTANCE_ID`, `EC2_INSTANCE_NAME`, `EC2_INSTANCE_PRIVATE_HOST_NAME`, `EC2_INSTANCE_PUBLIC_HOST_NAME`, `EC2_INSTANCE_TAGS`, `ENTERPRISE_APPLICATION_DECODER_TYPE`, `ENTERPRISE_APPLICATION_IP_ADDRESS`, `ENTERPRISE_APPLICATION_METADATA`, `ENTERPRISE_APPLICATION_NAME`, `ENTERPRISE_APPLICATION_PORT`, `ENTERPRISE_APPLICATION_TAGS`, `ESXI_HOST_CLUSTER_NAME`, `ESXI_HOST_HARDWARE_MODEL`, `ESXI_HOST_HARDWARE_VENDOR`, `ESXI_HOST_NAME`, `ESXI_HOST_PRODUCT_NAME`, `ESXI_HOST_PRODUCT_VERSION`, `ESXI_HOST_TAGS`, `EXTERNAL_MONITOR_ENGINE_DESCRIPTION`, `EXTERNAL_MONITOR_ENGINE_NAME`, `EXTERNAL_MONITOR_ENGINE_TYPE`, `EXTERNAL_MONITOR_NAME`, `EXTERNAL_MONITOR_TAGS`, `GEOLOCATION_SITE_NAME`, `GOOGLE_CLOUD_PLATFORM_ZONE_NAME`, `GOOGLE_COMPUTE_INSTANCE_ID`, `GOOGLE_COMPUTE_INSTANCE_MACHINE_TYPE`, `GOOGLE_COMPUTE_INSTANCE_NAME`, `GOOGLE_COMPUTE_INSTANCE_PROJECT`, `GOOGLE_COMPUTE_INSTANCE_PROJECT_ID`, `GOOGLE_COMPUTE_INSTANCE_PUBLIC_IP_ADDRESSES`, `HOST_AIX_LOGICAL_CPU_COUNT`, `HOST_AIX_SIMULTANEOUS_THREADS`, `HOST_AIX_VIRTUAL_CPU_COUNT`, `HOST_ARCHITECTURE`, `HOST_AWS_NAME_TAG`, `HOST_AZURE_COMPUTE_MODE`, `HOST_AZURE_SKU`, `HOST_AZURE_WEB_APPLICATION_HOST_NAMES`, `HOST_AZURE_WEB_APPLICATION_SITE_NAMES`, `HOST_BITNESS`, `HOST_BOSH_AVAILABILITY_ZONE`, `HOST_BOSH_DEPLOYMENT_ID`, `HOST_BOSH_INSTANCE_ID`, `HOST_BOSH_INSTANCE_NAME`, `HOST_BOSH_NAME`, `HOST_BOSH_STEMCELL_VERSION`, `HOST_CLOUD_TYPE`, `HOST_CPU_CORES`, `HOST_CUSTOM_METADATA`, `HOST_DETECTED_NAME`, `HOST_GROUP_ID`, `HOST_GROUP_NAME`, `HOST_HYPERVISOR_TYPE`, `HOST_IP_ADDRESS`, `HOST_KUBERNETES_LABELS`, `HOST_LOGICAL_CPU_CORES`, `HOST_NAME`, `HOST_ONEAGENT_CUSTOM_HOST_NAME`, `HOST_OS_TYPE`, `HOST_OS_VERSION`, `HOST_PAAS_MEMORY_LIMIT`, `HOST_PAAS_TYPE`, `HOST_TAGS`, `HOST_TECHNOLOGY`, `HTTP_MONITOR_NAME`, `HTTP_MONITOR_TAGS`, `KUBERNETES_CLUSTER_NAME`, `KUBERNETES_NODE_NAME`, `KUBERNETES_SERVICE_NAME`, `MOBILE_APPLICATION_NAME`, `MOBILE_APPLICATION_PLATFORM`, `MOBILE_APPLICATION_TAGS`, `NAME_OF_COMPUTE_NODE`, `OPENSTACK_ACCOUNT_NAME`, `OPENSTACK_ACCOUNT_PROJECT_NAME`, `OPENSTACK_AVAILABILITY_ZONE_NAME`, `OPENSTACK_PROJECT_NAME`, `OPENSTACK_REGION_NAME`, `OPENSTACK_VM_INSTANCE_TYPE`, `OPENSTACK_VM_NAME`, `OPENSTACK_VM_SECURITY_GROUP`, `PROCESS_GROUP_AZURE_HOST_NAME`, `PROCESS_GROUP_AZURE_SITE_NAME`, `PROCESS_GROUP_CUSTOM_METADATA`, `PROCESS_GROUP_DETECTED_NAME`, `PROCESS_GROUP_ID`, `PROCESS_GROUP_LISTEN_PORT`, `PROCESS_GROUP_NAME`, `PROCESS_GROUP_PREDEFINED_METADATA`, `PROCESS_GROUP_TAGS`, `PROCESS_GROUP_TECHNOLOGY`, `PROCESS_GROUP_TECHNOLOGY_EDITION`, `PROCESS_GROUP_TECHNOLOGY_VERSION`, `QUEUE_NAME`, `QUEUE_TECHNOLOGY`, `QUEUE_VENDOR`, `SERVICE_AKKA_ACTOR_SYSTEM`, `SERVICE_CTG_SERVICE_NAME`, `SERVICE_DATABASE_HOST_NAME`, `SERVICE_DATABASE_NAME`, `SERVICE_DATABASE_TOPOLOGY`, `SERVICE_DATABASE_VENDOR`, `SERVICE_DETECTED_NAME`, `SERVICE_ESB_APPLICATION_NAME`, `SERVICE_IBM_CTG_GATEWAY_URL`, `SERVICE_MESSAGING_LISTENER_CLASS_NAME`, `SERVICE_NAME`, `SERVICE_PORT`, `SERVICE_PUBLIC_DOMAIN_NAME`, `SERVICE_REMOTE_ENDPOINT`, `SERVICE_REMOTE_SERVICE_NAME`, `SERVICE_TAGS`, `SERVICE_TECHNOLOGY`, `SERVICE_TECHNOLOGY_EDITION`, `SERVICE_TECHNOLOGY_VERSION`, `SERVICE_TOPOLOGY`, `SERVICE_TYPE`, `SERVICE_WEB_APPLICATION_ID`, `SERVICE_WEB_CONTEXT_ROOT`, `SERVICE_WEB_SERVER_ENDPOINT`, `SERVICE_WEB_SERVER_NAME`, `SERVICE_WEB_SERVICE_NAME`, `SERVICE_WEB_SERVICE_NAMESPACE`, `VMWARE_DATACENTER_NAME`, `VMWARE_VM_NAME`, `WEB_APPLICATION_NAME`, `WEB_APPLICATION_NAME_PATTERN`, `WEB_APPLICATION_TAGS`, `WEB_APPLICATION_TYPE`",
			Required:    true,
		},
		"operator": {
			Type:        schema.TypeString,
			Description: "Possible Values: `BEGINS_WITH`, `CONTAINS`, `ENDS_WITH`, `EQUALS`, `EXISTS`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `IS_IP_IN_RANGE`, `LOWER_THAN`, `LOWER_THAN_OR_EQUAL`, `NOT_BEGINS_WITH`, `NOT_CONTAINS`, `NOT_ENDS_WITH`, `NOT_EQUALS`, `NOT_EXISTS`, `NOT_GREATER_THAN`, `NOT_GREATER_THAN_OR_EQUAL`, `NOT_IS_IP_IN_RANGE`, `NOT_LOWER_THAN`, `NOT_LOWER_THAN_OR_EQUAL`, `NOT_REGEX_MATCHES`, `NOT_TAG_KEY_EQUALS`, `REGEX_MATCHES`, `TAG_KEY_EQUALS`",
			Required:    true,
		},
		"string_value": {
			Type:        schema.TypeString,
			Description: "Value",
			Optional:    true, // precondition
		},
		"tag": {
			Type:        schema.TypeString,
			Description: "Format: `[CONTEXT]tagKey:tagValue`",
			Optional:    true, // precondition
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

func (me *AttributeCondition) HandlePreconditions() error {
	if me.DynamicKey == nil && slices.Contains([]string{"CLOUD_APPLICATION_LABELS", "CLOUD_APPLICATION_NAMESPACE_LABELS", "HOST_KUBERNETES_LABELS", "PROCESS_GROUP_PREDEFINED_METADATA", "CUSTOM_DEVICE_METADATA", "ENTERPRISE_APPLICATION_METADATA", "DATA_CENTER_SERVICE_METADATA", "HOST_CUSTOM_METADATA", "PROCESS_GROUP_CUSTOM_METADATA"}, string(me.Key)) {
		me.DynamicKey = opt.NewString("")
	}
	if me.DynamicKeySource == nil && slices.Contains([]string{"HOST_CUSTOM_METADATA", "PROCESS_GROUP_CUSTOM_METADATA"}, string(me.Key)) {
		return fmt.Errorf("'dynamic_key_source' must be specified if 'key' is set to '%v'", me.Key)
	}
	if me.EntityID == nil && slices.Contains([]string{"PROCESS_GROUP_ID", "HOST_GROUP_ID"}, string(me.Key)) {
		return fmt.Errorf("'entity_id' must be specified if 'key' is set to '%v'", me.Key)
	}
	// ---- CaseSensitive *bool -> {"preconditions":[{"expectedValues":["CLOUD_APPLICATION_LABELS","CLOUD_APPLICATION_NAMESPACE_LABELS","HOST_KUBERNETES_LABELS","PROCESS_GROUP_PREDEFINED_METADATA","CUSTOM_DEVICE_METADATA","ENTERPRISE_APPLICATION_METADATA","DATA_CENTER_SERVICE_METADATA","HOST_CUSTOM_METADATA","PROCESS_GROUP_CUSTOM_METADATA","KUBERNETES_SERVICE_NAME","HOST_AZURE_WEB_APPLICATION_HOST_NAMES","HOST_AZURE_WEB_APPLICATION_SITE_NAMES","HOST_DETECTED_NAME","HOST_NAME","HOST_OS_VERSION","HOST_BOSH_NAME","HOST_BOSH_INSTANCE_ID","HOST_BOSH_INSTANCE_NAME","HOST_BOSH_AVAILABILITY_ZONE","HOST_BOSH_DEPLOYMENT_ID","HOST_BOSH_STEMCELL_VERSION","HOST_AWS_NAME_TAG","HOST_ONEAGENT_CUSTOM_HOST_NAME","KUBERNETES_CLUSTER_NAME","KUBERNETES_NODE_NAME","CLOUD_APPLICATION_NAMESPACE_NAME","CLOUD_APPLICATION_NAME","PROCESS_GROUP_AZURE_HOST_NAME","PROCESS_GROUP_AZURE_SITE_NAME","PROCESS_GROUP_DETECTED_NAME","PROCESS_GROUP_NAME","PROCESS_GROUP_TECHNOLOGY_EDITION","PROCESS_GROUP_TECHNOLOGY_VERSION","SERVICE_AKKA_ACTOR_SYSTEM","SERVICE_DATABASE_NAME","SERVICE_DATABASE_VENDOR","SERVICE_DATABASE_HOST_NAME","SERVICE_DETECTED_NAME","SERVICE_WEB_SERVER_ENDPOINT","SERVICE_IBM_CTG_GATEWAY_URL","SERVICE_MESSAGING_LISTENER_CLASS_NAME","SERVICE_NAME","SERVICE_PUBLIC_DOMAIN_NAME","SERVICE_REMOTE_ENDPOINT","SERVICE_REMOTE_SERVICE_NAME","SERVICE_TECHNOLOGY_EDITION","SERVICE_TECHNOLOGY_VERSION","SERVICE_WEB_APPLICATION_ID","SERVICE_WEB_CONTEXT_ROOT","SERVICE_WEB_SERVER_NAME","SERVICE_WEB_SERVICE_NAME","SERVICE_WEB_SERVICE_NAMESPACE","SERVICE_CTG_SERVICE_NAME","SERVICE_ESB_APPLICATION_NAME","CUSTOM_DEVICE_DNS_ADDRESS","CUSTOM_DEVICE_NAME","CUSTOM_DEVICE_GROUP_NAME","WEB_APPLICATION_NAME","WEB_APPLICATION_NAME_PATTERN","MOBILE_APPLICATION_NAME","CUSTOM_APPLICATION_NAME","ENTERPRISE_APPLICATION_NAME","DATA_CENTER_SERVICE_NAME","BROWSER_MONITOR_NAME","EXTERNAL_MONITOR_NAME","EXTERNAL_MONITOR_ENGINE_NAME","EXTERNAL_MONITOR_ENGINE_DESCRIPTION","HTTP_MONITOR_NAME","DOCKER_CONTAINER_NAME","DOCKER_FULL_IMAGE_NAME","DOCKER_IMAGE_VERSION","ESXI_HOST_HARDWARE_MODEL","ESXI_HOST_HARDWARE_VENDOR","ESXI_HOST_NAME","ESXI_HOST_CLUSTER_NAME","ESXI_HOST_PRODUCT_NAME","ESXI_HOST_PRODUCT_VERSION","NAME_OF_COMPUTE_NODE","EC2_INSTANCE_NAME","EC2_INSTANCE_AMI_ID","EC2_INSTANCE_BEANSTALK_ENV_NAME","EC2_INSTANCE_AWS_INSTANCE_TYPE","EC2_INSTANCE_ID","EC2_INSTANCE_PRIVATE_HOST_NAME","EC2_INSTANCE_PUBLIC_HOST_NAME","EC2_INSTANCE_AWS_SECURITY_GROUP","OPENSTACK_VM_INSTANCE_TYPE","OPENSTACK_VM_NAME","OPENSTACK_VM_SECURITY_GROUP","VMWARE_VM_NAME","GOOGLE_COMPUTE_INSTANCE_ID","GOOGLE_COMPUTE_INSTANCE_NAME","GOOGLE_COMPUTE_INSTANCE_MACHINE_TYPE","GOOGLE_COMPUTE_INSTANCE_PROJECT","GOOGLE_COMPUTE_INSTANCE_PROJECT_ID","AWS_AVAILABILITY_ZONE_NAME","AZURE_REGION_NAME","OPENSTACK_REGION_NAME","OPENSTACK_AVAILABILITY_ZONE_NAME","GEOLOCATION_SITE_NAME","VMWARE_DATACENTER_NAME","GOOGLE_CLOUD_PLATFORM_ZONE_NAME","AWS_AUTO_SCALING_GROUP_NAME","AWS_CLASSIC_LOAD_BALANCER_NAME","AWS_APPLICATION_LOAD_BALANCER_NAME","AWS_NETWORK_LOAD_BALANCER_NAME","AWS_RELATIONAL_DATABASE_SERVICE_NAME","AWS_RELATIONAL_DATABASE_SERVICE_INSTANCE_CLASS","AWS_RELATIONAL_DATABASE_SERVICE_ENDPOINT","AWS_RELATIONAL_DATABASE_SERVICE_ENGINE","AWS_RELATIONAL_DATABASE_SERVICE_DB_NAME","AZURE_SCALE_SET_NAME","AZURE_VM_NAME","OPENSTACK_PROJECT_NAME","HOST_GROUP_NAME","AWS_ACCOUNT_ID","AWS_ACCOUNT_NAME","OPENSTACK_ACCOUNT_NAME","OPENSTACK_ACCOUNT_PROJECT_NAME","CLOUD_FOUNDRY_ORG_NAME","CLOUD_FOUNDRY_FOUNDATION_NAME","APPMON_SERVER_NAME","APPMON_SYSTEM_PROFILE_NAME","QUEUE_NAME","QUEUE_VENDOR"],"property":"key","type":"IN"},{"precondition":{"expectedValues":["EXISTS","NOT_EXISTS"],"property":"operator","type":"IN"},"type":"NOT"}],"type":"AND"}
	// ---- EnumValue *string -> {"preconditions":[{"expectedValues":["HOST_AZURE_COMPUTE_MODE","HOST_AZURE_SKU","HOST_BITNESS","HOST_CLOUD_TYPE","HOST_HYPERVISOR_TYPE","HOST_ARCHITECTURE","HOST_OS_TYPE","HOST_PAAS_TYPE","HOST_TECHNOLOGY","PROCESS_GROUP_TECHNOLOGY","SERVICE_DATABASE_TOPOLOGY","SERVICE_TOPOLOGY","SERVICE_TYPE","SERVICE_TECHNOLOGY","CUSTOM_DEVICE_TECHNOLOGY","WEB_APPLICATION_TYPE","MOBILE_APPLICATION_PLATFORM","CUSTOM_APPLICATION_TYPE","CUSTOM_APPLICATION_PLATFORM","ENTERPRISE_APPLICATION_DECODER_TYPE","DATA_CENTER_SERVICE_DECODER_TYPE","EXTERNAL_MONITOR_ENGINE_TYPE","QUEUE_TECHNOLOGY"],"property":"key","type":"IN"},{"precondition":{"expectedValues":["EXISTS","NOT_EXISTS"],"property":"operator","type":"IN"},"type":"NOT"}],"type":"AND"}
	// ---- IntegerValue *int -> {"preconditions":[{"expectedValues":["HOST_PAAS_MEMORY_LIMIT","HOST_AIX_VIRTUAL_CPU_COUNT","HOST_AIX_LOGICAL_CPU_COUNT","HOST_AIX_SIMULTANEOUS_THREADS","HOST_CPU_CORES","HOST_LOGICAL_CPU_CORES","PROCESS_GROUP_LISTEN_PORT","SERVICE_PORT","CUSTOM_DEVICE_PORT","ENTERPRISE_APPLICATION_PORT","DATA_CENTER_SERVICE_PORT","AWS_CLASSIC_LOAD_BALANCER_FRONTEND_PORTS","AWS_RELATIONAL_DATABASE_SERVICE_PORT"],"property":"key","type":"IN"},{"precondition":{"expectedValues":["EXISTS","NOT_EXISTS"],"property":"operator","type":"IN"},"type":"NOT"}],"type":"AND"}
	// ---- StringValue *string -> {"preconditions":[{"preconditions":[{"expectedValues":["CLOUD_APPLICATION_LABELS","CLOUD_APPLICATION_NAMESPACE_LABELS","HOST_KUBERNETES_LABELS","PROCESS_GROUP_PREDEFINED_METADATA","CUSTOM_DEVICE_METADATA","ENTERPRISE_APPLICATION_METADATA","DATA_CENTER_SERVICE_METADATA","HOST_CUSTOM_METADATA","PROCESS_GROUP_CUSTOM_METADATA","KUBERNETES_SERVICE_NAME","HOST_AZURE_WEB_APPLICATION_HOST_NAMES","HOST_AZURE_WEB_APPLICATION_SITE_NAMES","HOST_DETECTED_NAME","HOST_NAME","HOST_OS_VERSION","HOST_BOSH_NAME","HOST_BOSH_INSTANCE_ID","HOST_BOSH_INSTANCE_NAME","HOST_BOSH_AVAILABILITY_ZONE","HOST_BOSH_DEPLOYMENT_ID","HOST_BOSH_STEMCELL_VERSION","HOST_AWS_NAME_TAG","HOST_ONEAGENT_CUSTOM_HOST_NAME","KUBERNETES_CLUSTER_NAME","KUBERNETES_NODE_NAME","CLOUD_APPLICATION_NAMESPACE_NAME","CLOUD_APPLICATION_NAME","PROCESS_GROUP_AZURE_HOST_NAME","PROCESS_GROUP_AZURE_SITE_NAME","PROCESS_GROUP_DETECTED_NAME","PROCESS_GROUP_NAME","PROCESS_GROUP_TECHNOLOGY_EDITION","PROCESS_GROUP_TECHNOLOGY_VERSION","SERVICE_AKKA_ACTOR_SYSTEM","SERVICE_DATABASE_NAME","SERVICE_DATABASE_VENDOR","SERVICE_DATABASE_HOST_NAME","SERVICE_DETECTED_NAME","SERVICE_WEB_SERVER_ENDPOINT","SERVICE_IBM_CTG_GATEWAY_URL","SERVICE_MESSAGING_LISTENER_CLASS_NAME","SERVICE_NAME","SERVICE_PUBLIC_DOMAIN_NAME","SERVICE_REMOTE_ENDPOINT","SERVICE_REMOTE_SERVICE_NAME","SERVICE_TECHNOLOGY_EDITION","SERVICE_TECHNOLOGY_VERSION","SERVICE_WEB_APPLICATION_ID","SERVICE_WEB_CONTEXT_ROOT","SERVICE_WEB_SERVER_NAME","SERVICE_WEB_SERVICE_NAME","SERVICE_WEB_SERVICE_NAMESPACE","SERVICE_CTG_SERVICE_NAME","SERVICE_ESB_APPLICATION_NAME","CUSTOM_DEVICE_DNS_ADDRESS","CUSTOM_DEVICE_NAME","CUSTOM_DEVICE_GROUP_NAME","WEB_APPLICATION_NAME","WEB_APPLICATION_NAME_PATTERN","MOBILE_APPLICATION_NAME","CUSTOM_APPLICATION_NAME","ENTERPRISE_APPLICATION_NAME","DATA_CENTER_SERVICE_NAME","BROWSER_MONITOR_NAME","EXTERNAL_MONITOR_NAME","EXTERNAL_MONITOR_ENGINE_NAME","EXTERNAL_MONITOR_ENGINE_DESCRIPTION","HTTP_MONITOR_NAME","DOCKER_CONTAINER_NAME","DOCKER_FULL_IMAGE_NAME","DOCKER_IMAGE_VERSION","ESXI_HOST_HARDWARE_MODEL","ESXI_HOST_HARDWARE_VENDOR","ESXI_HOST_NAME","ESXI_HOST_CLUSTER_NAME","ESXI_HOST_PRODUCT_NAME","ESXI_HOST_PRODUCT_VERSION","NAME_OF_COMPUTE_NODE","EC2_INSTANCE_NAME","EC2_INSTANCE_AMI_ID","EC2_INSTANCE_BEANSTALK_ENV_NAME","EC2_INSTANCE_AWS_INSTANCE_TYPE","EC2_INSTANCE_ID","EC2_INSTANCE_PRIVATE_HOST_NAME","EC2_INSTANCE_PUBLIC_HOST_NAME","EC2_INSTANCE_AWS_SECURITY_GROUP","OPENSTACK_VM_INSTANCE_TYPE","OPENSTACK_VM_NAME","OPENSTACK_VM_SECURITY_GROUP","VMWARE_VM_NAME","GOOGLE_COMPUTE_INSTANCE_ID","GOOGLE_COMPUTE_INSTANCE_NAME","GOOGLE_COMPUTE_INSTANCE_MACHINE_TYPE","GOOGLE_COMPUTE_INSTANCE_PROJECT","GOOGLE_COMPUTE_INSTANCE_PROJECT_ID","AWS_AVAILABILITY_ZONE_NAME","AZURE_REGION_NAME","OPENSTACK_REGION_NAME","OPENSTACK_AVAILABILITY_ZONE_NAME","GEOLOCATION_SITE_NAME","VMWARE_DATACENTER_NAME","GOOGLE_CLOUD_PLATFORM_ZONE_NAME","AWS_AUTO_SCALING_GROUP_NAME","AWS_CLASSIC_LOAD_BALANCER_NAME","AWS_APPLICATION_LOAD_BALANCER_NAME","AWS_NETWORK_LOAD_BALANCER_NAME","AWS_RELATIONAL_DATABASE_SERVICE_NAME","AWS_RELATIONAL_DATABASE_SERVICE_INSTANCE_CLASS","AWS_RELATIONAL_DATABASE_SERVICE_ENDPOINT","AWS_RELATIONAL_DATABASE_SERVICE_ENGINE","AWS_RELATIONAL_DATABASE_SERVICE_DB_NAME","AZURE_SCALE_SET_NAME","AZURE_VM_NAME","OPENSTACK_PROJECT_NAME","HOST_GROUP_NAME","AWS_ACCOUNT_ID","AWS_ACCOUNT_NAME","OPENSTACK_ACCOUNT_NAME","OPENSTACK_ACCOUNT_PROJECT_NAME","CLOUD_FOUNDRY_ORG_NAME","CLOUD_FOUNDRY_FOUNDATION_NAME","APPMON_SERVER_NAME","APPMON_SYSTEM_PROFILE_NAME","QUEUE_NAME","QUEUE_VENDOR","AZURE_TENANT_UUID","AZURE_MGMT_GROUP_UUID","AZURE_SUBSCRIPTION_UUID","AZURE_ENTITY_NAME","AZURE_TENANT_NAME","AZURE_MGMT_GROUP_NAME","AZURE_SUBSCRIPTION_NAME","HOST_IP_ADDRESS","CUSTOM_DEVICE_IP_ADDRESS","ENTERPRISE_APPLICATION_IP_ADDRESS","DATA_CENTER_SERVICE_IP_ADDRESS","GOOGLE_COMPUTE_INSTANCE_PUBLIC_IP_ADDRESSES"],"property":"key","type":"IN"},{"preconditions":[{"expectedValue":"VERBATIM","property":"enumValue","type":"EQUALS"},{"expectedValues":["HOST_TECHNOLOGY","SERVICE_TECHNOLOGY","CUSTOM_DEVICE_TECHNOLOGY","PROCESS_GROUP_TECHNOLOGY","QUEUE_TECHNOLOGY"],"property":"key","type":"IN"}],"type":"AND"}],"type":"OR"},{"expectedValues":["BEGINS_WITH","CONTAINS","ENDS_WITH","EQUALS","REGEX_MATCHES","IS_IP_IN_RANGE","NOT_BEGINS_WITH","NOT_CONTAINS","NOT_ENDS_WITH","NOT_EQUALS","NOT_REGEX_MATCHES","NOT_IS_IP_IN_RANGE"],"property":"operator","type":"IN"}],"type":"AND"}
	// ---- Tag *string -> {"preconditions":[{"expectedValues":["AZURE_ENTITY_TAGS","HOST_TAGS","PROCESS_GROUP_TAGS","SERVICE_TAGS","CUSTOM_DEVICE_TAGS","CUSTOM_DEVICE_GROUP_TAGS","WEB_APPLICATION_TAGS","MOBILE_APPLICATION_TAGS","CUSTOM_APPLICATION_TAGS","ENTERPRISE_APPLICATION_TAGS","DATA_CENTER_SERVICE_TAGS","BROWSER_MONITOR_TAGS","EXTERNAL_MONITOR_TAGS","HTTP_MONITOR_TAGS","ESXI_HOST_TAGS","EC2_INSTANCE_TAGS","AWS_AUTO_SCALING_GROUP_TAGS","AWS_CLASSIC_LOAD_BALANCER_TAGS","AWS_APPLICATION_LOAD_BALANCER_TAGS","AWS_NETWORK_LOAD_BALANCER_TAGS","AWS_RELATIONAL_DATABASE_SERVICE_TAGS"],"property":"key","type":"IN"},{"precondition":{"expectedValues":["EXISTS","NOT_EXISTS"],"property":"operator","type":"IN"},"type":"NOT"}],"type":"AND"}
	return nil
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
