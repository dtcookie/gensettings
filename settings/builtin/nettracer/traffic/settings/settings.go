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

package traffic

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	NetTracer bool    `json:"netTracer"`       // When disabled, OneAgent won't use NetTracer to collect network data from containers. Disabled by default. Applies only to Linux hosts. Requires OneAgent 1.231+.
	Scope     *string `json:"-" scope:"scope"` // The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"net_tracer": {
			Type:        schema.TypeBool,
			Description: "When disabled, OneAgent won't use NetTracer to collect network data from containers. Disabled by default. Applies only to Linux hosts. Requires OneAgent 1.231+.",
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
		"net_tracer": me.NetTracer,
		"scope":      me.Scope,
	})
}

func (me *Settings) HandlePreconditions() {
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"net_tracer": &me.NetTracer,
		"scope":      &me.Scope,
	})
}
