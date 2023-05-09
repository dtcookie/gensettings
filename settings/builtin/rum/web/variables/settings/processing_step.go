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

package variables

import (
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/opt"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"
)

type ProcessingSteps []*ProcessingStep

func (me *ProcessingSteps) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"processing_step": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(ProcessingStep).Schema()},
		},
	}
}

func (me ProcessingSteps) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("processing_step", me)
}

func (me *ProcessingSteps) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("processing_step", me)
}

type ProcessingStep struct {
	FallbackToSelectedSource *bool              `json:"fallbackToSelectedSource,omitempty"` // Fall back to selected source
	PatternAfter             *string            `json:"patternAfter,omitempty"`             // Trailing delimiter (extract/replace ends here)
	PatternAfterSearchType   *PatternSearchType `json:"patternAfterSearchType,omitempty"`   // Possible Values: `FIRST`, `LAST`
	PatternBefore            *string            `json:"patternBefore,omitempty"`            // Leading delimiter (extract/replace starts here)
	PatternBeforeSearchType  *PatternSearchType `json:"patternBeforeSearchType,omitempty"`  // Possible Values: `FIRST`, `LAST`
	RegularExpression        *string            `json:"regularExpression,omitempty"`        // Regular expression
	Replacement              *string            `json:"replacement,omitempty"`              // Replace with – Optional
	SearchPattern            *string            `json:"searchPattern,omitempty"`            // Text to search
	Type                     ProcessingStepType `json:"type"`                               // Possible Values: `CONVERT_TO_LOWER_CASE`, `CONVERT_TO_UPPER_CASE`, `EXTRACT`, `EXTRACT_WITH_REGEX`, `REPLACE`, `REPLACE_IDS`, `REPLACE_TEXT`, `REPLACE_WITH_REGEX`
}

func (me *ProcessingStep) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"fallback_to_selected_source": {
			Type:        schema.TypeBool,
			Description: "Fall back to selected source",
			Optional:    true, // precondition
		},
		"pattern_after": {
			Type:        schema.TypeString,
			Description: "Trailing delimiter (extract/replace ends here)",
			Optional:    true, // precondition
		},
		"pattern_after_search_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `FIRST`, `LAST`",
			Optional:    true, // precondition
		},
		"pattern_before": {
			Type:        schema.TypeString,
			Description: "Leading delimiter (extract/replace starts here)",
			Optional:    true, // precondition
		},
		"pattern_before_search_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `FIRST`, `LAST`",
			Optional:    true, // precondition
		},
		"regular_expression": {
			Type:        schema.TypeString,
			Description: "Regular expression",
			Optional:    true, // precondition
		},
		"replacement": {
			Type:        schema.TypeString,
			Description: "Replace with – Optional",
			Optional:    true, // nullable & precondition
		},
		"search_pattern": {
			Type:        schema.TypeString,
			Description: "Text to search",
			Optional:    true, // precondition
		},
		"type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `CONVERT_TO_LOWER_CASE`, `CONVERT_TO_UPPER_CASE`, `EXTRACT`, `EXTRACT_WITH_REGEX`, `REPLACE`, `REPLACE_IDS`, `REPLACE_TEXT`, `REPLACE_WITH_REGEX`",
			Required:    true,
		},
	}
}

func (me *ProcessingStep) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"fallback_to_selected_source": me.FallbackToSelectedSource,
		"pattern_after":               me.PatternAfter,
		"pattern_after_search_type":   me.PatternAfterSearchType,
		"pattern_before":              me.PatternBefore,
		"pattern_before_search_type":  me.PatternBeforeSearchType,
		"regular_expression":          me.RegularExpression,
		"replacement":                 me.Replacement,
		"search_pattern":              me.SearchPattern,
		"type":                        me.Type,
	})
}

func (me *ProcessingStep) HandlePreconditions() error {
	if me.FallbackToSelectedSource == nil && !slices.Contains([]string{"CONVERT_TO_LOWER_CASE", "CONVERT_TO_UPPER_CASE"}, string(me.Type)) {
		me.FallbackToSelectedSource = opt.NewBool(false)
	}
	if me.PatternAfter == nil && slices.Contains([]string{"EXTRACT", "REPLACE"}, string(me.Type)) {
		me.PatternAfter = opt.NewString("")
	}
	if me.PatternBefore == nil && slices.Contains([]string{"EXTRACT", "REPLACE"}, string(me.Type)) {
		me.PatternBefore = opt.NewString("")
	}
	if me.RegularExpression == nil && slices.Contains([]string{"EXTRACT_WITH_REGEX", "REPLACE_WITH_REGEX"}, string(me.Type)) {
		me.RegularExpression = opt.NewString("")
	}
	if me.SearchPattern == nil && slices.Contains([]string{"REPLACE_TEXT"}, string(me.Type)) {
		me.SearchPattern = opt.NewString("")
	}
	if me.PatternAfterSearchType == nil && slices.Contains([]string{"EXTRACT", "REPLACE"}, string(me.Type)) {
		return fmt.Errorf("'pattern_after_search_type' must be specified if 'type' is set to '%v'", me.Type)
	}
	if me.PatternAfterSearchType != nil && !slices.Contains([]string{"EXTRACT", "REPLACE"}, string(me.Type)) {
		return fmt.Errorf("'pattern_after_search_type' must not be specified if 'type' is set to '%v'", me.Type)
	}
	if me.PatternBeforeSearchType == nil && slices.Contains([]string{"EXTRACT", "REPLACE"}, string(me.Type)) {
		return fmt.Errorf("'pattern_before_search_type' must be specified if 'type' is set to '%v'", me.Type)
	}
	if me.PatternBeforeSearchType != nil && !slices.Contains([]string{"EXTRACT", "REPLACE"}, string(me.Type)) {
		return fmt.Errorf("'pattern_before_search_type' must not be specified if 'type' is set to '%v'", me.Type)
	}
	if me.Replacement == nil && slices.Contains([]string{"REPLACE", "REPLACE_TEXT", "REPLACE_WITH_REGEX", "REPLACE_IDS"}, string(me.Type)) {
		return fmt.Errorf("'replacement' must be specified if 'type' is set to '%v'", me.Type)
	}
	return nil
}

func (me *ProcessingStep) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"fallback_to_selected_source": &me.FallbackToSelectedSource,
		"pattern_after":               &me.PatternAfter,
		"pattern_after_search_type":   &me.PatternAfterSearchType,
		"pattern_before":              &me.PatternBefore,
		"pattern_before_search_type":  &me.PatternBeforeSearchType,
		"regular_expression":          &me.RegularExpression,
		"replacement":                 &me.Replacement,
		"search_pattern":              &me.SearchPattern,
		"type":                        &me.Type,
	})
}
