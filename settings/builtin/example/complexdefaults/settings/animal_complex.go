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

package complexdefaults

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Animal. All animals are good bois
type AnimalComplex struct {
	Animal  AnimalEnum `json:"animal"`  // Possible Values: `Cat`, `Dog`, `Hamster`, `Snake`
	GoodBoi bool       `json:"goodBoi"` // Of course he is
}

func (me *AnimalComplex) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"animal": {
			Type:        schema.TypeString,
			Description: "Possible Values: `Cat`, `Dog`, `Hamster`, `Snake`",
			Required:    true,
		},
		"good_boi": {
			Type:        schema.TypeBool,
			Description: "Of course he is",
			Required:    true,
		},
	}
}

func (me *AnimalComplex) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"animal":   me.Animal,
		"good_boi": me.GoodBoi,
	})
}

func (me *AnimalComplex) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"animal":   &me.Animal,
		"good_boi": &me.GoodBoi,
	})
}
