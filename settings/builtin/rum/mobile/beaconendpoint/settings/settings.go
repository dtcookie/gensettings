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

package beaconendpoint

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Scope string             `json:"-" scope:"scope"` // The scope of this setting (MOBILE_APPLICATION, CUSTOM_APPLICATION)
	Type  BeaconEndpointType `json:"type"`            // Possible Values: `CLUSTER_ACTIVEGATE`, `ENVIRONMENT_ACTIVEGATE`, `INSTRUMENTED_WEBSERVER`
	Url   *string            `json:"url,omitempty"`   // This must be a valid beacon endpoint URL.
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (MOBILE_APPLICATION, CUSTOM_APPLICATION)",
			Required:    true,
		},
		"type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `CLUSTER_ACTIVEGATE`, `ENVIRONMENT_ACTIVEGATE`, `INSTRUMENTED_WEBSERVER`",
			Required:    true,
		},
		"url": {
			Type:        schema.TypeString,
			Description: "This must be a valid beacon endpoint URL.",
			Optional:    true, // precondition
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"scope": me.Scope,
		"type":  me.Type,
		"url":   me.Url,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"scope": &me.Scope,
		"type":  &me.Type,
		"url":   &me.Url,
	})
}
