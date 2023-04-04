package preconditions

import (
	"fmt"
	"reflect"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/opt"
)

type Precondition interface {
	Evaluate(Attribute) error
}

func In(property any, v any) Precondition {
	return nil
}

func Not(precondition Precondition) Precondition {
	return nil
}

func And(preconditions ...Precondition) Precondition {
	return nil
}

func Or(preconditions ...Precondition) Precondition {
	return nil
}

type Attribute struct {
	Value any
	Name  string
}

func Equals(attribute Attribute, v any) Precondition {
	return &equals{attribute: attribute, v: v}
}

type equals struct {
	attribute Attribute
	v         any
}

var boolType = reflect.TypeOf(true)
var boolPtrType = reflect.TypeOf(opt.NewBool(false))
var stringType = reflect.TypeOf("")
var stringPtrType = reflect.TypeOf(opt.NewString(""))

type err struct {
	message string
}

func (e *err) Error() string {
	return e.message
}

func (me *equals) Evaluate(attribute Attribute) error {
	targetValue := reflect.ValueOf(attribute.Value)
	targetIsNil := (attribute.Value == nil) || (targetValue.IsValid() && targetValue.IsNil())
	compValue := reflect.ValueOf(me.attribute.Value)
	compIsNil := me.attribute.Value == nil || (compValue.IsValid() && compValue.IsNil())
	if targetIsNil && compIsNil {
		return nil
	}
	if !targetIsNil && !targetValue.IsValid() {
		return &err{fmt.Sprintf("unable to check preconditions for attribute '%s': passed value is invalid", attribute.Name)}
	}

	if !compIsNil && !compValue.IsValid() {
		return &err{fmt.Sprintf("unable to check preconditions for attribute '%s': passed value is for attribute '%s' is invalid", attribute.Name, me.attribute.Name)}
	}
	if targetIsNil { // target is only allowed to be nil if the precondition is false
		switch me.v.(type) {
		case string:
			if compValue.CanConvert(stringType) {
				if me.v.(string) == compValue.Convert(stringType).Interface().(string) {
					return fmt.Errorf("'%s' must be specified if '%s' is set to '%s'", attribute.Name, me.attribute.Name, me.v.(string))
				} else {
					return nil
				}
			} else if compValue.Elem().CanConvert(stringType) {
				if me.v.(string) == compValue.Elem().Convert(stringType).Interface().(string) {
					return fmt.Errorf("'%s' must be specified if '%s' is set to '%s'", attribute.Name, me.attribute.Name, me.v.(string))
				} else {
					return nil
				}
			}
			return &err{fmt.Sprintf("unable to check preconditions for attribute '%s': the value of attribute '%s' was expected to be of kind string, but its type is %T", attribute.Name, me.attribute.Name, me.attribute.Value)}
		case bool:
			if compValue.CanConvert(boolType) {
				if me.v.(bool) == compValue.Convert(boolType).Interface().(bool) {
					return fmt.Errorf("'%s' must be specified if '%s' is set to '%v'", attribute.Name, me.attribute.Name, me.v.(bool))
				} else {
					return nil
				}
			} else if compValue.Elem().CanConvert(boolType) {
				if me.v.(bool) == compValue.Elem().Convert(boolType).Interface().(bool) {
					return fmt.Errorf("'%s' must be specified if '%s' is set to '%v'", attribute.Name, me.attribute.Name, me.v.(bool))
				} else {
					return nil
				}
			}
			return &err{fmt.Sprintf("unable to check preconditions for attribute '%s': the value of attribute '%s' was expected to be of kind bool, but its type is %T", attribute.Name, me.attribute.Name, me.attribute.Value)}
		default:
			return fmt.Errorf("the value of attribute '%s' is of type '%T' which is not supported", me.attribute.Name, me.v)
		}
	} else {
		if compIsNil {
			return fmt.Errorf("'%s' requires '%s' to be set to '%v'", attribute.Name, me.attribute.Name, me.v)
		}
		switch me.v.(type) {
		case string:
			if compValue.CanConvert(stringType) {
				if me.v.(string) != compValue.Convert(stringType).Interface().(string) {
					return fmt.Errorf("'%s' requires '%s' to be set to '%v'", attribute.Name, me.attribute.Name, me.v.(string))
				} else {
					return nil
				}
			} else if compValue.Elem().CanConvert(stringType) {
				if me.v.(string) != compValue.Elem().Convert(stringType).Interface().(string) {
					return fmt.Errorf("'%s' requires '%s' to be set to '%v'", attribute.Name, me.attribute.Name, me.v.(string))
				} else {
					return nil
				}
			}
			return &err{fmt.Sprintf("unable to check preconditions for attribute '%s': the value of attribute '%s' was expected to be of kind string, but its type is %T", attribute.Name, me.attribute.Name, me.attribute.Value)}
		case bool:
			if compValue.CanConvert(boolType) {
				if me.v.(bool) != compValue.Convert(boolType).Interface().(bool) {
					return fmt.Errorf("'%s' requires '%s' to be set to '%v'", attribute.Name, me.attribute.Name, me.v.(bool))
				} else {
					return nil
				}
			} else if compValue.Elem().CanConvert(boolType) {
				if me.v.(bool) != compValue.Elem().Convert(boolType).Interface().(bool) {
					return fmt.Errorf("'%s' requires '%s' to be set to '%v'", attribute.Name, me.attribute.Name, me.v.(bool))
				} else {
					return nil
				}
			}
			return &err{fmt.Sprintf("unable to check preconditions for attribute '%s': the value of attribute '%s' was expected to be of kind bool, but its type is %T", attribute.Name, me.attribute.Name, me.attribute.Value)}
		default:
			return fmt.Errorf("the value of attribute '%s' is of type '%T' which is not supported", me.attribute.Name, me.v)
		}
	}
}

// func (me *equals)
