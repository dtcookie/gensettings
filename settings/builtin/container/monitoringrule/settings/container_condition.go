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

package monitoringrule

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ContainerCondition struct {
	Operator ConditionOperator `json:"operator"`        // Possible Values: `ENDS`, `NOT_CONTAINS`, `EQUALS`, `NOT_EXISTS`, `STARTS`, `NOT_ENDS`, `CONTAINS`, `NOT_EQUALS`, `NOT_STARTS`, `EXISTS`
	Property ContainerItem     `json:"property"`        // Possible Values: `CONTAINER_NAME`, `IMAGE_NAME`, `KUBERNETES_NAMESPACE`, `KUBERNETES_CONTAINERNAME`, `KUBERNETES_BASEPODNAME`, `KUBERNETES_FULLPODNAME`, `KUBERNETES_PODUID`
	Value    *string           `json:"value,omitempty"` // Condition value
}

func (me *ContainerCondition) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"operator": {
			Type:        schema.TypeString,
			Description: "Possible Values: `ENDS`, `NOT_CONTAINS`, `EQUALS`, `NOT_EXISTS`, `STARTS`, `NOT_ENDS`, `CONTAINS`, `NOT_EQUALS`, `NOT_STARTS`, `EXISTS`",
			Required:    true,
		},
		"property": {
			Type:        schema.TypeString,
			Description: "Possible Values: `CONTAINER_NAME`, `IMAGE_NAME`, `KUBERNETES_NAMESPACE`, `KUBERNETES_CONTAINERNAME`, `KUBERNETES_BASEPODNAME`, `KUBERNETES_FULLPODNAME`, `KUBERNETES_PODUID`",
			Required:    true,
		},
		"value": {
			Type:        schema.TypeString,
			Description: "Condition value",
			Optional:    true,
		},
	}
}

func (me *ContainerCondition) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"operator": me.Operator,
		"property": me.Property,
		"value":    me.Value,
	})
}

func (me *ContainerCondition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"operator": &me.Operator,
		"property": &me.Property,
		"value":    &me.Value,
	})
}
