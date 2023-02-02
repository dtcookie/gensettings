package ui

// ButtonCustomization represents UI customization for defining a button that calls a function when pressed
type ButtonCustomization struct {
	InsertAfter string `json:"insertAfter"` // The path of a property where the button should be shown in the UI
	Function    string `json:"function"`    // The function to be called when the button is pressed
	Title       string `json:"title"`       // The title of the button
}
