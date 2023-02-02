package reflection

import "fmt"

type Type struct {
	Library    *Library
	ID         string `json:"id,omitempty"`
	Name       string
	Kind       Kind       `json:"kind"`
	Elem       *Type      `json:"elem,omitempty"`
	Properties Properties `json:"properties,omitempty"`
	Comment    string
}

func (me *Type) Equals(other *Type) bool {
	return me.ID == other.ID
}

func (me *Type) CollectTypes(m map[string]*Type) error {
	if _, exists := m[me.ID]; exists {
		return nil
	}
	m[me.ID] = me
	if me.Elem != nil {
		if err := me.Elem.CollectTypes(m); err != nil {
			return fmt.Errorf("%s. Type: %s", err.Error(), me.ID)
		}
	}
	if len(me.Properties) > 0 {
		if err := me.Properties.CollectTypes(m); err != nil {
			return fmt.Errorf("%s. Type: %s", err.Error(), me.ID)
		}
	}
	return nil
}

func (me *Type) Rewire(original *Type, replacement *Type) {
	if me.Elem != nil {
		if me.Elem.Equals(original) {
			me.Elem = replacement
		}
		me.Elem.Rewire(original, replacement)
	}
	if len(me.Properties) > 0 {
		me.Properties.Rewire(original, replacement)
	}
}

func (me *Type) List() *Type {
	return &Type{
		Library: me.Library,
		ID:      "[]" + me.ID,
		Kind:    ListKind,
		Elem:    me,
	}
}

func (me *Type) Set() *Type {
	return &Type{
		Library: me.Library,
		ID:      "()" + me.ID,
		Kind:    SetKind,
		Elem:    me,
	}
}

func (me *Type) Pointer() *Type {
	if me.Kind == PointerKind {
		return me
	}
	return &Type{
		Library: me.Library,
		ID:      "*" + me.ID,
		Kind:    PointerKind,
		Elem:    me,
	}
}

func (me *Type) UnrefPtr() *Type {
	if me.Kind == PointerKind {
		return me.Elem.Unref()
	}
	return me
}

func (me *Type) Unref() *Type {
	if me.Kind == PointerKind {
		return me.Elem.Unref()
	}
	if me.Kind == ListKind {
		return me.Elem.Unref()
	}
	if me.Kind == SetKind {
		return me.Elem.Unref()
	}
	return me
}
