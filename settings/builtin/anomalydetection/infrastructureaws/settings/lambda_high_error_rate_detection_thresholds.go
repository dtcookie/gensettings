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

package infrastructureaws

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// LambdaHighErrorRateDetectionThresholds. Alert if the condition is met in 3 out of 5 samples
type LambdaHighErrorRateDetectionThresholds struct {
	FailedInvocationsRate int `json:"failedInvocationsRate"` // Failed invocations rate is higher than
}

func (me *LambdaHighErrorRateDetectionThresholds) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"failed_invocations_rate": {
			Type:        schema.TypeInt,
			Description: "Failed invocations rate is higher than",
			Required:    true,
		},
	}
}

func (me *LambdaHighErrorRateDetectionThresholds) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"failed_invocations_rate": me.FailedInvocationsRate,
	})
}

func (me *LambdaHighErrorRateDetectionThresholds) HandlePreconditions() {
}

func (me *LambdaHighErrorRateDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"failed_invocations_rate": &me.FailedInvocationsRate,
	})
}
