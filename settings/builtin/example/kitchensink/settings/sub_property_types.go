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

type SubPropertyTypes struct {
	Credentials         string  `json:"credentials"`         // Sub property type: `credential`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints:\n* `PATTERN`: _Must match pattern ^CREDENTIALS_VAULT-.+$_\n* `PATTERN`: _Please use a format of TYPE-ID_
	CssSelector         string  `json:"cssSelector"`         // Sub property type: `cssSelector`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints: -
	Datasource_property string  `json:"datasource_property"` // datasource_property sub property type
	Email               string  `json:"email"`               // Sub property type: `email`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints:\n* `PATTERN`: _Email address is not valid_
	Float_percentage    float64 `json:"float_percentage"`    // Sub property type: `percentage`\n\nBase type: `float`\n\nJava class: `java.lang.Double`\n\nDefault constraints: -\n\nAdditional constraints: -
	Host_entity         string  `json:"host_entity"`         // Sub property type: `me`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints:\n* `PATTERN`: _Please use a format of TYPE-ID_\n* `PATTERN`: _Must match pattern ^HOST-.+$_
	Http_url            string  `json:"http_url"`            // Sub property type: `url`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints:\n* `PATTERN`: _Not a valid URI_\n* `PATTERN`: _Not a valid http(s) URL_
	Integer_percentage  int     `json:"integer_percentage"`  // Sub property type: `percentage`\n\nBase type: `integer`\n\nJava class: `java.lang.Long`\n\nDefault constraints: -\n\nAdditional constraints: -
	JsonCode            string  `json:"jsonCode"`            // Sub property type: `code/json`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints:\n* `CUSTOM_VALIDATOR_REF`: _Custom validation failed_
	Management_zone     string  `json:"management_zone"`     // management_zone sub property type
	Metric              string  `json:"metric"`              // Sub property type: `metric`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints:\n* `CUSTOM_VALIDATOR_REF`: _Custom validation failed_
	Monitored_entity    string  `json:"monitored_entity"`    // Sub property type: `me`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints:\n* `PATTERN`: _Please use a format of TYPE-ID_
	Multiline_text      string  `json:"multiline_text"`      // Sub property type: `multiline`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints: -
	Regex               string  `json:"regex"`               // Sub property type: `regex`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints:\n* `REGEX`: _Must be a valid regular expression_
	Tag                 string  `json:"tag"`                 // Sub property type: `tag`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints:\n* `CUSTOM_VALIDATOR_REF`: _Custom validation failed_
	Uri                 string  `json:"uri"`                 // Sub property type: `uri`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints:\n* `PATTERN`: _Not a valid URI_
}

func (me *SubPropertyTypes) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"credentials": {
			Type:        schema.TypeString,
			Description: "Sub property type: `credential`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints:\n* `PATTERN`: _Must match pattern ^CREDENTIALS_VAULT-.+$_\n* `PATTERN`: _Please use a format of TYPE-ID_",
			Required:    true,
		},
		"css_selector": {
			Type:        schema.TypeString,
			Description: "Sub property type: `cssSelector`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints: -",
			Required:    true,
		},
		"datasource_property": {
			Type:        schema.TypeString,
			Description: "datasource_property sub property type",
			Required:    true,
		},
		"email": {
			Type:        schema.TypeString,
			Description: "Sub property type: `email`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints:\n* `PATTERN`: _Email address is not valid_",
			Required:    true,
		},
		"float_percentage": {
			Type:        schema.TypeFloat,
			Description: "Sub property type: `percentage`\n\nBase type: `float`\n\nJava class: `java.lang.Double`\n\nDefault constraints: -\n\nAdditional constraints: -",
			Required:    true,
		},
		"host_entity": {
			Type:        schema.TypeString,
			Description: "Sub property type: `me`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints:\n* `PATTERN`: _Please use a format of TYPE-ID_\n* `PATTERN`: _Must match pattern ^HOST-.+$_",
			Required:    true,
		},
		"http_url": {
			Type:        schema.TypeString,
			Description: "Sub property type: `url`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints:\n* `PATTERN`: _Not a valid URI_\n* `PATTERN`: _Not a valid http(s) URL_",
			Required:    true,
		},
		"integer_percentage": {
			Type:        schema.TypeInt,
			Description: "Sub property type: `percentage`\n\nBase type: `integer`\n\nJava class: `java.lang.Long`\n\nDefault constraints: -\n\nAdditional constraints: -",
			Required:    true,
		},
		"json_code": {
			Type:        schema.TypeString,
			Description: "Sub property type: `code/json`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints:\n* `CUSTOM_VALIDATOR_REF`: _Custom validation failed_",
			Required:    true,
		},
		"management_zone": {
			Type:        schema.TypeString,
			Description: "management_zone sub property type",
			Required:    true,
		},
		"metric": {
			Type:        schema.TypeString,
			Description: "Sub property type: `metric`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints:\n* `CUSTOM_VALIDATOR_REF`: _Custom validation failed_",
			Required:    true,
		},
		"monitored_entity": {
			Type:        schema.TypeString,
			Description: "Sub property type: `me`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints:\n* `PATTERN`: _Please use a format of TYPE-ID_",
			Required:    true,
		},
		"multiline_text": {
			Type:        schema.TypeString,
			Description: "Sub property type: `multiline`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints: -",
			Required:    true,
		},
		"regex": {
			Type:        schema.TypeString,
			Description: "Sub property type: `regex`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints:\n* `REGEX`: _Must be a valid regular expression_",
			Required:    true,
		},
		"tag": {
			Type:        schema.TypeString,
			Description: "Sub property type: `tag`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints:\n* `CUSTOM_VALIDATOR_REF`: _Custom validation failed_",
			Required:    true,
		},
		"uri": {
			Type:        schema.TypeString,
			Description: "Sub property type: `uri`\n\nBase type: `text`\n\nJava class: `java.lang.String`\n\nDefault constraints:\n* `LENGTH`: _Size must be between 1 and 500_\n\nAdditional constraints:\n* `PATTERN`: _Not a valid URI_",
			Required:    true,
		},
	}
}

func (me *SubPropertyTypes) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"credentials":         me.Credentials,
		"css_selector":        me.CssSelector,
		"datasource_property": me.Datasource_property,
		"email":               me.Email,
		"float_percentage":    me.Float_percentage,
		"host_entity":         me.Host_entity,
		"http_url":            me.Http_url,
		"integer_percentage":  me.Integer_percentage,
		"json_code":           me.JsonCode,
		"management_zone":     me.Management_zone,
		"metric":              me.Metric,
		"monitored_entity":    me.Monitored_entity,
		"multiline_text":      me.Multiline_text,
		"regex":               me.Regex,
		"tag":                 me.Tag,
		"uri":                 me.Uri,
	})
}

func (me *SubPropertyTypes) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"credentials":         &me.Credentials,
		"css_selector":        &me.CssSelector,
		"datasource_property": &me.Datasource_property,
		"email":               &me.Email,
		"float_percentage":    &me.Float_percentage,
		"host_entity":         &me.Host_entity,
		"http_url":            &me.Http_url,
		"integer_percentage":  &me.Integer_percentage,
		"json_code":           &me.JsonCode,
		"management_zone":     &me.Management_zone,
		"metric":              &me.Metric,
		"monitored_entity":    &me.Monitored_entity,
		"multiline_text":      &me.Multiline_text,
		"regex":               &me.Regex,
		"tag":                 &me.Tag,
		"uri":                 &me.Uri,
	})
}
