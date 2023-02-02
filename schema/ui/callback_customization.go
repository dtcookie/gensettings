package ui

// CallbackCustomization represents UI customization options for defining custom callbacks
type CallbackCustomization struct {
	Buttons []ButtonCustomization `json:"buttons,omitempty"` // UI customization for defining buttons that call functions when pressed
}
