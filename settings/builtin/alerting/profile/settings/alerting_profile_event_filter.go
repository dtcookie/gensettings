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
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AlertingProfileEventFilters []*AlertingProfileEventFilter

func (me *AlertingProfileEventFilters) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"event_filter": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(AlertingProfileEventFilter).Schema()},
		},
	}
}

func (me AlertingProfileEventFilters) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("event_filter", me)
}

func (me *AlertingProfileEventFilters) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("event_filter", me)
}

type AlertingProfileEventFilter struct {
	CustomFilter     *CustomEventFilter             `json:"customFilter,omitempty"`
	PredefinedFilter *PredefinedEventFilter         `json:"predefinedFilter,omitempty"`
	Type             AlertingProfileEventFilterType `json:"type"` // Possible Values: `CUSTOM`, `PREDEFINED`
}

func (me *AlertingProfileEventFilter) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"custom_filter": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(CustomEventFilter).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"predefined_filter": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(PredefinedEventFilter).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `CUSTOM`, `PREDEFINED`",
			Required:    true,
		},
	}
}

func (me *AlertingProfileEventFilter) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"custom_filter":     me.CustomFilter,
		"predefined_filter": me.PredefinedFilter,
		"type":              me.Type,
	})
}

func (me *AlertingProfileEventFilter) HandlePreconditions() error {
	if me.CustomFilter == nil && string(me.Type) == "CUSTOM" {
		return fmt.Errorf("'custom_filter' must be specified if 'type' is set to '%v'", me.Type)
	}
	if me.CustomFilter != nil && string(me.Type) != "CUSTOM" {
		return fmt.Errorf("'custom_filter' must not be specified if 'type' is set to '%v'", me.Type)
	}
	if me.PredefinedFilter == nil && string(me.Type) == "PREDEFINED" {
		return fmt.Errorf("'predefined_filter' must be specified if 'type' is set to '%v'", me.Type)
	}
	if me.PredefinedFilter != nil && string(me.Type) != "PREDEFINED" {
		return fmt.Errorf("'predefined_filter' must not be specified if 'type' is set to '%v'", me.Type)
	}
	return nil
}

func (me *AlertingProfileEventFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_filter":     &me.CustomFilter,
		"predefined_filter": &me.PredefinedFilter,
		"type":              &me.Type,
	})
}
