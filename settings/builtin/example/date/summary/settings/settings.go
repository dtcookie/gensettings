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

package summary

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Local_date      *string `json:"local_date,omitempty"`      // local_date
	Local_date_time *string `json:"local_date_time,omitempty"` // local_date_time
	Local_time      *string `json:"local_time,omitempty"`      // local_time
	Zoned_date      *string `json:"zoned_date,omitempty"`      // zoned_date_time
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"local_date": {
			Type:        schema.TypeString,
			Description: "local_date",
			Optional:    true, // nullable
		},
		"local_date_time": {
			Type:        schema.TypeString,
			Description: "local_date_time",
			Optional:    true, // nullable
		},
		"local_time": {
			Type:        schema.TypeString,
			Description: "local_time",
			Optional:    true, // nullable
		},
		"zoned_date": {
			Type:        schema.TypeString,
			Description: "zoned_date_time",
			Optional:    true, // nullable
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"local_date":      me.Local_date,
		"local_date_time": me.Local_date_time,
		"local_time":      me.Local_time,
		"zoned_date":      me.Zoned_date,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"local_date":      &me.Local_date,
		"local_date_time": &me.Local_date_time,
		"local_time":      &me.Local_time,
		"zoned_date":      &me.Zoned_date,
	})
}
