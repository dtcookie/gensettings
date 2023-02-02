package reflection

type Kind int

const (
	EnumKind Kind = iota
	StructKind
	StringKind
	BooleanKind
	IntKind
	FloatKind
	ListKind
	SetKind
	PointerKind
	SliceAliasKind
	SetAliasKind
)
