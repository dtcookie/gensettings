package precondition

type Type string

var Types = struct {
	And        Type
	Equals     Type
	In         Type
	Not        Type
	Null       Type
	Or         Type
	RegexMatch Type
}{
	Type("AND"),
	Type("EQUALS"),
	Type("IN"),
	Type("NOT"),
	Type("NULL"),
	Type("OR"),
	Type("REGEX_MATCH"),
}
