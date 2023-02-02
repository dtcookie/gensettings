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

package queuemanagers

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AliasQueues []*AliasQueue

func (me *AliasQueues) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"aliasQueue": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(AliasQueue).Schema()},
		},
	}
}

func (me AliasQueues) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("aliasQueue", me)
}

func (me *AliasQueues) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("aliasQueue", me)
}

// AliasQueue. Alias queue
type AliasQueue struct {
	AliasQueue        string   `json:"aliasQueue"`        // Alias queue name
	BaseQueue         string   `json:"baseQueue"`         // Base queue name
	ClusterVisibility []string `json:"clusterVisibility"` // Name of the cluster(s) this alias should be visible in
}

func (me *AliasQueue) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"alias_queue": {
			Type:        schema.TypeString,
			Description: "Alias queue name",
			Required:    true,
		},
		"base_queue": {
			Type:        schema.TypeString,
			Description: "Base queue name",
			Required:    true,
		},
		"cluster_visibility": {
			Type:        schema.TypeSet,
			Description: "Name of the cluster(s) this alias should be visible in",
			Required:    true,

			Elem: &schema.Schema{Type: schema.TypeString},
		},
	}
}

func (me *AliasQueue) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"alias_queue":        me.AliasQueue,
		"base_queue":         me.BaseQueue,
		"cluster_visibility": me.ClusterVisibility,
	})
}

func (me *AliasQueue) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"alias_queue":        &me.AliasQueue,
		"base_queue":         &me.BaseQueue,
		"cluster_visibility": &me.ClusterVisibility,
	})
}
