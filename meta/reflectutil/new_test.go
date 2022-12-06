package reflectutil

import (
	"testing"

	"reflect"
)

func TestNewUntypedNils(t *testing.T) {
	untypedNil := NewUntypedNil()
	if u, n := NilType(untypedNil); !u || n {
		t.Errorf("IsNilValue(untypedNil) = %v, %v", u, n)
	}

	if !IsUntypedNil(untypedNil) {
		t.Errorf("IsValueUntypedNil(untypedNil) should be true")
	}

	if IsTypedNil(untypedNil) {
		t.Errorf("IsValueTypedNil(untypedNil) should be false")
	}
}

func TestNewTypedNil(t *testing.T) {
	{
		v := 42
		typedNil := NewTypedNil(reflect.TypeOf(&v))
		if u, n := NilType(typedNil); u || !n {
			t.Errorf("IsNilValue(typedNil) = %v, %v", u, n)
		}

		if !IsTypedNil(typedNil) {
			t.Errorf("IsValueTypedNil(typedNil) = true")
		}

		if IsUntypedNil(typedNil) {
			t.Errorf("IsValueUntypedNil(typedNil) = false")
		}
	}

	{
		var i *int
		typedNil := NewTypedNil(reflect.TypeOf(i))
		if u, n := NilType(typedNil); u || !n {
			t.Errorf("IsNilValue(typedNil) = %v, %v", u, n)
		}

		if !IsTypedNil(typedNil) {
			t.Error("IsValueTypedNil(typedNil) should be true")
		}

		if IsUntypedNil(typedNil) {
			t.Error("IsValueUntypedNil(typedNil) should be false")
		}
	}

	{
		var v []int
		if u, n := NilType(reflect.ValueOf(v)); u || !n {
			t.Errorf("IsNilValue(nil) = %v, %v", u, n)
		}
	}
}

func TestNewEmptyValueOfValue(t *testing.T) {
	{
		v := 42
		value := NewEmptyValueOfValue(reflect.ValueOf(v))
		if u, n := NilType(value); u || n {
			t.Errorf("IsNilValue(int) = %v, %v", u, n)
		}

		if value.Type() != reflect.TypeOf(v) {
			t.Errorf("new value type is not %v, got %v", reflect.TypeOf(v), value.Type())
		}
	}

	{
		v := 42
		value := NewEmptyValueOfValue(reflect.ValueOf(&v))
		if u, n := NilType(value); u || !n {
			t.Errorf("IsNilValue(*int) = %v, %v", u, n)
		}

		if value.Type() != reflect.TypeOf(&v) {
			t.Errorf("new value type is not %v, got %v", reflect.TypeOf(&v), value.Type())
		}
	}

	{
		var v *int
		value := NewEmptyValueOfValue(reflect.ValueOf(v))
		if u, n := NilType(value); u || !n {
			t.Errorf("IsNilValue(*int) = %v, %v", u, n)
		}

		if value.Type() != reflect.TypeOf(v) {
			t.Errorf("new value type is not %v, got %v", reflect.TypeOf(v), value.Type())
		}
	}

	{
		value := NewEmptyValueOfValue(reflect.ValueOf(nil))
		if u, n := NilType(value); !u || n {
			t.Errorf("IsNilValue(nil) = %v, %v", u, n)
		}

		if value.Kind() != reflect.Invalid {
			t.Errorf("new value type is not %v, got %v", reflect.Invalid, value.Kind())
		}
	}
}
