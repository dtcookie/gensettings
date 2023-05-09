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

package settingszones

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type EntitySelectorRules []*EntitySelectorRule

func (me *EntitySelectorRules) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"entity_selector_rule": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(EntitySelectorRule).Schema()},
		},
	}
}

func (me EntitySelectorRules) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("entity_selector_rule", me)
}

func (me *EntitySelectorRules) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("entity_selector_rule", me)
}

type EntitySelectorRule struct {
	Enabled        bool                `json:"enabled"`        // This setting is enabled (`true`) or disabled (`false`)
	EntitySelector string              `json:"entitySelector"` // The documentation of the entity selector can be found [here](https://dt-url.net/apientityselector). The type is optional here (already defined above).
	MeType         MonitoredEntityType `json:"meType"`         // Possible Values: `SERVICE`
}

func (me *EntitySelectorRule) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Type:        schema.TypeBool,
			Description: "This setting is enabled (`true`) or disabled (`false`)",
			Required:    true,
		},
		"entity_selector": {
			Type:        schema.TypeString,
			Description: "The documentation of the entity selector can be found [here](https://dt-url.net/apientityselector). The type is optional here (already defined above).",
			Required:    true,
		},
		"me_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `SERVICE`",
			Required:    true,
		},
	}
}

func (me *EntitySelectorRule) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"enabled":         me.Enabled,
		"entity_selector": me.EntitySelector,
		"me_type":         me.MeType,
	})
}

func (me *EntitySelectorRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":         &me.Enabled,
		"entity_selector": &me.EntitySelector,
		"me_type":         &me.MeType,
	})
}
