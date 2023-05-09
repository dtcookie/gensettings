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

package datasource

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type InnerType struct {
	FullContextDatasource string `json:"fullContextDatasource"` // Data source that expects full setting payload on validation\n\nWill return possible values based on type, only allows datasource values (with validation). Does not use api to filter values. Some of the values are visible only with devops access.
}

func (me *InnerType) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"full_context_datasource": {
			Type:        schema.TypeString,
			Description: "Data source that expects full setting payload on validation\n\nWill return possible values based on type, only allows datasource values (with validation). Does not use api to filter values. Some of the values are visible only with devops access.",
			Required:    true,
		},
	}
}

func (me *InnerType) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"full_context_datasource": me.FullContextDatasource,
	})
}

func (me *InnerType) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"full_context_datasource": &me.FullContextDatasource,
	})
}
