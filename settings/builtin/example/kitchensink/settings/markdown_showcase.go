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

type MarkdownShowcase struct {
	Markdown_code_samples   string `json:"markdown_code_samples"`   // There are many different ways to style code with GitHub's markdown. If you have inline code blocks, wrap them in backticks: `var example = true`. If you've got a longer block of code, you can indent with four spaces:\n\n    if (isAwesome){\n      return true\n    }\n\nemoji time!\n😀😎📆🐫
	Markdown_emoji          string `json:"markdown_emoji"`          // 😀😎📆🐫
	Markdown_links_external string `json:"markdown_links_external"` // Sometimes there is a need to link to [external documentation](https://www.dynatrace.com/ \"Visit Dynatrace Help\") or link to [Dynatrace.com!](https://dynatrace.com \"Visit dynatrace.com \")
	Markdown_links_internal string `json:"markdown_links_internal"` // To link to other [Angular screen](/ui/docs \"Visit Angular docs page\") on same environment, please start your URL with `/`.\nIf you need to link to [GWT screen](/#reports \"Visit GWT Reports page\") on same environment, please start your URL with `/#`.\n\nWe also support [anchors](#top) on same pages, just start your url with `#`.\n\n**Important!** As you see - you don't need to worry about environment in your href!. **Markdown** is supported inside info popover [Angular screen](/ui/docs \"Visit Angular docs page\")
}

func (me *MarkdownShowcase) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"markdown_code_samples": {
			Type:        schema.TypeString,
			Description: "There are many different ways to style code with GitHub's markdown. If you have inline code blocks, wrap them in backticks: `var example = true`. If you've got a longer block of code, you can indent with four spaces:\n\n    if (isAwesome){\n      return true\n    }\n\nemoji time!\n😀😎📆🐫",
			Required:    true,
		},
		"markdown_emoji": {
			Type:        schema.TypeString,
			Description: "😀😎📆🐫",
			Required:    true,
		},
		"markdown_links_external": {
			Type:        schema.TypeString,
			Description: "Sometimes there is a need to link to [external documentation](https://www.dynatrace.com/ \"Visit Dynatrace Help\") or link to [Dynatrace.com!](https://dynatrace.com \"Visit dynatrace.com \")",
			Required:    true,
		},
		"markdown_links_internal": {
			Type:        schema.TypeString,
			Description: "To link to other [Angular screen](/ui/docs \"Visit Angular docs page\") on same environment, please start your URL with `/`.\nIf you need to link to [GWT screen](/#reports \"Visit GWT Reports page\") on same environment, please start your URL with `/#`.\n\nWe also support [anchors](#top) on same pages, just start your url with `#`.\n\n**Important!** As you see - you don't need to worry about environment in your href!. **Markdown** is supported inside info popover [Angular screen](/ui/docs \"Visit Angular docs page\")",
			Required:    true,
		},
	}
}

func (me *MarkdownShowcase) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"markdown_code_samples":   me.Markdown_code_samples,
		"markdown_emoji":          me.Markdown_emoji,
		"markdown_links_external": me.Markdown_links_external,
		"markdown_links_internal": me.Markdown_links_internal,
	})
}

func (me *MarkdownShowcase) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"markdown_code_samples":   &me.Markdown_code_samples,
		"markdown_emoji":          &me.Markdown_emoji,
		"markdown_links_external": &me.Markdown_links_external,
		"markdown_links_internal": &me.Markdown_links_internal,
	})
}
