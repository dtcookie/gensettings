package reflection

import "fmt"

type Library struct {
	Types map[string]*Type
	JSON  []byte
}

func NewLibrary(jsonSchema []byte) *Library {
	return &Library{Types: map[string]*Type{}, JSON: jsonSchema}
}

func (me *Library) AllTypes() []*Type {
	m := map[string]*Type{}
	if err := me.CollectTypes(m); err != nil {
		panic(err.Error())
	}
	result := []*Type{}
	for _, t := range m {
		result = append(result, t)
	}
	return result
}

func (me *Library) CollectTypes(m map[string]*Type) error {
	for _, t := range me.Types {
		if err := t.CollectTypes(m); err != nil {
			return err
		}
	}
	return nil
}

func (me *Library) Define(t *Type, panicIfDuplicate bool) *Type {
	if len(t.ID) == 0 {
		panic("tried to define a type without ID")
	}
	if stored, found := me.Types[t.ID]; found {
		if panicIfDuplicate {
			panic(fmt.Sprintf("type %s already defined", t.ID))
		}
		return stored
	}
	t.Library = me
	me.Types[t.ID] = t
	return t
}

func (me *Library) Rewire(original *Type, replacement *Type) {
	for _, storedType := range me.Types {
		storedType.Rewire(original, replacement)
	}
}
