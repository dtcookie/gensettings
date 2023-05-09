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

type GraphChartConfig struct {
	ConnectGaps        *bool                 `json:"connectGaps,omitempty"`        // When true, gaps in charts will be connected even if there is truly no data
	DefaultAggregation *AggregationType      `json:"defaultAggregation,omitempty"` // Possible Values: `AVG`, `COUNT`, `MAX`, `MEDIAN`, `MIN`, `SUM`, `VALUE`
	Metrics            MetricDefinitions     `json:"metrics,omitempty"`            // Defines how to fetch data for a single line in the chart
	Stacked            *bool                 `json:"stacked,omitempty"`            // When true, series on chart are stacked.
	Thresholds         Thresholds            `json:"thresholds,omitempty"`         // List of defined thresholds for the chart
	Visualization      *VisualizationOptions `json:"visualization,omitempty"`
	XAxis              *XAxis                `json:"xAxis,omitempty"` // You can control how X axis of your graph are displayed
	YAxes              YAxiss                `json:"yAxes,omitempty"` // You can control how each Y axis of your graph are displayed
}

func (me *GraphChartConfig) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"connect_gaps": {
			Type:        schema.TypeBool,
			Description: "When true, gaps in charts will be connected even if there is truly no data",
			Optional:    true, // nullable
		},
		"default_aggregation": {
			Type:        schema.TypeString,
			Description: "Possible Values: `AVG`, `COUNT`, `MAX`, `MEDIAN`, `MIN`, `SUM`, `VALUE`",
			Optional:    true, // nullable
		},
		"metrics": {
			Type:        schema.TypeList,
			Description: "Defines how to fetch data for a single line in the chart",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(MetricDefinitions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"stacked": {
			Type:        schema.TypeBool,
			Description: "When true, series on chart are stacked.",
			Optional:    true, // nullable
		},
		"thresholds": {
			Type:        schema.TypeList,
			Description: "List of defined thresholds for the chart",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(Thresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"visualization": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // nullable
			Elem:        &schema.Resource{Schema: new(VisualizationOptions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"x_axis": {
			Type:        schema.TypeList,
			Description: "You can control how X axis of your graph are displayed",
			Optional:    true, // nullable
			Elem:        &schema.Resource{Schema: new(XAxis).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"y_axes": {
			Type:        schema.TypeList,
			Description: "You can control how each Y axis of your graph are displayed",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(YAxiss).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *GraphChartConfig) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"connect_gaps":        me.ConnectGaps,
		"default_aggregation": me.DefaultAggregation,
		"metrics":             me.Metrics,
		"stacked":             me.Stacked,
		"thresholds":          me.Thresholds,
		"visualization":       me.Visualization,
		"x_axis":              me.XAxis,
		"y_axes":              me.YAxes,
	})
}

func (me *GraphChartConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"connect_gaps":        &me.ConnectGaps,
		"default_aggregation": &me.DefaultAggregation,
		"metrics":             &me.Metrics,
		"stacked":             &me.Stacked,
		"thresholds":          &me.Thresholds,
		"visualization":       &me.Visualization,
		"x_axis":              &me.XAxis,
		"y_axes":              &me.YAxes,
	})
}
