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

package entitieslist

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SplittingChartConfig struct {
	DisplayChart bool `json:"displayChart"` // When true, splitting chart for this filter will be displayed
}

func (me *SplittingChartConfig) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"display_chart": {
			Type:        schema.TypeBool,
			Description: "When true, splitting chart for this filter will be displayed",
			Required:    true,
		},
	}
}

func (me *SplittingChartConfig) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"display_chart": me.DisplayChart,
	})
}

func (me *SplittingChartConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"display_chart": &me.DisplayChart,
	})
}
