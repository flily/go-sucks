package meta

import (
	"reflect"
)

// IsNil returns result whether the value is nil. It always returns true or false for every value,
// and never panics like Value.IsNil(). It returns true when it is untyped nil and typed nil.
func IsNil(value interface{}) bool {
	v := reflect.ValueOf(value)
	if !v.IsValid() {
		return true
	}

	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer:
		return v.IsNil()

	case reflect.Interface, reflect.Slice:
		return v.IsNil()

	default:
		return false
	}
}
