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

package privacypreferences

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type BlockListRules []*BlockListRule

func (me *BlockListRules) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"playback_masking_block_list_rule": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(BlockListRule).Schema()},
		},
	}
}

func (me BlockListRules) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("playback_masking_block_list_rule", me)
}

func (me *BlockListRules) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("playback_masking_block_list_rule", me)
}

type BlockListRule struct {
	AttributeExpression *string           `json:"attributeExpression,omitempty"` // Attribute masking can be applied to web applications that store data within attributes, typically data-NAME attributes in HTML5. When you define attributes, their values are masked while recording but not removed.
	CssExpression       *string           `json:"cssExpression,omitempty"`       // Content masking can be applied to webpages where personal data is displayed. When content masking is applied to parent elements, all child elements are masked by default.
	HideUserInteraction *bool             `json:"hideUserInteraction,omitempty"` // Hide user interactions with these elements, including clicks that expand elements, highlighting that results from hovering a cursor over an option, and selection of specific form options.
	Target              MaskingTargetType `json:"target"`                        // Possible Values: `ATTRIBUTE`, `ELEMENT`
}

func (me *BlockListRule) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attribute_expression": {
			Type:        schema.TypeString,
			Description: "Attribute masking can be applied to web applications that store data within attributes, typically data-NAME attributes in HTML5. When you define attributes, their values are masked while recording but not removed.",
			Optional:    true, // precondition
		},
		"css_expression": {
			Type:        schema.TypeString,
			Description: "Content masking can be applied to webpages where personal data is displayed. When content masking is applied to parent elements, all child elements are masked by default.",
			Optional:    true, // precondition
		},
		"hide_user_interaction": {
			Type:        schema.TypeBool,
			Description: "Hide user interactions with these elements, including clicks that expand elements, highlighting that results from hovering a cursor over an option, and selection of specific form options.",
			Optional:    true, // precondition
		},
		"target": {
			Type:        schema.TypeString,
			Description: "Possible Values: `ATTRIBUTE`, `ELEMENT`",
			Required:    true,
		},
	}
}

func (me *BlockListRule) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"attribute_expression":  me.AttributeExpression,
		"css_expression":        me.CssExpression,
		"hide_user_interaction": me.HideUserInteraction,
		"target":                me.Target,
	})
}

func (me *BlockListRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"attribute_expression":  &me.AttributeExpression,
		"css_expression":        &me.CssExpression,
		"hide_user_interaction": &me.HideUserInteraction,
		"target":                &me.Target,
	})
}
