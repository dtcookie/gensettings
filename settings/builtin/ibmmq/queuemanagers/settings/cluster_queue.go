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

type ClusterQueues []*ClusterQueue

func (me *ClusterQueues) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"clusterQueue": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(ClusterQueue).Schema()},
		},
	}
}

func (me ClusterQueues) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("clusterQueue", me)
}

func (me *ClusterQueues) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("clusterQueue", me)
}

type ClusterQueue struct {
	ClusterVisibility []string `json:"clusterVisibility"` // Name of the cluster(s) this local queue should be visible in
	LocalQueue        string   `json:"localQueue"`        // Local queue name
}

func (me *ClusterQueue) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cluster_visibility": {
			Type:        schema.TypeSet,
			Description: "Name of the cluster(s) this local queue should be visible in",
			Required:    true,

			Elem: &schema.Schema{Type: schema.TypeString},
		},
		"local_queue": {
			Type:        schema.TypeString,
			Description: "Local queue name",
			Required:    true,
		},
	}
}

func (me *ClusterQueue) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"cluster_visibility": me.ClusterVisibility,
		"local_queue":        me.LocalQueue,
	})
}

func (me *ClusterQueue) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"cluster_visibility": &me.ClusterVisibility,
		"local_queue":        &me.LocalQueue,
	})
}
