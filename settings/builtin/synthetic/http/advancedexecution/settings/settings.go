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

package advancedexecution

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	ConnectTimeout                  *int   `json:"connectTimeout,omitempty"`                  // Supported values are between 1 000 and 60 000 ms
	DnsQueryTimeout                 *int   `json:"dnsQueryTimeout,omitempty"`                 // Supported values are between 1 000 and 60 000 ms
	MaxCustomScriptSize             *int   `json:"maxCustomScriptSize,omitempty"`             // Supported values are between 10 240 and 102 400 bytes
	MaxHeaderSize                   *int   `json:"maxHeaderSize,omitempty"`                   // Supported values are between 10 240 and 61 440 bytes
	MaxRequestBodySize              *int   `json:"maxRequestBodySize,omitempty"`              // Supported values are between 10 240 and 102 400 bytes
	MaxResponseBodyReadByScriptSize *int   `json:"maxResponseBodyReadByScriptSize,omitempty"` // Supported values are between 10 240 and 102 400 bytes
	MaxResponseBodySize             *int   `json:"maxResponseBodySize,omitempty"`             // Supported values are between 51 200 and 20 971 520 bytes
	MonitorExecutionTimeout         *int   `json:"monitorExecutionTimeout,omitempty"`         // Supported values are between 10 000 and 300 000 ms
	RequestTimeout                  *int   `json:"requestTimeout,omitempty"`                  // Supported values are between 1 000 and 60 000 ms
	Scope                           string `json:"-" scope:"scope"`                           // The scope of this setting (HTTP_CHECK)
	ScriptExecutionTimeout          *int   `json:"scriptExecutionTimeout,omitempty"`          // Supported values are between 1 000 and 10 000 ms
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"connect_timeout": {
			Type:        schema.TypeInt,
			Description: "Supported values are between 1 000 and 60 000 ms",
			Optional:    true, // nullable
		},
		"dns_query_timeout": {
			Type:        schema.TypeInt,
			Description: "Supported values are between 1 000 and 60 000 ms",
			Optional:    true, // nullable
		},
		"max_custom_script_size": {
			Type:        schema.TypeInt,
			Description: "Supported values are between 10 240 and 102 400 bytes",
			Optional:    true, // nullable
		},
		"max_header_size": {
			Type:        schema.TypeInt,
			Description: "Supported values are between 10 240 and 61 440 bytes",
			Optional:    true, // nullable
		},
		"max_request_body_size": {
			Type:        schema.TypeInt,
			Description: "Supported values are between 10 240 and 102 400 bytes",
			Optional:    true, // nullable
		},
		"max_response_body_read_by_script_size": {
			Type:        schema.TypeInt,
			Description: "Supported values are between 10 240 and 102 400 bytes",
			Optional:    true, // nullable
		},
		"max_response_body_size": {
			Type:        schema.TypeInt,
			Description: "Supported values are between 51 200 and 20 971 520 bytes",
			Optional:    true, // nullable
		},
		"monitor_execution_timeout": {
			Type:        schema.TypeInt,
			Description: "Supported values are between 10 000 and 300 000 ms",
			Optional:    true, // nullable
		},
		"request_timeout": {
			Type:        schema.TypeInt,
			Description: "Supported values are between 1 000 and 60 000 ms",
			Optional:    true, // nullable
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (HTTP_CHECK)",
			Required:    true,
		},
		"script_execution_timeout": {
			Type:        schema.TypeInt,
			Description: "Supported values are between 1 000 and 10 000 ms",
			Optional:    true, // nullable
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"connect_timeout":                       me.ConnectTimeout,
		"dns_query_timeout":                     me.DnsQueryTimeout,
		"max_custom_script_size":                me.MaxCustomScriptSize,
		"max_header_size":                       me.MaxHeaderSize,
		"max_request_body_size":                 me.MaxRequestBodySize,
		"max_response_body_read_by_script_size": me.MaxResponseBodyReadByScriptSize,
		"max_response_body_size":                me.MaxResponseBodySize,
		"monitor_execution_timeout":             me.MonitorExecutionTimeout,
		"request_timeout":                       me.RequestTimeout,
		"scope":                                 me.Scope,
		"script_execution_timeout":              me.ScriptExecutionTimeout,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"connect_timeout":                       &me.ConnectTimeout,
		"dns_query_timeout":                     &me.DnsQueryTimeout,
		"max_custom_script_size":                &me.MaxCustomScriptSize,
		"max_header_size":                       &me.MaxHeaderSize,
		"max_request_body_size":                 &me.MaxRequestBodySize,
		"max_response_body_read_by_script_size": &me.MaxResponseBodyReadByScriptSize,
		"max_response_body_size":                &me.MaxResponseBodySize,
		"monitor_execution_timeout":             &me.MonitorExecutionTimeout,
		"request_timeout":                       &me.RequestTimeout,
		"scope":                                 &me.Scope,
		"script_execution_timeout":              &me.ScriptExecutionTimeout,
	})
}
