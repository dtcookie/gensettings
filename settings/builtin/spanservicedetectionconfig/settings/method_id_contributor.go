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

package spanservicedetectionconfig

import (
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type MethodIdContributors []*MethodIdContributor

func (me *MethodIdContributors) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"method_id_contributor": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(MethodIdContributor).Schema()},
		},
	}
}

func (me MethodIdContributors) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("method_id_contributor", me)
}

func (me *MethodIdContributors) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("method_id_contributor", me)
}

type MethodIdContributor struct {
	ExtractionValue      *string          `json:"extractionValue,omitempty"`      // Extraction id
	ResourceAttributeKey *string          `json:"resourceAttributeKey,omitempty"` // Key
	Source               KnownMethodInput `json:"source"`                         // Possible Values: `EXTRACTION_VALUE`, `INSTRUMENTATION_SCOPE_NAME`, `INSTRUMENTATION_SCOPE_VERSION`, `RESOURCE_ATTRIBUTE`, `SPAN_ATTRIBUTE`, `SPAN_KIND`, `SPAN_NAME`
	SourceKey            *string          `json:"sourceKey,omitempty"`            // Key
}

func (me *MethodIdContributor) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"extraction_value": {
			Type:        schema.TypeString,
			Description: "Extraction id",
			Optional:    true, // precondition
		},
		"resource_attribute_key": {
			Type:        schema.TypeString,
			Description: "Key",
			Optional:    true, // precondition
		},
		"source": {
			Type:        schema.TypeString,
			Description: "Possible Values: `EXTRACTION_VALUE`, `INSTRUMENTATION_SCOPE_NAME`, `INSTRUMENTATION_SCOPE_VERSION`, `RESOURCE_ATTRIBUTE`, `SPAN_ATTRIBUTE`, `SPAN_KIND`, `SPAN_NAME`",
			Required:    true,
		},
		"source_key": {
			Type:        schema.TypeString,
			Description: "Key",
			Optional:    true, // precondition
		},
	}
}

func (me *MethodIdContributor) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"extraction_value":       me.ExtractionValue,
		"resource_attribute_key": me.ResourceAttributeKey,
		"source":                 me.Source,
		"source_key":             me.SourceKey,
	})
}

func (me *MethodIdContributor) HandlePreconditions() error {
	if me.ExtractionValue == nil && (string(me.Source) == "EXTRACTION_VALUE") {
		return fmt.Errorf("'extraction_value' must be specified if 'source' is set to '%v'", me.Source)
	}
	if me.ResourceAttributeKey == nil && (string(me.Source) == "RESOURCE_ATTRIBUTE") {
		return fmt.Errorf("'resource_attribute_key' must be specified if 'source' is set to '%v'", me.Source)
	}
	if me.SourceKey == nil && (string(me.Source) == "SPAN_ATTRIBUTE") {
		return fmt.Errorf("'source_key' must be specified if 'source' is set to '%v'", me.Source)
	}
	return nil
}

func (me *MethodIdContributor) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"extraction_value":       &me.ExtractionValue,
		"resource_attribute_key": &me.ResourceAttributeKey,
		"source":                 &me.Source,
		"source_key":             &me.SourceKey,
	})
}
