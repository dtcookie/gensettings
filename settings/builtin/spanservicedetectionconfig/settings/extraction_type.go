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

package spanservicedetectionconfig

import (
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ExtractionTypes []*ExtractionType

func (me *ExtractionTypes) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"transformation": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(ExtractionType).Schema()},
		},
	}
}

func (me ExtractionTypes) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("transformation", me)
}

func (me *ExtractionTypes) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("transformation", me)
}

type ExtractionType struct {
	Marker               string                `json:"marker"`
	Marker2              *string               `json:"marker2,omitempty"`              // Marker_2
	SplitAndSelectMarker *string               `json:"splitAndSelectMarker,omitempty"` // selectAtIndex
	Type                 ExtractionMatcherType `json:"type"`                           // Possible Values: `AFTER`, `BEFORE`, `BETWEEN`, `SPLIT_AND_SELECT`
}

func (me *ExtractionType) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"marker": {
			Type:        schema.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"marker_2": {
			Type:        schema.TypeString,
			Description: "Marker_2",
			Optional:    true, // precondition
		},
		"split_and_select_marker": {
			Type:        schema.TypeString,
			Description: "selectAtIndex",
			Optional:    true, // precondition
		},
		"type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `AFTER`, `BEFORE`, `BETWEEN`, `SPLIT_AND_SELECT`",
			Required:    true,
		},
	}
}

func (me *ExtractionType) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"marker":                  me.Marker,
		"marker_2":                me.Marker2,
		"split_and_select_marker": me.SplitAndSelectMarker,
		"type":                    me.Type,
	})
}

func (me *ExtractionType) HandlePreconditions() error {
	if me.Marker2 == nil && (string(me.Type) == "BETWEEN") {
		return fmt.Errorf("'marker_2' must be specified if 'type' is set to '%v'", me.Type)
	}
	if me.SplitAndSelectMarker == nil && (string(me.Type) == "SPLIT_AND_SELECT") {
		return fmt.Errorf("'split_and_select_marker' must be specified if 'type' is set to '%v'", me.Type)
	}
	return nil
}

func (me *ExtractionType) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"marker":                  &me.Marker,
		"marker_2":                &me.Marker2,
		"split_and_select_marker": &me.SplitAndSelectMarker,
		"type":                    &me.Type,
	})
}
