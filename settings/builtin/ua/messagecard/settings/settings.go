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
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Card    *CardType         `json:"card,omitempty"`    // Card
	Key     string            `json:"key"`               // Unique key, which is used to map to this message card in the screen layout config
	Message *MessageType      `json:"message,omitempty"` // Message
	Scope   string            `json:"-" scope:"scope"`   // The scope of this setting (ua-screen)
	Type    VisualizationType `json:"type"`              // Possible Values: `CARD`, `MESSAGE`
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"card": {
			Type:        schema.TypeList,
			Description: "Card",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(CardType).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"key": {
			Type:        schema.TypeString,
			Description: "Unique key, which is used to map to this message card in the screen layout config",
			Required:    true,
		},
		"message": {
			Type:        schema.TypeList,
			Description: "Message",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(MessageType).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (ua-screen)",
			Required:    true,
		},
		"type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `CARD`, `MESSAGE`",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"card":    me.Card,
		"key":     me.Key,
		"message": me.Message,
		"scope":   me.Scope,
		"type":    me.Type,
	})
}

func (me *Settings) HandlePreconditions() error {
	if me.Card == nil && (string(me.Type) == "CARD") {
		return fmt.Errorf("'card' must be specified if 'type' is set to '%v'", me.Type)
	}
	if me.Card != nil && (string(me.Type) != "CARD") {
		return fmt.Errorf("'card' must not be specified if 'type' is set to '%v'", me.Type)
	}
	if me.Message == nil && (string(me.Type) == "MESSAGE") {
		return fmt.Errorf("'message' must be specified if 'type' is set to '%v'", me.Type)
	}
	if me.Message != nil && (string(me.Type) != "MESSAGE") {
		return fmt.Errorf("'message' must not be specified if 'type' is set to '%v'", me.Type)
	}
	return nil
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"card":    &me.Card,
		"key":     &me.Key,
		"message": &me.Message,
		"scope":   &me.Scope,
		"type":    &me.Type,
	})
}
