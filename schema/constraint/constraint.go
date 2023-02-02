package constraint

// Constraint represents a constraint on the values accepted for a settings property
type Constraint struct {
	MaxLength         int32    `json:"maxLength,omitempty"`         // The maximum allowed length of string values
	MinLength         int32    `json:"minLength,omitempty"`         // The minimum required length of string values
	CustomMessage     string   `json:"customMessage,omitempty"`     // A custom message for invalid values
	CustomValidatorID string   `json:"customValidatorId,omitempty"` // The ID of a custom validator
	UniqueProperties  []string `json:"uniqueProperties,omitempty"`  // A list of properties for which the combination of values must be unique
	Maximum           any      `json:"maximum,omitempty"`           // The maximum allowed value
	Minimum           any      `json:"minimum,omitempty"`           // The minimum allowed value
	Type              Type     `json:"type"`                        // The ID of a custom validator
	Pattern           string   `json:"pattern,omitempty"`           // The regular expression pattern for valid string values
}
