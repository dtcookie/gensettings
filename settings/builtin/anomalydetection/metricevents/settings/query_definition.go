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

package metricevents

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type QueryDefinition struct {
	Aggregation     *Aggregation     `json:"aggregation,omitempty"`     // Possible Values: `AVG`, `COUNT`, `MAX`, `MEDIAN`, `MIN`, `PERCENTILE90`, `SUM`, `VALUE`
	DimensionFilter DimensionFilters `json:"dimensionFilter,omitempty"` // Dimension filter
	EntityFilter    *EntityFilter    `json:"entityFilter,omitempty"`    // Use rule-based filters to define the scope this event monitors.
	ManagementZone  *string          `json:"managementZone,omitempty"`  // Management zone
	MetricKey       *string          `json:"metricKey,omitempty"`       // Metric key
	MetricSelector  *string          `json:"metricSelector,omitempty"`  // To learn more, visit [Metric Selector](https://dt-url.net/metselad)
	QueryOffset     *int             `json:"queryOffset,omitempty"`     // Minute offset of sliding evaluation window for metrics with latency
	Type            Type             `json:"type"`                      // Possible Values: `METRIC_KEY`, `METRIC_SELECTOR`
}

func (me *QueryDefinition) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"aggregation": {
			Type:        schema.TypeString,
			Description: "Possible Values: `AVG`, `COUNT`, `MAX`, `MEDIAN`, `MIN`, `PERCENTILE90`, `SUM`, `VALUE`",
			Optional:    true,
		},
		"dimension_filter": {
			Type:        schema.TypeList,
			Description: "Dimension filter",
			Optional:    true,

			Elem:     &schema.Resource{Schema: new(DimensionFilters).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"entity_filter": {
			Type:        schema.TypeList,
			Description: "Use rule-based filters to define the scope this event monitors.",
			Optional:    true,

			Elem:     &schema.Resource{Schema: new(EntityFilter).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"management_zone": {
			Type:        schema.TypeString,
			Description: "Management zone",
			Optional:    true,
		},
		"metric_key": {
			Type:        schema.TypeString,
			Description: "Metric key",
			Optional:    true,
		},
		"metric_selector": {
			Type:        schema.TypeString,
			Description: "To learn more, visit [Metric Selector](https://dt-url.net/metselad)",
			Optional:    true,
		},
		"query_offset": {
			Type:        schema.TypeInt,
			Description: "Minute offset of sliding evaluation window for metrics with latency",
			Optional:    true,
		},
		"type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `METRIC_KEY`, `METRIC_SELECTOR`",
			Required:    true,
		},
	}
}

func (me *QueryDefinition) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"aggregation":      me.Aggregation,
		"dimension_filter": me.DimensionFilter,
		"entity_filter":    me.EntityFilter,
		"management_zone":  me.ManagementZone,
		"metric_key":       me.MetricKey,
		"metric_selector":  me.MetricSelector,
		"query_offset":     me.QueryOffset,
		"type":             me.Type,
	})
}

func (me *QueryDefinition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"aggregation":      &me.Aggregation,
		"dimension_filter": &me.DimensionFilter,
		"entity_filter":    &me.EntityFilter,
		"management_zone":  &me.ManagementZone,
		"metric_key":       &me.MetricKey,
		"metric_selector":  &me.MetricSelector,
		"query_offset":     &me.QueryOffset,
		"type":             &me.Type,
	})
}
