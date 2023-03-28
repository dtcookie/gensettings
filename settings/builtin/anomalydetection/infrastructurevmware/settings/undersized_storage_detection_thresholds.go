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

package infrastructurevmware

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// UndersizedStorageDetectionThresholds. Alert if **any** condition is met in 3 out of 5 samples
type UndersizedStorageDetectionThresholds struct {
	AverageQueueCommandLatency int `json:"averageQueueCommandLatency"` // Average queue command latency is higher than
	PeakQueueCommandLatency    int `json:"peakQueueCommandLatency"`    // Peak queue command latency is higher than
}

func (me *UndersizedStorageDetectionThresholds) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"average_queue_command_latency": {
			Type:        schema.TypeInt,
			Description: "Average queue command latency is higher than",
			Required:    true,
		},
		"peak_queue_command_latency": {
			Type:        schema.TypeInt,
			Description: "Peak queue command latency is higher than",
			Required:    true,
		},
	}
}

func (me *UndersizedStorageDetectionThresholds) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"average_queue_command_latency": me.AverageQueueCommandLatency,
		"peak_queue_command_latency":    me.PeakQueueCommandLatency,
	})
}

func (me *UndersizedStorageDetectionThresholds) HandlePreconditions() {
}

func (me *UndersizedStorageDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"average_queue_command_latency": &me.AverageQueueCommandLatency,
		"peak_queue_command_latency":    &me.PeakQueueCommandLatency,
	})
}
