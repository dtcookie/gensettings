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

package maturity

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Outer1 *string    `json:"outer1,omitempty"` // Property 1
	Outer3 InnerTypes `json:"outer3,omitempty"` // Some List
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"outer_1": {
			Type:        schema.TypeString,
			Description: "Property 1",
			Optional:    true, // nullable
		},
		"outer_3": {
			Type:        schema.TypeList,
			Description: "Some List",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(InnerTypes).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"outer_1": me.Outer1,
		"outer_3": me.Outer3,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"outer_1": &me.Outer1,
		"outer_3": &me.Outer3,
	})
}
