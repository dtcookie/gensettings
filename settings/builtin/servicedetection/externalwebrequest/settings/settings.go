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

package externalwebrequest

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Conditions      Conditions          `json:"conditions"`            // A list of conditions for the rule.\nIf multiple conditions are specified, they *all* must match for the rule to apply.
	Description     *string             `json:"description,omitempty"` // Description
	Enabled         bool                `json:"enabled"`               // This setting is enabled (`true`) or disabled (`false`)
	IdContributors  *IdContributorsType `json:"idContributors"`        // All contributors of the service identifier calculation.
	ManagementZones []string            `json:"managementZones"`       // Define a management zone filter for this service detection rule.
	Name            string              `json:"name"`                  // Rule name
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"conditions": {
			Type:        schema.TypeList,
			Description: "A list of conditions for the rule.\nIf multiple conditions are specified, they *all* must match for the rule to apply.",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(Conditions).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"description": {
			Type:        schema.TypeString,
			Description: "Description",
			Optional:    true,
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "This setting is enabled (`true`) or disabled (`false`)",
			Required:    true,
		},
		"id_contributors": {
			Type:        schema.TypeList,
			Description: "All contributors of the service identifier calculation.",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(IdContributorsType).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"management_zones": {
			Type:        schema.TypeSet,
			Description: "Define a management zone filter for this service detection rule.",
			Required:    true,

			Elem: &schema.Schema{Type: schema.TypeString},
		},
		"name": {
			Type:        schema.TypeString,
			Description: "Rule name",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"conditions":       me.Conditions,
		"description":      me.Description,
		"enabled":          me.Enabled,
		"id_contributors":  me.IdContributors,
		"management_zones": me.ManagementZones,
		"name":             me.Name,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"conditions":       &me.Conditions,
		"description":      &me.Description,
		"enabled":          &me.Enabled,
		"id_contributors":  &me.IdContributors,
		"management_zones": &me.ManagementZones,
		"name":             &me.Name,
	})
}
