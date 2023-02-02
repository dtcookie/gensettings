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

type Attribute string

var Attributes = struct {
	AppmonServerName                          Attribute
	AppmonSystemProfileName                   Attribute
	AwsAccountId                              Attribute
	AwsAccountName                            Attribute
	AwsApplicationLoadBalancerName            Attribute
	AwsApplicationLoadBalancerTags            Attribute
	AwsAutoScalingGroupName                   Attribute
	AwsAutoScalingGroupTags                   Attribute
	AwsAvailabilityZoneName                   Attribute
	AwsClassicLoadBalancerFrontendPorts       Attribute
	AwsClassicLoadBalancerName                Attribute
	AwsClassicLoadBalancerTags                Attribute
	AwsNetworkLoadBalancerName                Attribute
	AwsNetworkLoadBalancerTags                Attribute
	AwsRelationalDatabaseServiceDbName        Attribute
	AwsRelationalDatabaseServiceEndpoint      Attribute
	AwsRelationalDatabaseServiceEngine        Attribute
	AwsRelationalDatabaseServiceInstanceClass Attribute
	AwsRelationalDatabaseServiceName          Attribute
	AwsRelationalDatabaseServicePort          Attribute
	AwsRelationalDatabaseServiceTags          Attribute
	AzureEntityName                           Attribute
	AzureEntityTags                           Attribute
	AzureMgmtGroupName                        Attribute
	AzureMgmtGroupUuid                        Attribute
	AzureRegionName                           Attribute
	AzureScaleSetName                         Attribute
	AzureSubscriptionName                     Attribute
	AzureSubscriptionUuid                     Attribute
	AzureTenantName                           Attribute
	AzureTenantUuid                           Attribute
	AzureVmName                               Attribute
	BrowserMonitorName                        Attribute
	BrowserMonitorTags                        Attribute
	CloudApplicationLabels                    Attribute
	CloudApplicationName                      Attribute
	CloudApplicationNamespaceLabels           Attribute
	CloudApplicationNamespaceName             Attribute
	CloudFoundryFoundationName                Attribute
	CloudFoundryOrgName                       Attribute
	CustomApplicationName                     Attribute
	CustomApplicationPlatform                 Attribute
	CustomApplicationTags                     Attribute
	CustomApplicationType                     Attribute
	CustomDeviceDnsAddress                    Attribute
	CustomDeviceGroupName                     Attribute
	CustomDeviceGroupTags                     Attribute
	CustomDeviceIpAddress                     Attribute
	CustomDeviceMetadata                      Attribute
	CustomDeviceName                          Attribute
	CustomDevicePort                          Attribute
	CustomDeviceTags                          Attribute
	CustomDeviceTechnology                    Attribute
	DataCenterServiceDecoderType              Attribute
	DataCenterServiceIpAddress                Attribute
	DataCenterServiceMetadata                 Attribute
	DataCenterServiceName                     Attribute
	DataCenterServicePort                     Attribute
	DataCenterServiceTags                     Attribute
	DockerContainerName                       Attribute
	DockerFullImageName                       Attribute
	DockerImageVersion                        Attribute
	Ec2InstanceAmiId                          Attribute
	Ec2InstanceAwsInstanceType                Attribute
	Ec2InstanceAwsSecurityGroup               Attribute
	Ec2InstanceBeanstalkEnvName               Attribute
	Ec2InstanceId                             Attribute
	Ec2InstanceName                           Attribute
	Ec2InstancePrivateHostName                Attribute
	Ec2InstancePublicHostName                 Attribute
	Ec2InstanceTags                           Attribute
	EnterpriseApplicationDecoderType          Attribute
	EnterpriseApplicationIpAddress            Attribute
	EnterpriseApplicationMetadata             Attribute
	EnterpriseApplicationName                 Attribute
	EnterpriseApplicationPort                 Attribute
	EnterpriseApplicationTags                 Attribute
	EsxiHostClusterName                       Attribute
	EsxiHostHardwareModel                     Attribute
	EsxiHostHardwareVendor                    Attribute
	EsxiHostName                              Attribute
	EsxiHostProductName                       Attribute
	EsxiHostProductVersion                    Attribute
	EsxiHostTags                              Attribute
	ExternalMonitorEngineDescription          Attribute
	ExternalMonitorEngineName                 Attribute
	ExternalMonitorEngineType                 Attribute
	ExternalMonitorName                       Attribute
	ExternalMonitorTags                       Attribute
	GeolocationSiteName                       Attribute
	GoogleCloudPlatformZoneName               Attribute
	GoogleComputeInstanceId                   Attribute
	GoogleComputeInstanceMachineType          Attribute
	GoogleComputeInstanceName                 Attribute
	GoogleComputeInstanceProject              Attribute
	GoogleComputeInstanceProjectId            Attribute
	GoogleComputeInstancePublicIpAddresses    Attribute
	HostAixLogicalCpuCount                    Attribute
	HostAixSimultaneousThreads                Attribute
	HostAixVirtualCpuCount                    Attribute
	HostArchitecture                          Attribute
	HostAwsNameTag                            Attribute
	HostAzureComputeMode                      Attribute
	HostAzureSku                              Attribute
	HostAzureWebApplicationHostNames          Attribute
	HostAzureWebApplicationSiteNames          Attribute
	HostBitness                               Attribute
	HostBoshAvailabilityZone                  Attribute
	HostBoshDeploymentId                      Attribute
	HostBoshInstanceId                        Attribute
	HostBoshInstanceName                      Attribute
	HostBoshName                              Attribute
	HostBoshStemcellVersion                   Attribute
	HostCloudType                             Attribute
	HostCpuCores                              Attribute
	HostCustomMetadata                        Attribute
	HostDetectedName                          Attribute
	HostGroupId                               Attribute
	HostGroupName                             Attribute
	HostHypervisorType                        Attribute
	HostIpAddress                             Attribute
	HostKubernetesLabels                      Attribute
	HostLogicalCpuCores                       Attribute
	HostName                                  Attribute
	HostOneagentCustomHostName                Attribute
	HostOsType                                Attribute
	HostOsVersion                             Attribute
	HostPaasMemoryLimit                       Attribute
	HostPaasType                              Attribute
	HostTags                                  Attribute
	HostTechnology                            Attribute
	HttpMonitorName                           Attribute
	HttpMonitorTags                           Attribute
	KubernetesClusterName                     Attribute
	KubernetesNodeName                        Attribute
	KubernetesServiceName                     Attribute
	MobileApplicationName                     Attribute
	MobileApplicationPlatform                 Attribute
	MobileApplicationTags                     Attribute
	NameOfComputeNode                         Attribute
	OpenstackAccountName                      Attribute
	OpenstackAccountProjectName               Attribute
	OpenstackAvailabilityZoneName             Attribute
	OpenstackProjectName                      Attribute
	OpenstackRegionName                       Attribute
	OpenstackVmInstanceType                   Attribute
	OpenstackVmName                           Attribute
	OpenstackVmSecurityGroup                  Attribute
	ProcessGroupAzureHostName                 Attribute
	ProcessGroupAzureSiteName                 Attribute
	ProcessGroupCustomMetadata                Attribute
	ProcessGroupDetectedName                  Attribute
	ProcessGroupId                            Attribute
	ProcessGroupListenPort                    Attribute
	ProcessGroupName                          Attribute
	ProcessGroupPredefinedMetadata            Attribute
	ProcessGroupTags                          Attribute
	ProcessGroupTechnology                    Attribute
	ProcessGroupTechnologyEdition             Attribute
	ProcessGroupTechnologyVersion             Attribute
	QueueName                                 Attribute
	QueueTechnology                           Attribute
	QueueVendor                               Attribute
	ServiceAkkaActorSystem                    Attribute
	ServiceCtgServiceName                     Attribute
	ServiceDatabaseHostName                   Attribute
	ServiceDatabaseName                       Attribute
	ServiceDatabaseTopology                   Attribute
	ServiceDatabaseVendor                     Attribute
	ServiceDetectedName                       Attribute
	ServiceEsbApplicationName                 Attribute
	ServiceIbmCtgGatewayUrl                   Attribute
	ServiceMessagingListenerClassName         Attribute
	ServiceName                               Attribute
	ServicePort                               Attribute
	ServicePublicDomainName                   Attribute
	ServiceRemoteEndpoint                     Attribute
	ServiceRemoteServiceName                  Attribute
	ServiceTags                               Attribute
	ServiceTechnology                         Attribute
	ServiceTechnologyEdition                  Attribute
	ServiceTechnologyVersion                  Attribute
	ServiceTopology                           Attribute
	ServiceType                               Attribute
	ServiceWebApplicationId                   Attribute
	ServiceWebContextRoot                     Attribute
	ServiceWebServerEndpoint                  Attribute
	ServiceWebServerName                      Attribute
	ServiceWebServiceName                     Attribute
	ServiceWebServiceNamespace                Attribute
	VmwareDatacenterName                      Attribute
	VmwareVmName                              Attribute
	WebApplicationName                        Attribute
	WebApplicationNamePattern                 Attribute
	WebApplicationTags                        Attribute
	WebApplicationType                        Attribute
}{
	Attribute("APPMON_SERVER_NAME"),
	Attribute("APPMON_SYSTEM_PROFILE_NAME"),
	Attribute("AWS_ACCOUNT_ID"),
	Attribute("AWS_ACCOUNT_NAME"),
	Attribute("AWS_APPLICATION_LOAD_BALANCER_NAME"),
	Attribute("AWS_APPLICATION_LOAD_BALANCER_TAGS"),
	Attribute("AWS_AUTO_SCALING_GROUP_NAME"),
	Attribute("AWS_AUTO_SCALING_GROUP_TAGS"),
	Attribute("AWS_AVAILABILITY_ZONE_NAME"),
	Attribute("AWS_CLASSIC_LOAD_BALANCER_FRONTEND_PORTS"),
	Attribute("AWS_CLASSIC_LOAD_BALANCER_NAME"),
	Attribute("AWS_CLASSIC_LOAD_BALANCER_TAGS"),
	Attribute("AWS_NETWORK_LOAD_BALANCER_NAME"),
	Attribute("AWS_NETWORK_LOAD_BALANCER_TAGS"),
	Attribute("AWS_RELATIONAL_DATABASE_SERVICE_DB_NAME"),
	Attribute("AWS_RELATIONAL_DATABASE_SERVICE_ENDPOINT"),
	Attribute("AWS_RELATIONAL_DATABASE_SERVICE_ENGINE"),
	Attribute("AWS_RELATIONAL_DATABASE_SERVICE_INSTANCE_CLASS"),
	Attribute("AWS_RELATIONAL_DATABASE_SERVICE_NAME"),
	Attribute("AWS_RELATIONAL_DATABASE_SERVICE_PORT"),
	Attribute("AWS_RELATIONAL_DATABASE_SERVICE_TAGS"),
	Attribute("AZURE_ENTITY_NAME"),
	Attribute("AZURE_ENTITY_TAGS"),
	Attribute("AZURE_MGMT_GROUP_NAME"),
	Attribute("AZURE_MGMT_GROUP_UUID"),
	Attribute("AZURE_REGION_NAME"),
	Attribute("AZURE_SCALE_SET_NAME"),
	Attribute("AZURE_SUBSCRIPTION_NAME"),
	Attribute("AZURE_SUBSCRIPTION_UUID"),
	Attribute("AZURE_TENANT_NAME"),
	Attribute("AZURE_TENANT_UUID"),
	Attribute("AZURE_VM_NAME"),
	Attribute("BROWSER_MONITOR_NAME"),
	Attribute("BROWSER_MONITOR_TAGS"),
	Attribute("CLOUD_APPLICATION_LABELS"),
	Attribute("CLOUD_APPLICATION_NAME"),
	Attribute("CLOUD_APPLICATION_NAMESPACE_LABELS"),
	Attribute("CLOUD_APPLICATION_NAMESPACE_NAME"),
	Attribute("CLOUD_FOUNDRY_FOUNDATION_NAME"),
	Attribute("CLOUD_FOUNDRY_ORG_NAME"),
	Attribute("CUSTOM_APPLICATION_NAME"),
	Attribute("CUSTOM_APPLICATION_PLATFORM"),
	Attribute("CUSTOM_APPLICATION_TAGS"),
	Attribute("CUSTOM_APPLICATION_TYPE"),
	Attribute("CUSTOM_DEVICE_DNS_ADDRESS"),
	Attribute("CUSTOM_DEVICE_GROUP_NAME"),
	Attribute("CUSTOM_DEVICE_GROUP_TAGS"),
	Attribute("CUSTOM_DEVICE_IP_ADDRESS"),
	Attribute("CUSTOM_DEVICE_METADATA"),
	Attribute("CUSTOM_DEVICE_NAME"),
	Attribute("CUSTOM_DEVICE_PORT"),
	Attribute("CUSTOM_DEVICE_TAGS"),
	Attribute("CUSTOM_DEVICE_TECHNOLOGY"),
	Attribute("DATA_CENTER_SERVICE_DECODER_TYPE"),
	Attribute("DATA_CENTER_SERVICE_IP_ADDRESS"),
	Attribute("DATA_CENTER_SERVICE_METADATA"),
	Attribute("DATA_CENTER_SERVICE_NAME"),
	Attribute("DATA_CENTER_SERVICE_PORT"),
	Attribute("DATA_CENTER_SERVICE_TAGS"),
	Attribute("DOCKER_CONTAINER_NAME"),
	Attribute("DOCKER_FULL_IMAGE_NAME"),
	Attribute("DOCKER_IMAGE_VERSION"),
	Attribute("EC2_INSTANCE_AMI_ID"),
	Attribute("EC2_INSTANCE_AWS_INSTANCE_TYPE"),
	Attribute("EC2_INSTANCE_AWS_SECURITY_GROUP"),
	Attribute("EC2_INSTANCE_BEANSTALK_ENV_NAME"),
	Attribute("EC2_INSTANCE_ID"),
	Attribute("EC2_INSTANCE_NAME"),
	Attribute("EC2_INSTANCE_PRIVATE_HOST_NAME"),
	Attribute("EC2_INSTANCE_PUBLIC_HOST_NAME"),
	Attribute("EC2_INSTANCE_TAGS"),
	Attribute("ENTERPRISE_APPLICATION_DECODER_TYPE"),
	Attribute("ENTERPRISE_APPLICATION_IP_ADDRESS"),
	Attribute("ENTERPRISE_APPLICATION_METADATA"),
	Attribute("ENTERPRISE_APPLICATION_NAME"),
	Attribute("ENTERPRISE_APPLICATION_PORT"),
	Attribute("ENTERPRISE_APPLICATION_TAGS"),
	Attribute("ESXI_HOST_CLUSTER_NAME"),
	Attribute("ESXI_HOST_HARDWARE_MODEL"),
	Attribute("ESXI_HOST_HARDWARE_VENDOR"),
	Attribute("ESXI_HOST_NAME"),
	Attribute("ESXI_HOST_PRODUCT_NAME"),
	Attribute("ESXI_HOST_PRODUCT_VERSION"),
	Attribute("ESXI_HOST_TAGS"),
	Attribute("EXTERNAL_MONITOR_ENGINE_DESCRIPTION"),
	Attribute("EXTERNAL_MONITOR_ENGINE_NAME"),
	Attribute("EXTERNAL_MONITOR_ENGINE_TYPE"),
	Attribute("EXTERNAL_MONITOR_NAME"),
	Attribute("EXTERNAL_MONITOR_TAGS"),
	Attribute("GEOLOCATION_SITE_NAME"),
	Attribute("GOOGLE_CLOUD_PLATFORM_ZONE_NAME"),
	Attribute("GOOGLE_COMPUTE_INSTANCE_ID"),
	Attribute("GOOGLE_COMPUTE_INSTANCE_MACHINE_TYPE"),
	Attribute("GOOGLE_COMPUTE_INSTANCE_NAME"),
	Attribute("GOOGLE_COMPUTE_INSTANCE_PROJECT"),
	Attribute("GOOGLE_COMPUTE_INSTANCE_PROJECT_ID"),
	Attribute("GOOGLE_COMPUTE_INSTANCE_PUBLIC_IP_ADDRESSES"),
	Attribute("HOST_AIX_LOGICAL_CPU_COUNT"),
	Attribute("HOST_AIX_SIMULTANEOUS_THREADS"),
	Attribute("HOST_AIX_VIRTUAL_CPU_COUNT"),
	Attribute("HOST_ARCHITECTURE"),
	Attribute("HOST_AWS_NAME_TAG"),
	Attribute("HOST_AZURE_COMPUTE_MODE"),
	Attribute("HOST_AZURE_SKU"),
	Attribute("HOST_AZURE_WEB_APPLICATION_HOST_NAMES"),
	Attribute("HOST_AZURE_WEB_APPLICATION_SITE_NAMES"),
	Attribute("HOST_BITNESS"),
	Attribute("HOST_BOSH_AVAILABILITY_ZONE"),
	Attribute("HOST_BOSH_DEPLOYMENT_ID"),
	Attribute("HOST_BOSH_INSTANCE_ID"),
	Attribute("HOST_BOSH_INSTANCE_NAME"),
	Attribute("HOST_BOSH_NAME"),
	Attribute("HOST_BOSH_STEMCELL_VERSION"),
	Attribute("HOST_CLOUD_TYPE"),
	Attribute("HOST_CPU_CORES"),
	Attribute("HOST_CUSTOM_METADATA"),
	Attribute("HOST_DETECTED_NAME"),
	Attribute("HOST_GROUP_ID"),
	Attribute("HOST_GROUP_NAME"),
	Attribute("HOST_HYPERVISOR_TYPE"),
	Attribute("HOST_IP_ADDRESS"),
	Attribute("HOST_KUBERNETES_LABELS"),
	Attribute("HOST_LOGICAL_CPU_CORES"),
	Attribute("HOST_NAME"),
	Attribute("HOST_ONEAGENT_CUSTOM_HOST_NAME"),
	Attribute("HOST_OS_TYPE"),
	Attribute("HOST_OS_VERSION"),
	Attribute("HOST_PAAS_MEMORY_LIMIT"),
	Attribute("HOST_PAAS_TYPE"),
	Attribute("HOST_TAGS"),
	Attribute("HOST_TECHNOLOGY"),
	Attribute("HTTP_MONITOR_NAME"),
	Attribute("HTTP_MONITOR_TAGS"),
	Attribute("KUBERNETES_CLUSTER_NAME"),
	Attribute("KUBERNETES_NODE_NAME"),
	Attribute("KUBERNETES_SERVICE_NAME"),
	Attribute("MOBILE_APPLICATION_NAME"),
	Attribute("MOBILE_APPLICATION_PLATFORM"),
	Attribute("MOBILE_APPLICATION_TAGS"),
	Attribute("NAME_OF_COMPUTE_NODE"),
	Attribute("OPENSTACK_ACCOUNT_NAME"),
	Attribute("OPENSTACK_ACCOUNT_PROJECT_NAME"),
	Attribute("OPENSTACK_AVAILABILITY_ZONE_NAME"),
	Attribute("OPENSTACK_PROJECT_NAME"),
	Attribute("OPENSTACK_REGION_NAME"),
	Attribute("OPENSTACK_VM_INSTANCE_TYPE"),
	Attribute("OPENSTACK_VM_NAME"),
	Attribute("OPENSTACK_VM_SECURITY_GROUP"),
	Attribute("PROCESS_GROUP_AZURE_HOST_NAME"),
	Attribute("PROCESS_GROUP_AZURE_SITE_NAME"),
	Attribute("PROCESS_GROUP_CUSTOM_METADATA"),
	Attribute("PROCESS_GROUP_DETECTED_NAME"),
	Attribute("PROCESS_GROUP_ID"),
	Attribute("PROCESS_GROUP_LISTEN_PORT"),
	Attribute("PROCESS_GROUP_NAME"),
	Attribute("PROCESS_GROUP_PREDEFINED_METADATA"),
	Attribute("PROCESS_GROUP_TAGS"),
	Attribute("PROCESS_GROUP_TECHNOLOGY"),
	Attribute("PROCESS_GROUP_TECHNOLOGY_EDITION"),
	Attribute("PROCESS_GROUP_TECHNOLOGY_VERSION"),
	Attribute("QUEUE_NAME"),
	Attribute("QUEUE_TECHNOLOGY"),
	Attribute("QUEUE_VENDOR"),
	Attribute("SERVICE_AKKA_ACTOR_SYSTEM"),
	Attribute("SERVICE_CTG_SERVICE_NAME"),
	Attribute("SERVICE_DATABASE_HOST_NAME"),
	Attribute("SERVICE_DATABASE_NAME"),
	Attribute("SERVICE_DATABASE_TOPOLOGY"),
	Attribute("SERVICE_DATABASE_VENDOR"),
	Attribute("SERVICE_DETECTED_NAME"),
	Attribute("SERVICE_ESB_APPLICATION_NAME"),
	Attribute("SERVICE_IBM_CTG_GATEWAY_URL"),
	Attribute("SERVICE_MESSAGING_LISTENER_CLASS_NAME"),
	Attribute("SERVICE_NAME"),
	Attribute("SERVICE_PORT"),
	Attribute("SERVICE_PUBLIC_DOMAIN_NAME"),
	Attribute("SERVICE_REMOTE_ENDPOINT"),
	Attribute("SERVICE_REMOTE_SERVICE_NAME"),
	Attribute("SERVICE_TAGS"),
	Attribute("SERVICE_TECHNOLOGY"),
	Attribute("SERVICE_TECHNOLOGY_EDITION"),
	Attribute("SERVICE_TECHNOLOGY_VERSION"),
	Attribute("SERVICE_TOPOLOGY"),
	Attribute("SERVICE_TYPE"),
	Attribute("SERVICE_WEB_APPLICATION_ID"),
	Attribute("SERVICE_WEB_CONTEXT_ROOT"),
	Attribute("SERVICE_WEB_SERVER_ENDPOINT"),
	Attribute("SERVICE_WEB_SERVER_NAME"),
	Attribute("SERVICE_WEB_SERVICE_NAME"),
	Attribute("SERVICE_WEB_SERVICE_NAMESPACE"),
	Attribute("VMWARE_DATACENTER_NAME"),
	Attribute("VMWARE_VM_NAME"),
	Attribute("WEB_APPLICATION_NAME"),
	Attribute("WEB_APPLICATION_NAME_PATTERN"),
	Attribute("WEB_APPLICATION_TAGS"),
	Attribute("WEB_APPLICATION_TYPE"),
}

type AutoTagMeType string

var AutoTagMeTypes = struct {
	Application                  AutoTagMeType
	AwsApplicationLoadBalancer   AutoTagMeType
	AwsClassicLoadBalancer       AutoTagMeType
	AwsNetworkLoadBalancer       AutoTagMeType
	AwsRelationalDatabaseService AutoTagMeType
	Azure                        AutoTagMeType
	CustomApplication            AutoTagMeType
	CustomDevice                 AutoTagMeType
	DcrumApplication             AutoTagMeType
	EsxiHost                     AutoTagMeType
	ExternalSyntheticTest        AutoTagMeType
	Host                         AutoTagMeType
	HttpCheck                    AutoTagMeType
	MobileApplication            AutoTagMeType
	ProcessGroup                 AutoTagMeType
	Service                      AutoTagMeType
	SyntheticTest                AutoTagMeType
}{
	AutoTagMeType("APPLICATION"),
	AutoTagMeType("AWS_APPLICATION_LOAD_BALANCER"),
	AutoTagMeType("AWS_CLASSIC_LOAD_BALANCER"),
	AutoTagMeType("AWS_NETWORK_LOAD_BALANCER"),
	AutoTagMeType("AWS_RELATIONAL_DATABASE_SERVICE"),
	AutoTagMeType("AZURE"),
	AutoTagMeType("CUSTOM_APPLICATION"),
	AutoTagMeType("CUSTOM_DEVICE"),
	AutoTagMeType("DCRUM_APPLICATION"),
	AutoTagMeType("ESXI_HOST"),
	AutoTagMeType("EXTERNAL_SYNTHETIC_TEST"),
	AutoTagMeType("HOST"),
	AutoTagMeType("HTTP_CHECK"),
	AutoTagMeType("MOBILE_APPLICATION"),
	AutoTagMeType("PROCESS_GROUP"),
	AutoTagMeType("SERVICE"),
	AutoTagMeType("SYNTHETIC_TEST"),
}

type Normalization string

var Normalizations = struct {
	LeavetextasIs Normalization
	Tolowercase   Normalization
	Touppercase   Normalization
}{
	Normalization("Leavetextas_is"),
	Normalization("Tolowercase"),
	Normalization("Touppercase"),
}

type Operator string

var Operators = struct {
	BeginsWith            Operator
	Contains              Operator
	EndsWith              Operator
	Equals                Operator
	Exists                Operator
	GreaterThan           Operator
	GreaterThanOrEqual    Operator
	IsIpInRange           Operator
	LowerThan             Operator
	LowerThanOrEqual      Operator
	NotBeginsWith         Operator
	NotContains           Operator
	NotEndsWith           Operator
	NotEquals             Operator
	NotExists             Operator
	NotGreaterThan        Operator
	NotGreaterThanOrEqual Operator
	NotIsIpInRange        Operator
	NotLowerThan          Operator
	NotLowerThanOrEqual   Operator
	NotRegexMatches       Operator
	NotTagKeyEquals       Operator
	RegexMatches          Operator
	TagKeyEquals          Operator
}{
	Operator("BEGINS_WITH"),
	Operator("CONTAINS"),
	Operator("ENDS_WITH"),
	Operator("EQUALS"),
	Operator("EXISTS"),
	Operator("GREATER_THAN"),
	Operator("GREATER_THAN_OR_EQUAL"),
	Operator("IS_IP_IN_RANGE"),
	Operator("LOWER_THAN"),
	Operator("LOWER_THAN_OR_EQUAL"),
	Operator("NOT_BEGINS_WITH"),
	Operator("NOT_CONTAINS"),
	Operator("NOT_ENDS_WITH"),
	Operator("NOT_EQUALS"),
	Operator("NOT_EXISTS"),
	Operator("NOT_GREATER_THAN"),
	Operator("NOT_GREATER_THAN_OR_EQUAL"),
	Operator("NOT_IS_IP_IN_RANGE"),
	Operator("NOT_LOWER_THAN"),
	Operator("NOT_LOWER_THAN_OR_EQUAL"),
	Operator("NOT_REGEX_MATCHES"),
	Operator("NOT_TAG_KEY_EQUALS"),
	Operator("REGEX_MATCHES"),
	Operator("TAG_KEY_EQUALS"),
}

type RuleType string

var RuleTypes = struct {
	Me       RuleType
	Selector RuleType
}{
	RuleType("ME"),
	RuleType("SELECTOR"),
}
