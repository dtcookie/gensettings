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

// MetricVisualizationOptions. Set visualization options for timeseries. You can set it globally for whole chart, or individually for each metric selector.// // Metric selector options override options for charts.
type MetricVisualizationOptions struct {
	ColorOverride ColorOverridess `json:"colorOverride,omitempty"` // There is possibility to override color of single series, even single dimension of a metric
	DisplayName   *string         `json:"displayName,omitempty"`   // Display name
	SeriesType    *SeriesType     `json:"seriesType,omitempty"`    // Possible Values: `AREA`, `COLUMN`, `LINE`
	ThemeColor    *ThemeColorEnum `json:"themeColor,omitempty"`    // Possible Values: `BLUE`, `DEFAULT`, `GRAY`, `GREEN`, `ORANGE`, `PURPLE`, `RED`, `ROYALBLUE`, `TURQUOISE`, `YELLOW`
}

func (me *MetricVisualizationOptions) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"color_override": {
			Type:        schema.TypeList,
			Description: "There is possibility to override color of single series, even single dimension of a metric",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(ColorOverridess).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"display_name": {
			Type:        schema.TypeString,
			Description: "Display name",
			Optional:    true, // nullable
		},
		"series_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `AREA`, `COLUMN`, `LINE`",
			Optional:    true, // nullable
		},
		"theme_color": {
			Type:        schema.TypeString,
			Description: "Possible Values: `BLUE`, `DEFAULT`, `GRAY`, `GREEN`, `ORANGE`, `PURPLE`, `RED`, `ROYALBLUE`, `TURQUOISE`, `YELLOW`",
			Optional:    true, // nullable
		},
	}
}

func (me *MetricVisualizationOptions) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"color_override": me.ColorOverride,
		"display_name":   me.DisplayName,
		"series_type":    me.SeriesType,
		"theme_color":    me.ThemeColor,
	})
}

func (me *MetricVisualizationOptions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"color_override": &me.ColorOverride,
		"display_name":   &me.DisplayName,
		"series_type":    &me.SeriesType,
		"theme_color":    &me.ThemeColor,
	})
}
