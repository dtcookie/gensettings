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

type SingleValueConfig struct {
	DefaultAggregation *AggregationType    `json:"defaultAggregation,omitempty"` // Possible Values: `AVG`, `COUNT`, `MAX`, `MEDIAN`, `MIN`, `SUM`, `VALUE`
	DisplayName        *string             `json:"displayName,omitempty"`        // Display name for value
	FoldTransformation *FoldTransformation `json:"foldTransformation,omitempty"` // Possible Values: `AUTO`, `LAST_VALUE`, `MAX`, `MEDIAN`, `MIN`, `PERCENTILE_10`, `PERCENTILE_75`, `PERCENTILE_90`, `SUM`, `VALUE`
	Metric             *MetricSelector     `json:"metric"`                       // Defines how to fetch data for the single value visualization, please note that it shows only the first enabled metric from the query and also various fold transformations needs to be explicitly defined here in the selector
	ShowSparkline      *bool               `json:"showSparkline,omitempty"`      // When true, the sparkline is visible
	ShowTrend          *bool               `json:"showTrend,omitempty"`          // When true, the trend is visible
	Thresholds         Thresholds          `json:"thresholds,omitempty"`         // List of defined thresholds for the chart
}

func (me *SingleValueConfig) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"default_aggregation": {
			Type:        schema.TypeString,
			Description: "Possible Values: `AVG`, `COUNT`, `MAX`, `MEDIAN`, `MIN`, `SUM`, `VALUE`",
			Optional:    true, // nullable
		},
		"display_name": {
			Type:        schema.TypeString,
			Description: "Display name for value",
			Optional:    true, // nullable
		},
		"fold_transformation": {
			Type:        schema.TypeString,
			Description: "Possible Values: `AUTO`, `LAST_VALUE`, `MAX`, `MEDIAN`, `MIN`, `PERCENTILE_10`, `PERCENTILE_75`, `PERCENTILE_90`, `SUM`, `VALUE`",
			Optional:    true, // nullable
		},
		"metric": {
			Type:        schema.TypeList,
			Description: "Defines how to fetch data for the single value visualization, please note that it shows only the first enabled metric from the query and also various fold transformations needs to be explicitly defined here in the selector",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(MetricSelector).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"show_sparkline": {
			Type:        schema.TypeBool,
			Description: "When true, the sparkline is visible",
			Optional:    true, // nullable
		},
		"show_trend": {
			Type:        schema.TypeBool,
			Description: "When true, the trend is visible",
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
	}
}

func (me *SingleValueConfig) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"default_aggregation": me.DefaultAggregation,
		"display_name":        me.DisplayName,
		"fold_transformation": me.FoldTransformation,
		"metric":              me.Metric,
		"show_sparkline":      me.ShowSparkline,
		"show_trend":          me.ShowTrend,
		"thresholds":          me.Thresholds,
	})
}

func (me *SingleValueConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"default_aggregation": &me.DefaultAggregation,
		"display_name":        &me.DisplayName,
		"fold_transformation": &me.FoldTransformation,
		"metric":              &me.Metric,
		"show_sparkline":      &me.ShowSparkline,
		"show_trend":          &me.ShowTrend,
		"thresholds":          &me.Thresholds,
	})
}
