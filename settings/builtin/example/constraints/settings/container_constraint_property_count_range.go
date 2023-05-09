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

type ContainerConstraintPropertyCountRange struct {
	FieldA *float64 `json:"fieldA,omitempty"` // At least one of this fields has to be set
	FieldB *float64 `json:"fieldB,omitempty"` // At least one of this fields has to be set
}

func (me *ContainerConstraintPropertyCountRange) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"field_a": {
			Type:        schema.TypeFloat,
			Description: "At least one of this fields has to be set",
			Optional:    true, // nullable
		},
		"field_b": {
			Type:        schema.TypeFloat,
			Description: "At least one of this fields has to be set",
			Optional:    true, // nullable
		},
	}
}

func (me *ContainerConstraintPropertyCountRange) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"field_a": me.FieldA,
		"field_b": me.FieldB,
	})
}

func (me *ContainerConstraintPropertyCountRange) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"field_a": &me.FieldA,
		"field_b": &me.FieldB,
	})
}
