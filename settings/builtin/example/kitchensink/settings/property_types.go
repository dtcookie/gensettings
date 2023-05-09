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

package kitchensink

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type PropertyTypes struct {
	AnimalEnum      AnimalEnum `json:"AnimalEnum"`      // Possible Values: `Cat`, `Dog`, `Hamster`, `Snake`
	Boolean         bool       `json:"boolean"`         // Property type: `boolean`\n\nJava class: `java.lang.Boolean`\n\nDefault constraints: -. Property type: `boolean`\n\nJava class: `java.lang.Boolean`\n\nDefault constraints: -
	Float           float64    `json:"float"`           // Property type: `float`\n\nJava class: `java.lang.Double`\n\nDefault constraints: -. Property type: `float`\n\nJava class: `java.lang.Double`\n\nDefault constraints: -
	Integer         int        `json:"integer"`         // Property type: `integer`\n\nJava class: `java.lang.Long`\n\nDefault constraints: -. Property type: `integer`\n\nJava class: `java.lang.Long`\n\nDefault constraints: -
	Local_date_time string     `json:"local_date_time"` // Property type: `local_date_time`\n\nJava class: `java.time.LocalDateTime`\n\nDefault constraints: -. Property type: `local_date_time`\n\nJava class: `java.time.LocalDateTime`\n\nDefault constraints: -
	Secret          string     `json:"secret"`          // Property type: `secret`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_. Property type: `secret`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_
	Text            string     `json:"text"`            // Property type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_. Property type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_
	Zoned_date_time string     `json:"zoned_date_time"` // Property type: `zoned_date_time`\n\nJava class: `java.time.ZonedDateTime`\n\nDefault constraints: -. Property type: `zoned_date_time`\n\nJava class: `java.time.ZonedDateTime`\n\nDefault constraints: -
}

func (me *PropertyTypes) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"animal_enum": {
			Type:        schema.TypeString,
			Description: "Possible Values: `Cat`, `Dog`, `Hamster`, `Snake`",
			Required:    true,
		},
		"boolean": {
			Type:        schema.TypeBool,
			Description: "Property type: `boolean`\n\nJava class: `java.lang.Boolean`\n\nDefault constraints: -. Property type: `boolean`\n\nJava class: `java.lang.Boolean`\n\nDefault constraints: -",
			Required:    true,
		},
		"float": {
			Type:        schema.TypeFloat,
			Description: "Property type: `float`\n\nJava class: `java.lang.Double`\n\nDefault constraints: -. Property type: `float`\n\nJava class: `java.lang.Double`\n\nDefault constraints: -",
			Required:    true,
		},
		"integer": {
			Type:        schema.TypeInt,
			Description: "Property type: `integer`\n\nJava class: `java.lang.Long`\n\nDefault constraints: -. Property type: `integer`\n\nJava class: `java.lang.Long`\n\nDefault constraints: -",
			Required:    true,
		},
		"local_date_time": {
			Type:        schema.TypeString,
			Description: "Property type: `local_date_time`\n\nJava class: `java.time.LocalDateTime`\n\nDefault constraints: -. Property type: `local_date_time`\n\nJava class: `java.time.LocalDateTime`\n\nDefault constraints: -",
			Required:    true,
		},
		"secret": {
			Type:        schema.TypeString,
			Description: "Property type: `secret`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_. Property type: `secret`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_",
			Required:    true,
			Sensitive:   true,
		},
		"text": {
			Type:        schema.TypeString,
			Description: "Property type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_. Property type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_",
			Required:    true,
		},
		"zoned_date_time": {
			Type:        schema.TypeString,
			Description: "Property type: `zoned_date_time`\n\nJava class: `java.time.ZonedDateTime`\n\nDefault constraints: -. Property type: `zoned_date_time`\n\nJava class: `java.time.ZonedDateTime`\n\nDefault constraints: -",
			Required:    true,
		},
	}
}

func (me *PropertyTypes) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"animal_enum":     me.AnimalEnum,
		"boolean":         me.Boolean,
		"float":           me.Float,
		"integer":         me.Integer,
		"local_date_time": me.Local_date_time,
		"secret":          me.Secret,
		"text":            me.Text,
		"zoned_date_time": me.Zoned_date_time,
	})
}

func (me *PropertyTypes) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"animal_enum":     &me.AnimalEnum,
		"boolean":         &me.Boolean,
		"float":           &me.Float,
		"integer":         &me.Integer,
		"local_date_time": &me.Local_date_time,
		"secret":          &me.Secret,
		"text":            &me.Text,
		"zoned_date_time": &me.Zoned_date_time,
	})
}
