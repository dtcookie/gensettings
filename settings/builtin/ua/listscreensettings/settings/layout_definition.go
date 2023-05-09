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

package listscreensettings

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type LayoutDefinition struct {
	AutoGenerate bool           `json:"autoGenerate"`    // If true the screen content is auto generated. Disable to configure your own layout.
	Cards        CardReferences `json:"cards,omitempty"` // A list of cards that should be shown on the screen. Reference cards that should be shown on the screen.  \nCards will be displayed in order defined on this list.
}

func (me *LayoutDefinition) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"auto_generate": {
			Type:        schema.TypeBool,
			Description: "If true the screen content is auto generated. Disable to configure your own layout.",
			Required:    true,
		},
		"cards": {
			Type:        schema.TypeList,
			Description: "A list of cards that should be shown on the screen. Reference cards that should be shown on the screen.  \nCards will be displayed in order defined on this list.",
			Optional:    true, // precondition & minobjects == 0
			Elem:        &schema.Resource{Schema: new(CardReferences).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *LayoutDefinition) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"auto_generate": me.AutoGenerate,
		"cards":         me.Cards,
	})
}

func (me *LayoutDefinition) HandlePreconditions() error {
	// ---- Cards CardReferences -> {"expectedValue":false,"property":"autoGenerate","type":"EQUALS"}
	return nil
}

func (me *LayoutDefinition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"auto_generate": &me.AutoGenerate,
		"cards":         &me.Cards,
	})
}
