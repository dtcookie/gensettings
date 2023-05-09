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

type Settings struct {
	Icu_number_formatting IcuNumberFormattingItems `json:"icu_number_formatting,omitempty"` // This summary makes use of some of ICU's number formatting capabilities
	Icu_plural            IcuPluralItems           `json:"icu_plural,omitempty"`            // This summary makes use of ICU's `plural` to display some text depending on an integer value
	Icu_select            IcuSelectItems           `json:"icu_select,omitempty"`            // This summary makes use of ICU's `select` to display some text depending on a boolean flag
	Non_empty_text        NonEmptyTextItems        `json:"non_empty_text,omitempty"`        // This summary shows how to use pipes, in particular the `valueSet` pipe to check if a value is set
	Truncated             TruncatedTextItems       `json:"truncated,omitempty"`             // This summary shows how to use pipes, in particular the `truncate` pipe to limit the length of the displayed text to 10 characters
	TruncatedCollection   TruncatedCollections     `json:"truncatedCollection,omitempty"`   // This summary shows how to use pipes, in particular the `truncateCollection` pipe to limit the size of the collection to 3 elements and the length of each text item in it to 4 characters
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"icu_number_formatting": {
			Type:        schema.TypeList,
			Description: "This summary makes use of some of ICU's number formatting capabilities",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(IcuNumberFormattingItems).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"icu_plural": {
			Type:        schema.TypeList,
			Description: "This summary makes use of ICU's `plural` to display some text depending on an integer value",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(IcuPluralItems).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"icu_select": {
			Type:        schema.TypeList,
			Description: "This summary makes use of ICU's `select` to display some text depending on a boolean flag",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(IcuSelectItems).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"non_empty_text": {
			Type:        schema.TypeList,
			Description: "This summary shows how to use pipes, in particular the `valueSet` pipe to check if a value is set",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(NonEmptyTextItems).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"truncated": {
			Type:        schema.TypeList,
			Description: "This summary shows how to use pipes, in particular the `truncate` pipe to limit the length of the displayed text to 10 characters",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(TruncatedTextItems).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"truncated_collection": {
			Type:        schema.TypeList,
			Description: "This summary shows how to use pipes, in particular the `truncateCollection` pipe to limit the size of the collection to 3 elements and the length of each text item in it to 4 characters",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(TruncatedCollections).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"icu_number_formatting": me.Icu_number_formatting,
		"icu_plural":            me.Icu_plural,
		"icu_select":            me.Icu_select,
		"non_empty_text":        me.Non_empty_text,
		"truncated":             me.Truncated,
		"truncated_collection":  me.TruncatedCollection,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"icu_number_formatting": &me.Icu_number_formatting,
		"icu_plural":            &me.Icu_plural,
		"icu_select":            &me.Icu_select,
		"non_empty_text":        &me.Non_empty_text,
		"truncated":             &me.Truncated,
		"truncated_collection":  &me.TruncatedCollection,
	})
}
