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

package rumweb

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ResponseTimeAuto struct {
	OverAlertingProtection *OverAlertingProtectionAuto `json:"overAlertingProtection"` // Avoid over-alerting
	ResponseTimeAll        *ResponseTimeAutoAll        `json:"responseTimeAll"`        // Alert if the median response time of all user actions degrades beyond **both** the absolute and relative thresholds:
	ResponseTimeSlowest    *ResponseTimeAutoSlowest    `json:"responseTimeSlowest"`    // Alert if the response time of the slowest 10% of requests degrades beyond **both** the absolute and relative thresholds:
}

func (me *ResponseTimeAuto) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"over_alerting_protection": {
			Type:        schema.TypeList,
			Description: "Avoid over-alerting",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(OverAlertingProtectionAuto).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"response_time_all": {
			Type:        schema.TypeList,
			Description: "Alert if the median response time of all user actions degrades beyond **both** the absolute and relative thresholds:",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(ResponseTimeAutoAll).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"response_time_slowest": {
			Type:        schema.TypeList,
			Description: "Alert if the response time of the slowest 10% of requests degrades beyond **both** the absolute and relative thresholds:",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(ResponseTimeAutoSlowest).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *ResponseTimeAuto) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"over_alerting_protection": me.OverAlertingProtection,
		"response_time_all":        me.ResponseTimeAll,
		"response_time_slowest":    me.ResponseTimeSlowest,
	})
}

func (me *ResponseTimeAuto) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"over_alerting_protection": &me.OverAlertingProtection,
		"response_time_all":        &me.ResponseTimeAll,
		"response_time_slowest":    &me.ResponseTimeSlowest,
	})
}
