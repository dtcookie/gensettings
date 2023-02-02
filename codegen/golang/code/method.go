package code

type Methods map[string]*Method

type Method struct {
	Name      string
	Owner     Type
	Comment   Comment
	Arguments []*Argument
	Returns   []Type
	Body      string
}

func NewMethod(name string, owner Type, comment Comment, arguments []*Argument, returns []Type, body string) *Method {
	return &Method{name, owner, comment, arguments, returns, body}
}

func (me *Method) RenameTo(name string) bool {
	var methods map[string]*Method
	switch owner := me.Owner.(type) {
	case *Struct:
		methods = owner.Methods
	case *Alias:
		methods = owner.Methods
	}
	if _, found := methods[name]; found {
		return false
	}

	delete(methods, me.Name)

	me.Name = name
	methods[name] = me

	return true
}
