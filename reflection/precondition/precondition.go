package precondition

// Precondition is a precondition for visibility of a property
type Precondition struct {
	Precondition   *Precondition  `json:"precondition,omitempty"`   // A precondition for visibility of a property
	ExpectedValue  any            `json:"expectedValue,omitempty"`  // The expected value of the property. \n\nOnly applicable to properties of the `EQUALS` type
	ExpectedValues []any          `json:"expectedValues,omitempty"` // A list of valid values of the property. \n\nOnly applicable to properties of the `IN` type
	Preconditions  []Precondition `json:"preconditions,omitempty"`  // A list of child preconditions to be evaluated. \n\nOnly applicable to properties of the `AND` and `OR` types
	Property       string         `json:"property,omitempty"`       // The property to be evaluated
	Type           Type           `json:"type"`                     // The type of the precondition
	Pattern        string         `json:"pattern,omitempty"`        // The Regular expression which is matched against the property. \n\nOnly applicable to properties of the `REGEX_MATCH` type
}
