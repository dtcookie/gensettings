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

package customunit

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	AdvancedSettings      bool             `json:"advancedSettings"`      // If not enabled a simple custom unit with the information above is created. If the advanced custom unit option is enabled, dependencies to already existing units can be defined.
	FirstUnit             string           `json:"firstUnit"`             // The new unit depends on existing unit.
	NewUnitDescription    string           `json:"newUnitDescription"`    // Unit description should provide additional information about the new unit
	NewUnitName           string           `json:"newUnitName"`           // Unit name has to be unique and is used as identifier.
	NewUnitPluralName     string           `json:"newUnitPluralName"`     // Unit plural name represent the plural form of the unit name.
	NewUnitSymbol         string           `json:"newUnitSymbol"`         // Unit symbol has to be unique.
	PowerFactor           int              `json:"powerFactor"`           // (existing unit) ^ (exponent) = (new unit)
	ScalingFactor         float64          `json:"scalingFactor"`         // (existing unit) * (Scaling factor) = new unit
	SecondUnit            string           `json:"secondUnit"`            // New unit depends on existing unit (based on selection at unit composition).
	UnitCombinationSelect UnitCombinations `json:"unitCombinationSelect"` // Possible Values: `PRODUCT`, `POWER`, `SCALAR`, `QUOTIENT`
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"advanced_settings": {
			Type:        schema.TypeBool,
			Description: "If not enabled a simple custom unit with the information above is created. If the advanced custom unit option is enabled, dependencies to already existing units can be defined.",
			Required:    true,
		},
		"first_unit": {
			Type:        schema.TypeString,
			Description: "The new unit depends on existing unit.",
			Required:    true,
		},
		"new_unit_description": {
			Type:        schema.TypeString,
			Description: "Unit description should provide additional information about the new unit",
			Required:    true,
		},
		"new_unit_name": {
			Type:        schema.TypeString,
			Description: "Unit name has to be unique and is used as identifier.",
			Required:    true,
		},
		"new_unit_plural_name": {
			Type:        schema.TypeString,
			Description: "Unit plural name represent the plural form of the unit name.",
			Required:    true,
		},
		"new_unit_symbol": {
			Type:        schema.TypeString,
			Description: "Unit symbol has to be unique.",
			Required:    true,
		},
		"power_factor": {
			Type:        schema.TypeInt,
			Description: "(existing unit) ^ (exponent) = (new unit)",
			Required:    true,
		},
		"scaling_factor": {
			Type:        schema.TypeFloat,
			Description: "(existing unit) * (Scaling factor) = new unit",
			Required:    true,
		},
		"second_unit": {
			Type:        schema.TypeString,
			Description: "New unit depends on existing unit (based on selection at unit composition).",
			Required:    true,
		},
		"unit_combination_select": {
			Type:        schema.TypeString,
			Description: "Possible Values: `PRODUCT`, `POWER`, `SCALAR`, `QUOTIENT`",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"advanced_settings":       me.AdvancedSettings,
		"first_unit":              me.FirstUnit,
		"new_unit_description":    me.NewUnitDescription,
		"new_unit_name":           me.NewUnitName,
		"new_unit_plural_name":    me.NewUnitPluralName,
		"new_unit_symbol":         me.NewUnitSymbol,
		"power_factor":            me.PowerFactor,
		"scaling_factor":          me.ScalingFactor,
		"second_unit":             me.SecondUnit,
		"unit_combination_select": me.UnitCombinationSelect,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"advanced_settings":       &me.AdvancedSettings,
		"first_unit":              &me.FirstUnit,
		"new_unit_description":    &me.NewUnitDescription,
		"new_unit_name":           &me.NewUnitName,
		"new_unit_plural_name":    &me.NewUnitPluralName,
		"new_unit_symbol":         &me.NewUnitSymbol,
		"power_factor":            &me.PowerFactor,
		"scaling_factor":          &me.ScalingFactor,
		"second_unit":             &me.SecondUnit,
		"unit_combination_select": &me.UnitCombinationSelect,
	})
}
