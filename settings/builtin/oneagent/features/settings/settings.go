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

package features

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Enabled         bool   `json:"enabled"`                   // Enabled
	Forcible        *bool  `json:"forcible,omitempty"`        // Activate this feature also in OneAgents only fulfilling the minimum Opt-In version
	Instrumentation *bool  `json:"instrumentation,omitempty"` // Instrumentation enabled (change needs a process restart)
	Key             string `json:"key"`                       // Feature
	Scope           string `json:"-" scope:"scope"`           // The scope of this setting (PROCESS_GROUP_INSTANCE PROCESS_GROUP environment)
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Type:        schema.TypeBool,
			Description: "Enabled",
			Required:    true,
		},
		"forcible": {
			Type:        schema.TypeBool,
			Description: "Activate this feature also in OneAgents only fulfilling the minimum Opt-In version",
			Optional:    true,
		},
		"instrumentation": {
			Type:        schema.TypeBool,
			Description: "Instrumentation enabled (change needs a process restart)",
			Optional:    true,
		},
		"key": {
			Type:        schema.TypeString,
			Description: "Feature",
			Required:    true,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (PROCESS_GROUP_INSTANCE PROCESS_GROUP environment)",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"enabled":         me.Enabled,
		"forcible":        me.Forcible,
		"instrumentation": me.Instrumentation,
		"key":             me.Key,
		"scope":           me.Scope,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":         &me.Enabled,
		"forcible":        &me.Forcible,
		"instrumentation": &me.Instrumentation,
		"key":             &me.Key,
		"scope":           &me.Scope,
	})
}
