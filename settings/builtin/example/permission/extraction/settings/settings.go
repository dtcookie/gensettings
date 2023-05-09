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

package extraction

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Name string `json:"name"` // Setting value can be deleted only by devops user if its `Name` is `DEVOPS_DELETE`.\n\n Setting value can be added only be devops user if its `Name` is `DEVOPS_CREATE`.\n\n Setting value can be updated only by devops user if its `Name` is `DEVOPS_UPDATE`.
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Description: "Setting value can be deleted only by devops user if its `Name` is `DEVOPS_DELETE`.\n\n Setting value can be added only be devops user if its `Name` is `DEVOPS_CREATE`.\n\n Setting value can be updated only by devops user if its `Name` is `DEVOPS_UPDATE`.",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"name": me.Name,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"name": &me.Name,
	})
}
