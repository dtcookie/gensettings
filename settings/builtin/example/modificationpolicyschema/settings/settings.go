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

package modificationpolicyschema

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	AlwaysModifiable  string `json:"alwaysModifiable"`  // Value of this property can be always modified
	DefaultModifiable string `json:"defaultModifiable"` // Modifiability of this value is defined by ModificationFlags used on the creation of the setting item
	ModifiableFlag    bool   `json:"modifiableFlag"`    // Flag indicating the value of ModificationFlags.isModifiable, that was used on the creation of the whole setting item
	NeverModifiable   string `json:"neverModifiable"`   // Value of this property can be never modified once it was persisted
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"always_modifiable": {
			Type:        schema.TypeString,
			Description: "Value of this property can be always modified",
			Required:    true,
		},
		"default_modifiable": {
			Type:        schema.TypeString,
			Description: "Modifiability of this value is defined by ModificationFlags used on the creation of the setting item",
			Required:    true,
		},
		"modifiable_flag": {
			Type:        schema.TypeBool,
			Description: "Flag indicating the value of ModificationFlags.isModifiable, that was used on the creation of the whole setting item",
			Required:    true,
		},
		"never_modifiable": {
			Type:        schema.TypeString,
			Description: "Value of this property can be never modified once it was persisted",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"always_modifiable":  me.AlwaysModifiable,
		"default_modifiable": me.DefaultModifiable,
		"modifiable_flag":    me.ModifiableFlag,
		"never_modifiable":   me.NeverModifiable,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"always_modifiable":  &me.AlwaysModifiable,
		"default_modifiable": &me.DefaultModifiable,
		"modifiable_flag":    &me.ModifiableFlag,
		"never_modifiable":   &me.NeverModifiable,
	})
}
