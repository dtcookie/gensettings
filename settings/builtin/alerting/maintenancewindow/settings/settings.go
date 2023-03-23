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

package maintenancewindow

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Enabled           bool               `json:"enabled"`           // This setting is enabled (`true`) or disabled (`false`)
	Filters           Filters            `json:"filters,omitempty"` // ## Filters\nAdd filters to limit the scope of maintenance to only select matching entities. If no filter is defined, the maintenance window is valid for the whole environment. Each filter is evaluated separately (**OR**).
	GeneralProperties *GeneralProperties `json:"generalProperties"` // ## Properties
	Schedule          *Schedule          `json:"schedule"`          // ## Schedule
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Type:        schema.TypeBool,
			Description: "This setting is enabled (`true`) or disabled (`false`)",
			Required:    true,
		},
		"filters": {
			Type:        schema.TypeList,
			Description: "## Filters\nAdd filters to limit the scope of maintenance to only select matching entities. If no filter is defined, the maintenance window is valid for the whole environment. Each filter is evaluated separately (**OR**).",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(Filters).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"general_properties": {
			Type:        schema.TypeList,
			Description: "## Properties",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(GeneralProperties).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"schedule": {
			Type:        schema.TypeList,
			Description: "## Schedule",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(Schedule).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"enabled":            me.Enabled,
		"filters":            me.Filters,
		"general_properties": me.GeneralProperties,
		"schedule":           me.Schedule,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":            &me.Enabled,
		"filters":            &me.Filters,
		"general_properties": &me.GeneralProperties,
		"schedule":           &me.Schedule,
	})
}
