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

package entitieslist

import (
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/opt"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type EntityFilterDefinitions []*EntityFilterDefinition

func (me *EntityFilterDefinitions) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"filter": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(EntityFilterDefinition).Schema()},
		},
	}
}

func (me EntityFilterDefinitions) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("filter", me)
}

func (me *EntityFilterDefinitions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("filter", me)
}

type EntityFilterDefinition struct {
	DefaultSearch  *bool                 `json:"defaultSearch,omitempty"`  // When true, as user is typing it will suggest this key as default for searching. There can be only 1 default search filter in filtering.
	DisplayName    string                `json:"displayName"`              // Display name
	Distinct       bool                  `json:"distinct"`                 // When true, there can be only a single instance of this filter.
	EntityTypes    []string              `json:"entityTypes,omitempty"`    // List of entity types that are filtered by this filter. All of them must have defined relationship to entity.
	FreeText       bool                  `json:"freeText"`                 // When true, user can type any text value. When false, user can only select one of the suggestions.
	Hidden         *bool                 `json:"hidden,omitempty"`         // When true, this filter will be hidden.
	Modifier       *ModifierType         `json:"modifier,omitempty"`       // Possible Values: `Contains`, `Equals`, `StartsWith`
	SplittingChart *SplittingChartConfig `json:"splittingChart,omitempty"` // Configure splitting chart for this filter.
	Type           string                `json:"type"`
}

func (me *EntityFilterDefinition) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"default_search": {
			Type:        schema.TypeBool,
			Description: "When true, as user is typing it will suggest this key as default for searching. There can be only 1 default search filter in filtering.",
			Optional:    true, // precondition
		},
		"display_name": {
			Type:        schema.TypeString,
			Description: "Display name",
			Required:    true,
		},
		"distinct": {
			Type:        schema.TypeBool,
			Description: "When true, there can be only a single instance of this filter.",
			Required:    true,
		},
		"entity_types": {
			Type:        schema.TypeSet,
			Description: "List of entity types that are filtered by this filter. All of them must have defined relationship to entity.",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"free_text": {
			Type:        schema.TypeBool,
			Description: "When true, user can type any text value. When false, user can only select one of the suggestions.",
			Required:    true,
		},
		"hidden": {
			Type:        schema.TypeBool,
			Description: "When true, this filter will be hidden.",
			Optional:    true, // nullable
		},
		"modifier": {
			Type:        schema.TypeString,
			Description: "Possible Values: `Contains`, `Equals`, `StartsWith`",
			Optional:    true, // precondition
		},
		"splitting_chart": {
			Type:        schema.TypeList,
			Description: "Configure splitting chart for this filter.",
			Optional:    true, // nullable
			Elem:        &schema.Resource{Schema: new(SplittingChartConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"type": {
			Type:        schema.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *EntityFilterDefinition) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"default_search":  me.DefaultSearch,
		"display_name":    me.DisplayName,
		"distinct":        me.Distinct,
		"entity_types":    me.EntityTypes,
		"free_text":       me.FreeText,
		"hidden":          me.Hidden,
		"modifier":        me.Modifier,
		"splitting_chart": me.SplittingChart,
		"type":            me.Type,
	})
}

func (me *EntityFilterDefinition) HandlePreconditions() error {
	if me.DefaultSearch == nil && me.FreeText {
		me.DefaultSearch = opt.NewBool(false)
	}
	if me.Modifier == nil && me.FreeText {
		return fmt.Errorf("'modifier' must be specified if 'free_text' is set to '%v'", me.FreeText)
	}
	if me.Modifier != nil && !me.FreeText {
		return fmt.Errorf("'modifier' must not be specified if 'free_text' is set to '%v'", me.FreeText)
	}
	return nil
}

func (me *EntityFilterDefinition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"default_search":  &me.DefaultSearch,
		"display_name":    &me.DisplayName,
		"distinct":        &me.Distinct,
		"entity_types":    &me.EntityTypes,
		"free_text":       &me.FreeText,
		"hidden":          &me.Hidden,
		"modifier":        &me.Modifier,
		"splitting_chart": &me.SplittingChart,
		"type":            &me.Type,
	})
}
