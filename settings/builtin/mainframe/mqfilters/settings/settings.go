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

package mqfilters

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	CicsMqQueueIdExcludes []string `json:"cicsMqQueueIdExcludes"` // CICS: Excluded MQ queues
	CicsMqQueueIdIncludes []string `json:"cicsMqQueueIdIncludes"` // CICS: Included MQ queues
	ImsCrTrnIdExcludes    []string `json:"imsCrTrnIdExcludes"`    // When you add a transaction ID to the exclude list remaining transactions are still monitored.
	ImsCrTrnIdIncludes    []string `json:"imsCrTrnIdIncludes"`    // When you add a transaction ID to the include list, all the remaining transactions are ignored.
	ImsMqQueueIdExcludes  []string `json:"imsMqQueueIdExcludes"`  // IMS: Excluded MQ queues
	ImsMqQueueIdIncludes  []string `json:"imsMqQueueIdIncludes"`  // IMS: Included MQ queues
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cics_mq_queue_id_excludes": {
			Type:        schema.TypeSet,
			Description: "CICS: Excluded MQ queues",
			Required:    true,

			Elem: &schema.Schema{Type: schema.TypeString},
		},
		"cics_mq_queue_id_includes": {
			Type:        schema.TypeSet,
			Description: "CICS: Included MQ queues",
			Required:    true,

			Elem: &schema.Schema{Type: schema.TypeString},
		},
		"ims_cr_trn_id_excludes": {
			Type:        schema.TypeSet,
			Description: "When you add a transaction ID to the exclude list remaining transactions are still monitored.",
			Required:    true,

			Elem: &schema.Schema{Type: schema.TypeString},
		},
		"ims_cr_trn_id_includes": {
			Type:        schema.TypeSet,
			Description: "When you add a transaction ID to the include list, all the remaining transactions are ignored.",
			Required:    true,

			Elem: &schema.Schema{Type: schema.TypeString},
		},
		"ims_mq_queue_id_excludes": {
			Type:        schema.TypeSet,
			Description: "IMS: Excluded MQ queues",
			Required:    true,

			Elem: &schema.Schema{Type: schema.TypeString},
		},
		"ims_mq_queue_id_includes": {
			Type:        schema.TypeSet,
			Description: "IMS: Included MQ queues",
			Required:    true,

			Elem: &schema.Schema{Type: schema.TypeString},
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"cics_mq_queue_id_excludes": me.CicsMqQueueIdExcludes,
		"cics_mq_queue_id_includes": me.CicsMqQueueIdIncludes,
		"ims_cr_trn_id_excludes":    me.ImsCrTrnIdExcludes,
		"ims_cr_trn_id_includes":    me.ImsCrTrnIdIncludes,
		"ims_mq_queue_id_excludes":  me.ImsMqQueueIdExcludes,
		"ims_mq_queue_id_includes":  me.ImsMqQueueIdIncludes,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"cics_mq_queue_id_excludes": &me.CicsMqQueueIdExcludes,
		"cics_mq_queue_id_includes": &me.CicsMqQueueIdIncludes,
		"ims_cr_trn_id_excludes":    &me.ImsCrTrnIdExcludes,
		"ims_cr_trn_id_includes":    &me.ImsCrTrnIdIncludes,
		"ims_mq_queue_id_excludes":  &me.ImsMqQueueIdExcludes,
		"ims_mq_queue_id_includes":  &me.ImsMqQueueIdIncludes,
	})
}
