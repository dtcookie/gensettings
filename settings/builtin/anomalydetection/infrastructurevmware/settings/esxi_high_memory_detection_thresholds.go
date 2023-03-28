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

// EsxiHighMemoryDetectionThresholds. Alert if the condition is met in 3 out of 5 samples
type EsxiHighMemoryDetectionThresholds struct {
	CompressionDecompressionRate float64 `json:"compressionDecompressionRate"` // ESXi host swap IN/OUT or compression/decompression rate is higher than
}

func (me *EsxiHighMemoryDetectionThresholds) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"compression_decompression_rate": {
			Type:        schema.TypeFloat,
			Description: "ESXi host swap IN/OUT or compression/decompression rate is higher than",
			Required:    true,
		},
	}
}

func (me *EsxiHighMemoryDetectionThresholds) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"compression_decompression_rate": me.CompressionDecompressionRate,
	})
}

func (me *EsxiHighMemoryDetectionThresholds) HandlePreconditions() {
}

func (me *EsxiHighMemoryDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"compression_decompression_rate": &me.CompressionDecompressionRate,
	})
}
