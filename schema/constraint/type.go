package constraint

type Type string

var Types = struct {
	CustomValidatorRef Type
	Length             Type
	NotBlank           Type
	NotEmpty           Type
	Pattern            Type
	Range              Type
	Regex              Type
	Trimmed            Type
	Unique             Type
	Unknown            Type
}{
	Type("CUSTOM_VALIDATOR_REF"),
	Type("LENGTH"),
	Type("NOT_BLANK"),
	Type("NOT_EMPTY"),
	Type("PATTERN"),
	Type("RANGE"),
	Type("REGEX"),
	Type("TRIMMED"),
	Type("UNIQUE"),
	Type("UNKNOWN"),
}
