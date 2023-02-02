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

package local

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Enabled            bool               `json:"enabled"`            // Enable Extension Execution Controller
	IngestActive       bool               `json:"ingestActive"`       // Enable local PIPE/HTTP metric and Log Ingest API
	PerformanceProfile PerformanceProfile `json:"performanceProfile"` // Possible Values: `DEFAULT`, `HIGH`
	Scope              *string            `json:"-" scope:"scope"`    // The scope of this setting (HOST HOST_GROUP environment)
	StatsdActive       bool               `json:"statsdActive"`       // Enable Dynatrace StatsD
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Type:        schema.TypeBool,
			Description: "Enable Extension Execution Controller",
			Required:    true,
		},
		"ingest_active": {
			Type:        schema.TypeBool,
			Description: "Enable local PIPE/HTTP metric and Log Ingest API",
			Required:    true,
		},
		"performance_profile": {
			Type:        schema.TypeString,
			Description: "Possible Values: `DEFAULT`, `HIGH`",
			Required:    true,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (HOST HOST_GROUP environment)",
			Optional:    true,
			Default:     "environment",
		},
		"statsd_active": {
			Type:        schema.TypeBool,
			Description: "Enable Dynatrace StatsD",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"enabled":             me.Enabled,
		"ingest_active":       me.IngestActive,
		"performance_profile": me.PerformanceProfile,
		"scope":               me.Scope,
		"statsd_active":       me.StatsdActive,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":             &me.Enabled,
		"ingest_active":       &me.IngestActive,
		"performance_profile": &me.PerformanceProfile,
		"scope":               &me.Scope,
		"statsd_active":       &me.StatsdActive,
	})
}
