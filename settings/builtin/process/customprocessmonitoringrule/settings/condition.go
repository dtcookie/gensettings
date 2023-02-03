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

package customprocessmonitoringrule

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Condition struct {
	EnvVar   *string           `json:"envVar,omitempty"` // supported only with OneAgent 1.167+
	Item     AgentItemName     `json:"item"`             // Possible Values: `WEBLOGIC_NAME`, `UNKNOWN`, `WEBSPHERE_SERVER_NAME`, `TIBCO_BUSINESSWORKS_CE_VERSION`, `MSSQL_INSTANCE_NAME`, `EXE_PATH`, `WEBSPHERE_NODE_NAME`, `SOFTWAREAG_INSTALL_ROOT`, `JAVA_JAR_PATH`, `CATALINA_BASE`, `CLOUD_FOUNDRY_APP_NAME`, `CLOUD_FOUNDRY_INSTANCE_INDEX`, `COLDFUSION_JVM_CONFIG_FILE`, `APACHE_SPARK_MASTER_IP_ADDRESS`, `CONTAINER_NAME`, `AWS_REGION`, `COMMAND_LINE_ARGS`, `EQUINOX_CONFIG_PATH`, `IBM_CICS_IMS_JOBNAME`, `JBOSS_SERVER_NAME`, `TIBCO_BUSINESSWORKS_DOMAIN_NAME`, `KUBERNETES_NAMESPACE`, `DOTNET_COMMAND_PATH`, `KUBERNETES_FULLPODNAME`, `KUBERNETES_BASEPODNAME`, `DECLARATIVE_ID`, `IIS_ROLE_NAME`, `IBM_IMS_CONTROL`, `IBM_IMS_CONNECT`, `GLASSFISH_DOMAIN_NAME`, `IBM_IMS_SOAP_GW_NAME`, `TIPCO_BUSINESSWORKS_PROPERTY_FILE_PATH`, `NODEJS_APP_NAME`, `TIBCO_BUSINESSWORKS_APP_SPACE_NAME`, `AWS_ECS_CONTAINERNAME`, `CONTAINER_ID`, `IIS_APP_POOL`, `WEBSPHERE_CLUSTER_NAME`, `GAE_SERVICE`, `SOFTWAREAG_PRODUCTPROPNAME`, `SPRINGBOOT_STARTUP_CLASS`, `VARNISH_INSTANCE_NAME`, `PHP_CLI_WORKING_DIR`, `WEBLOGIC_CLUSTER_NAME`, `GOOGLE_CLOUD_PROJECT`, `APACHE_CONFIG_PATH`, `WEBLOGIC_HOME`, `HYBRIS_BIN_DIR`, `JAVA_MAIN_CLASS`, `IBM_CICS_REGION`, `SERVICE_NAME`, `TIBCO_BUSINESSWORKS_CE_APP_NAME`, `AWS_ECR_REGION`, `IIB_BROKER_NAME`, `SPRINGBOOT_APP_NAME`, `ELASTIC_SEARCH_CLUSTER_NAME`, `SPRINGBOOT_PROFILE_NAME`, `NODEJS_SCRIPT_NAME`, `JBOSS_HOME`, `PHP_CLI_SCRIPT_PATH`, `JBOSS_MODE`, `ORACLE_SID`, `ELASTIC_SEARCH_NODE_NAME`, `CLOUD_FOUNDRY_APPLICATION_ID`, `HYBRIS_DATA_DIR`, `AWS_ECS_CLUSTER`, `PG_ID_CALC_INPUT_KEY_LINKAGE`, `TIBCO_BUSINESSWORKS_APP_NODE_NAME`, `AWS_LAMBDA_FUNCTION_NAME`, `TIBCO_BUSINESSWORKS_HOME`, `CLOUD_FOUNDRY_SPACE_ID`, `CONTAINER_IMAGE_VERSION`, `GAE_INSTANCE`, `CLOUD_FOUNDRY_SPACE_NAME`, `WEBSPHERE_CELL_NAME`, `ASP_NET_CORE_APPLICATION_PATH`, `RUXIT_NODE_ID`, `GLASSFISH_INSTANCE_NAME`, `AWS_ECS_FAMILY`, `CONTAINER_IMAGE_NAME`, `IBM_CTG_NAME`, `RUXIT_CLUSTER_ID`, `DOTNET_COMMAND`, `HYBRIS_CONFIG_DIR`, `EXE_NAME`, `TIPCO_BUSINESSWORKS_PROPERTY_FILE`, `KUBERNETES_PODUID`, `IIB_EXECUTION_GROUP_NAME`, `IBM_CICS_IMS_APPLID`, `AWS_ECR_ACCOUNT_ID`, `NODEJS_APP_BASE_DIR`, `IBM_IMS_MPR`, `WEBLOGIC_DOMAIN_NAME`, `JAVA_JAR_FILE`, `KUBERNETES_CONTAINERNAME`, `WEBSPHERE_LIBERTY_SERVER_NAME`, `AWS_ECS_REVISION`, `CATALINA_HOME`
	Operator ConditionOperator `json:"operator"`         // Possible Values: `NOT_ENDS`, `CONTAINS`, `ENDS`, `EQUALS`, `EXISTS`, `NOT_CONTAINS`, `NOT_EQUALS`, `STARTS`, `NOT_EXISTS`, `NOT_STARTS`
	Value    *string           `json:"value,omitempty"`  // Condition value
}

func (me *Condition) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"env_var": {
			Type:        schema.TypeString,
			Description: "supported only with OneAgent 1.167+",
			Optional:    true,
		},
		"item": {
			Type:        schema.TypeString,
			Description: "Possible Values: `WEBLOGIC_NAME`, `UNKNOWN`, `WEBSPHERE_SERVER_NAME`, `TIBCO_BUSINESSWORKS_CE_VERSION`, `MSSQL_INSTANCE_NAME`, `EXE_PATH`, `WEBSPHERE_NODE_NAME`, `SOFTWAREAG_INSTALL_ROOT`, `JAVA_JAR_PATH`, `CATALINA_BASE`, `CLOUD_FOUNDRY_APP_NAME`, `CLOUD_FOUNDRY_INSTANCE_INDEX`, `COLDFUSION_JVM_CONFIG_FILE`, `APACHE_SPARK_MASTER_IP_ADDRESS`, `CONTAINER_NAME`, `AWS_REGION`, `COMMAND_LINE_ARGS`, `EQUINOX_CONFIG_PATH`, `IBM_CICS_IMS_JOBNAME`, `JBOSS_SERVER_NAME`, `TIBCO_BUSINESSWORKS_DOMAIN_NAME`, `KUBERNETES_NAMESPACE`, `DOTNET_COMMAND_PATH`, `KUBERNETES_FULLPODNAME`, `KUBERNETES_BASEPODNAME`, `DECLARATIVE_ID`, `IIS_ROLE_NAME`, `IBM_IMS_CONTROL`, `IBM_IMS_CONNECT`, `GLASSFISH_DOMAIN_NAME`, `IBM_IMS_SOAP_GW_NAME`, `TIPCO_BUSINESSWORKS_PROPERTY_FILE_PATH`, `NODEJS_APP_NAME`, `TIBCO_BUSINESSWORKS_APP_SPACE_NAME`, `AWS_ECS_CONTAINERNAME`, `CONTAINER_ID`, `IIS_APP_POOL`, `WEBSPHERE_CLUSTER_NAME`, `GAE_SERVICE`, `SOFTWAREAG_PRODUCTPROPNAME`, `SPRINGBOOT_STARTUP_CLASS`, `VARNISH_INSTANCE_NAME`, `PHP_CLI_WORKING_DIR`, `WEBLOGIC_CLUSTER_NAME`, `GOOGLE_CLOUD_PROJECT`, `APACHE_CONFIG_PATH`, `WEBLOGIC_HOME`, `HYBRIS_BIN_DIR`, `JAVA_MAIN_CLASS`, `IBM_CICS_REGION`, `SERVICE_NAME`, `TIBCO_BUSINESSWORKS_CE_APP_NAME`, `AWS_ECR_REGION`, `IIB_BROKER_NAME`, `SPRINGBOOT_APP_NAME`, `ELASTIC_SEARCH_CLUSTER_NAME`, `SPRINGBOOT_PROFILE_NAME`, `NODEJS_SCRIPT_NAME`, `JBOSS_HOME`, `PHP_CLI_SCRIPT_PATH`, `JBOSS_MODE`, `ORACLE_SID`, `ELASTIC_SEARCH_NODE_NAME`, `CLOUD_FOUNDRY_APPLICATION_ID`, `HYBRIS_DATA_DIR`, `AWS_ECS_CLUSTER`, `PG_ID_CALC_INPUT_KEY_LINKAGE`, `TIBCO_BUSINESSWORKS_APP_NODE_NAME`, `AWS_LAMBDA_FUNCTION_NAME`, `TIBCO_BUSINESSWORKS_HOME`, `CLOUD_FOUNDRY_SPACE_ID`, `CONTAINER_IMAGE_VERSION`, `GAE_INSTANCE`, `CLOUD_FOUNDRY_SPACE_NAME`, `WEBSPHERE_CELL_NAME`, `ASP_NET_CORE_APPLICATION_PATH`, `RUXIT_NODE_ID`, `GLASSFISH_INSTANCE_NAME`, `AWS_ECS_FAMILY`, `CONTAINER_IMAGE_NAME`, `IBM_CTG_NAME`, `RUXIT_CLUSTER_ID`, `DOTNET_COMMAND`, `HYBRIS_CONFIG_DIR`, `EXE_NAME`, `TIPCO_BUSINESSWORKS_PROPERTY_FILE`, `KUBERNETES_PODUID`, `IIB_EXECUTION_GROUP_NAME`, `IBM_CICS_IMS_APPLID`, `AWS_ECR_ACCOUNT_ID`, `NODEJS_APP_BASE_DIR`, `IBM_IMS_MPR`, `WEBLOGIC_DOMAIN_NAME`, `JAVA_JAR_FILE`, `KUBERNETES_CONTAINERNAME`, `WEBSPHERE_LIBERTY_SERVER_NAME`, `AWS_ECS_REVISION`, `CATALINA_HOME`",
			Required:    true,
		},
		"operator": {
			Type:        schema.TypeString,
			Description: "Possible Values: `NOT_ENDS`, `CONTAINS`, `ENDS`, `EQUALS`, `EXISTS`, `NOT_CONTAINS`, `NOT_EQUALS`, `STARTS`, `NOT_EXISTS`, `NOT_STARTS`",
			Required:    true,
		},
		"value": {
			Type:        schema.TypeString,
			Description: "Condition value",
			Optional:    true,
		},
	}
}

func (me *Condition) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"env_var":  me.EnvVar,
		"item":     me.Item,
		"operator": me.Operator,
		"value":    me.Value,
	})
}

func (me *Condition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"env_var":  &me.EnvVar,
		"item":     &me.Item,
		"operator": &me.Operator,
		"value":    &me.Value,
	})
}
