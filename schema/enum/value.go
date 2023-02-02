package enum

// Value is an allowed value for an enum property
type Value struct {
	EnumInstance string `json:"enumInstance,omitempty"` // The name of the value in an existing Java enum class
	Icon         string `json:"icon,omitempty"`         // The icon of the value
	DisplayName  string `json:"displayName"`            // The display name of the value
	Value        any    `json:"value"`                  // The allowed value of the enum
	Description  string `json:"description,omitempty"`  // A short description of the value
}
