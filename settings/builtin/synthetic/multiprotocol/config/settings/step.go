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

package config

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Steps []*Step

func (me *Steps) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"step": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(Step).Schema()},
		},
	}
}

func (me Steps) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("step", me)
}

func (me *Steps) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("step", me)
}

type Step struct {
	Constraints           Constraints           `json:"constraints,omitempty"` // List of step constraints
	PostScript            *string               `json:"postScript,omitempty"`  // Post-execution script
	PreScript             *string               `json:"preScript,omitempty"`   // Pre-execution script
	Properties            Properties            `json:"properties,omitempty"`
	RequestConfigurations RequestConfigurations `json:"requestConfigurations,omitempty"` // List of request configurations
	RequestType           RequestType           `json:"requestType"`                     // Possible Values: `ICMP`, `TCP`
	TargetFilter          *string               `json:"targetFilter,omitempty"`          // Hosts filter
	TargetList            []string              `json:"targetList"`                      // List of hosts
}

func (me *Step) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"constraints": {
			Type:        schema.TypeList,
			Description: "List of step constraints",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(Constraints).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"post_script": {
			Type:        schema.TypeString,
			Description: "Post-execution script",
			Optional:    true, // nullable
		},
		"pre_script": {
			Type:        schema.TypeString,
			Description: "Pre-execution script",
			Optional:    true, // nullable
		},
		"properties": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(Properties).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"request_configurations": {
			Type:        schema.TypeList,
			Description: "List of request configurations",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(RequestConfigurations).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"request_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `ICMP`, `TCP`",
			Required:    true,
		},
		"target_filter": {
			Type:        schema.TypeString,
			Description: "Hosts filter",
			Optional:    true, // nullable
		},
		"target_list": {
			Type:        schema.TypeSet,
			Description: "List of hosts",
			Required:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
	}
}

func (me *Step) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"constraints":            me.Constraints,
		"post_script":            me.PostScript,
		"pre_script":             me.PreScript,
		"properties":             me.Properties,
		"request_configurations": me.RequestConfigurations,
		"request_type":           me.RequestType,
		"target_filter":          me.TargetFilter,
		"target_list":            me.TargetList,
	})
}

func (me *Step) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"constraints":            &me.Constraints,
		"post_script":            &me.PostScript,
		"pre_script":             &me.PreScript,
		"properties":             &me.Properties,
		"request_configurations": &me.RequestConfigurations,
		"request_type":           &me.RequestType,
		"target_filter":          &me.TargetFilter,
		"target_list":            &me.TargetList,
	})
}
