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

type Settings struct {
	ContainerConstraintPropertyCountRange *ContainerConstraintPropertyCountRange `json:"containerConstraintPropertyCountRange"` // Container Constraint propertyCountRange
	ContainerConstraintsGreaterThan       *ContainerConstraintsGreaterThan       `json:"containerConstraintsGreaterThan"`       // Container Constraints greaterThanOrEqual & greaterThan
	ContainerConstraintsLessThan          *ContainerConstraintsLessThan          `json:"containerConstraintsLessThan"`          // Container Constraints lessThanOrEqual & lessThan
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"container_constraint_property_count_range": {
			Type:        schema.TypeList,
			Description: "Container Constraint propertyCountRange",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(ContainerConstraintPropertyCountRange).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"container_constraints_greater_than": {
			Type:        schema.TypeList,
			Description: "Container Constraints greaterThanOrEqual & greaterThan",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(ContainerConstraintsGreaterThan).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"container_constraints_less_than": {
			Type:        schema.TypeList,
			Description: "Container Constraints lessThanOrEqual & lessThan",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(ContainerConstraintsLessThan).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"container_constraint_property_count_range": me.ContainerConstraintPropertyCountRange,
		"container_constraints_greater_than":        me.ContainerConstraintsGreaterThan,
		"container_constraints_less_than":           me.ContainerConstraintsLessThan,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"container_constraint_property_count_range": &me.ContainerConstraintPropertyCountRange,
		"container_constraints_greater_than":        &me.ContainerConstraintsGreaterThan,
		"container_constraints_less_than":           &me.ContainerConstraintsLessThan,
	})
}
