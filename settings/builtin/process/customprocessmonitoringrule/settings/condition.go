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
	Item     AgentItemName     `json:"item"`             // Possible Values: `CONTAINER_IMAGE_NAME`, `EXE_PATH`, `WEBSPHERE_CLUSTER_NAME`, `TIBCO_BUSINESSWORKS_APP_NODE_NAME`, `PHP_CLI_SCRIPT_PATH`, `HYBRIS_DATA_DIR`, `WEBLOGIC_HOME`, `KUBERNETES_NAMESPACE`, `EXE_NAME`, `AWS_LAMBDA_FUNCTION_NAME`, `DOTNET_COMMAND`, `IBM_IMS_CONNECT`, `WEBLOGIC_DOMAIN_NAME`, `HYBRIS_BIN_DIR`, `CLOUD_FOUNDRY_INSTANCE_INDEX`, `DOTNET_COMMAND_PATH`, `TIBCO_BUSINESSWORKS_CE_APP_NAME`, `JBOSS_SERVER_NAME`, `COLDFUSION_JVM_CONFIG_FILE`, `WEBLOGIC_NAME`, `IIB_EXECUTION_GROUP_NAME`, `IBM_CICS_IMS_APPLID`, `MSSQL_INSTANCE_NAME`, `CLOUD_FOUNDRY_SPACE_NAME`, `AWS_ECS_FAMILY`, `IIB_BROKER_NAME`, `APACHE_CONFIG_PATH`, `TIBCO_BUSINESSWORKS_CE_VERSION`, `CONTAINER_ID`, `CONTAINER_NAME`, `NODEJS_SCRIPT_NAME`, `WEBLOGIC_CLUSTER_NAME`, `AWS_ECR_ACCOUNT_ID`, `IIS_APP_POOL`, `CATALINA_HOME`, `CLOUD_FOUNDRY_APP_NAME`, `GLASSFISH_INSTANCE_NAME`, `IIS_ROLE_NAME`, `ELASTIC_SEARCH_CLUSTER_NAME`, `JAVA_MAIN_CLASS`, `TIPCO_BUSINESSWORKS_PROPERTY_FILE_PATH`, `CONTAINER_IMAGE_VERSION`, `GLASSFISH_DOMAIN_NAME`, `ASP_NET_CORE_APPLICATION_PATH`, `AWS_ECS_REVISION`, `GAE_SERVICE`, `GOOGLE_CLOUD_PROJECT`, `KUBERNETES_FULLPODNAME`, `PG_ID_CALC_INPUT_KEY_LINKAGE`, `IBM_CICS_REGION`, `TIBCO_BUSINESSWORKS_DOMAIN_NAME`, `CATALINA_BASE`, `APACHE_SPARK_MASTER_IP_ADDRESS`, `GAE_INSTANCE`, `AWS_ECS_CONTAINERNAME`, `RUXIT_NODE_ID`, `IBM_IMS_MPR`, `AWS_ECR_REGION`, `CLOUD_FOUNDRY_SPACE_ID`, `WEBSPHERE_CELL_NAME`, `COMMAND_LINE_ARGS`, `WEBSPHERE_LIBERTY_SERVER_NAME`, `SOFTWAREAG_INSTALL_ROOT`, `RUXIT_CLUSTER_ID`, `KUBERNETES_PODUID`, `ORACLE_SID`, `JAVA_JAR_FILE`, `WEBSPHERE_NODE_NAME`, `SPRINGBOOT_STARTUP_CLASS`, `EQUINOX_CONFIG_PATH`, `JAVA_JAR_PATH`, `WEBSPHERE_SERVER_NAME`, `IBM_CICS_IMS_JOBNAME`, `CLOUD_FOUNDRY_APPLICATION_ID`, `TIPCO_BUSINESSWORKS_PROPERTY_FILE`, `SERVICE_NAME`, `IBM_IMS_CONTROL`, `SPRINGBOOT_APP_NAME`, `HYBRIS_CONFIG_DIR`, `KUBERNETES_BASEPODNAME`, `AWS_REGION`, `PHP_CLI_WORKING_DIR`, `IBM_IMS_SOAP_GW_NAME`, `ELASTIC_SEARCH_NODE_NAME`, `UNKNOWN`, `TIBCO_BUSINESSWORKS_APP_SPACE_NAME`, `IBM_CTG_NAME`, `NODEJS_APP_NAME`, `JBOSS_MODE`, `JBOSS_HOME`, `VARNISH_INSTANCE_NAME`, `SOFTWAREAG_PRODUCTPROPNAME`, `TIBCO_BUSINESSWORKS_HOME`, `DECLARATIVE_ID`, `SPRINGBOOT_PROFILE_NAME`, `AWS_ECS_CLUSTER`, `NODEJS_APP_BASE_DIR`, `KUBERNETES_CONTAINERNAME`
	Operator ConditionOperator `json:"operator"`         // Possible Values: `EXISTS`, `NOT_STARTS`, `CONTAINS`, `NOT_EXISTS`, `STARTS`, `NOT_CONTAINS`, `EQUALS`, `NOT_EQUALS`, `ENDS`, `NOT_ENDS`
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
			Description: "Possible Values: `CONTAINER_IMAGE_NAME`, `EXE_PATH`, `WEBSPHERE_CLUSTER_NAME`, `TIBCO_BUSINESSWORKS_APP_NODE_NAME`, `PHP_CLI_SCRIPT_PATH`, `HYBRIS_DATA_DIR`, `WEBLOGIC_HOME`, `KUBERNETES_NAMESPACE`, `EXE_NAME`, `AWS_LAMBDA_FUNCTION_NAME`, `DOTNET_COMMAND`, `IBM_IMS_CONNECT`, `WEBLOGIC_DOMAIN_NAME`, `HYBRIS_BIN_DIR`, `CLOUD_FOUNDRY_INSTANCE_INDEX`, `DOTNET_COMMAND_PATH`, `TIBCO_BUSINESSWORKS_CE_APP_NAME`, `JBOSS_SERVER_NAME`, `COLDFUSION_JVM_CONFIG_FILE`, `WEBLOGIC_NAME`, `IIB_EXECUTION_GROUP_NAME`, `IBM_CICS_IMS_APPLID`, `MSSQL_INSTANCE_NAME`, `CLOUD_FOUNDRY_SPACE_NAME`, `AWS_ECS_FAMILY`, `IIB_BROKER_NAME`, `APACHE_CONFIG_PATH`, `TIBCO_BUSINESSWORKS_CE_VERSION`, `CONTAINER_ID`, `CONTAINER_NAME`, `NODEJS_SCRIPT_NAME`, `WEBLOGIC_CLUSTER_NAME`, `AWS_ECR_ACCOUNT_ID`, `IIS_APP_POOL`, `CATALINA_HOME`, `CLOUD_FOUNDRY_APP_NAME`, `GLASSFISH_INSTANCE_NAME`, `IIS_ROLE_NAME`, `ELASTIC_SEARCH_CLUSTER_NAME`, `JAVA_MAIN_CLASS`, `TIPCO_BUSINESSWORKS_PROPERTY_FILE_PATH`, `CONTAINER_IMAGE_VERSION`, `GLASSFISH_DOMAIN_NAME`, `ASP_NET_CORE_APPLICATION_PATH`, `AWS_ECS_REVISION`, `GAE_SERVICE`, `GOOGLE_CLOUD_PROJECT`, `KUBERNETES_FULLPODNAME`, `PG_ID_CALC_INPUT_KEY_LINKAGE`, `IBM_CICS_REGION`, `TIBCO_BUSINESSWORKS_DOMAIN_NAME`, `CATALINA_BASE`, `APACHE_SPARK_MASTER_IP_ADDRESS`, `GAE_INSTANCE`, `AWS_ECS_CONTAINERNAME`, `RUXIT_NODE_ID`, `IBM_IMS_MPR`, `AWS_ECR_REGION`, `CLOUD_FOUNDRY_SPACE_ID`, `WEBSPHERE_CELL_NAME`, `COMMAND_LINE_ARGS`, `WEBSPHERE_LIBERTY_SERVER_NAME`, `SOFTWAREAG_INSTALL_ROOT`, `RUXIT_CLUSTER_ID`, `KUBERNETES_PODUID`, `ORACLE_SID`, `JAVA_JAR_FILE`, `WEBSPHERE_NODE_NAME`, `SPRINGBOOT_STARTUP_CLASS`, `EQUINOX_CONFIG_PATH`, `JAVA_JAR_PATH`, `WEBSPHERE_SERVER_NAME`, `IBM_CICS_IMS_JOBNAME`, `CLOUD_FOUNDRY_APPLICATION_ID`, `TIPCO_BUSINESSWORKS_PROPERTY_FILE`, `SERVICE_NAME`, `IBM_IMS_CONTROL`, `SPRINGBOOT_APP_NAME`, `HYBRIS_CONFIG_DIR`, `KUBERNETES_BASEPODNAME`, `AWS_REGION`, `PHP_CLI_WORKING_DIR`, `IBM_IMS_SOAP_GW_NAME`, `ELASTIC_SEARCH_NODE_NAME`, `UNKNOWN`, `TIBCO_BUSINESSWORKS_APP_SPACE_NAME`, `IBM_CTG_NAME`, `NODEJS_APP_NAME`, `JBOSS_MODE`, `JBOSS_HOME`, `VARNISH_INSTANCE_NAME`, `SOFTWAREAG_PRODUCTPROPNAME`, `TIBCO_BUSINESSWORKS_HOME`, `DECLARATIVE_ID`, `SPRINGBOOT_PROFILE_NAME`, `AWS_ECS_CLUSTER`, `NODEJS_APP_BASE_DIR`, `KUBERNETES_CONTAINERNAME`",
			Required:    true,
		},
		"operator": {
			Type:        schema.TypeString,
			Description: "Possible Values: `EXISTS`, `NOT_STARTS`, `CONTAINS`, `NOT_EXISTS`, `STARTS`, `NOT_CONTAINS`, `EQUALS`, `NOT_EQUALS`, `ENDS`, `NOT_ENDS`",
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
