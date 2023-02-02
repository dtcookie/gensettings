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

package kubernetes

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	ActiveGateGroup                 *string           `json:"activeGateGroup,omitempty"`       // ActiveGate Group
	AuthToken                       string            `json:"authToken"`                       // Create a bearer token for [Kubernetes](https://dt-url.net/og43szq \"Kubernetes\") or [OpenShift](https://dt-url.net/7l43xtp \"OpenShift\").
	CertificateCheckEnabled         bool              `json:"certificateCheckEnabled"`         // Require valid certificates for communication with API server (recommended)
	CloudApplicationPipelineEnabled bool              `json:"cloudApplicationPipelineEnabled"` // Monitor Kubernetes namespaces, services, workloads, and pods
	ClusterID                       string            `json:"clusterId"`                       // Unique ID of the cluster, the containerized ActiveGate is deployed to. Defaults to the UUID of the kube-system namespace. The cluster ID of containerized ActiveGates is shown on the Deployment status screen.
	ClusterIdEnabled                bool              `json:"clusterIdEnabled"`                // This is required for monitoring persistent volume claims. For more information on local Kubernetes API monitoring, see the [documentation](https://dt-url.net/6q62uep).
	Enabled                         bool              `json:"enabled"`                         // Enabled
	EndpointUrl                     string            `json:"endpointUrl"`                     // Get the API URL for [Kubernetes](https://dt-url.net/kz23snj \"Kubernetes\") or [OpenShift](https://dt-url.net/d623xgw \"OpenShift\").
	EventPatterns                   EventComplexTypes `json:"eventPatterns"`                   // Define Kubernetes event filters to ingest events into your environment. For more details, see the [documentation](https://dt-url.net/2201p0u).
	EventProcessingActive           bool              `json:"eventProcessingActive"`           // All events are monitored by default unless event filters are specified.\n\nKubernetes events are subject to Davis data units (DDU) licensing.\nSee [DDUs for events](https://dt-url.net/5n03vcu) for details.
	FilterEvents                    bool              `json:"filterEvents"`                    // Include only events specified by Events Field Selectors
	HostnameVerificationEnabled     bool              `json:"hostnameVerificationEnabled"`     // Verify hostname in certificate against Kubernetes API URL
	IncludeAllFdiEvents             bool              `json:"includeAllFdiEvents"`             // For a list of included events, see the [documentation](https://dt-url.net/l61d02no).
	Label                           string            `json:"label"`                           // Renaming the cluster breaks configurations that are based on its name (e.g., management zones, and alerting).
	OpenMetricsPipelineEnabled      bool              `json:"openMetricsPipelineEnabled"`      // For annotation guidance, see the [documentation](https://dt-url.net/g42i0ppw).
	PvcMonitoringEnabled            bool              `json:"pvcMonitoringEnabled"`            // To enable dashboards and alerts, add the [Kubernetes persistent volume claims](ui/hub/ext/com.dynatrace.extension.kubernetes-pvc) extension to your environment.
	Scope                           string            `json:"-" scope:"scope"`                 // The scope of this setting (KUBERNETES_CLUSTER)
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active_gate_group": {
			Type:        schema.TypeString,
			Description: "ActiveGate Group",
			Optional:    true,
		},
		"auth_token": {
			Type:        schema.TypeString,
			Description: "Create a bearer token for [Kubernetes](https://dt-url.net/og43szq \"Kubernetes\") or [OpenShift](https://dt-url.net/7l43xtp \"OpenShift\").",
			Required:    true,
		},
		"certificate_check_enabled": {
			Type:        schema.TypeBool,
			Description: "Require valid certificates for communication with API server (recommended)",
			Required:    true,
		},
		"cloud_application_pipeline_enabled": {
			Type:        schema.TypeBool,
			Description: "Monitor Kubernetes namespaces, services, workloads, and pods",
			Required:    true,
		},
		"cluster_id": {
			Type:        schema.TypeString,
			Description: "Unique ID of the cluster, the containerized ActiveGate is deployed to. Defaults to the UUID of the kube-system namespace. The cluster ID of containerized ActiveGates is shown on the Deployment status screen.",
			Required:    true,
		},
		"cluster_id_enabled": {
			Type:        schema.TypeBool,
			Description: "This is required for monitoring persistent volume claims. For more information on local Kubernetes API monitoring, see the [documentation](https://dt-url.net/6q62uep).",
			Required:    true,
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "Enabled",
			Required:    true,
		},
		"endpoint_url": {
			Type:        schema.TypeString,
			Description: "Get the API URL for [Kubernetes](https://dt-url.net/kz23snj \"Kubernetes\") or [OpenShift](https://dt-url.net/d623xgw \"OpenShift\").",
			Required:    true,
		},
		"event_patterns": {
			Type:        schema.TypeList,
			Description: "Define Kubernetes event filters to ingest events into your environment. For more details, see the [documentation](https://dt-url.net/2201p0u).",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(EventComplexTypes).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"event_processing_active": {
			Type:        schema.TypeBool,
			Description: "All events are monitored by default unless event filters are specified.\n\nKubernetes events are subject to Davis data units (DDU) licensing.\nSee [DDUs for events](https://dt-url.net/5n03vcu) for details.",
			Required:    true,
		},
		"filter_events": {
			Type:        schema.TypeBool,
			Description: "Include only events specified by Events Field Selectors",
			Required:    true,
		},
		"hostname_verification_enabled": {
			Type:        schema.TypeBool,
			Description: "Verify hostname in certificate against Kubernetes API URL",
			Required:    true,
		},
		"include_all_fdi_events": {
			Type:        schema.TypeBool,
			Description: "For a list of included events, see the [documentation](https://dt-url.net/l61d02no).",
			Required:    true,
		},
		"label": {
			Type:        schema.TypeString,
			Description: "Renaming the cluster breaks configurations that are based on its name (e.g., management zones, and alerting).",
			Required:    true,
		},
		"open_metrics_pipeline_enabled": {
			Type:        schema.TypeBool,
			Description: "For annotation guidance, see the [documentation](https://dt-url.net/g42i0ppw).",
			Required:    true,
		},
		"pvc_monitoring_enabled": {
			Type:        schema.TypeBool,
			Description: "To enable dashboards and alerts, add the [Kubernetes persistent volume claims](ui/hub/ext/com.dynatrace.extension.kubernetes-pvc) extension to your environment.",
			Required:    true,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (KUBERNETES_CLUSTER)",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"active_gate_group":                  me.ActiveGateGroup,
		"auth_token":                         me.AuthToken,
		"certificate_check_enabled":          me.CertificateCheckEnabled,
		"cloud_application_pipeline_enabled": me.CloudApplicationPipelineEnabled,
		"cluster_id":                         me.ClusterID,
		"cluster_id_enabled":                 me.ClusterIdEnabled,
		"enabled":                            me.Enabled,
		"endpoint_url":                       me.EndpointUrl,
		"event_patterns":                     me.EventPatterns,
		"event_processing_active":            me.EventProcessingActive,
		"filter_events":                      me.FilterEvents,
		"hostname_verification_enabled":      me.HostnameVerificationEnabled,
		"include_all_fdi_events":             me.IncludeAllFdiEvents,
		"label":                              me.Label,
		"open_metrics_pipeline_enabled":      me.OpenMetricsPipelineEnabled,
		"pvc_monitoring_enabled":             me.PvcMonitoringEnabled,
		"scope":                              me.Scope,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"active_gate_group":                  &me.ActiveGateGroup,
		"auth_token":                         &me.AuthToken,
		"certificate_check_enabled":          &me.CertificateCheckEnabled,
		"cloud_application_pipeline_enabled": &me.CloudApplicationPipelineEnabled,
		"cluster_id":                         &me.ClusterID,
		"cluster_id_enabled":                 &me.ClusterIdEnabled,
		"enabled":                            &me.Enabled,
		"endpoint_url":                       &me.EndpointUrl,
		"event_patterns":                     &me.EventPatterns,
		"event_processing_active":            &me.EventProcessingActive,
		"filter_events":                      &me.FilterEvents,
		"hostname_verification_enabled":      &me.HostnameVerificationEnabled,
		"include_all_fdi_events":             &me.IncludeAllFdiEvents,
		"label":                              &me.Label,
		"open_metrics_pipeline_enabled":      &me.OpenMetricsPipelineEnabled,
		"pvc_monitoring_enabled":             &me.PvcMonitoringEnabled,
		"scope":                              &me.Scope,
	})
}
