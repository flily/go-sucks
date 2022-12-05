package meta

import (
	"reflect"

	"github.com/flily/go-sucks/meta/reflectutil"
)

// IsNil returns result whether the value is nil. It always returns true or false for every value,
// and never panics like Value.IsNil(). It returns true when it is untyped nil and typed nil.
func IsNil(value interface{}) (bool, bool) {
	v := reflect.ValueOf(value)
	return reflectutil.NilType(v)
}

// IsUntypedNil returns result whether the value is an untyped nil.
// An untyped nil is the nil literal, or an interface variable which is not bound to a value.
func IsUntypedNil(value interface{}) bool {
	return ValueOf(value).IsUntypedNil()
}

// IsTypedNil returns result whether the value is a typed nil.
// A typed nil is a nil pointer that the type of the instance that points to can be determined.
func IsTypedNil(value interface{}) bool {
	return ValueOf(value).IsTypedNil()
}

// IsPointer returns result whether the value is a pointer.
func IsPointer(v interface{}) bool {
	value := reflect.ValueOf(v)
	return value.Kind() == reflect.Ptr
}

// IsStruct returns result whether the value is a struct.
func IsStruct(v interface{}) bool {
	instance := reflectutil.InstanceOf(v)

	return instance.Kind() == reflect.Struct
}
