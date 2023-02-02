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

type PredefinedEventFilter struct {
	EventType EventType `json:"eventType"` // Filter problems by a Dynatrace event type
	Negate    bool      `json:"negate"`
}

func (me *PredefinedEventFilter) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"event_type": {
			Type:        schema.TypeString,
			Description: "Filter problems by a Dynatrace event type",
			Required:    true,
		},
		"negate": {
			Type:        schema.TypeBool,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *PredefinedEventFilter) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"event_type": me.EventType,
		"negate":     me.Negate,
	})
}

func (me *PredefinedEventFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"event_type": &me.EventType,
		"negate":     &me.Negate,
	})
}
