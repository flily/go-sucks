package reflectutil

import (
	"reflect"
)

// NewUntypedNil creates a new untyped nil value, e.g. nil literal.
func NewUntypedNil() reflect.Value {
	return reflect.Value{}
}

// NewTypedNilOf creates a new typed nil value, e.g. nil pointer to a type.
func NewTypedNil(t reflect.Type) reflect.Value {
	return reflect.Zero(t)
}

// NewPointerOf create a new pointer of specified type.
func NewPointerOf(t reflect.Type) reflect.Value {
	pointer := reflect.New(t)
	return pointer
}

// NewValueOfType creates a new value of specified type.
func NewValueOfType(t reflect.Type) reflect.Value {
	pointer := NewPointerOf(t)
	return pointer.Elem()
}

// NewEmptyValueOfValue create a new value container with same type of specified value.
func NewEmptyValueOfValue(v reflect.Value) reflect.Value {
	isUntypedNil, isTypedNil := NilType(v)
	if isUntypedNil {
		return NewUntypedNil()
	}

	if isTypedNil {
		return NewTypedNil(v.Type())
	}

	return NewValueOfType(v.Type())
}

// NewPointerOfValue create a new pointer of specified value.
func NewPointerTo(value reflect.Value) reflect.Value {
	vt := value.Type()
	addr := reflect.New(vt)
	addr.Elem().Set(value)
	return addr
}
