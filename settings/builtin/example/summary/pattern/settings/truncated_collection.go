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

package pattern

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type TruncatedCollections []*TruncatedCollection

func (me *TruncatedCollections) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"truncated_collection": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(TruncatedCollection).Schema()},
		},
	}
}

func (me TruncatedCollections) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("truncated_collection", me)
}

func (me *TruncatedCollections) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("truncated_collection", me)
}

type TruncatedCollection struct {
	CollProperty []string `json:"collProperty,omitempty"`
}

func (me *TruncatedCollection) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"coll_property": {
			Type:        schema.TypeSet,
			Description: "no documentation available",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
	}
}

func (me *TruncatedCollection) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"coll_property": me.CollProperty,
	})
}

func (me *TruncatedCollection) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"coll_property": &me.CollProperty,
	})
}
