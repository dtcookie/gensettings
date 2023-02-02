package schema

import (
	"github.com/dtcookie/dynatrace/gensettings/collections"
	"github.com/dtcookie/dynatrace/gensettings/schema/constraint"
	"github.com/dtcookie/dynatrace/gensettings/schema/enum"
	"github.com/dtcookie/dynatrace/gensettings/schema/property"
	"github.com/dtcookie/dynatrace/gensettings/schema/ui"
)

type Definition struct {
	Dynatrace       string                         `json:"dynatrace"`                 // The version of the data format
	SchemaID        string                         `json:"schemaId"`                  // The ID of the schema
	DisplayName     string                         `json:"displayName"`               // The display name of the schema
	Description     string                         `json:"description"`               // A short description of the schema
	Documentation   string                         `json:"documentation"`             // An extended description of the schema and/or links to documentation
	SchemaGroups    collections.Set[string]        `json:"schemaGroups,omitempty"`    // Names of the groups, which the schema belongs to
	Version         string                         `json:"version"`                   // The version of the schema
	MultiObject     bool                           `json:"multiObject"`               // Multiple (`true`) objects per scope are permitted or a single (`false``) object per scope is permitted
	Ordered         *bool                          `json:"ordered,omitempty"`         // If `true`` the order of objects has semantic significance
	MaxObjects      int32                          `json:"maxObjects"`                // The maximum amount of objects per scope. Only applicable when `multiObject`` is set to true
	AllowedScopes   collections.Set[string]        `json:"allowedScopes"`             // A list of scopes where the schema can be used
	Enums           map[string]enum.Type           `json:"enums"`                     // A list of definitions of enum properties
	Types           map[string]Type                `json:"types"`                     // A list of definitions of types. \n\n A type is a complex property that contains its own set of subproperties
	Properties      map[string]property.Definition `json:"properties"`                // A list of schema's properties
	Constraints     []constraint.Complex           `json:"constraints,omitempty"`     // A list of constrains limiting the values to be accepted by the schema
	Metadata        map[string]any                 `json:"metadata,omitempty"`        // Metadata of the setting
	UICustomization *ui.Customization              `json:"uiCustomization,omitempty"` // Customization for UI elements
}
