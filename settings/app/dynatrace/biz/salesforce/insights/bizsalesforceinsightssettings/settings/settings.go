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

package bizsalesforceinsightssettings

import (
	"fmt"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	Client_id     string        `json:"client_id"`            // Client Id
	Client_secret string        `json:"client_secret"`        // Client secret
	EventTypes    []EeventTypes `json:"eventTypes,omitempty"` // Event Types
	Grant_type    EgrantType    `json:"grant_type"`           // Possible Values: `Client_credentials`, `Password`
	Name          string        `json:"name"`                 // Configuration Name
	Password      *string       `json:"password,omitempty"`   // Password
	Url           string        `json:"url"`                  // URL of salesforce instance
	Username      *string       `json:"username,omitempty"`   // Username
	WorkflowID    *string       `json:"workflowId,omitempty"` // Changing this will break the app
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"client_id": {
			Type:        schema.TypeString,
			Description: "Client Id",
			Required:    true,
			Sensitive:   true,
		},
		"client_secret": {
			Type:        schema.TypeString,
			Description: "Client secret",
			Required:    true,
			Sensitive:   true,
		},
		"event_types": {
			Type:        schema.TypeSet,
			Description: "Event Types",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"grant_type": {
			Type:        schema.TypeString,
			Description: "Possible Values: `Client_credentials`, `Password`",
			Required:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Description: "Configuration Name",
			Required:    true,
		},
		"password": {
			Type:        schema.TypeString,
			Description: "Password",
			Optional:    true, // precondition
			Sensitive:   true,
		},
		"url": {
			Type:        schema.TypeString,
			Description: "URL of salesforce instance",
			Required:    true,
		},
		"username": {
			Type:        schema.TypeString,
			Description: "Username",
			Optional:    true, // precondition
		},
		"workflow_id": {
			Type:        schema.TypeString,
			Description: "Changing this will break the app",
			Optional:    true, // nullable
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"client_id":     me.Client_id,
		"client_secret": me.Client_secret,
		"event_types":   me.EventTypes,
		"grant_type":    me.Grant_type,
		"name":          me.Name,
		"password":      me.Password,
		"url":           me.Url,
		"username":      me.Username,
		"workflow_id":   me.WorkflowID,
	})
}

func (me *Settings) HandlePreconditions() error {
	if me.Password == nil && (string(me.Grant_type) == "password") {
		return fmt.Errorf("'password' must be specified if 'grant_type' is set to '%v'", me.Grant_type)
	}
	if me.Username == nil && (string(me.Grant_type) == "password") {
		return fmt.Errorf("'username' must be specified if 'grant_type' is set to '%v'", me.Grant_type)
	}
	return nil
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"client_id":     &me.Client_id,
		"client_secret": &me.Client_secret,
		"event_types":   &me.EventTypes,
		"grant_type":    &me.Grant_type,
		"name":          &me.Name,
		"password":      &me.Password,
		"url":           &me.Url,
		"username":      &me.Username,
		"workflow_id":   &me.WorkflowID,
	})
}
