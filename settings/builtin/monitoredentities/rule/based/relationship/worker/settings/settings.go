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

package worker

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	MaxNewRelationshipsPerRule int `json:"maxNewRelationshipsPerRule"` // Maximum number of new relationships created for one entity per rule
	PartitionSize              int `json:"partitionSize"`              // Size of the partition, which defines how many entities are handled at once
	ResultLimit                int `json:"resultLimit"`                // Query limit for the number of entities which are considered
	RunInterval                int `json:"runInterval"`                // Time in between worker runs (in minutes)
	Timeframe                  int `json:"timeframe"`                  // Timeframe (in minutes) to look back to. Only entities which are last seen in this timeframe are considered
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"max_new_relationships_per_rule": {
			Type:        schema.TypeInt,
			Description: "Maximum number of new relationships created for one entity per rule",
			Required:    true,
		},
		"partition_size": {
			Type:        schema.TypeInt,
			Description: "Size of the partition, which defines how many entities are handled at once",
			Required:    true,
		},
		"result_limit": {
			Type:        schema.TypeInt,
			Description: "Query limit for the number of entities which are considered",
			Required:    true,
		},
		"run_interval": {
			Type:        schema.TypeInt,
			Description: "Time in between worker runs (in minutes)",
			Required:    true,
		},
		"timeframe": {
			Type:        schema.TypeInt,
			Description: "Timeframe (in minutes) to look back to. Only entities which are last seen in this timeframe are considered",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"max_new_relationships_per_rule": me.MaxNewRelationshipsPerRule,
		"partition_size":                 me.PartitionSize,
		"result_limit":                   me.ResultLimit,
		"run_interval":                   me.RunInterval,
		"timeframe":                      me.Timeframe,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"max_new_relationships_per_rule": &me.MaxNewRelationshipsPerRule,
		"partition_size":                 &me.PartitionSize,
		"result_limit":                   &me.ResultLimit,
		"run_interval":                   &me.RunInterval,
		"timeframe":                      &me.Timeframe,
	})
}
