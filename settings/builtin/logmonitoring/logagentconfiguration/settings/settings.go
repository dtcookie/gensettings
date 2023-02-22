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

package logagentconfiguration

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	LAConfigContainerTimezoneHeuristicEnabled bool    `json:"LAConfigContainerTimezoneHeuristicEnabled"` // Enables automatic detection of timezone in container's logs if it is not explicitly defined in content or configured
	LAConfigContainersLogsDetectionEnabled    bool    `json:"LAConfigContainersLogsDetectionEnabled"`    // If set to FALSE - logs from containers won't be detected
	LAConfigDateSearchLimit_Bytes             int     `json:"LAConfigDateSearchLimit_Bytes"`             // Defines the number of characters in every log line (starting from the first character in the line) where the timestamp is searched.
	LAConfigDefaultTimezone                   string  `json:"LAConfigDefaultTimezone"`                   // Default timezone for agent if more specific configurations is not defined.
	LAConfigEventLogQueryTimeout_Sec          int     `json:"LAConfigEventLogQueryTimeout_Sec"`          // Defines the maximum timeout value, in seconds, for the query extracting Windows Event Logs
	LAConfigIISDetectionEnabled               bool    `json:"LAConfigIISDetectionEnabled"`               // If set to FALSE - IIS log detection mechanism will be disabled
	LAConfigLogScannerLinuxNfsEnabled         bool    `json:"LAConfigLogScannerLinuxNfsEnabled"`         // If set to FALSE - logs on NFS won't be detected
	LAConfigMaxLgisPerEntityCount             int     `json:"LAConfigMaxLgisPerEntityCount"`             // Defines the maximum number of log group instances per entity after which, the new automatic ones wouldn't be added.
	LAConfigMinBinaryDetectionLimit_Bytes     int     `json:"LAConfigMinBinaryDetectionLimit_Bytes"`     // Defines the minimum number of bytes in log file required for binary detection.
	LAConfigMonitorOwnLogsEnabled             bool    `json:"LAConfigMonitorOwnLogsEnabled"`             // If set to FALSE - logs produced by OneAgent won't be monitored. Enabling this option may affect your DDU consumption.
	LAConfigOpenLogFilesDetectionEnabled      bool    `json:"LAConfigOpenLogFilesDetectionEnabled"`      // Enables auto-detection of log files on hosts. If set to false, logs won't be auto-detected.
	LAConfigSeverityDetectionLimit_Bytes      int     `json:"LAConfigSeverityDetectionLimit_Bytes"`      // Defines the number of characters in every log line (starting from the first character in the line) where severity is searched.
	LAConfigSeverityDetectionLinesLimit       int     `json:"LAConfigSeverityDetectionLinesLimit"`       // Defines the number of the first lines of every log entry where severity is searched.
	LAConfigSystemLogsDetectionEnabled        bool    `json:"LAConfigSystemLogsDetectionEnabled"`        // If set to FALSE - system logs detection mechanism will be disabled (Linux: syslog, message log) (Windows: System, Application, Security event logs)
	LAConfigUTCAsDefaultContainerTimezone     bool    `json:"LAConfigUTCAsDefaultContainerTimezone"`     // Defines if UTC is used as a default timezone in containers. This option is deprecated for agents in version 1.247 or newer.
	Scope                                     *string `json:"-" scope:"scope"`                           // The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"laconfig_container_timezone_heuristic_enabled": {
			Type:        schema.TypeBool,
			Description: "Enables automatic detection of timezone in container's logs if it is not explicitly defined in content or configured",
			Required:    true,
		},
		"laconfig_containers_logs_detection_enabled": {
			Type:        schema.TypeBool,
			Description: "If set to FALSE - logs from containers won't be detected",
			Required:    true,
		},
		"laconfig_date_search_limit_bytes": {
			Type:        schema.TypeInt,
			Description: "Defines the number of characters in every log line (starting from the first character in the line) where the timestamp is searched.",
			Required:    true,
		},
		"laconfig_default_timezone": {
			Type:        schema.TypeString,
			Description: "Default timezone for agent if more specific configurations is not defined.",
			Required:    true,
		},
		"laconfig_event_log_query_timeout_sec": {
			Type:        schema.TypeInt,
			Description: "Defines the maximum timeout value, in seconds, for the query extracting Windows Event Logs",
			Required:    true,
		},
		"laconfig_iisdetection_enabled": {
			Type:        schema.TypeBool,
			Description: "If set to FALSE - IIS log detection mechanism will be disabled",
			Required:    true,
		},
		"laconfig_log_scanner_linux_nfs_enabled": {
			Type:        schema.TypeBool,
			Description: "If set to FALSE - logs on NFS won't be detected",
			Required:    true,
		},
		"laconfig_max_lgis_per_entity_count": {
			Type:        schema.TypeInt,
			Description: "Defines the maximum number of log group instances per entity after which, the new automatic ones wouldn't be added.",
			Required:    true,
		},
		"laconfig_min_binary_detection_limit_bytes": {
			Type:        schema.TypeInt,
			Description: "Defines the minimum number of bytes in log file required for binary detection.",
			Required:    true,
		},
		"laconfig_monitor_own_logs_enabled": {
			Type:        schema.TypeBool,
			Description: "If set to FALSE - logs produced by OneAgent won't be monitored. Enabling this option may affect your DDU consumption.",
			Required:    true,
		},
		"laconfig_open_log_files_detection_enabled": {
			Type:        schema.TypeBool,
			Description: "Enables auto-detection of log files on hosts. If set to false, logs won't be auto-detected.",
			Required:    true,
		},
		"laconfig_severity_detection_limit_bytes": {
			Type:        schema.TypeInt,
			Description: "Defines the number of characters in every log line (starting from the first character in the line) where severity is searched.",
			Required:    true,
		},
		"laconfig_severity_detection_lines_limit": {
			Type:        schema.TypeInt,
			Description: "Defines the number of the first lines of every log entry where severity is searched.",
			Required:    true,
		},
		"laconfig_system_logs_detection_enabled": {
			Type:        schema.TypeBool,
			Description: "If set to FALSE - system logs detection mechanism will be disabled (Linux: syslog, message log) (Windows: System, Application, Security event logs)",
			Required:    true,
		},
		"laconfig_utcas_default_container_timezone": {
			Type:        schema.TypeBool,
			Description: "Defines if UTC is used as a default timezone in containers. This option is deprecated for agents in version 1.247 or newer.",
			Required:    true,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.",
			Optional:    true,
			Default:     "environment",
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"laconfig_container_timezone_heuristic_enabled": me.LAConfigContainerTimezoneHeuristicEnabled,
		"laconfig_containers_logs_detection_enabled":    me.LAConfigContainersLogsDetectionEnabled,
		"laconfig_date_search_limit_bytes":              me.LAConfigDateSearchLimit_Bytes,
		"laconfig_default_timezone":                     me.LAConfigDefaultTimezone,
		"laconfig_event_log_query_timeout_sec":          me.LAConfigEventLogQueryTimeout_Sec,
		"laconfig_iisdetection_enabled":                 me.LAConfigIISDetectionEnabled,
		"laconfig_log_scanner_linux_nfs_enabled":        me.LAConfigLogScannerLinuxNfsEnabled,
		"laconfig_max_lgis_per_entity_count":            me.LAConfigMaxLgisPerEntityCount,
		"laconfig_min_binary_detection_limit_bytes":     me.LAConfigMinBinaryDetectionLimit_Bytes,
		"laconfig_monitor_own_logs_enabled":             me.LAConfigMonitorOwnLogsEnabled,
		"laconfig_open_log_files_detection_enabled":     me.LAConfigOpenLogFilesDetectionEnabled,
		"laconfig_severity_detection_limit_bytes":       me.LAConfigSeverityDetectionLimit_Bytes,
		"laconfig_severity_detection_lines_limit":       me.LAConfigSeverityDetectionLinesLimit,
		"laconfig_system_logs_detection_enabled":        me.LAConfigSystemLogsDetectionEnabled,
		"laconfig_utcas_default_container_timezone":     me.LAConfigUTCAsDefaultContainerTimezone,
		"scope": me.Scope,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"laconfig_container_timezone_heuristic_enabled": &me.LAConfigContainerTimezoneHeuristicEnabled,
		"laconfig_containers_logs_detection_enabled":    &me.LAConfigContainersLogsDetectionEnabled,
		"laconfig_date_search_limit_bytes":              &me.LAConfigDateSearchLimit_Bytes,
		"laconfig_default_timezone":                     &me.LAConfigDefaultTimezone,
		"laconfig_event_log_query_timeout_sec":          &me.LAConfigEventLogQueryTimeout_Sec,
		"laconfig_iisdetection_enabled":                 &me.LAConfigIISDetectionEnabled,
		"laconfig_log_scanner_linux_nfs_enabled":        &me.LAConfigLogScannerLinuxNfsEnabled,
		"laconfig_max_lgis_per_entity_count":            &me.LAConfigMaxLgisPerEntityCount,
		"laconfig_min_binary_detection_limit_bytes":     &me.LAConfigMinBinaryDetectionLimit_Bytes,
		"laconfig_monitor_own_logs_enabled":             &me.LAConfigMonitorOwnLogsEnabled,
		"laconfig_open_log_files_detection_enabled":     &me.LAConfigOpenLogFilesDetectionEnabled,
		"laconfig_severity_detection_limit_bytes":       &me.LAConfigSeverityDetectionLimit_Bytes,
		"laconfig_severity_detection_lines_limit":       &me.LAConfigSeverityDetectionLinesLimit,
		"laconfig_system_logs_detection_enabled":        &me.LAConfigSystemLogsDetectionEnabled,
		"laconfig_utcas_default_container_timezone":     &me.LAConfigUTCAsDefaultContainerTimezone,
		"scope": &me.Scope,
	})
}
