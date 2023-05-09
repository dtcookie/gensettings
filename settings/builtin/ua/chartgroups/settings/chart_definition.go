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

package chartgroups

import (
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ChartDefinitions []*ChartDefinition

func (me *ChartDefinitions) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"chart": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(ChartDefinition).Schema()},
		},
	}
}

func (me ChartDefinitions) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("chart", me)
}

func (me *ChartDefinitions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("chart", me)
}

type ChartDefinition struct {
	Conditions        []string               `json:"conditions,omitempty"`     // All conditions from the list need to be fulfilled for the chart to be visible
	DetailsMessage    *string                `json:"detailsMessage,omitempty"` // The message displayed when chart-group choice of charts and aggregation is fixed
	DisplayName       string                 `json:"displayName"`              // Display name
	GraphChartConfig  *GraphChartConfig      `json:"graphChartConfig,omitempty"`
	PieChartConfig    *PieChartConfig        `json:"pieChartConfig,omitempty"`
	SingleValueConfig *SingleValueConfig     `json:"singleValueConfig,omitempty"`
	VisualizationType ChartVisualizationType `json:"visualizationType"` // Possible Values: `GRAPH_CHART`, `PIE_CHART`, `SINGLE_VALUE`
}

func (me *ChartDefinition) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"conditions": {
			Type:        schema.TypeSet,
			Description: "All conditions from the list need to be fulfilled for the chart to be visible",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"details_message": {
			Type:        schema.TypeString,
			Description: "The message displayed when chart-group choice of charts and aggregation is fixed",
			Optional:    true, // nullable
		},
		"display_name": {
			Type:        schema.TypeString,
			Description: "Display name",
			Required:    true,
		},
		"graph_chart_config": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(GraphChartConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"pie_chart_config": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(PieChartConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"single_value_config": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // precondition
			Elem:        &schema.Resource{Schema: new(SingleValueConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"visualization_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `GRAPH_CHART`, `PIE_CHART`, `SINGLE_VALUE`",
			Required:    true,
		},
	}
}

func (me *ChartDefinition) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"conditions":          me.Conditions,
		"details_message":     me.DetailsMessage,
		"display_name":        me.DisplayName,
		"graph_chart_config":  me.GraphChartConfig,
		"pie_chart_config":    me.PieChartConfig,
		"single_value_config": me.SingleValueConfig,
		"visualization_type":  me.VisualizationType,
	})
}

func (me *ChartDefinition) HandlePreconditions() error {
	if me.GraphChartConfig == nil && (string(me.VisualizationType) == "GRAPH_CHART") {
		return fmt.Errorf("'graph_chart_config' must be specified if 'visualization_type' is set to '%v'", me.VisualizationType)
	}
	if me.GraphChartConfig != nil && (string(me.VisualizationType) != "GRAPH_CHART") {
		return fmt.Errorf("'graph_chart_config' must not be specified if 'visualization_type' is set to '%v'", me.VisualizationType)
	}
	if me.PieChartConfig == nil && (string(me.VisualizationType) == "PIE_CHART") {
		return fmt.Errorf("'pie_chart_config' must be specified if 'visualization_type' is set to '%v'", me.VisualizationType)
	}
	if me.PieChartConfig != nil && (string(me.VisualizationType) != "PIE_CHART") {
		return fmt.Errorf("'pie_chart_config' must not be specified if 'visualization_type' is set to '%v'", me.VisualizationType)
	}
	if me.SingleValueConfig == nil && (string(me.VisualizationType) == "SINGLE_VALUE") {
		return fmt.Errorf("'single_value_config' must be specified if 'visualization_type' is set to '%v'", me.VisualizationType)
	}
	if me.SingleValueConfig != nil && (string(me.VisualizationType) != "SINGLE_VALUE") {
		return fmt.Errorf("'single_value_config' must not be specified if 'visualization_type' is set to '%v'", me.VisualizationType)
	}
	return nil
}

func (me *ChartDefinition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"conditions":          &me.Conditions,
		"details_message":     &me.DetailsMessage,
		"display_name":        &me.DisplayName,
		"graph_chart_config":  &me.GraphChartConfig,
		"pie_chart_config":    &me.PieChartConfig,
		"single_value_config": &me.SingleValueConfig,
		"visualization_type":  &me.VisualizationType,
	})
}
