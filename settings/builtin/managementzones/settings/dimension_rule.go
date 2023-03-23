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

package managementzones

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type DimensionRule struct {
	AppliesTo  DimensionType       `json:"appliesTo"` // Possible Values: `ANY`, `LOG`, `METRIC`
	Conditions DimensionConditions `json:"conditions,omitempty"`
}

func (me *DimensionRule) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"applies_to": {
			Type:        schema.TypeString,
			Description: "Possible Values: `ANY`, `LOG`, `METRIC`",
			Required:    true,
		},
		"conditions": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(DimensionConditions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *DimensionRule) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"applies_to": me.AppliesTo,
		"conditions": me.Conditions,
	})
}

func (me *DimensionRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"applies_to": &me.AppliesTo,
		"conditions": &me.Conditions,
	})
}
