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

type ContainerConstraintsGreaterThan struct {
	FieldA *float64 `json:"fieldA,omitempty"` // Should be greaterThanOrEqual Field B
	FieldB *float64 `json:"fieldB,omitempty"` // Should be greaterThan Field C
	FieldC *float64 `json:"fieldC,omitempty"` // Should be greaterThanOrEqual Field D
	FieldD *float64 `json:"fieldD,omitempty"` // Field D
}

func (me *ContainerConstraintsGreaterThan) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"field_a": {
			Type:        schema.TypeFloat,
			Description: "Should be greaterThanOrEqual Field B",
			Optional:    true, // nullable
		},
		"field_b": {
			Type:        schema.TypeFloat,
			Description: "Should be greaterThan Field C",
			Optional:    true, // nullable
		},
		"field_c": {
			Type:        schema.TypeFloat,
			Description: "Should be greaterThanOrEqual Field D",
			Optional:    true, // nullable
		},
		"field_d": {
			Type:        schema.TypeFloat,
			Description: "Field D",
			Optional:    true, // nullable
		},
	}
}

func (me *ContainerConstraintsGreaterThan) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"field_a": me.FieldA,
		"field_b": me.FieldB,
		"field_c": me.FieldC,
		"field_d": me.FieldD,
	})
}

func (me *ContainerConstraintsGreaterThan) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"field_a": &me.FieldA,
		"field_b": &me.FieldB,
		"field_c": &me.FieldC,
		"field_d": &me.FieldD,
	})
}
