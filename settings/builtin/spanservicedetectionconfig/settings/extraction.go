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

package spanservicedetectionconfig

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Extractions []*Extraction

func (me *Extractions) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"extraction": {
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Description: "",
			Elem:        &schema.Resource{Schema: new(Extraction).Schema()},
		},
	}
}

func (me Extractions) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeSlice("extraction", me)
}

func (me *Extractions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("extraction", me)
}

type Extraction struct {
	Transformations ExtractionTypes `json:"transformations"`
	UniqueID        string          `json:"uniqueId"`    // This unique identifier can be used to reference to the result of the extraction in later steps (f.ex. Service detection & Service method detection). The unique identifier must start with a letter and allows the usage of `numbers` and the following special characters [**.**  **_**  **-**]
	ValueSource     *ValueSource    `json:"valueSource"` // Value source
}

func (me *Extraction) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"transformations": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(ExtractionTypes).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"unique_id": {
			Type:        schema.TypeString,
			Description: "This unique identifier can be used to reference to the result of the extraction in later steps (f.ex. Service detection & Service method detection). The unique identifier must start with a letter and allows the usage of `numbers` and the following special characters [**.**  **_**  **-**]",
			Required:    true,
		},
		"value_source": {
			Type:        schema.TypeList,
			Description: "Value source",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(ValueSource).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Extraction) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"transformations": me.Transformations,
		"unique_id":       me.UniqueID,
		"value_source":    me.ValueSource,
	})
}

func (me *Extraction) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"transformations": &me.Transformations,
		"unique_id":       &me.UniqueID,
		"value_source":    &me.ValueSource,
	})
}
