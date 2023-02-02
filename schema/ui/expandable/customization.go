package expandable

// Customization represents UI customization for expandable section
type Customization struct {
	Expanded    bool                   `json:"expanded,omitempty"`    // Defines if the item should be expanded by default
	DisplayName string                 `json:"displayName,omitempty"` // The display name
	Sections    []SectionCustomization `json:"sections,omitempty"`    // A list of sections

}
