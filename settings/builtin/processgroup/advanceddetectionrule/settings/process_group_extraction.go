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

package advanceddetectionrule

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ProcessGroupExtraction struct {
	Delimiter      *Delimiter `json:"delimiter"` // Optionally delimit this property between *From* and *To*.
	Property       string     `json:"property"`
	StandaloneRule bool       `json:"standaloneRule"` // If this option is selected, the default Dynatrace behavior is disabled for these detected processes. Only this rule is used to separate the process group.\n\nIf this option is not selected, this rule contributes to the default Dynatrace process group detection. \n\n[See our help page for examples.](https://dt-url.net/1722wrz)
}

func (me *ProcessGroupExtraction) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"delimiter": {
			Type:        schema.TypeList,
			Description: "Optionally delimit this property between *From* and *To*.",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(Delimiter).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"property": {
			Type:        schema.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
		"standalone_rule": {
			Type:        schema.TypeBool,
			Description: "If this option is selected, the default Dynatrace behavior is disabled for these detected processes. Only this rule is used to separate the process group.\n\nIf this option is not selected, this rule contributes to the default Dynatrace process group detection. \n\n[See our help page for examples.](https://dt-url.net/1722wrz)",
			Required:    true,
		},
	}
}

func (me *ProcessGroupExtraction) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"delimiter":       me.Delimiter,
		"property":        me.Property,
		"standalone_rule": me.StandaloneRule,
	})
}

func (me *ProcessGroupExtraction) HandlePreconditions() {
}

func (me *ProcessGroupExtraction) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"delimiter":       &me.Delimiter,
		"property":        &me.Property,
		"standalone_rule": &me.StandaloneRule,
	})
}
