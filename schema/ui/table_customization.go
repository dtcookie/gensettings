package ui

// TableCustomization represents Customization for UI tables
type TableCustomization struct {
	EmptyState *EmptyStateCustomization   `json:"emptyState"` // UI customization for empty state in a table
	Columns    []TableColumnCustomization `json:"columns"`    // A list of columns for the UI table
}
