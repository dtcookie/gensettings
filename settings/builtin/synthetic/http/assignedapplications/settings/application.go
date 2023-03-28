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

package assignedapplications

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Applications []*Application

func (me *Applications) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(Application).Schema()},
		},
	}
}

func (me Applications) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("application", me)
}

func (me *Applications) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("application", me)
}

type Application struct {
	Application string `json:"application"`
}

func (me *Application) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application": {
			Type:        schema.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *Application) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"application": me.Application,
	})
}

func (me *Application) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"application": &me.Application,
	})
}
