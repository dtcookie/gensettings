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

type AlertingProfileSeverityRules []*AlertingProfileSeverityRule

func (me *AlertingProfileSeverityRules) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"severity_rule": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(AlertingProfileSeverityRule).Schema()},
		},
	}
}

func (me AlertingProfileSeverityRules) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("severity_rule", me)
}

func (me *AlertingProfileSeverityRules) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("severity_rule", me)
}

type AlertingProfileSeverityRule struct {
	DelayInMinutes       int                  `json:"delayInMinutes"`       // Send a notification if a problem remains open longer than X minutes.
	SeverityLevel        SeverityLevel        `json:"severityLevel"`        // Possible Values: `AVAILABILITY`, `CUSTOM_ALERT`, `ERRORS`, `MONITORING_UNAVAILABLE`, `PERFORMANCE`, `RESOURCE_CONTENTION`
	TagFilter            []string             `json:"tagFilter,omitempty"`  // Tags
	TagFilterIncludeMode TagFilterIncludeMode `json:"tagFilterIncludeMode"` // Possible Values: `INCLUDE_ALL`, `INCLUDE_ANY`, `NONE`
}

func (me *AlertingProfileSeverityRule) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"delay_in_minutes": {
			Type:        schema.TypeInt,
			Description: "Send a notification if a problem remains open longer than X minutes.",
			Required:    true,
		},
		"severity_level": {
			Type:        schema.TypeString,
			Description: "Possible Values: `AVAILABILITY`, `CUSTOM_ALERT`, `ERRORS`, `MONITORING_UNAVAILABLE`, `PERFORMANCE`, `RESOURCE_CONTENTION`",
			Required:    true,
		},
		"tag_filter": {
			Type:        schema.TypeSet,
			Description: "Tags",
			Optional:    true, // precondition
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"tag_filter_include_mode": {
			Type:        schema.TypeString,
			Description: "Possible Values: `INCLUDE_ALL`, `INCLUDE_ANY`, `NONE`",
			Required:    true,
		},
	}
}

func (me *AlertingProfileSeverityRule) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"delay_in_minutes":        me.DelayInMinutes,
		"severity_level":          me.SeverityLevel,
		"tag_filter":              me.TagFilter,
		"tag_filter_include_mode": me.TagFilterIncludeMode,
	})
}

func (me *AlertingProfileSeverityRule) HandlePreconditions() error {
	// ---- TagFilter []string -> {"expectedValues":["INCLUDE_ANY","INCLUDE_ALL"],"property":"tagFilterIncludeMode","type":"IN"}
	return nil
}

func (me *AlertingProfileSeverityRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"delay_in_minutes":        &me.DelayInMinutes,
		"severity_level":          &me.SeverityLevel,
		"tag_filter":              &me.TagFilter,
		"tag_filter_include_mode": &me.TagFilterIncludeMode,
	})
}
