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

package profile

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type TextFilter struct {
	CaseSensitive bool               `json:"caseSensitive"` // Case sensitive
	Enabled       bool               `json:"enabled"`
	Negate        bool               `json:"negate"`
	Operator      ComparisonOperator `json:"operator"` // Operator of the comparison
	Value         string             `json:"value"`
}

func (me *TextFilter) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"case_sensitive": {
			Type:        schema.TypeBool,
			Description: "Case sensitive",
			Required:    true,
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "no documentation available",
			Required:    true,
		},
		"negate": {
			Type:        schema.TypeBool,
			Description: "no documentation available",
			Required:    true,
		},
		"operator": {
			Type:        schema.TypeString,
			Description: "Operator of the comparison",
			Required:    true,
		},
		"value": {
			Type:        schema.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *TextFilter) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"case_sensitive": me.CaseSensitive,
		"enabled":        me.Enabled,
		"negate":         me.Negate,
		"operator":       me.Operator,
		"value":          me.Value,
	})
}

func (me *TextFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"case_sensitive": &me.CaseSensitive,
		"enabled":        &me.Enabled,
		"negate":         &me.Negate,
		"operator":       &me.Operator,
		"value":          &me.Value,
	})
}
