package meta

import (
	"reflect"
)

// IsNilValue returns result whether the value is nil. The first return value is true when the value
// is an untyped nil. The second return value is true when the value is a typed nil.
func IsNilValue(value reflect.Value) (bool, bool) {
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
	return IsNilValue(v)
}

// IsValueUntypedNil returns result whether the value is an untyped nil.
// An untyped nil is the nil literal, or an interface variable which is not bound to a value.
func IsValueUntypedNil(value reflect.Value) bool {
	isUntypedNil, _ := IsNilValue(value)
	return isUntypedNil
}

// IsUntypedNil returns result whether the value is an untyped nil.
// An untyped nil is the nil literal, or an interface variable which is not bound to a value.
func IsUntypedNil(value interface{}) bool {
	v := reflect.ValueOf(value)
	return IsValueUntypedNil(v)
}

// IsValueTypedNil returns result whether the value is a typed nil.
// A typed nil is a nil pointer that the type of the instance that points to can be determined.
func IsValueTypedNil(value reflect.Value) bool {
	_, typedNil := IsNilValue(value)
	return typedNil
}

// IsTypedNil returns result whether the value is a typed nil.
// A typed nil is a nil pointer that the type of the instance that points to can be determined.
func IsTypedNil(value interface{}) bool {
	v := reflect.ValueOf(value)
	return IsValueTypedNil(v)
}

// IsPointer returns result whether the value is a pointer.
func IsPointer(v interface{}) bool {
	value := reflect.ValueOf(v)
	return value.Kind() == reflect.Ptr
}

// IsStruct returns result whether the value is a struct.
func IsStruct(v interface{}) bool {
	instance := InstanceOf(v)

	return instance.Kind() == reflect.Struct
}
