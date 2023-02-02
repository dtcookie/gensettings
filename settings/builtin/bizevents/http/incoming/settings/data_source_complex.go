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

package incoming

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type DataSourceComplex struct {
	DataSource DataSourceEnum `json:"dataSource"`     // Possible Values: `Response_statusCode`, `Request_path`, `Request_method`, `Request_headers`, `Request_parameters`, `Request_body`, `Response_body`, `Response_headers`
	Path       *string        `json:"path,omitempty"` // [See our documentation](https://dt-url.net/ei034bx)
}

func (me *DataSourceComplex) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"data_source": {
			Type:        schema.TypeString,
			Description: "Possible Values: `Response_statusCode`, `Request_path`, `Request_method`, `Request_headers`, `Request_parameters`, `Request_body`, `Response_body`, `Response_headers`",
			Required:    true,
		},
		"path": {
			Type:        schema.TypeString,
			Description: "[See our documentation](https://dt-url.net/ei034bx)",
			Optional:    true,
		},
	}
}

func (me *DataSourceComplex) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"data_source": me.DataSource,
		"path":        me.Path,
	})
}

func (me *DataSourceComplex) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"data_source": &me.DataSource,
		"path":        &me.Path,
	})
}
