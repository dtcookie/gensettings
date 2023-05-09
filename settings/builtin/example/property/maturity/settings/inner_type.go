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

package maturity

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type InnerTypes []*InnerType

func (me *InnerTypes) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"inner_type": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(InnerType).Schema()},
		},
	}
}

func (me InnerTypes) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("inner_type", me)
}

func (me *InnerTypes) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("inner_type", me)
}

// Inner type. Inner type containing a property depending on a feature flag
type InnerType struct {
	Inner1 *string `json:"inner1,omitempty"` // Property 1
}

func (me *InnerType) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"inner_1": {
			Type:        schema.TypeString,
			Description: "Property 1",
			Optional:    true, // nullable
		},
	}
}

func (me *InnerType) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"inner_1": me.Inner1,
	})
}

func (me *InnerType) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"inner_1": &me.Inner1,
	})
}
