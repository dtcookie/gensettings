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

package unifiedservicesfromsubpathsoptin

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	OptInEnabled bool   `json:"optInEnabled"` // Unified service detection is enabled for the PG.
	PgMeID       string `json:"pgMeId"`       // The ME identifier of the PG, something like PROCESS_GROUP-0442D9F8990AF9C4
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"opt_in_enabled": {
			Type:        schema.TypeBool,
			Description: "Unified service detection is enabled for the PG.",
			Required:    true,
		},
		"pg_me_id": {
			Type:        schema.TypeString,
			Description: "The ME identifier of the PG, something like PROCESS_GROUP-0442D9F8990AF9C4",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"opt_in_enabled": me.OptInEnabled,
		"pg_me_id":       me.PgMeID,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"opt_in_enabled": &me.OptInEnabled,
		"pg_me_id":       &me.PgMeID,
	})
}
