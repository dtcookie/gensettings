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

package workload

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	ContainerRestarts        *ContainerRestarts        `json:"containerRestarts"`
	NotAllPodsReady          *NotAllPodsReady          `json:"notAllPodsReady"`
	PendingPods              *PendingPods              `json:"pendingPods"`
	Scope                    *string                   `json:"-" scope:"scope"` // The scope of this setting (CLOUD_APPLICATION_NAMESPACE KUBERNETES_CLUSTER environment)
	WorkloadWithoutReadyPods *WorkloadWithoutReadyPods `json:"workloadWithoutReadyPods"`
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"container_restarts": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(ContainerRestarts).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"not_all_pods_ready": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(NotAllPodsReady).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"pending_pods": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(PendingPods).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"scope": {
			Type:        schema.TypeString,
			Description: "The scope of this setting (CLOUD_APPLICATION_NAMESPACE KUBERNETES_CLUSTER environment)",
			Optional:    true,
			Default:     "environment",
		},
		"workload_without_ready_pods": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Required:    true,

			Elem:     &schema.Resource{Schema: new(WorkloadWithoutReadyPods).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"container_restarts":          me.ContainerRestarts,
		"not_all_pods_ready":          me.NotAllPodsReady,
		"pending_pods":                me.PendingPods,
		"scope":                       me.Scope,
		"workload_without_ready_pods": me.WorkloadWithoutReadyPods,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"container_restarts":          &me.ContainerRestarts,
		"not_all_pods_ready":          &me.NotAllPodsReady,
		"pending_pods":                &me.PendingPods,
		"scope":                       &me.Scope,
		"workload_without_ready_pods": &me.WorkloadWithoutReadyPods,
	})
}
