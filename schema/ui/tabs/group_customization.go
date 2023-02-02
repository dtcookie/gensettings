package tabs

// GroupCustomization represents tab group customization for UI
type GroupCustomization struct {
	DisplayName string   `json:"displayName"`           // The display name
	Properties  []string `json:"properties"`            // A list of properties
	Description string   `json:"description,omitempty"` // The description
}
