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
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Authentication     *Authentication     `json:"authentication"`     // Dynatrace can automatically send bulk data to Elasticsearch. You can use an SSL certificate, basic authentication or OAuth 2.0 to secure access. For Dynatrace SaaS installations, the Elasticsearch instance must be publicly reachable by the Dynatrace Cluster.
	EndpointDefinition *EndpointDefinition `json:"endpointDefinition"` // Dynatrace will send JSON data periodically to this endpoint. You can also pause and disable an endpoint to stop the data stream from Dynatrace to your endpoint.
	ExportBehavior     *ExportBehavior     `json:"exportBehavior"`     // Define the scope of your export by using a specific management zone. You can also disable UI notifications for failing exports, or add special settings provided by Dynatrace support for troubleshooting.
	SendDirect         *SendDirect         `json:"sendDirect"`         // Activate this if you want to export user session data to your own Elasticsearch cluster. If you run Elasticsearch 7, make sure to enter <em>_doc</em> as the type. For Elasticsearch 8 omit the type. If you really want to use a type, then you have to add <em>include_type_name=true</em> when creating your Elasticsearch index. See more information in the Dynatrace help.
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"authentication": {
			Type:        schema.TypeList,
			Description: "Dynatrace can automatically send bulk data to Elasticsearch. You can use an SSL certificate, basic authentication or OAuth 2.0 to secure access. For Dynatrace SaaS installations, the Elasticsearch instance must be publicly reachable by the Dynatrace Cluster.",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(Authentication).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"endpoint_definition": {
			Type:        schema.TypeList,
			Description: "Dynatrace will send JSON data periodically to this endpoint. You can also pause and disable an endpoint to stop the data stream from Dynatrace to your endpoint.",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(EndpointDefinition).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"export_behavior": {
			Type:        schema.TypeList,
			Description: "Define the scope of your export by using a specific management zone. You can also disable UI notifications for failing exports, or add special settings provided by Dynatrace support for troubleshooting.",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(ExportBehavior).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"send_direct": {
			Type:        schema.TypeList,
			Description: "Activate this if you want to export user session data to your own Elasticsearch cluster. If you run Elasticsearch 7, make sure to enter <em>_doc</em> as the type. For Elasticsearch 8 omit the type. If you really want to use a type, then you have to add <em>include_type_name=true</em> when creating your Elasticsearch index. See more information in the Dynatrace help.",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(SendDirect).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"authentication":      me.Authentication,
		"endpoint_definition": me.EndpointDefinition,
		"export_behavior":     me.ExportBehavior,
		"send_direct":         me.SendDirect,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"authentication":      &me.Authentication,
		"endpoint_definition": &me.EndpointDefinition,
		"export_behavior":     &me.ExportBehavior,
		"send_direct":         &me.SendDirect,
	})
}
