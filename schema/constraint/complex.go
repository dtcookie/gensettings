package constraint

// Complex represents a constraint on the values accepted for a complex settings property
type Complex struct {
	CustomMessage        string      `json:"customMessage,omitempty"`        // A custom message for invalid values
	CustomValidatorID    string      `json:"customValidatorId,omitempty"`    // The ID of a custom validator
	MinimumPropertyCount int32       `json:"minimumPropertyCount,omitempty"` // The minimum number of properties that must be set
	MaximumPropertyCount int32       `json:"maximumPropertyCount,omitempty"` // The maximum number of properties that can be set
	Properties           []string    `json:"properties,omitempty"`           // A list of properties (defined by IDs) that are used to check the constraint
	Type                 ComplexType `json:"type"`                           // The type of the constraint
}
