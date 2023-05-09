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

package healthcards

import (
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Tiles []*Tile

func (me *Tiles) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"tile": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(Tile).Schema()},
		},
	}
}

func (me Tiles) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("tile", me)
}

func (me *Tiles) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("tile", me)
}

type Tile struct {
	Anchor             *Anchor             `json:"anchor,omitempty"`             // Define element to scroll to on click.
	Conditions         []string            `json:"conditions,omitempty"`         // All conditions from the list need to be fulfilled for the Tile to be visible
	DisplayMetricValue *bool               `json:"displayMetricValue,omitempty"` // When true, value of the selected metric will be displayed in the tile
	DisplayName        string              `json:"displayName"`                  // Display name
	FoldTransformation *FoldTransformation `json:"foldTransformation,omitempty"` // Possible Values: `AUTO`, `LAST_VALUE`, `MAX`, `MEDIAN`, `MIN`, `PERCENTILE_10`, `PERCENTILE_75`, `PERCENTILE_90`, `SUM`, `VALUE`
	MetricSelector     string              `json:"metricSelector"`               // A metric selector expression for a dedicated metric. For example, if you want to show the available disk space of a host, use  \n`builtin:host.disk.free`. \n\n It also supports the ability to split by dimension, aggregate the selected metric and a lot more, for example:  \n`builtin:host.disk.free:splitBy(dt.entity.disk):max`  \nThe filter for the ME id will be applied automatically.\n\nThe [Metric Browser](/ui/metrics \"Visit Metrics\") gives an overview of all the metrics an the available operations for every metric.
}

func (me *Tile) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"anchor": {
			Type:        schema.TypeList,
			Description: "Define element to scroll to on click.",
			Optional:    true, // nullable
			Elem:        &schema.Resource{Schema: new(Anchor).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"conditions": {
			Type:        schema.TypeSet,
			Description: "All conditions from the list need to be fulfilled for the Tile to be visible",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"display_metric_value": {
			Type:        schema.TypeBool,
			Description: "When true, value of the selected metric will be displayed in the tile",
			Optional:    true, // nullable
		},
		"display_name": {
			Type:        schema.TypeString,
			Description: "Display name",
			Required:    true,
		},
		"fold_transformation": {
			Type:        schema.TypeString,
			Description: "Possible Values: `AUTO`, `LAST_VALUE`, `MAX`, `MEDIAN`, `MIN`, `PERCENTILE_10`, `PERCENTILE_75`, `PERCENTILE_90`, `SUM`, `VALUE`",
			Optional:    true, // precondition
		},
		"metric_selector": {
			Type:        schema.TypeString,
			Description: "A metric selector expression for a dedicated metric. For example, if you want to show the available disk space of a host, use  \n`builtin:host.disk.free`. \n\n It also supports the ability to split by dimension, aggregate the selected metric and a lot more, for example:  \n`builtin:host.disk.free:splitBy(dt.entity.disk):max`  \nThe filter for the ME id will be applied automatically.\n\nThe [Metric Browser](/ui/metrics \"Visit Metrics\") gives an overview of all the metrics an the available operations for every metric.",
			Required:    true,
		},
	}
}

func (me *Tile) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"anchor":               me.Anchor,
		"conditions":           me.Conditions,
		"display_metric_value": me.DisplayMetricValue,
		"display_name":         me.DisplayName,
		"fold_transformation":  me.FoldTransformation,
		"metric_selector":      me.MetricSelector,
	})
}

func (me *Tile) HandlePreconditions() error {
	if me.FoldTransformation == nil && me.DisplayMetricValue != nil && *me.DisplayMetricValue {
		return fmt.Errorf("'fold_transformation' must be specified if 'display_metric_value' is set to '%v'", me.DisplayMetricValue)
	}
	if me.FoldTransformation != nil && me.DisplayMetricValue != nil && !*me.DisplayMetricValue {
		return fmt.Errorf("'fold_transformation' must not be specified if 'display_metric_value' is set to '%v'", me.DisplayMetricValue)
	}
	return nil
}

func (me *Tile) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"anchor":               &me.Anchor,
		"conditions":           &me.Conditions,
		"display_metric_value": &me.DisplayMetricValue,
		"display_name":         &me.DisplayName,
		"fold_transformation":  &me.FoldTransformation,
		"metric_selector":      &me.MetricSelector,
	})
}
