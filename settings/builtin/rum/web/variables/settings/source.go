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

package variables

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/opt"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Source struct {
	CookieValue      *string    `json:"cookieValue,omitempty"`      // Specify a cookie name to capture its value.
	CssSelector      *string    `json:"cssSelector,omitempty"`      // Specify a CSS selector. This mechanism will capture the first match's innerText/textContent value. To retrieve a specific attribute value of the element, append the '@' symbol, followed by the attribute name. i.e. '#someDomElement@someAttribute'
	JsVariable       *string    `json:"jsVariable,omitempty"`       // Specify a JavaScript variable. This variable must be available globally and the capture expression must use dot notation for property reference. i.e. 'someVar.version'. See documentation for details.
	MetaTag          *string    `json:"metaTag,omitempty"`          // Specify a meta tag name to capture its 'content' value.
	QueryString      *string    `json:"queryString,omitempty"`      // Specify a query string parameter to capture its value.
	RequestAttribute *string    `json:"requestAttribute,omitempty"` // Select an existing request attribute or [create a new one](/#settings/requestattributes;reqAttrNew=true).
	Type             SourceType `json:"type"`                       // Possible Values: `COOKIE_VALUE`, `CSS_SELECTOR`, `CUSTOM_ACTION_NAME`, `ELEMENT_IDENTIFIER`, `JAVASCRIPT_VARIABLE`, `KEYPRESS_EVENT`, `META_TAG`, `PAGE_TITLE`, `PAGE_URL`, `PAGE_URL_ANCHOR`, `PAGE_URL_PATH`, `QUERY_STRING`, `REQUEST_ATTRIBUTE`, `SOURCE_URL`, `SOURCE_URL_ANCHOR`, `SOURCE_URL_PATH`, `TOP_XHR_URL`, `TOP_XHR_URL_ANCHOR`, `TOP_XHR_URL_PATH`, `USER_INTERACTION`, `XHR_URL`, `XHR_URL_ANCHOR`, `XHR_URL_PATH`
}

func (me *Source) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cookie_value": {
			Type:        schema.TypeString,
			Description: "Specify a cookie name to capture its value.",
			Optional:    true, // precondition
		},
		"css_selector": {
			Type:        schema.TypeString,
			Description: "Specify a CSS selector. This mechanism will capture the first match's innerText/textContent value. To retrieve a specific attribute value of the element, append the '@' symbol, followed by the attribute name. i.e. '#someDomElement@someAttribute'",
			Optional:    true, // precondition
		},
		"js_variable": {
			Type:        schema.TypeString,
			Description: "Specify a JavaScript variable. This variable must be available globally and the capture expression must use dot notation for property reference. i.e. 'someVar.version'. See documentation for details.",
			Optional:    true, // precondition
		},
		"meta_tag": {
			Type:        schema.TypeString,
			Description: "Specify a meta tag name to capture its 'content' value.",
			Optional:    true, // precondition
		},
		"query_string": {
			Type:        schema.TypeString,
			Description: "Specify a query string parameter to capture its value.",
			Optional:    true, // precondition
		},
		"request_attribute": {
			Type:        schema.TypeString,
			Description: "Select an existing request attribute or [create a new one](/#settings/requestattributes;reqAttrNew=true).",
			Optional:    true, // precondition
		},
		"type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `COOKIE_VALUE`, `CSS_SELECTOR`, `CUSTOM_ACTION_NAME`, `ELEMENT_IDENTIFIER`, `JAVASCRIPT_VARIABLE`, `KEYPRESS_EVENT`, `META_TAG`, `PAGE_TITLE`, `PAGE_URL`, `PAGE_URL_ANCHOR`, `PAGE_URL_PATH`, `QUERY_STRING`, `REQUEST_ATTRIBUTE`, `SOURCE_URL`, `SOURCE_URL_ANCHOR`, `SOURCE_URL_PATH`, `TOP_XHR_URL`, `TOP_XHR_URL_ANCHOR`, `TOP_XHR_URL_PATH`, `USER_INTERACTION`, `XHR_URL`, `XHR_URL_ANCHOR`, `XHR_URL_PATH`",
			Required:    true,
		},
	}
}

func (me *Source) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"cookie_value":      me.CookieValue,
		"css_selector":      me.CssSelector,
		"js_variable":       me.JsVariable,
		"meta_tag":          me.MetaTag,
		"query_string":      me.QueryString,
		"request_attribute": me.RequestAttribute,
		"type":              me.Type,
	})
}

func (me *Source) HandlePreconditions() error {
	if me.CookieValue == nil && (string(me.Type) == "COOKIE_VALUE") {
		me.CookieValue = opt.NewString("")
	}
	if me.CssSelector == nil && (string(me.Type) == "CSS_SELECTOR") {
		me.CssSelector = opt.NewString("")
	}
	if me.JsVariable == nil && (string(me.Type) == "JAVASCRIPT_VARIABLE") {
		me.JsVariable = opt.NewString("")
	}
	if me.MetaTag == nil && (string(me.Type) == "META_TAG") {
		me.MetaTag = opt.NewString("")
	}
	if me.QueryString == nil && (string(me.Type) == "QUERY_STRING") {
		me.QueryString = opt.NewString("")
	}
	if me.RequestAttribute == nil && (string(me.Type) == "REQUEST_ATTRIBUTE") {
		me.RequestAttribute = opt.NewString("")
	}
	return nil
}

func (me *Source) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"cookie_value":      &me.CookieValue,
		"css_selector":      &me.CssSelector,
		"js_variable":       &me.JsVariable,
		"meta_tag":          &me.MetaTag,
		"query_string":      &me.QueryString,
		"request_attribute": &me.RequestAttribute,
		"type":              &me.Type,
	})
}
