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

package php

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Enabled        bool    `json:"enabled"`             // Monitor PHP
	EnabledFastCGI bool    `json:"enabledFastCGI"`      // Requires PHP monitoring enabled and from Dynatrace OneAgent version 1.191 it's ignored and permanently enabled
	ServiceID      *string `json:"-" scope:"serviceId"` // The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Type:        schema.TypeBool,
			Description: "Monitor PHP",
			Required:    true,
		},
		"enabled_fast_cgi": {
			Type:        schema.TypeBool,
			Description: "Requires PHP monitoring enabled and from Dynatrace OneAgent version 1.191 it's ignored and permanently enabled",
			Required:    true,
		},
		"service_id": {
			Type:        schema.TypeString,
			Description: "The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.",
			Optional:    true,
			Default:     "environment",
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"enabled":          me.Enabled,
		"enabled_fast_cgi": me.EnabledFastCGI,
		"service_id":       me.ServiceID,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":          &me.Enabled,
		"enabled_fast_cgi": &me.EnabledFastCGI,
		"service_id":       &me.ServiceID,
	})
}
