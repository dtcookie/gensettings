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

package profile

type AlertingProfileEventFilterType string

var AlertingProfileEventFilterTypes = struct {
	Custom     AlertingProfileEventFilterType
	Predefined AlertingProfileEventFilterType
}{
	AlertingProfileEventFilterType("CUSTOM"),
	AlertingProfileEventFilterType("PREDEFINED"),
}

type ComparisonOperator string

var ComparisonOperators = struct {
	BeginsWith   ComparisonOperator
	Contains     ComparisonOperator
	EndsWith     ComparisonOperator
	RegexMatches ComparisonOperator
	StringEquals ComparisonOperator
}{
	ComparisonOperator("BEGINS_WITH"),
	ComparisonOperator("CONTAINS"),
	ComparisonOperator("ENDS_WITH"),
	ComparisonOperator("REGEX_MATCHES"),
	ComparisonOperator("STRING_EQUALS"),
}

type EventType string

var EventTypes = struct {
	ApplicationErrorRateIncreased       EventType
	ApplicationSlowdown                 EventType
	ApplicationUnexpectedHighLoad       EventType
	ApplicationUnexpectedLowLoad        EventType
	AwsLambdaHighErrorRate              EventType
	CustomAppCrashRateIncreased         EventType
	CustomApplicationErrorRateIncreased EventType
	CustomApplicationSlowdown           EventType
	CustomApplicationUnexpectedHighLoad EventType
	CustomApplicationUnexpectedLowLoad  EventType
	DatabaseConnectionFailure           EventType
	EbsVolumeHighLatency                EventType
	Ec2HighCpu                          EventType
	ElbHighBackendErrorRate             EventType
	EsxiGuestActiveSwapWait             EventType
	EsxiGuestCpuLimitReached            EventType
	EsxiHostCpuSaturation               EventType
	EsxiHostDatastoreLowDiskSpace       EventType
	EsxiHostDiskQueueSlow               EventType
	EsxiHostDiskSlow                    EventType
	EsxiHostMemorySaturation            EventType
	EsxiHostNetworkProblems             EventType
	EsxiHostOverloadedStorage           EventType
	EsxiVmImpactHostCpuSaturation       EventType
	EsxiVmImpactHostMemorySaturation    EventType
	ExternalSyntheticTestOutage         EventType
	ExternalSyntheticTestSlowdown       EventType
	HostOfServiceUnavailable            EventType
	HttpCheckGlobalOutage               EventType
	HttpCheckLocalOutage                EventType
	HttpCheckTestLocationSlowdown       EventType
	MobileAppCrashRateIncreased         EventType
	MobileApplicationErrorRateIncreased EventType
	MobileApplicationSlowdown           EventType
	MobileApplicationUnexpectedHighLoad EventType
	MobileApplicationUnexpectedLowLoad  EventType
	MonitoringUnavailable               EventType
	OsiDiskLowInodes                    EventType
	OsiGracefullyShutdown               EventType
	OsiHighCpu                          EventType
	OsiHighMemory                       EventType
	OsiLowDiskSpace                     EventType
	OsiNicDroppedPacketsHigh            EventType
	OsiNicErrorsHigh                    EventType
	OsiNicUtilizationHigh               EventType
	OsiSlowDisk                         EventType
	OsiUnexpectedlyUnavailable          EventType
	PgLowInstanceCount                  EventType
	PgiOfServiceUnavailable             EventType
	PgiUnavailable                      EventType
	ProcessCrashed                      EventType
	ProcessHighGcActivity               EventType
	ProcessMemoryResourceExhausted      EventType
	ProcessNaHighConnFailRate           EventType
	ProcessNaHighLossRate               EventType
	ProcessThreadsResourceExhausted     EventType
	RdsHighCpu                          EventType
	RdsHighLatency                      EventType
	RdsLowMemory                        EventType
	RdsLowStorageSpace                  EventType
	RdsOfServiceUnavailable             EventType
	RdsRestartSequence                  EventType
	ServiceErrorRateIncreased           EventType
	ServiceSlowdown                     EventType
	ServiceUnexpectedHighLoad           EventType
	ServiceUnexpectedLowLoad            EventType
	SyntheticGlobalOutage               EventType
	SyntheticLocalOutage                EventType
	SyntheticNodeOutage                 EventType
	SyntheticPrivateLocationOutage      EventType
	SyntheticTestLocationSlowdown       EventType
}{
	EventType("APPLICATION_ERROR_RATE_INCREASED"),
	EventType("APPLICATION_SLOWDOWN"),
	EventType("APPLICATION_UNEXPECTED_HIGH_LOAD"),
	EventType("APPLICATION_UNEXPECTED_LOW_LOAD"),
	EventType("AWS_LAMBDA_HIGH_ERROR_RATE"),
	EventType("CUSTOM_APP_CRASH_RATE_INCREASED"),
	EventType("CUSTOM_APPLICATION_ERROR_RATE_INCREASED"),
	EventType("CUSTOM_APPLICATION_SLOWDOWN"),
	EventType("CUSTOM_APPLICATION_UNEXPECTED_HIGH_LOAD"),
	EventType("CUSTOM_APPLICATION_UNEXPECTED_LOW_LOAD"),
	EventType("DATABASE_CONNECTION_FAILURE"),
	EventType("EBS_VOLUME_HIGH_LATENCY"),
	EventType("EC2_HIGH_CPU"),
	EventType("ELB_HIGH_BACKEND_ERROR_RATE"),
	EventType("ESXI_GUEST_ACTIVE_SWAP_WAIT"),
	EventType("ESXI_GUEST_CPU_LIMIT_REACHED"),
	EventType("ESXI_HOST_CPU_SATURATION"),
	EventType("ESXI_HOST_DATASTORE_LOW_DISK_SPACE"),
	EventType("ESXI_HOST_DISK_QUEUE_SLOW"),
	EventType("ESXI_HOST_DISK_SLOW"),
	EventType("ESXI_HOST_MEMORY_SATURATION"),
	EventType("ESXI_HOST_NETWORK_PROBLEMS"),
	EventType("ESXI_HOST_OVERLOADED_STORAGE"),
	EventType("ESXI_VM_IMPACT_HOST_CPU_SATURATION"),
	EventType("ESXI_VM_IMPACT_HOST_MEMORY_SATURATION"),
	EventType("EXTERNAL_SYNTHETIC_TEST_OUTAGE"),
	EventType("EXTERNAL_SYNTHETIC_TEST_SLOWDOWN"),
	EventType("HOST_OF_SERVICE_UNAVAILABLE"),
	EventType("HTTP_CHECK_GLOBAL_OUTAGE"),
	EventType("HTTP_CHECK_LOCAL_OUTAGE"),
	EventType("HTTP_CHECK_TEST_LOCATION_SLOWDOWN"),
	EventType("MOBILE_APP_CRASH_RATE_INCREASED"),
	EventType("MOBILE_APPLICATION_ERROR_RATE_INCREASED"),
	EventType("MOBILE_APPLICATION_SLOWDOWN"),
	EventType("MOBILE_APPLICATION_UNEXPECTED_HIGH_LOAD"),
	EventType("MOBILE_APPLICATION_UNEXPECTED_LOW_LOAD"),
	EventType("MONITORING_UNAVAILABLE"),
	EventType("OSI_DISK_LOW_INODES"),
	EventType("OSI_GRACEFULLY_SHUTDOWN"),
	EventType("OSI_HIGH_CPU"),
	EventType("OSI_HIGH_MEMORY"),
	EventType("OSI_LOW_DISK_SPACE"),
	EventType("OSI_NIC_DROPPED_PACKETS_HIGH"),
	EventType("OSI_NIC_ERRORS_HIGH"),
	EventType("OSI_NIC_UTILIZATION_HIGH"),
	EventType("OSI_SLOW_DISK"),
	EventType("OSI_UNEXPECTEDLY_UNAVAILABLE"),
	EventType("PG_LOW_INSTANCE_COUNT"),
	EventType("PGI_OF_SERVICE_UNAVAILABLE"),
	EventType("PGI_UNAVAILABLE"),
	EventType("PROCESS_CRASHED"),
	EventType("PROCESS_HIGH_GC_ACTIVITY"),
	EventType("PROCESS_MEMORY_RESOURCE_EXHAUSTED"),
	EventType("PROCESS_NA_HIGH_CONN_FAIL_RATE"),
	EventType("PROCESS_NA_HIGH_LOSS_RATE"),
	EventType("PROCESS_THREADS_RESOURCE_EXHAUSTED"),
	EventType("RDS_HIGH_CPU"),
	EventType("RDS_HIGH_LATENCY"),
	EventType("RDS_LOW_MEMORY"),
	EventType("RDS_LOW_STORAGE_SPACE"),
	EventType("RDS_OF_SERVICE_UNAVAILABLE"),
	EventType("RDS_RESTART_SEQUENCE"),
	EventType("SERVICE_ERROR_RATE_INCREASED"),
	EventType("SERVICE_SLOWDOWN"),
	EventType("SERVICE_UNEXPECTED_HIGH_LOAD"),
	EventType("SERVICE_UNEXPECTED_LOW_LOAD"),
	EventType("SYNTHETIC_GLOBAL_OUTAGE"),
	EventType("SYNTHETIC_LOCAL_OUTAGE"),
	EventType("SYNTHETIC_NODE_OUTAGE"),
	EventType("SYNTHETIC_PRIVATE_LOCATION_OUTAGE"),
	EventType("SYNTHETIC_TEST_LOCATION_SLOWDOWN"),
}

type SeverityLevel string

var SeverityLevels = struct {
	Availability          SeverityLevel
	CustomAlert           SeverityLevel
	Errors                SeverityLevel
	MonitoringUnavailable SeverityLevel
	Performance           SeverityLevel
	ResourceContention    SeverityLevel
}{
	SeverityLevel("AVAILABILITY"),
	SeverityLevel("CUSTOM_ALERT"),
	SeverityLevel("ERRORS"),
	SeverityLevel("MONITORING_UNAVAILABLE"),
	SeverityLevel("PERFORMANCE"),
	SeverityLevel("RESOURCE_CONTENTION"),
}

type TagFilterIncludeMode string

var TagFilterIncludeModes = struct {
	IncludeAll TagFilterIncludeMode
	IncludeAny TagFilterIncludeMode
	None       TagFilterIncludeMode
}{
	TagFilterIncludeMode("INCLUDE_ALL"),
	TagFilterIncludeMode("INCLUDE_ANY"),
	TagFilterIncludeMode("NONE"),
}
