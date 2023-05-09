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

type InitialSortingType struct {
	ColumnID string    `json:"columnId"` // Column ID is a unique identifier of column. There are several ID types for different column types: \n - Entity column - 'name'\n\n - Metric with selectable aggregation - 'metric-{metricDisplayName}-{aggregation}'\n Example: 'metric-CPU Usage-Average'\n\n - metric without selectable aggregation - 'metric-{metricDisplayName}'            \n Example: 'metric-CPU Usage'\n\n - Attribute - 'attribute-{attributeKey}'                                               \n Example: 'attribute-osType'\n\n - Relation - 'attribute-{relationEntityType}-{relationDisplayName}'                     \n Example: 'attribute-CLOUD_APPLICATION-Workloads'\n\n - Metric dimension (Metric Table only) - 'metricDimension-{dimensionKey}'                     \n Example: 'metricDimension-dt.entity.host'
	Order    OrderType `json:"order"`    // Possible Values: `ASCENDING`, `DESCENDING`
}

func (me *InitialSortingType) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"column_id": {
			Type:        schema.TypeString,
			Description: "Column ID is a unique identifier of column. There are several ID types for different column types: \n - Entity column - 'name'\n\n - Metric with selectable aggregation - 'metric-{metricDisplayName}-{aggregation}'\n Example: 'metric-CPU Usage-Average'\n\n - metric without selectable aggregation - 'metric-{metricDisplayName}'            \n Example: 'metric-CPU Usage'\n\n - Attribute - 'attribute-{attributeKey}'                                               \n Example: 'attribute-osType'\n\n - Relation - 'attribute-{relationEntityType}-{relationDisplayName}'                     \n Example: 'attribute-CLOUD_APPLICATION-Workloads'\n\n - Metric dimension (Metric Table only) - 'metricDimension-{dimensionKey}'                     \n Example: 'metricDimension-dt.entity.host'",
			Required:    true,
		},
		"order": {
			Type:        schema.TypeString,
			Description: "Possible Values: `ASCENDING`, `DESCENDING`",
			Required:    true,
		},
	}
}

func (me *InitialSortingType) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"column_id": me.ColumnID,
		"order":     me.Order,
	})
}

func (me *InitialSortingType) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"column_id": &me.ColumnID,
		"order":     &me.Order,
	})
}
