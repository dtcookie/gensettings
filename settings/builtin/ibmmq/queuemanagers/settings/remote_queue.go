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

type RemoteQueues []*RemoteQueue

func (me *RemoteQueues) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"remote_queue": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(RemoteQueue).Schema()},
		},
	}
}

func (me RemoteQueues) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("remote_queue", me)
}

func (me *RemoteQueues) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("remote_queue", me)
}

type RemoteQueue struct {
	ClusterVisibility  []string `json:"clusterVisibility,omitempty"` // Name of the cluster(s) this local definition of the remote queue should be visible in
	LocalQueue         string   `json:"localQueue"`                  // Local queue name
	RemoteQueue        string   `json:"remoteQueue"`                 // Remote queue name
	RemoteQueueManager string   `json:"remoteQueueManager"`          // Remote queue manager name
}

func (me *RemoteQueue) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cluster_visibility": {
			Type:        schema.TypeSet,
			Description: "Name of the cluster(s) this local definition of the remote queue should be visible in",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"local_queue": {
			Type:        schema.TypeString,
			Description: "Local queue name",
			Required:    true,
		},
		"remote_queue": {
			Type:        schema.TypeString,
			Description: "Remote queue name",
			Required:    true,
		},
		"remote_queue_manager": {
			Type:        schema.TypeString,
			Description: "Remote queue manager name",
			Required:    true,
		},
	}
}

func (me *RemoteQueue) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"cluster_visibility":   me.ClusterVisibility,
		"local_queue":          me.LocalQueue,
		"remote_queue":         me.RemoteQueue,
		"remote_queue_manager": me.RemoteQueueManager,
	})
}

func (me *RemoteQueue) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"cluster_visibility":   &me.ClusterVisibility,
		"local_queue":          &me.LocalQueue,
		"remote_queue":         &me.RemoteQueue,
		"remote_queue_manager": &me.RemoteQueueManager,
	})
}
