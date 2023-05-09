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

// Message. A message visualization type
type MessageType struct {
	Text  string           `json:"text"`  // The text that will be displayed as a message. Markdown syntax can be used to enrich the message. To define a link, you can only use seaOtter links (https://dt-url.net/XXXXXXX) or action expression \n\n The format of a single action expression is: actionName|key=param|key2=param2
	Theme MessageColorType `json:"theme"` // Possible Values: `ERROR`, `INFO`, `WARNING`
}

func (me *MessageType) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"text": {
			Type:        schema.TypeString,
			Description: "The text that will be displayed as a message. Markdown syntax can be used to enrich the message. To define a link, you can only use seaOtter links (https://dt-url.net/XXXXXXX) or action expression \n\n The format of a single action expression is: actionName|key=param|key2=param2",
			Required:    true,
		},
		"theme": {
			Type:        schema.TypeString,
			Description: "Possible Values: `ERROR`, `INFO`, `WARNING`",
			Required:    true,
		},
	}
}

func (me *MessageType) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"text":  me.Text,
		"theme": me.Theme,
	})
}

func (me *MessageType) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"text":  &me.Text,
		"theme": &me.Theme,
	})
}
