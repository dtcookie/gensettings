package ui

// TableColumnCustomization represent Customization for UI table columns
type TableColumnCustomization struct {
	BuiltinColumnRef string `json:"builtinColumnRef,omitempty"` // The ui specific builtin column-implementation for this column
	PropertyRef      string `json:"propertyRef,omitempty"`      // The referenced property for this column
	DisplayName      string `json:"displayName,omitempty"`      // The display name for this column
	Type             string `json:"type,omitempty"`             // The ui specific type for this column
}
