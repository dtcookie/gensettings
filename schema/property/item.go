package property

import (
	"github.com/dtcookie/dynatrace/gensettings/schema/constraint"
	"github.com/dtcookie/dynatrace/gensettings/schema/datasource"
	"github.com/dtcookie/dynatrace/gensettings/schema/ui"
)

// Item represents an item of a collection property
type Item struct {
	ReferencedType  string                  `json:"referencedType,omitempty"`  // The type referenced by the item's value
	Documentation   *string                 `json:"documentation,omitempty"`   // An extended description and/or links to documentation
	UICustomization *ui.Customization       `json:"uiCustomization,omitempty"` // Customization for UI elements
	Datasource      *datasource.Definition  `json:"datasource,omitempty"`      // represents configuration of a datasource for a property
	SubType         string                  `json:"subType,omitempty"`         // The subtype of the item's value
	DisplayName     *string                 `json:"displayName,omitempty"`     // The display name of the item
	Type            any                     `json:"type"`                      // The type of the item's value
	Description     *string                 `json:"description,omitempty"`     // A short description of the item
	Metadata        *map[string]any         `json:"metadata,omitempty"`        // Metadata of the items
	Constraints     []constraint.Constraint `json:"constraints,omitempty"`     // A list of constraints limiting the values to be accepted
}
