package code

type Fields map[string]*Field

type Field struct {
	Name     string
	Owner    *Struct
	Comment  Comment
	Type     Type
	Optional bool
	JSON     string
}

func NewField(name string, owner *Struct, comment Comment, typ Type, optional bool, json string) *Field {
	return &Field{name, owner, comment, typ, optional, json}
}

func (me *Field) RenameTo(name string) bool {
	if _, found := me.Owner.Fields[name]; found {
		return false
	}

	delete(me.Owner.Fields, me.Name)

	me.Name = name
	me.Owner.Fields[name] = me

	return true
}
