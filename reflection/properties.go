package reflection

import (
	"fmt"

	"github.com/dtcookie/dynatrace/gensettings/reflection/precondition"
)

type Property struct {
	Name            string `json:"name"`
	RawName         string
	Type            *Type  `json:"type,omitempty"`
	Comment         string `json:"comment,omitempty"`
	Optional        bool   `json:"optional"`
	OptionalComment string
	Preconditions   []*precondition.Precondition `json:"preconditions"`
	Scope           string                       `json:"scope"`
}

func (me *Property) Rewire(original *Type, replacement *Type) {
	if me.Type.Equals(original) {
		me.Type = replacement
	}
	// me.Type.Rewire(original, replacement)
}

func (me *Property) CollectTypes(m map[string]*Type) error {
	if me.Type == nil {
		return fmt.Errorf("property %s has no type", me.Name)
	}
	return me.Type.CollectTypes(m)
}

type Properties map[string]*Property

func (me Properties) Rewire(original *Type, replacement *Type) {
	for _, property := range me {
		property.Rewire(original, replacement)
	}
}

func (me Properties) CollectTypes(m map[string]*Type) error {
	for propertyName, property := range me {
		if err := property.CollectTypes(m); err != nil {
			return fmt.Errorf("%s. Poperty: %s", err.Error(), propertyName)
		}
	}
	return nil
}
