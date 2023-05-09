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

type MetricDefinitions []*MetricDefinition

func (me *MetricDefinitions) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"metric": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(MetricDefinition).Schema()},
		},
	}
}

func (me MetricDefinitions) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("metric", me)
}

func (me *MetricDefinitions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("metric", me)
}

type MetricDefinition struct {
	MetricSelector         string                      `json:"metricSelector"`                   // A metric selector expression for a dedicated metric line. For example, if you want to show the available disk space of a host, use  \n`builtin:host.disk.free`. \n\n It also supports the ability to split by dimension, aggregate the selected metric and a lot more, for example:  \n`builtin:host.disk.free:splitBy(dt.entity.disk):max`  \nThe filter for the ME id will be applied automatically.\n\nThe [Metric Browser](/ui/metrics \"Visit Metrics\") gives an overview of all the metrics an the available operations for every metric.
	MetricSelectorDetailed *string                     `json:"metricSelectorDetailed,omitempty"` // A metric selector expression for a dedicated metric line. This selector is used to gather data for details section in table row.\n\n By default, that selector is created based on main one. e.g.\n for `builtin:host.disk.free` selector, per-page one will be\n `builtin:host.disk.free$(entityFilter)$(userFilter):splitBy(\"dt.entity.disk\")$(aggregation)`\n\n For other examples, examine `detailsSelector` field of metric object from screen definition generated on the entity screen.\n\n When providing custom selector, following placeholders needs to be included:\n - *$(entityFilter)* - it is replaced with `filter` operator that adds context of currently visible rows. It is recommended to use directly after each metric key. e.g. for disk list it will be evaluated to `:filter(in(\"dt.entity.disk\",entitySelector(type(DISK) AND fromRelationship.isDiskOf(type(HOST),entityId(HOST-1)))))`,\n - *$(userFilter)* - it is replaced with `filter` operator that corresponds to filters from filter bar above the table. It is recommended to use before last splitBy operator.\n - *$(aggregation)* - it is replaced with aggregation operator based on selection above the chart. Should be as close to the end of expression as possible.
	MetricSelectorSort     *string                     `json:"metricSelectorSort,omitempty"`     // A metric selector expression for a dedicated metric line. This selector is used to gather data for currently sorted column. It should point to the single value. It queries data for whole table for given metric.\n\n By default, that selector is created based on main one. e.g.\n for `builtin:host.disk.free` selector, per-page one will be\n `builtin:host.disk.free$(entityFilter)$(userFilter):splitBy(\"dt.entity.disk\"):names$(aggregation):last:sort(value(avg,$(sortOrder))):limit(X)`\n\n For other examples, examine `sortSelector` field of metric object from screen definition generated on the entity screen.\n\n When providing custom selector, following placeholders needs to be included:\n - *$(entityFilter)* - it is replaced with `filter` operator that adds context of currently visible rows. It is recommended to use directly after each metric key. e.g. for disk list it will be evaluated to `:filter(in(\"dt.entity.disk\",entitySelector(type(DISK) AND fromRelationship.isDiskOf(type(HOST),entityId(HOST-1)))))`,\n - *$(userFilter)* - it is replaced with `filter` operator that corresponds to filters from filter bar above the table. It is recommended to use before last splitBy operator.\n - *$(aggregation)* - it is replaced with aggregation operator based on selection above the chart. Should be as close to the end of expression as possible.\n - *$(sortOrder)* - it is replaced with sort order constant from metric selector `descending` or `ascending`. \n\n Do not include limit operator. It will be appended automatically according to card configuration.
	Visualization          *MetricVisualizationOptions `json:"visualization,omitempty"`
	YAxisKey               *string                     `json:"yAxisKey,omitempty"` // Select Y-axis to be matched with this metric by its key
}

func (me *MetricDefinition) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"metric_selector": {
			Type:        schema.TypeString,
			Description: "A metric selector expression for a dedicated metric line. For example, if you want to show the available disk space of a host, use  \n`builtin:host.disk.free`. \n\n It also supports the ability to split by dimension, aggregate the selected metric and a lot more, for example:  \n`builtin:host.disk.free:splitBy(dt.entity.disk):max`  \nThe filter for the ME id will be applied automatically.\n\nThe [Metric Browser](/ui/metrics \"Visit Metrics\") gives an overview of all the metrics an the available operations for every metric.",
			Required:    true,
		},
		"metric_selector_detailed": {
			Type:        schema.TypeString,
			Description: "A metric selector expression for a dedicated metric line. This selector is used to gather data for details section in table row.\n\n By default, that selector is created based on main one. e.g.\n for `builtin:host.disk.free` selector, per-page one will be\n `builtin:host.disk.free$(entityFilter)$(userFilter):splitBy(\"dt.entity.disk\")$(aggregation)`\n\n For other examples, examine `detailsSelector` field of metric object from screen definition generated on the entity screen.\n\n When providing custom selector, following placeholders needs to be included:\n - *$(entityFilter)* - it is replaced with `filter` operator that adds context of currently visible rows. It is recommended to use directly after each metric key. e.g. for disk list it will be evaluated to `:filter(in(\"dt.entity.disk\",entitySelector(type(DISK) AND fromRelationship.isDiskOf(type(HOST),entityId(HOST-1)))))`,\n - *$(userFilter)* - it is replaced with `filter` operator that corresponds to filters from filter bar above the table. It is recommended to use before last splitBy operator.\n - *$(aggregation)* - it is replaced with aggregation operator based on selection above the chart. Should be as close to the end of expression as possible.",
			Optional:    true, // nullable
		},
		"metric_selector_sort": {
			Type:        schema.TypeString,
			Description: "A metric selector expression for a dedicated metric line. This selector is used to gather data for currently sorted column. It should point to the single value. It queries data for whole table for given metric.\n\n By default, that selector is created based on main one. e.g.\n for `builtin:host.disk.free` selector, per-page one will be\n `builtin:host.disk.free$(entityFilter)$(userFilter):splitBy(\"dt.entity.disk\"):names$(aggregation):last:sort(value(avg,$(sortOrder))):limit(X)`\n\n For other examples, examine `sortSelector` field of metric object from screen definition generated on the entity screen.\n\n When providing custom selector, following placeholders needs to be included:\n - *$(entityFilter)* - it is replaced with `filter` operator that adds context of currently visible rows. It is recommended to use directly after each metric key. e.g. for disk list it will be evaluated to `:filter(in(\"dt.entity.disk\",entitySelector(type(DISK) AND fromRelationship.isDiskOf(type(HOST),entityId(HOST-1)))))`,\n - *$(userFilter)* - it is replaced with `filter` operator that corresponds to filters from filter bar above the table. It is recommended to use before last splitBy operator.\n - *$(aggregation)* - it is replaced with aggregation operator based on selection above the chart. Should be as close to the end of expression as possible.\n - *$(sortOrder)* - it is replaced with sort order constant from metric selector `descending` or `ascending`. \n\n Do not include limit operator. It will be appended automatically according to card configuration.",
			Optional:    true, // nullable
		},
		"visualization": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true, // nullable
			Elem:        &schema.Resource{Schema: new(MetricVisualizationOptions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"y_axis_key": {
			Type:        schema.TypeString,
			Description: "Select Y-axis to be matched with this metric by its key",
			Optional:    true, // nullable
		},
	}
}

func (me *MetricDefinition) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"metric_selector":          me.MetricSelector,
		"metric_selector_detailed": me.MetricSelectorDetailed,
		"metric_selector_sort":     me.MetricSelectorSort,
		"visualization":            me.Visualization,
		"y_axis_key":               me.YAxisKey,
	})
}

func (me *MetricDefinition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"metric_selector":          &me.MetricSelector,
		"metric_selector_detailed": &me.MetricSelectorDetailed,
		"metric_selector_sort":     &me.MetricSelectorSort,
		"visualization":            &me.Visualization,
		"y_axis_key":               &me.YAxisKey,
	})
}
