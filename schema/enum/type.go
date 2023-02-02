package enum

// Type is the definition of an enum property
type Type struct {
	EnumClass     string  `json:"enumClass,omitempty"` // An existing Java enum class that holds the allowed values of the enum
	Items         []Value `json:"items"`               // An existing Java enum class that holds the allowed values of the enum
	Type          string  `json:"type"`                // The type of the property
	Documentation string  `json:"documentation"`       // An extended description and/or links to documentation
	DisplayName   string  `json:"displayName"`         // The display name of the property
	Description   string  `json:"description"`         // A short description of the property
}
