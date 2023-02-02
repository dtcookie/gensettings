package golang

type CodeContributorKind int

const (
	AliasContrib CodeContributorKind = iota
	StructContrib
	EnumContrib
)

type CodeContributor interface {
	Bytes() ([]byte, error)
	FileName() string
	Imports() []string
	Prettify()
	Kind() CodeContributorKind
	SortName() string
}
