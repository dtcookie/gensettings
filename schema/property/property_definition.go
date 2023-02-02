package property

import (
	"github.com/dtcookie/dynatrace/gensettings/schema/constraint"
	"github.com/dtcookie/dynatrace/gensettings/schema/datasource"
	"github.com/dtcookie/dynatrace/gensettings/schema/precondition"
	"github.com/dtcookie/dynatrace/gensettings/schema/ui"
)

// Definition represents configuration of a property in a settings schema
type Definition struct {
	ReferencedType     string                     `json:"referencedType,omitempty"`     // The type referenced by the property value
	Documentation      string                     `json:"documentation"`                // An extended description and/or links to documentation
	MaxObjects         *int32                     `json:"maxObjects,omitempty"`         // The maximum number of **objects** in a collection property. \n\n Has the value of `1` for singletons
	UICustomization    *ui.Customization          `json:"uiCustomization,omitempty"`    // No documentation available
	Precondition       *precondition.Precondition `json:"precondition,omitempty"`       // No documentation available
	Datasource         *datasource.Definition     `json:"datasource,omitempty"`         // No documentation available
	MinObjects         *int32                     `json:"minObjects,omitempty"`         // The minimum number of **objects** in a collection property
	ModificationPolicy ModificationPolicy         `json:"modificationPolicy,omitempty"` // Modification policy of the property
	Items              *Item                      `json:"items,omitempty"`              // No documentation available
	SubType            string                     `json:"subType,omitempty"`            // The subtype of the property's value
	DisplayName        string                     `json:"displayName"`                  // The display name of the property
	Default            any                        `json:"default,omitempty"`            // The default value to be used when no value is provided. \n\nIf a non-singleton has the value of `null`, it means an empty collection
	Type               any                        `json:"type"`                         // The type of the property's value
	Description        string                     `json:"description"`                  // A short description of the property
	Metadata           map[string]any             `json:"metadata,omitempty"`           // Metadata of the property
	Constraints        []constraint.Constraint    `json:"constraints,omitempty"`        // A list of constraints limiting the values to be accepted
	Nullable           bool                       `json:"nullable"`                     // The value can (`true`) or can't (`false`) be `null`

}

type ModificationPolicy string

var ModificationPolicies = struct {
	Always  ModificationPolicy
	Default ModificationPolicy
	Never   ModificationPolicy
}{
	ModificationPolicy("ALWAYS"),
	ModificationPolicy("DEFAULT"),
	ModificationPolicy("NEVER"),
}
