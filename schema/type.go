package schema

import (
	"github.com/dtcookie/dynatrace/gensettings/schema/constraint"
	"github.com/dtcookie/dynatrace/gensettings/schema/property"
)

// Type is a list of definitions of types. \n\n A type is a complex property that contains its own set of subproperties
type Type struct {
	SummaryPattern string                         `json:"summaryPattern"`          // The pattern for the summary (for example, \"Alert after *X* minutes.\") of the configuration in the UI
	SearchPattern  string                         `json:"searchPattern,omitempty"` // The pattern for the summary search(for example, \"Alert after *X* minutes.\") of the configuration in the UI
	VersionInfo    string                         `json:"versionInfo"`             // A short description of the version
	Version        string                         `json:"version"`                 // The version of the type
	Properties     map[string]property.Definition `json:"properties"`              // Definition of properties that can be persisted
	Constraints    []constraint.Complex           `json:"constraints,omitempty"`   // A list of constraints limiting the values to be accepted
	Documentation  string                         `json:"documentation"`           // An extended description and/or links to documentation
	DisplayName    string                         `json:"displayName"`             // The display name of the type
	Description    string                         `json:"description"`             // A short description of the property
	Type           string                         `json:"type"`                    // No documentation available
}
