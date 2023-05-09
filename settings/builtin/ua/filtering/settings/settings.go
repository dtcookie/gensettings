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

package filtering

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	EntityFilters EntityFilterGroupDefinitions `json:"entityFilters,omitempty"` // Entity filters list
	Relationships []string                     `json:"relationships,omitempty"` // To filter related entities you need to define relationship between them. \nFor example to filter processes of a host you need to add entity selector with relationship like this: `type(PROCESS_GROUP_INSTANCE), fromRelationship.isProcessOf(type(HOST))`. \nInstead of `type(HOST)` you can use variable `$(entityConditions)`. \nRelationships should cover all possible pairs of entity types in filters.
	Scope         string                       `json:"-" scope:"scope"`         // The scope of this setting (ua-screen)
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"entity_filters": {
			Type:        schema.TypeList,
			Description: "Entity filters list",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(EntityFilterGroupDefinitions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"relationships": {
			Type:        schema.TypeSet,
			Description: "To filter related entities you need to define relationship between them. \nFor example to filter processes of a host you need to add entity selector with relationship like this: `type(PROCESS_GROUP_INSTANCE), fromRelationship.isProcessOf(type(HOST))`. \nInstead of `type(HOST)` you can use variable `$(entityConditions)`. \nRelationships should cover all possible pairs of entity types in filters.",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (ua-screen)",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"entity_filters": me.EntityFilters,
		"relationships":  me.Relationships,
		"scope":          me.Scope,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"entity_filters": &me.EntityFilters,
		"relationships":  &me.Relationships,
		"scope":          &me.Scope,
	})
}
