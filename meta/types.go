package meta

import (
	"reflect"
)

// isNilValue returns result whether the value is nil. The first return value is true when the value
// is an untyped nil. The second return value is true when the value is a typed nil.
func isNilValue(value reflect.Value) (bool, bool) {
	switch value.Kind() {
	case reflect.Invalid:
		return true, false

	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer:
		return false, value.IsNil()

	case reflect.Interface, reflect.Slice:
		return false, value.IsNil()

	default:
		return false, false
	}
}

// IsNil returns result whether the value is nil. It always returns true or false for every value,
// and never panics like Value.IsNil(). It returns true when it is untyped nil and typed nil.
func IsNil(value interface{}) (bool, bool) {
	v := reflect.ValueOf(value)
	return isNilValue(v)
}

// IsNilValue returns result whether the value is an untyped nil.
// An untyped nil is the nil literal, or an interface variable which is not bound to a value.
func IsUntypedNil(value interface{}) bool {
	v := reflect.ValueOf(value)
	untypedNil, _ := isNilValue(v)
	return untypedNil
}

// IsNilValue returns result whether the value is a typed nil.
// A typed nil is a nil pointer that the type of the instance that points to can be determined.
func IsTypedNil(value interface{}) bool {
	v := reflect.ValueOf(value)
	_, typedNil := isNilValue(v)
	return typedNil
}
