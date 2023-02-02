package datasource

import "github.com/dtcookie/dynatrace/gensettings/collections"

// Definition represents configuration of a datasource for a property
type Definition struct {
	FilterProperties collections.Set[string] `json:"filterProperties"`     // The properties to filter the datasource options on
	Validate         bool                    `json:"validate"`             // Whether to validate input to only allow values returned by the datasource
	FullContext      bool                    `json:"fullContext"`          // Whether this datasource expects full setting payload as the context
	UseApiSearch     bool                    `json:"useApiSearch"`         // If true, the datasource should use the api to filter the results instead of client-side filtering
	ResetValue       ResetValue              `json:"resetValue,omitempty"` // When to reset datasource value in the UI on filter change
	Identifier       string                  `json:"identifier"`           // The identifier of a custom data source of the property's value
}

type ResetValue string

var ResetValues = struct {
	Always      ResetValue
	InvalidOnly ResetValue
	Never       ResetValue
}{
	ResetValue("ALWAYS"),
	ResetValue("INVALID_ONLY"),
	ResetValue("NEVER"),
}
