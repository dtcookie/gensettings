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

package fullwebservice

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ContextRoot struct {
	ContributionType ContributionType        `json:"contributionType"`       // Possible Values: `OriginalValue`, `TransformValue`, `TransformURL`
	SegmentCount     *int                    `json:"segmentCount,omitempty"` // The number of segments of the URL to be kept.\nThe URL is divided by slashes (/), the indexing starts with 1 at context root.\nFor example, if you specify 2 for the www.dynatrace.com/support/help/dynatrace-api/ URL, the value of support/help is used.
	Transformations  *ReducedTransformations `json:"transformations,omitempty"`
}

func (me *ContextRoot) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"contribution_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `OriginalValue`, `TransformValue`, `TransformURL`",
			Required:    true,
		},
		"segment_count": {
			Type:        schema.TypeInt,
			Description: "The number of segments of the URL to be kept.\nThe URL is divided by slashes (/), the indexing starts with 1 at context root.\nFor example, if you specify 2 for the www.dynatrace.com/support/help/dynatrace-api/ URL, the value of support/help is used.",
			Optional:    true,
		},
		"transformations": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true,

			Elem:     &schema.Resource{Schema: new(ReducedTransformations).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
	}
}

func (me *ContextRoot) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"contribution_type": me.ContributionType,
		"segment_count":     me.SegmentCount,
		"transformations":   me.Transformations,
	})
}

func (me *ContextRoot) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"contribution_type": &me.ContributionType,
		"segment_count":     &me.SegmentCount,
		"transformations":   &me.Transformations,
	})
}
