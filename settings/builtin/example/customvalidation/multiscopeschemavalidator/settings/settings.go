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

package multiscopeschemavalidator

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	HostID                              string `json:"-" scope:"hostId"`                    // The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	MultiScopeUniqueCaseInsensitiveName string `json:"multiScopeUniqueCaseInsensitiveName"` // A unique (case-insensitive) name in the whole environment.
	UniqueName                          string `json:"uniqueName"`                          // A unique name in the whole environment.
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"host_id": {
			Type:        schema.TypeString,
			Description: "The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.",
			Required:    true,
		},
		"multi_scope_unique_case_insensitive_name": {
			Type:        schema.TypeString,
			Description: "A unique (case-insensitive) name in the whole environment.",
			Required:    true,
		},
		"unique_name": {
			Type:        schema.TypeString,
			Description: "A unique name in the whole environment.",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"host_id": me.HostID,
		"multi_scope_unique_case_insensitive_name": me.MultiScopeUniqueCaseInsensitiveName,
		"unique_name": me.UniqueName,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"host_id": &me.HostID,
		"multi_scope_unique_case_insensitive_name": &me.MultiScopeUniqueCaseInsensitiveName,
		"unique_name": &me.UniqueName,
	})
}
