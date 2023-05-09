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

package kitchensink

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	CollectionPropertyTypes *CollectionPropertyTypes `json:"collectionPropertyTypes"` // All supported collection property types
	MarkdownShowcase        *MarkdownShowcase        `json:"MarkdownShowcase"`        // It's very easy to make some words **bold** and other words *italic* with Markdown.\n\nSometimes you want numbered lists:\n\n1. One\n2. Two\n3. Three\n\nSometimes you want bullet points:\n\n* Start a line with a star\n* Profit!\n\nAlternatively,\n\n- Dashes work just as well\n- And if you have sub points, put two spaces before the dash or star:\n  - Like this\n  - And this\n\n## Structured documents\n\nSometimes it's useful to have different levels of headings to structure your documents. Start lines with a `#` to create headings. Multiple `##` in a row denote smaller heading sizes.\n\n### This is a third-tier heading\n\nYou can use one `#` all the way up to `######` six for different heading sizes.\n\nIf you'd like to quote someone, use the > character before the line:\n\n> Coffee. The finest organic suspension ever devised... I beat the Borg with it.\n> - Captain Janeway
	PropertyTypes           *PropertyTypes           `json:"propertyTypes"`           // All supported property types
	SubPropertyTypes        *SubPropertyTypes        `json:"subPropertyTypes"`        // All supported sub property types
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"collection_property_types": {
			Type:        schema.TypeList,
			Description: "All supported collection property types",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(CollectionPropertyTypes).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"markdown_showcase": {
			Type:        schema.TypeList,
			Description: "It's very easy to make some words **bold** and other words *italic* with Markdown.\n\nSometimes you want numbered lists:\n\n1. One\n2. Two\n3. Three\n\nSometimes you want bullet points:\n\n* Start a line with a star\n* Profit!\n\nAlternatively,\n\n- Dashes work just as well\n- And if you have sub points, put two spaces before the dash or star:\n  - Like this\n  - And this\n\n## Structured documents\n\nSometimes it's useful to have different levels of headings to structure your documents. Start lines with a `#` to create headings. Multiple `##` in a row denote smaller heading sizes.\n\n### This is a third-tier heading\n\nYou can use one `#` all the way up to `######` six for different heading sizes.\n\nIf you'd like to quote someone, use the > character before the line:\n\n> Coffee. The finest organic suspension ever devised... I beat the Borg with it.\n> - Captain Janeway",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(MarkdownShowcase).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"property_types": {
			Type:        schema.TypeList,
			Description: "All supported property types",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(PropertyTypes).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"sub_property_types": {
			Type:        schema.TypeList,
			Description: "All supported sub property types",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(SubPropertyTypes).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"collection_property_types": me.CollectionPropertyTypes,
		"markdown_showcase":         me.MarkdownShowcase,
		"property_types":            me.PropertyTypes,
		"sub_property_types":        me.SubPropertyTypes,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"collection_property_types": &me.CollectionPropertyTypes,
		"markdown_showcase":         &me.MarkdownShowcase,
		"property_types":            &me.PropertyTypes,
		"sub_property_types":        &me.SubPropertyTypes,
	})
}
