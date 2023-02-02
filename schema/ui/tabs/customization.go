package tabs

// Customization represents UI customization for tabs
type Customization struct {
	Groups []GroupCustomization `json:"groups,omitempty"` // A list of tab group customizations for UI
}
