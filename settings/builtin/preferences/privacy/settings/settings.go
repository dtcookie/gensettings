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

package privacy

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	DataCollection *DataCollection `json:"dataCollection"` // To provide your end users with the ability to decide for themselves if their activities should be tracked to measure application performance and usage, enable opt-in mode.
	DoNotTrack     *DoNotTrack     `json:"doNotTrack"`     // Most modern web browsers have a privacy feature called [\"Do Not Track\"](https://dt-url.net/sb3n0pnl) that individual users may have enabled on their devices. Customize how Dynatrace should behave when it encounters this setting.
	Masking        *Masking        `json:"masking"`
	Scope          string          `json:"-" scope:"scope"` // The scope of this setting (APPLICATION environment)
	UserTracking   *UserTracking   `json:"userTracking"`    // User tracking
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"data_collection": {
			Type:        schema.TypeList,
			Description: "To provide your end users with the ability to decide for themselves if their activities should be tracked to measure application performance and usage, enable opt-in mode.",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(DataCollection).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"do_not_track": {
			Type:        schema.TypeList,
			Description: "Most modern web browsers have a privacy feature called [\"Do Not Track\"](https://dt-url.net/sb3n0pnl) that individual users may have enabled on their devices. Customize how Dynatrace should behave when it encounters this setting.",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(DoNotTrack).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"masking": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(Masking).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (APPLICATION environment)",
			Required:    true,
		},
		"user_tracking": {
			Type:        schema.TypeList,
			Description: "User tracking",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(UserTracking).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"data_collection": me.DataCollection,
		"do_not_track":    me.DoNotTrack,
		"masking":         me.Masking,
		"scope":           me.Scope,
		"user_tracking":   me.UserTracking,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"data_collection": &me.DataCollection,
		"do_not_track":    &me.DoNotTrack,
		"masking":         &me.Masking,
		"scope":           &me.Scope,
		"user_tracking":   &me.UserTracking,
	})
}
