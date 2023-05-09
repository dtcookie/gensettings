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
	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/opt"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	ApplicationID    string                `json:"-" scope:"applicationId"`    // The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	InstrumentServer *bool                 `json:"instrumentServer,omitempty"` // Instrument web and app server
	Type             WebBeaconEndpointType `json:"type"`                       // Possible Values: `ACTIVEGATE`, `CUSTOM`, `ONEAGENT`
	Url              *string               `json:"url,omitempty"`              // For OneAgent settings, specify either path segments or an absolute URL. If you enter an absolute URL, data will be sent using CORS.\n\nFor Custom settings, specify an absolute URL.
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_id": {
			Type:        schema.TypeString,
			Description: "The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.",
			Required:    true,
		},
		"instrument_server": {
			Type:        schema.TypeBool,
			Description: "Instrument web and app server",
			Optional:    true, // precondition
		},
		"type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `ACTIVEGATE`, `CUSTOM`, `ONEAGENT`",
			Required:    true,
		},
		"url": {
			Type:        schema.TypeString,
			Description: "For OneAgent settings, specify either path segments or an absolute URL. If you enter an absolute URL, data will be sent using CORS.\n\nFor Custom settings, specify an absolute URL.",
			Optional:    true, // precondition
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"application_id":    me.ApplicationID,
		"instrument_server": me.InstrumentServer,
		"type":              me.Type,
		"url":               me.Url,
	})
}

func (me *Settings) HandlePreconditions() error {
	if me.InstrumentServer == nil && (string(me.Type) == "CUSTOM") {
		me.InstrumentServer = opt.NewBool(false)
	}
	if me.Url == nil && (string(me.Type) != "ACTIVEGATE") {
		me.Url = opt.NewString("")
	}
	return nil
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"application_id":    &me.ApplicationID,
		"instrument_server": &me.InstrumentServer,
		"type":              &me.Type,
		"url":               &me.Url,
	})
}
