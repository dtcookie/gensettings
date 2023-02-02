package constraint

type ComplexType string

var ComplexTypes = struct {
	CustomValidatorRef ComplexType
	GreaterThan        ComplexType
	GreaterThanOrEqual ComplexType
	LessThan           ComplexType
	LessThanOrEqual    ComplexType
	PropertyCountRange ComplexType
	Unknown            ComplexType
}{
	ComplexType("CUSTOM_VALIDATOR_REF"),
	ComplexType("GREATER_THAN"),
	ComplexType("GREATER_THAN_OR_EQUAL"),
	ComplexType("LESS_THAN"),
	ComplexType("LESS_THAN_OR_EQUAL"),
	ComplexType("PROPERTY_COUNT_RANGE"),
	ComplexType("UNKNOWN"),
}
