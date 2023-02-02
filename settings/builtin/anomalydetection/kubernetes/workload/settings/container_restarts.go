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

package workload

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ContainerRestarts struct {
	Configuration *ContainerRestartsConfig `json:"configuration,omitempty"` // Alert if
	Enabled       bool                     `json:"enabled"`                 // Detect container restarts
}

func (me *ContainerRestarts) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configuration": {
			Type:        schema.TypeList,
			Description: "Alert if",
			Optional:    true,

			Elem:     &schema.Resource{Schema: new(ContainerRestartsConfig).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "Detect container restarts",
			Required:    true,
		},
	}
}

func (me *ContainerRestarts) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"configuration": me.Configuration,
		"enabled":       me.Enabled,
	})
}

func (me *ContainerRestarts) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"configuration": &me.Configuration,
		"enabled":       &me.Enabled,
	})
}
