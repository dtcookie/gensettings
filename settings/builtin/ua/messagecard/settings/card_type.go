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

package messagecard

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Card. A card visualization type
type CardType struct {
	Buttons     ButtonTypes   `json:"buttons,omitempty"`     // Buttons visible on the card
	DisplayName *string       `json:"displayName,omitempty"` // The title of the card
	Icon        *string       `json:"icon,omitempty"`        // The icon of the card. Only names of icons from Barista icon set can be used
	Text        string        `json:"text"`                  // The text of the card. Markdown syntax can be used to enrich the message. To define a link, you can only use seaOtter links (https://dt-url.net/XXXXXXX) or action expression \n\n The format of a single action expression is: actionName|key=param|key2=param2
	Theme       CardThemeType `json:"theme"`                 // Possible Values: `CTA`, `MAIN`, `WARNING`
}

func (me *CardType) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"buttons": {
			Type:        schema.TypeList,
			Description: "Buttons visible on the card",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(ButtonTypes).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"display_name": {
			Type:        schema.TypeString,
			Description: "The title of the card",
			Optional:    true, // nullable
		},
		"icon": {
			Type:        schema.TypeString,
			Description: "The icon of the card. Only names of icons from Barista icon set can be used",
			Optional:    true, // nullable
		},
		"text": {
			Type:        schema.TypeString,
			Description: "The text of the card. Markdown syntax can be used to enrich the message. To define a link, you can only use seaOtter links (https://dt-url.net/XXXXXXX) or action expression \n\n The format of a single action expression is: actionName|key=param|key2=param2",
			Required:    true,
		},
		"theme": {
			Type:        schema.TypeString,
			Description: "Possible Values: `CTA`, `MAIN`, `WARNING`",
			Required:    true,
		},
	}
}

func (me *CardType) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"buttons":      me.Buttons,
		"display_name": me.DisplayName,
		"icon":         me.Icon,
		"text":         me.Text,
		"theme":        me.Theme,
	})
}

func (me *CardType) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"buttons":      &me.Buttons,
		"display_name": &me.DisplayName,
		"icon":         &me.Icon,
		"text":         &me.Text,
		"theme":        &me.Theme,
	})
}
