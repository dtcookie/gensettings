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

type Settings struct {
	AliasQueues   AliasQueues   `json:"aliasQueues"`   // Alias queues
	ClusterQueues ClusterQueues `json:"clusterQueues"` // Cluster queues
	Clusters      []string      `json:"clusters"`      // Name of the cluster(s) this queue manager is part of
	Name          string        `json:"name"`          // Queue manager name
	RemoteQueues  RemoteQueues  `json:"remoteQueues"`  // Remote queues
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"alias_queues": {
			Type:        schema.TypeList,
			Description: "Alias queues",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(AliasQueues).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"cluster_queues": {
			Type:        schema.TypeList,
			Description: "Cluster queues",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(ClusterQueues).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"clusters": {
			Type:        schema.TypeSet,
			Description: "Name of the cluster(s) this queue manager is part of",
			Required:    true,

			Elem: &schema.Schema{Type: schema.TypeString},
		},
		"name": {
			Type:        schema.TypeString,
			Description: "Queue manager name",
			Required:    true,
		},
		"remote_queues": {
			Type:        schema.TypeList,
			Description: "Remote queues",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(RemoteQueues).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"alias_queues":   me.AliasQueues,
		"cluster_queues": me.ClusterQueues,
		"clusters":       me.Clusters,
		"name":           me.Name,
		"remote_queues":  me.RemoteQueues,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"alias_queues":   &me.AliasQueues,
		"cluster_queues": &me.ClusterQueues,
		"clusters":       &me.Clusters,
		"name":           &me.Name,
		"remote_queues":  &me.RemoteQueues,
	})
}
