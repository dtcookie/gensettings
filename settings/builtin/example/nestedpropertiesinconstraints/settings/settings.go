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

package nestedpropertiesinconstraints

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	CollectionWithConstraint Type4s `json:"collectionWithConstraint,omitempty"` // List with unique properties
	ComplexInteger           *Type2 `json:"complexInteger"`                     // Nested integer value used in a container constraint
	IntegerValue             int    `json:"integerValue"`                       // Integer value used in a container constraint
	SchemaUniqueName         string `json:"schemaUniqueName"`                   // Name, unique on the schema level in a combination with `Nested name` defined in the inner type
	SchemaUniqueNameComplex  *Type1 `json:"schemaUniqueNameComplex"`            // Nested unique name
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"collection_with_constraint": {
			Type:        schema.TypeList,
			Description: "List with unique properties",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(Type4s).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"complex_integer": {
			Type:        schema.TypeList,
			Description: "Nested integer value used in a container constraint",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(Type2).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"integer_value": {
			Type:        schema.TypeInt,
			Description: "Integer value used in a container constraint",
			Required:    true,
		},
		"schema_unique_name": {
			Type:        schema.TypeString,
			Description: "Name, unique on the schema level in a combination with `Nested name` defined in the inner type",
			Required:    true,
		},
		"schema_unique_name_complex": {
			Type:        schema.TypeList,
			Description: "Nested unique name",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(Type1).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"collection_with_constraint": me.CollectionWithConstraint,
		"complex_integer":            me.ComplexInteger,
		"integer_value":              me.IntegerValue,
		"schema_unique_name":         me.SchemaUniqueName,
		"schema_unique_name_complex": me.SchemaUniqueNameComplex,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"collection_with_constraint": &me.CollectionWithConstraint,
		"complex_integer":            &me.ComplexInteger,
		"integer_value":              &me.IntegerValue,
		"schema_unique_name":         &me.SchemaUniqueName,
		"schema_unique_name_complex": &me.SchemaUniqueNameComplex,
	})
}
