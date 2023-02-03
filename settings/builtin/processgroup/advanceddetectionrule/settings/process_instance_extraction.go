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

package advanceddetectionrule

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ProcessInstanceExtraction struct {
	Delimiter *Delimiter `json:"delimiter,omitempty"` // Optionally delimit this property between *From* and *To*.
	Property  *string    `json:"property,omitempty"`
}

func (me *ProcessInstanceExtraction) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"delimiter": {
			Type:        schema.TypeList,
			Description: "Optionally delimit this property between *From* and *To*.",
			Optional:    true,

			Elem:     &schema.Resource{Schema: new(Delimiter).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"property": {
			Type:        schema.TypeString,
			Description: "no documentation available",
			Optional:    true,
		},
	}
}

func (me *ProcessInstanceExtraction) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"delimiter": me.Delimiter,
		"property":  me.Property,
	})
}

func (me *ProcessInstanceExtraction) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"delimiter": &me.Delimiter,
		"property":  &me.Property,
	})
}
