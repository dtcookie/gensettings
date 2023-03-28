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

package scheduling

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Locations []*Location

func (me *Locations) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"location": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(Location).Schema()},
		},
	}
}

func (me Locations) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("location", me)
}

func (me *Locations) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("location", me)
}

type Location struct {
	Location *string `json:"location,omitempty"`
}

func (me *Location) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"location": {
			Type:        schema.TypeString,
			Description: "no documentation available",
			Optional:    true, // nullable
		},
	}
}

func (me *Location) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"location": me.Location,
	})
}

func (me *Location) HandlePreconditions() {
}

func (me *Location) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"location": &me.Location,
	})
}
