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

package constraints

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ContainerConstraintsLessThan struct {
	DateA *string `json:"dateA,omitempty"` // Should be lessThanOrEqual Date B
	DateB *string `json:"dateB,omitempty"` // Should be lessThan Date C
	DateC *string `json:"dateC,omitempty"` // Should be lessThanOrEqual Date D
	DateD *string `json:"dateD,omitempty"` // Date D
}

func (me *ContainerConstraintsLessThan) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"date_a": {
			Type:        schema.TypeString,
			Description: "Should be lessThanOrEqual Date B",
			Optional:    true, // nullable
		},
		"date_b": {
			Type:        schema.TypeString,
			Description: "Should be lessThan Date C",
			Optional:    true, // nullable
		},
		"date_c": {
			Type:        schema.TypeString,
			Description: "Should be lessThanOrEqual Date D",
			Optional:    true, // nullable
		},
		"date_d": {
			Type:        schema.TypeString,
			Description: "Date D",
			Optional:    true, // nullable
		},
	}
}

func (me *ContainerConstraintsLessThan) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"date_a": me.DateA,
		"date_b": me.DateB,
		"date_c": me.DateC,
		"date_d": me.DateD,
	})
}

func (me *ContainerConstraintsLessThan) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"date_a": &me.DateA,
		"date_b": &me.DateB,
		"date_c": &me.DateC,
		"date_d": &me.DateD,
	})
}
