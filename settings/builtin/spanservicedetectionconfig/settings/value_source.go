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

type ValueSource struct {
	AttributeKey         *string         `json:"attributeKey,omitempty"`         // Span attribute key
	ResourceAttributeKey *string         `json:"resourceAttributeKey,omitempty"` // Resource attribute key
	Source               ValueSourceType `json:"source"`                         // Possible Values: `INSTRUMENTATION_SCOPE_NAME`, `INSTRUMENTATION_SCOPE_VERSION`, `RESOURCE_ATTRIBUTE`, `SPAN_ATTRIBUTE`, `SPAN_NAME`
}

func (me *ValueSource) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attribute_key": {
			Type:        schema.TypeString,
			Description: "Span attribute key",
			Optional:    true, // precondition
		},
		"resource_attribute_key": {
			Type:        schema.TypeString,
			Description: "Resource attribute key",
			Optional:    true, // precondition
		},
		"source": {
			Type:        schema.TypeString,
			Description: "Possible Values: `INSTRUMENTATION_SCOPE_NAME`, `INSTRUMENTATION_SCOPE_VERSION`, `RESOURCE_ATTRIBUTE`, `SPAN_ATTRIBUTE`, `SPAN_NAME`",
			Required:    true,
		},
	}
}

func (me *ValueSource) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"attribute_key":          me.AttributeKey,
		"resource_attribute_key": me.ResourceAttributeKey,
		"source":                 me.Source,
	})
}

func (me *ValueSource) HandlePreconditions() error {
	if me.AttributeKey == nil && (string(me.Source) == "SPAN_ATTRIBUTE") {
		return fmt.Errorf("'attribute_key' must be specified if 'source' is set to '%v'", me.Source)
	}
	if me.ResourceAttributeKey == nil && (string(me.Source) == "RESOURCE_ATTRIBUTE") {
		return fmt.Errorf("'resource_attribute_key' must be specified if 'source' is set to '%v'", me.Source)
	}
	return nil
}

func (me *ValueSource) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"attribute_key":          &me.AttributeKey,
		"resource_attribute_key": &me.ResourceAttributeKey,
		"source":                 &me.Source,
	})
}
