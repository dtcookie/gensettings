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

package frequentissues

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	DetectFrequentIssuesInApplications            bool `json:"detectFrequentIssuesInApplications"`            // Detect frequent issues within applications
	DetectFrequentIssuesInInfrastructure          bool `json:"detectFrequentIssuesInInfrastructure"`          // Detect frequent issues within infrastructure
	DetectFrequentIssuesInTransactionsAndServices bool `json:"detectFrequentIssuesInTransactionsAndServices"` // Detect frequent issues within transactions and services
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"detect_frequent_issues_in_applications": {
			Type:        schema.TypeBool,
			Description: "Detect frequent issues within applications",
			Required:    true,
		},
		"detect_frequent_issues_in_infrastructure": {
			Type:        schema.TypeBool,
			Description: "Detect frequent issues within infrastructure",
			Required:    true,
		},
		"detect_frequent_issues_in_transactions_and_services": {
			Type:        schema.TypeBool,
			Description: "Detect frequent issues within transactions and services",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"detect_frequent_issues_in_applications":              me.DetectFrequentIssuesInApplications,
		"detect_frequent_issues_in_infrastructure":            me.DetectFrequentIssuesInInfrastructure,
		"detect_frequent_issues_in_transactions_and_services": me.DetectFrequentIssuesInTransactionsAndServices,
	})
}

func (me *Settings) HandlePreconditions() {
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"detect_frequent_issues_in_applications":              &me.DetectFrequentIssuesInApplications,
		"detect_frequent_issues_in_infrastructure":            &me.DetectFrequentIssuesInInfrastructure,
		"detect_frequent_issues_in_transactions_and_services": &me.DetectFrequentIssuesInTransactionsAndServices,
	})
}
