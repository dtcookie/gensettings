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

package analyzer

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Definition string `json:"definition"` // The analyzer in its plain serialized form. Only edit if you know what you are doing!
	Name       string `json:"name"`       // Analyzer name
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"definition": {
			Type:        schema.TypeString,
			Description: "The analyzer in its plain serialized form. Only edit if you know what you are doing!",
			Required:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Description: "Analyzer name",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"definition": me.Definition,
		"name":       me.Name,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"definition": &me.Definition,
		"name":       &me.Name,
	})
}
