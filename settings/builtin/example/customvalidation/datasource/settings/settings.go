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

package datasource

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Advanced string     `json:"advanced"`        // Will return possible values based on type and limit, only allows datasource values (with validation). Does not use api to filter values. Some of the values are visible only with devops access.
	Complex  *InnerType `json:"complex"`         // Inner datasource works with the full setting payload on validation
	Limit    int        `json:"limit"`           // Controls the max number (inclusive) returned by the advanced datasource
	Scope    *string    `json:"-" scope:"scope"` // The scope of this setting (environment-default). Omit this property if you want to cover the whole environment.
	Simple   string     `json:"simple"`          // Will return possible values based on type, but also allows custom defined values (no validation). Uses api to filter values. Some of the values are visible only with devops access.
	Type     *Type      `json:"type,omitempty"`  // Controls the type of values the datasouces will deliver
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"advanced": {
			Type:        schema.TypeString,
			Description: "Will return possible values based on type and limit, only allows datasource values (with validation). Does not use api to filter values. Some of the values are visible only with devops access.",
			Required:    true,
		},
		"complex": {
			Type:        schema.TypeList,
			Description: "Inner datasource works with the full setting payload on validation",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(InnerType).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"limit": {
			Type:        schema.TypeInt,
			Description: "Controls the max number (inclusive) returned by the advanced datasource",
			Required:    true,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (environment-default). Omit this property if you want to cover the whole environment.",
			Optional:    true,
			Default:     "environment",
		},
		"simple": {
			Type:        schema.TypeString,
			Description: "Will return possible values based on type, but also allows custom defined values (no validation). Uses api to filter values. Some of the values are visible only with devops access.",
			Required:    true,
		},
		"type": {
			Type:        schema.TypeString,
			Description: "Controls the type of values the datasouces will deliver",
			Optional:    true, // nullable
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"advanced": me.Advanced,
		"complex":  me.Complex,
		"limit":    me.Limit,
		"scope":    me.Scope,
		"simple":   me.Simple,
		"type":     me.Type,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"advanced": &me.Advanced,
		"complex":  &me.Complex,
		"limit":    &me.Limit,
		"scope":    &me.Scope,
		"simple":   &me.Simple,
		"type":     &me.Type,
	})
}
