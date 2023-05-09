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

package edgehttptargets

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type HostLocationMappingSet []*HostLocationMapping

func (me *HostLocationMappingSet) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"mapping": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(HostLocationMapping).Schema()},
		},
	}
}

func (me HostLocationMappingSet) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("mapping", me)
}

func (me *HostLocationMappingSet) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("mapping", me)
}

type HostLocationMapping struct {
	Host     string `json:"host"`     // Either an IP, or a domain name which may contain a wildcard ('\\*') as the leftmost label. If a '\\*' is used, it must be followed by at least two labels.
	Location string `json:"location"` // EdgeConnect name
}

func (me *HostLocationMapping) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"host": {
			Type:        schema.TypeString,
			Description: "Either an IP, or a domain name which may contain a wildcard ('\\*') as the leftmost label. If a '\\*' is used, it must be followed by at least two labels.",
			Required:    true,
		},
		"location": {
			Type:        schema.TypeString,
			Description: "EdgeConnect name",
			Required:    true,
		},
	}
}

func (me *HostLocationMapping) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"host":     me.Host,
		"location": me.Location,
	})
}

func (me *HostLocationMapping) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"host":     &me.Host,
		"location": &me.Location,
	})
}
