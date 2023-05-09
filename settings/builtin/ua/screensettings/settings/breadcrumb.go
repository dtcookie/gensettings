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

package screensettings

import (
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"
)

type Breadcrumbs []*Breadcrumb

func (me *Breadcrumbs) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"breadcrumb": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(Breadcrumb).Schema()},
		},
	}
}

func (me Breadcrumbs) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("breadcrumb", me)
}

func (me *Breadcrumbs) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("breadcrumb", me)
}

type Breadcrumb struct {
	DisplayName            *string        `json:"displayName,omitempty"`            // A display name for a non-clickable or entity list breadcrumb
	EntitySelectorTemplate *string        `json:"entitySelectorTemplate,omitempty"` // An entity selector that is used to reference related entity screen. For example for a process screen where you want to show the host it runs on, this might look like this:  \n`type(HOST), toRelationships.isProcessOf($(entityConditions))`\n\nPlease mind that the `$(entityConditions)` is a placeholder for the relation condition.
	EntityType             *string        `json:"entityType,omitempty"`             // An entity type of entities list screen
	Link                   *string        `json:"link,omitempty"`                   // A key of server-defined static link
	Type                   BreadcrumbType `json:"type"`                             // Possible Values: `ENTITY_LIST_REF`, `ENTITY_REF`, `NOOP`, `STATIC_LINK`
}

func (me *Breadcrumb) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"display_name": {
			Type:        schema.TypeString,
			Description: "A display name for a non-clickable or entity list breadcrumb",
			Optional:    true, // nullable & precondition
		},
		"entity_selector_template": {
			Type:        schema.TypeString,
			Description: "An entity selector that is used to reference related entity screen. For example for a process screen where you want to show the host it runs on, this might look like this:  \n`type(HOST), toRelationships.isProcessOf($(entityConditions))`\n\nPlease mind that the `$(entityConditions)` is a placeholder for the relation condition.",
			Optional:    true, // precondition
		},
		"entity_type": {
			Type:        schema.TypeString,
			Description: "An entity type of entities list screen",
			Optional:    true, // precondition
		},
		"link": {
			Type:        schema.TypeString,
			Description: "A key of server-defined static link",
			Optional:    true, // precondition
		},
		"type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `ENTITY_LIST_REF`, `ENTITY_REF`, `NOOP`, `STATIC_LINK`",
			Required:    true,
		},
	}
}

func (me *Breadcrumb) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"display_name":             me.DisplayName,
		"entity_selector_template": me.EntitySelectorTemplate,
		"entity_type":              me.EntityType,
		"link":                     me.Link,
		"type":                     me.Type,
	})
}

func (me *Breadcrumb) HandlePreconditions() error {
	if me.DisplayName == nil && slices.Contains([]string{"NOOP", "ENTITY_LIST_REF"}, string(me.Type)) {
		return fmt.Errorf("'display_name' must be specified if 'type' is set to '%v'", me.Type)
	}
	if me.EntitySelectorTemplate == nil && (string(me.Type) == "ENTITY_REF") {
		return fmt.Errorf("'entity_selector_template' must be specified if 'type' is set to '%v'", me.Type)
	}
	if me.EntityType == nil && (string(me.Type) == "ENTITY_LIST_REF") {
		return fmt.Errorf("'entity_type' must be specified if 'type' is set to '%v'", me.Type)
	}
	if me.Link == nil && (string(me.Type) == "STATIC_LINK") {
		return fmt.Errorf("'link' must be specified if 'type' is set to '%v'", me.Type)
	}
	return nil
}

func (me *Breadcrumb) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"display_name":             &me.DisplayName,
		"entity_selector_template": &me.EntitySelectorTemplate,
		"entity_type":              &me.EntityType,
		"link":                     &me.Link,
		"type":                     &me.Type,
	})
}
