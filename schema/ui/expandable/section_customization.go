package expandable

// SectionCustomization represents expandable section customization for UI
type SectionCustomization struct {
	Expanded    bool     `json:"expanded,omitempty"`    // Defines if the section should be expanded by default
	DisplayName string   `json:"displayName"`           // The display name of the section
	Properties  []string `json:"properties"`            // A list of properties
	Description string   `json:"description,omitempty"` // The description
}
