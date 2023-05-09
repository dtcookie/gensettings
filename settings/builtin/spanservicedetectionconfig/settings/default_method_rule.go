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

type DefaultMethodRule struct {
	ID         *string         `json:"id,omitempty"`
	Name       *string         `json:"name,omitempty"`
	OutputType MethodRulesType `json:"outputType"` // Possible Values: `METHOD_ID_AND_NAME`, `NO_SERVICE`
}

func (me *DefaultMethodRule) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Description: "no documentation available",
			Optional:    true, // precondition
		},
		"name": {
			Type:        schema.TypeString,
			Description: "no documentation available",
			Optional:    true, // precondition
		},
		"output_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `METHOD_ID_AND_NAME`, `NO_SERVICE`",
			Required:    true,
		},
	}
}

func (me *DefaultMethodRule) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"id":          me.ID,
		"name":        me.Name,
		"output_type": me.OutputType,
	})
}

func (me *DefaultMethodRule) HandlePreconditions() error {
	if me.ID == nil && (string(me.OutputType) == "METHOD_ID_AND_NAME") {
		return fmt.Errorf("'id' must be specified if 'output_type' is set to '%v'", me.OutputType)
	}
	if me.Name == nil && (string(me.OutputType) == "METHOD_ID_AND_NAME") {
		return fmt.Errorf("'name' must be specified if 'output_type' is set to '%v'", me.OutputType)
	}
	return nil
}

func (me *DefaultMethodRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"id":          &me.ID,
		"name":        &me.Name,
		"output_type": &me.OutputType,
	})
}
