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

type Settings struct {
	EventFilters   AlertingProfileEventFilters  `json:"eventFilters,omitempty"`   // Define event filters for profile. A maximum of 100 event filters is allowed.
	ManagementZone *string                      `json:"managementZone,omitempty"` // Define management zone filter for profile
	Name           string                       `json:"name"`                     // Name
	SeverityRules  AlertingProfileSeverityRules `json:"severityRules,omitempty"`  // Define severity rules for profile. A maximum of 100 severity rules is allowed.
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"event_filters": {
			Type:        schema.TypeList,
			Description: "Define event filters for profile. A maximum of 100 event filters is allowed.",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(AlertingProfileEventFilters).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"management_zone": {
			Type:        schema.TypeString,
			Description: "Define management zone filter for profile",
			Optional:    true, // nullable
		},
		"name": {
			Type:        schema.TypeString,
			Description: "Name",
			Required:    true,
		},
		"severity_rules": {
			Type:        schema.TypeList,
			Description: "Define severity rules for profile. A maximum of 100 severity rules is allowed.",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(AlertingProfileSeverityRules).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"event_filters":   me.EventFilters,
		"management_zone": me.ManagementZone,
		"name":            me.Name,
		"severity_rules":  me.SeverityRules,
	})
}

func (me *Settings) HandlePreconditions() {
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"event_filters":   &me.EventFilters,
		"management_zone": &me.ManagementZone,
		"name":            &me.Name,
		"severity_rules":  &me.SeverityRules,
	})
}
