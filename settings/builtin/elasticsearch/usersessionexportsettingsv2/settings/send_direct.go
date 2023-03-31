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

package usersessionexportsettingsv2

import (
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SendDirect struct {
	Active    bool    `json:"active"`              // Activate
	DocType   *string `json:"docType,omitempty"`   // Type of documents in the Elasticsearch index
	IndexName *string `json:"indexName,omitempty"` // Name of the index where data is sent
}

func (me *SendDirect) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active": {
			Type:        schema.TypeBool,
			Description: "Activate",
			Required:    true,
		},
		"doc_type": {
			Type:        schema.TypeString,
			Description: "Type of documents in the Elasticsearch index",
			Optional:    true, // nullable & precondition
		},
		"index_name": {
			Type:        schema.TypeString,
			Description: "Name of the index where data is sent",
			Optional:    true, // precondition
		},
	}
}

func (me *SendDirect) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"active":     me.Active,
		"doc_type":   me.DocType,
		"index_name": me.IndexName,
	})
}

func (me *SendDirect) HandlePreconditions() error {
	if me.DocType == nil && me.Active {
		return fmt.Errorf("'doc_type' must be specified if 'active' is set to '%v'", me.Active)
	}
	if me.IndexName == nil && me.Active {
		return fmt.Errorf("'index_name' must be specified if 'active' is set to '%v'", me.Active)
	}
	return nil
}

func (me *SendDirect) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"active":     &me.Active,
		"doc_type":   &me.DocType,
		"index_name": &me.IndexName,
	})
}
