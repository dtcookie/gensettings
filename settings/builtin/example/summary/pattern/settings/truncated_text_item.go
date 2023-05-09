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

package pattern

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type TruncatedTextItems []*TruncatedTextItem

func (me *TruncatedTextItems) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"item": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(TruncatedTextItem).Schema()},
		},
	}
}

func (me TruncatedTextItems) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("item", me)
}

func (me *TruncatedTextItems) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("item", me)
}

type TruncatedTextItem struct {
	TextProperty *string `json:"textProperty,omitempty"` // Text
}

func (me *TruncatedTextItem) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"text_property": {
			Type:        schema.TypeString,
			Description: "Text",
			Optional:    true, // nullable
		},
	}
}

func (me *TruncatedTextItem) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"text_property": me.TextProperty,
	})
}

func (me *TruncatedTextItem) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"text_property": &me.TextProperty,
	})
}
