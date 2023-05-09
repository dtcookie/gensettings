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

package metrictable

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// VisualizationOptions. Set visualization options for timeseries. You can set it globally for whole chart, or individually for each metric selector.// // Metric selector options override options for charts.
type VisualizationOptions struct {
	SeriesType *SeriesType     `json:"seriesType,omitempty"` // Possible Values: `AREA`, `COLUMN`, `LINE`
	ShowLegend *bool           `json:"showLegend,omitempty"` // When true, axis legend on the whole chart is visible
	ThemeColor *ThemeColorEnum `json:"themeColor,omitempty"` // Possible Values: `BLUE`, `DEFAULT`, `GRAY`, `GREEN`, `ORANGE`, `PURPLE`, `RED`, `ROYALBLUE`, `TURQUOISE`, `YELLOW`
}

func (me *VisualizationOptions) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"series_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `AREA`, `COLUMN`, `LINE`",
			Optional:    true, // nullable
		},
		"show_legend": {
			Type:        schema.TypeBool,
			Description: "When true, axis legend on the whole chart is visible",
			Optional:    true, // nullable
		},
		"theme_color": {
			Type:        schema.TypeString,
			Description: "Possible Values: `BLUE`, `DEFAULT`, `GRAY`, `GREEN`, `ORANGE`, `PURPLE`, `RED`, `ROYALBLUE`, `TURQUOISE`, `YELLOW`",
			Optional:    true, // nullable
		},
	}
}

func (me *VisualizationOptions) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"series_type": me.SeriesType,
		"show_legend": me.ShowLegend,
		"theme_color": me.ThemeColor,
	})
}

func (me *VisualizationOptions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"series_type": &me.SeriesType,
		"show_legend": &me.ShowLegend,
		"theme_color": &me.ThemeColor,
	})
}
