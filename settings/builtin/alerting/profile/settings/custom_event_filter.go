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

type CustomEventFilter struct {
	DescriptionFilter *TextFilter     `json:"descriptionFilter,omitempty"` // Description filter
	MetadataFilter    *MetadataFilter `json:"metadataFilter,omitempty"`    // Property filters
	TitleFilter       *TextFilter     `json:"titleFilter,omitempty"`       // Title filter
}

func (me *CustomEventFilter) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description_filter": {
			Type:        schema.TypeList,
			Description: "Description filter",
			Optional:    true, // nullable
			Elem:        &schema.Resource{Schema: new(TextFilter).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"metadata_filter": {
			Type:        schema.TypeList,
			Description: "Property filters",
			Optional:    true, // nullable
			Elem:        &schema.Resource{Schema: new(MetadataFilter).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"title_filter": {
			Type:        schema.TypeList,
			Description: "Title filter",
			Optional:    true, // nullable
			Elem:        &schema.Resource{Schema: new(TextFilter).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *CustomEventFilter) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"description_filter": me.DescriptionFilter,
		"metadata_filter":    me.MetadataFilter,
		"title_filter":       me.TitleFilter,
	})
}

func (me *CustomEventFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"description_filter": &me.DescriptionFilter,
		"metadata_filter":    &me.MetadataFilter,
		"title_filter":       &me.TitleFilter,
	})
}
