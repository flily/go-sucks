package meta

import (
	"reflect"
)

// Value is a wrapper of reflect.Value. Provide methods that make sense for meta programming
// and don't panic.
type Value struct {
	value reflect.Value
}

func ValueOf(v interface{}) Value {
	r := Value{
		value: reflect.ValueOf(v),
	}
	return r
}

// Value returns a instance of reflect.Value.
func (v Value) Value() reflect.Value {
	return v.value
}

// IsNilV returns result whether the value is nil. The first return value is true when the value
// is an untyped nil. The second return value is true when the value is a typed nil.
func (v Value) NilType() (bool, bool) {
	switch v.value.Kind() {
	case reflect.Invalid:
		return true, false

	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer:
		return false, v.value.IsNil()

	case reflect.Interface, reflect.Slice:
		return false, v.value.IsNil()

	default:
		return false, false
	}
}

// IsNil returns result whether the value is nil, including both untyped nil and typed nil.
func (v Value) IsNil() bool {
	u, t := v.NilType()
	return u || t
}

// IsUntypedNil returns result whether the value is an untyped nil.
// An untyped nil is the nil literal, or an interface variable which is not bound to a value.
func (v Value) IsUntypedNil() bool {
	u, _ := v.NilType()
	return u
}

// IsTypedNil returns result whether the value is a typed nil.
// A typed nil is a nil pointer that the type of the instance that points to can be determined.
func (v Value) IsTypedNil() bool {
	_, t := v.NilType()
	return t
}
