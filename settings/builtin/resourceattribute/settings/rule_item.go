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

package resourceattribute

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type RuleItems []*RuleItem

func (me *RuleItems) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attributeKey": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(RuleItem).Schema()},
		},
	}
}

func (me RuleItems) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("attributeKey", me)
}

func (me *RuleItems) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("attributeKey", me)
}

type RuleItem struct {
	AttributeKey string      `json:"attributeKey"` // Attribute key **service.name** is automatically captured by default
	Enabled      bool        `json:"enabled"`      // If this is true, the value of the specified key is stored.
	Masking      MaskingType `json:"masking"`      // If this attribute contains confidential data, turn on masking to conceal its value from users. Introduce more granular control over the visibility of attribute values.  \nChoose **Do not mask** to allow every user to see the actual value and use it in defining other configurations.  \nChoose **Mask entire value** to hide the whole value of this attribute from everyone who does not have 'View sensitive request data' permission. These attributes can't be used to define other configurations. \nChoose **Mask only confidential data** to apply automatic masking strategies to your data. These strategies include, for example, credit card numbers, IBAN, IPs, email-addresses, etc. It may not be possible to recognize all sensitive data so please always make sure to verify that sensitive data is actually masked. If sensitive data is not recognized, please use **Mask entire value** instead. Users with 'View sensitive request data' permission will be able to see the entire value, others only the non-sensitive parts. These attributes can't be used to define other configurations.
}

func (me *RuleItem) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attribute_key": {
			Type:        schema.TypeString,
			Description: "Attribute key **service.name** is automatically captured by default",
			Required:    true,
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "If this is true, the value of the specified key is stored.",
			Required:    true,
		},
		"masking": {
			Type:        schema.TypeString,
			Description: "If this attribute contains confidential data, turn on masking to conceal its value from users. Introduce more granular control over the visibility of attribute values.  \nChoose **Do not mask** to allow every user to see the actual value and use it in defining other configurations.  \nChoose **Mask entire value** to hide the whole value of this attribute from everyone who does not have 'View sensitive request data' permission. These attributes can't be used to define other configurations. \nChoose **Mask only confidential data** to apply automatic masking strategies to your data. These strategies include, for example, credit card numbers, IBAN, IPs, email-addresses, etc. It may not be possible to recognize all sensitive data so please always make sure to verify that sensitive data is actually masked. If sensitive data is not recognized, please use **Mask entire value** instead. Users with 'View sensitive request data' permission will be able to see the entire value, others only the non-sensitive parts. These attributes can't be used to define other configurations.",
			Required:    true,
		},
	}
}

func (me *RuleItem) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"attribute_key": me.AttributeKey,
		"enabled":       me.Enabled,
		"masking":       me.Masking,
	})
}

func (me *RuleItem) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"attribute_key": &me.AttributeKey,
		"enabled":       &me.Enabled,
		"masking":       &me.Masking,
	})
}
