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

package custommetrics

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Filters []*Filter

func (me *Filters) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"filter": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(Filter).Schema()},
		},
	}
}

func (me Filters) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("filter", me)
}

func (me *Filters) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("filter", me)
}

type Filter struct {
	FieldName string    `json:"fieldName"` // Field name
	Operator  Operator  `json:"operator"`  // Possible Values: `IS_NULL`, `LESS_THAN_OR_EQUAL_TO`, `GREATER_THAN_OR_EQUAL_TO`, `IS_NOT_NULL`, `LESS_THAN`, `NOT_EQUAL`, `IN`, `LIKE`, `NOT_LIKE`, `EQUALS`, `STARTS_WITH`, `GREATER_THAN`
	Value     *string   `json:"value,omitempty"`
	ValueIn   *[]string `json:"valueIn,omitempty"` // Values
}

func (me *Filter) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"field_name": {
			Type:        schema.TypeString,
			Description: "Field name",
			Required:    true,
		},
		"operator": {
			Type:        schema.TypeString,
			Description: "Possible Values: `IS_NULL`, `LESS_THAN_OR_EQUAL_TO`, `GREATER_THAN_OR_EQUAL_TO`, `IS_NOT_NULL`, `LESS_THAN`, `NOT_EQUAL`, `IN`, `LIKE`, `NOT_LIKE`, `EQUALS`, `STARTS_WITH`, `GREATER_THAN`",
			Required:    true,
		},
		"value": {
			Type:        schema.TypeString,
			Description: "no documentation available",
			Optional:    true,
		},
		"value_in": {
			Type:        schema.TypeList,
			Description: "Values",
			Optional:    true,

			Elem: &schema.Schema{Type: schema.TypeString},
		},
	}
}

func (me *Filter) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"field_name": me.FieldName,
		"operator":   me.Operator,
		"value":      me.Value,
		"value_in":   me.ValueIn,
	})
}

func (me *Filter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"field_name": &me.FieldName,
		"operator":   &me.Operator,
		"value":      &me.Value,
		"value_in":   &me.ValueIn,
	})
}
