package ui

import (
	"github.com/dtcookie/dynatrace/gensettings/schema/ui/expandable"
	"github.com/dtcookie/dynatrace/gensettings/schema/ui/tabs"
)

// Customization represents Customization for UI elements
type Customization struct {
	Tabs       *tabs.Customization       `json:"tabs,omitempty"`       // UI customization for tabs
	Expandable *expandable.Customization `json:"expandable,omitempty"` // UI customization for expandable section
	Callback   *CallbackCustomization    `json:"callback,omitempty"`   // UI customization options for defining custom callbacks
	Table      *TableCustomization       `json:"table,omitempty"`      // No documentation available
}
