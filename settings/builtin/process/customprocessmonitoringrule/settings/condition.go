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
	Item     AgentItemName     `json:"item"`             // Possible Values: `CONTAINER_IMAGE_NAME`, `WEBSPHERE_LIBERTY_SERVER_NAME`, `EXE_NAME`, `WEBLOGIC_NAME`, `AWS_ECR_REGION`, `CLOUD_FOUNDRY_APPLICATION_ID`, `IBM_IMS_SOAP_GW_NAME`, `RUXIT_NODE_ID`, `GAE_SERVICE`, `EQUINOX_CONFIG_PATH`, `WEBLOGIC_HOME`, `AWS_ECR_ACCOUNT_ID`, `GAE_INSTANCE`, `JAVA_MAIN_CLASS`, `JBOSS_MODE`, `TIBCO_BUSINESSWORKS_HOME`, `CATALINA_HOME`, `TIBCO_BUSINESSWORKS_APP_NODE_NAME`, `IIS_ROLE_NAME`, `GLASSFISH_INSTANCE_NAME`, `IIS_APP_POOL`, `SPRINGBOOT_STARTUP_CLASS`, `TIPCO_BUSINESSWORKS_PROPERTY_FILE`, `IBM_CICS_IMS_JOBNAME`, `CLOUD_FOUNDRY_SPACE_NAME`, `TIBCO_BUSINESSWORKS_CE_APP_NAME`, `DECLARATIVE_ID`, `DOTNET_COMMAND_PATH`, `KUBERNETES_PODUID`, `HYBRIS_CONFIG_DIR`, `AWS_ECS_CLUSTER`, `APACHE_CONFIG_PATH`, `PHP_CLI_SCRIPT_PATH`, `WEBSPHERE_CLUSTER_NAME`, `HYBRIS_BIN_DIR`, `ELASTIC_SEARCH_CLUSTER_NAME`, `GOOGLE_CLOUD_PROJECT`, `PHP_CLI_WORKING_DIR`, `IBM_CTG_NAME`, `IIB_BROKER_NAME`, `IBM_IMS_MPR`, `WEBLOGIC_DOMAIN_NAME`, `KUBERNETES_CONTAINERNAME`, `CLOUD_FOUNDRY_INSTANCE_INDEX`, `HYBRIS_DATA_DIR`, `IBM_CICS_IMS_APPLID`, `JAVA_JAR_FILE`, `WEBSPHERE_CELL_NAME`, `PG_ID_CALC_INPUT_KEY_LINKAGE`, `MSSQL_INSTANCE_NAME`, `CLOUD_FOUNDRY_APP_NAME`, `JBOSS_SERVER_NAME`, `COMMAND_LINE_ARGS`, `AWS_LAMBDA_FUNCTION_NAME`, `DOTNET_COMMAND`, `WEBSPHERE_SERVER_NAME`, `TIBCO_BUSINESSWORKS_APP_SPACE_NAME`, `RUXIT_CLUSTER_ID`, `ELASTIC_SEARCH_NODE_NAME`, `TIBCO_BUSINESSWORKS_CE_VERSION`, `UNKNOWN`, `ORACLE_SID`, `AWS_ECS_CONTAINERNAME`, `KUBERNETES_FULLPODNAME`, `WEBLOGIC_CLUSTER_NAME`, `JAVA_JAR_PATH`, `CLOUD_FOUNDRY_SPACE_ID`, `KUBERNETES_NAMESPACE`, `KUBERNETES_BASEPODNAME`, `TIBCO_BUSINESSWORKS_DOMAIN_NAME`, `ASP_NET_CORE_APPLICATION_PATH`, `SOFTWAREAG_INSTALL_ROOT`, `NODEJS_SCRIPT_NAME`, `IIB_EXECUTION_GROUP_NAME`, `GLASSFISH_DOMAIN_NAME`, `APACHE_SPARK_MASTER_IP_ADDRESS`, `CATALINA_BASE`, `NODEJS_APP_NAME`, `SPRINGBOOT_APP_NAME`, `SOFTWAREAG_PRODUCTPROPNAME`, `AWS_ECS_FAMILY`, `SERVICE_NAME`, `CONTAINER_NAME`, `AWS_REGION`, `NODEJS_APP_BASE_DIR`, `IBM_CICS_REGION`, `IBM_IMS_CONTROL`, `VARNISH_INSTANCE_NAME`, `CONTAINER_ID`, `WEBSPHERE_NODE_NAME`, `SPRINGBOOT_PROFILE_NAME`, `EXE_PATH`, `IBM_IMS_CONNECT`, `TIPCO_BUSINESSWORKS_PROPERTY_FILE_PATH`, `CONTAINER_IMAGE_VERSION`, `JBOSS_HOME`, `AWS_ECS_REVISION`, `COLDFUSION_JVM_CONFIG_FILE`
	Operator ConditionOperator `json:"operator"`         // Possible Values: `NOT_EXISTS`, `CONTAINS`, `EXISTS`, `NOT_EQUALS`, `EQUALS`, `STARTS`, `NOT_STARTS`, `ENDS`, `NOT_ENDS`, `NOT_CONTAINS`
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
			Description: "Possible Values: `CONTAINER_IMAGE_NAME`, `WEBSPHERE_LIBERTY_SERVER_NAME`, `EXE_NAME`, `WEBLOGIC_NAME`, `AWS_ECR_REGION`, `CLOUD_FOUNDRY_APPLICATION_ID`, `IBM_IMS_SOAP_GW_NAME`, `RUXIT_NODE_ID`, `GAE_SERVICE`, `EQUINOX_CONFIG_PATH`, `WEBLOGIC_HOME`, `AWS_ECR_ACCOUNT_ID`, `GAE_INSTANCE`, `JAVA_MAIN_CLASS`, `JBOSS_MODE`, `TIBCO_BUSINESSWORKS_HOME`, `CATALINA_HOME`, `TIBCO_BUSINESSWORKS_APP_NODE_NAME`, `IIS_ROLE_NAME`, `GLASSFISH_INSTANCE_NAME`, `IIS_APP_POOL`, `SPRINGBOOT_STARTUP_CLASS`, `TIPCO_BUSINESSWORKS_PROPERTY_FILE`, `IBM_CICS_IMS_JOBNAME`, `CLOUD_FOUNDRY_SPACE_NAME`, `TIBCO_BUSINESSWORKS_CE_APP_NAME`, `DECLARATIVE_ID`, `DOTNET_COMMAND_PATH`, `KUBERNETES_PODUID`, `HYBRIS_CONFIG_DIR`, `AWS_ECS_CLUSTER`, `APACHE_CONFIG_PATH`, `PHP_CLI_SCRIPT_PATH`, `WEBSPHERE_CLUSTER_NAME`, `HYBRIS_BIN_DIR`, `ELASTIC_SEARCH_CLUSTER_NAME`, `GOOGLE_CLOUD_PROJECT`, `PHP_CLI_WORKING_DIR`, `IBM_CTG_NAME`, `IIB_BROKER_NAME`, `IBM_IMS_MPR`, `WEBLOGIC_DOMAIN_NAME`, `KUBERNETES_CONTAINERNAME`, `CLOUD_FOUNDRY_INSTANCE_INDEX`, `HYBRIS_DATA_DIR`, `IBM_CICS_IMS_APPLID`, `JAVA_JAR_FILE`, `WEBSPHERE_CELL_NAME`, `PG_ID_CALC_INPUT_KEY_LINKAGE`, `MSSQL_INSTANCE_NAME`, `CLOUD_FOUNDRY_APP_NAME`, `JBOSS_SERVER_NAME`, `COMMAND_LINE_ARGS`, `AWS_LAMBDA_FUNCTION_NAME`, `DOTNET_COMMAND`, `WEBSPHERE_SERVER_NAME`, `TIBCO_BUSINESSWORKS_APP_SPACE_NAME`, `RUXIT_CLUSTER_ID`, `ELASTIC_SEARCH_NODE_NAME`, `TIBCO_BUSINESSWORKS_CE_VERSION`, `UNKNOWN`, `ORACLE_SID`, `AWS_ECS_CONTAINERNAME`, `KUBERNETES_FULLPODNAME`, `WEBLOGIC_CLUSTER_NAME`, `JAVA_JAR_PATH`, `CLOUD_FOUNDRY_SPACE_ID`, `KUBERNETES_NAMESPACE`, `KUBERNETES_BASEPODNAME`, `TIBCO_BUSINESSWORKS_DOMAIN_NAME`, `ASP_NET_CORE_APPLICATION_PATH`, `SOFTWAREAG_INSTALL_ROOT`, `NODEJS_SCRIPT_NAME`, `IIB_EXECUTION_GROUP_NAME`, `GLASSFISH_DOMAIN_NAME`, `APACHE_SPARK_MASTER_IP_ADDRESS`, `CATALINA_BASE`, `NODEJS_APP_NAME`, `SPRINGBOOT_APP_NAME`, `SOFTWAREAG_PRODUCTPROPNAME`, `AWS_ECS_FAMILY`, `SERVICE_NAME`, `CONTAINER_NAME`, `AWS_REGION`, `NODEJS_APP_BASE_DIR`, `IBM_CICS_REGION`, `IBM_IMS_CONTROL`, `VARNISH_INSTANCE_NAME`, `CONTAINER_ID`, `WEBSPHERE_NODE_NAME`, `SPRINGBOOT_PROFILE_NAME`, `EXE_PATH`, `IBM_IMS_CONNECT`, `TIPCO_BUSINESSWORKS_PROPERTY_FILE_PATH`, `CONTAINER_IMAGE_VERSION`, `JBOSS_HOME`, `AWS_ECS_REVISION`, `COLDFUSION_JVM_CONFIG_FILE`",
			Required:    true,
		},
		"operator": {
			Type:        schema.TypeString,
			Description: "Possible Values: `NOT_EXISTS`, `CONTAINS`, `EXISTS`, `NOT_EQUALS`, `EQUALS`, `STARTS`, `NOT_STARTS`, `ENDS`, `NOT_ENDS`, `NOT_CONTAINS`",
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
