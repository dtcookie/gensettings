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

type PieChartConfig struct {
	ColorOverride      ColorOverridess  `json:"colorOverride,omitempty"`      // There is possibility to override color of single series, even single dimension of a metric
	DefaultAggregation *AggregationType `json:"defaultAggregation,omitempty"` // Possible Values: `AVG`, `COUNT`, `MAX`, `MEDIAN`, `MIN`, `SUM`, `VALUE`
	Metric             *MetricSelector  `json:"metric"`                       // Defines how to fetch data for the pie chart, please note that various fold transformations needs to be explicitly defined here in the selector
	ShowLegend         *bool            `json:"showLegend,omitempty"`         // When true, axis legend on the whole chart is visible
	ThemeColor         *ThemeColorEnum  `json:"themeColor,omitempty"`         // Possible Values: `BLUE`, `DEFAULT`, `GRAY`, `GREEN`, `ORANGE`, `PURPLE`, `RED`, `ROYALBLUE`, `TURQUOISE`, `YELLOW`
}

func (me *PieChartConfig) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"color_override": {
			Type:        schema.TypeList,
			Description: "There is possibility to override color of single series, even single dimension of a metric",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(ColorOverridess).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"default_aggregation": {
			Type:        schema.TypeString,
			Description: "Possible Values: `AVG`, `COUNT`, `MAX`, `MEDIAN`, `MIN`, `SUM`, `VALUE`",
			Optional:    true, // nullable
		},
		"metric": {
			Type:        schema.TypeList,
			Description: "Defines how to fetch data for the pie chart, please note that various fold transformations needs to be explicitly defined here in the selector",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(MetricSelector).Schema()},
			MinItems:    1,
			MaxItems:    1,
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

func (me *PieChartConfig) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"color_override":      me.ColorOverride,
		"default_aggregation": me.DefaultAggregation,
		"metric":              me.Metric,
		"show_legend":         me.ShowLegend,
		"theme_color":         me.ThemeColor,
	})
}

func (me *PieChartConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"color_override":      &me.ColorOverride,
		"default_aggregation": &me.DefaultAggregation,
		"metric":              &me.Metric,
		"show_legend":         &me.ShowLegend,
		"theme_color":         &me.ThemeColor,
	})
}
