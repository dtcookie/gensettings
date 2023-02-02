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

package metricevents

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type EntityFilterConditions []*EntityFilterCondition

func (me *EntityFilterConditions) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"condition": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(EntityFilterCondition).Schema()},
		},
	}
}

func (me EntityFilterConditions) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("condition", me)
}

func (me *EntityFilterConditions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("condition", me)
}

type EntityFilterCondition struct {
	Operator EntityFilterOperator `json:"operator"` // Possible Values: `CONTAINS_CASE_INSENSITIVE`, `CONTAINS_CASE_SENSITIVE`, `EQUALS`
	Type     EntityFilterType     `json:"type"`     // Possible Values: `TAG`, `CUSTOM_DEVICE_GROUP_NAME`, `ENTITY_ID`, `PROCESS_GROUP_ID`, `MANAGEMENT_ZONE`, `NAME`, `PROCESS_GROUP_NAME`, `HOST_NAME`, `HOST_GROUP_NAME`
	Value    string               `json:"value"`
}

func (me *EntityFilterCondition) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"operator": {
			Type:        schema.TypeString,
			Description: "Possible Values: `CONTAINS_CASE_INSENSITIVE`, `CONTAINS_CASE_SENSITIVE`, `EQUALS`",
			Required:    true,
		},
		"type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `TAG`, `CUSTOM_DEVICE_GROUP_NAME`, `ENTITY_ID`, `PROCESS_GROUP_ID`, `MANAGEMENT_ZONE`, `NAME`, `PROCESS_GROUP_NAME`, `HOST_NAME`, `HOST_GROUP_NAME`",
			Required:    true,
		},
		"value": {
			Type:        schema.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *EntityFilterCondition) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"operator": me.Operator,
		"type":     me.Type,
		"value":    me.Value,
	})
}

func (me *EntityFilterCondition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"operator": &me.Operator,
		"type":     &me.Type,
		"value":    &me.Value,
	})
}
