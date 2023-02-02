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

package txstartfilters

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	IncludedCicsTerminalTransactionIds []string `json:"includedCicsTerminalTransactionIds"` // You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	IncludedCicsTransactionIds         []string `json:"includedCicsTransactionIds"`         // You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
	IncludedImsTransactionIds          []string `json:"includedImsTransactionIds"`          // You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"included_cics_terminal_transaction_ids": {
			Type:        schema.TypeSet,
			Description: "You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.",
			Required:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"included_cics_transaction_ids": {
			Type:        schema.TypeSet,
			Description: "You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.",
			Required:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"included_ims_transaction_ids": {
			Type:        schema.TypeSet,
			Description: "You can use * as wildcard. For example use A* to trace all transaction IDs that start with A.",
			Required:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"included_cics_terminal_transaction_ids": me.IncludedCicsTerminalTransactionIds,
		"included_cics_transaction_ids":          me.IncludedCicsTransactionIds,
		"included_ims_transaction_ids":           me.IncludedImsTransactionIds,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"included_cics_terminal_transaction_ids": &me.IncludedCicsTerminalTransactionIds,
		"included_cics_transaction_ids":          &me.IncludedCicsTransactionIds,
		"included_ims_transaction_ids":           &me.IncludedImsTransactionIds,
	})
}
