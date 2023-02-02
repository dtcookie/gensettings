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

package rumweb

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AppTrafficSpikes struct {
	Enabled       bool           `json:"enabled"`                 // Detect traffic spikes
	TrafficSpikes *TrafficSpikes `json:"trafficSpikes,omitempty"` // Dynatrace learns your typical application traffic over an observation period of one week.\n\nDepending on this expected value Dynatrace detects abnormal traffic spikes within your application.
}

func (me *AppTrafficSpikes) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Type:        schema.TypeBool,
			Description: "Detect traffic spikes",
			Required:    true,
		},
		"traffic_spikes": {
			Type:        schema.TypeList,
			Description: "Dynatrace learns your typical application traffic over an observation period of one week.\n\nDepending on this expected value Dynatrace detects abnormal traffic spikes within your application.",
			Optional:    true,
			Elem:        &schema.Resource{Schema: new(TrafficSpikes).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *AppTrafficSpikes) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"enabled":        me.Enabled,
		"traffic_spikes": me.TrafficSpikes,
	})
}

func (me *AppTrafficSpikes) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":        &me.Enabled,
		"traffic_spikes": &me.TrafficSpikes,
	})
}
