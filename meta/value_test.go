package meta

import (
	"testing"

	"reflect"
)

func TestNewUntypedNils(t *testing.T) {
	untypedNil := NewUntypedNil()
	if u, n := IsNilValue(untypedNil); !u || n {
		t.Errorf("IsNilValue(untypedNil) = %v, %v", u, n)
	}

	if !IsValueUntypedNil(untypedNil) {
		t.Errorf("IsValueUntypedNil(untypedNil) should be true")
	}

	if IsValueTypedNil(untypedNil) {
		t.Errorf("IsValueTypedNil(untypedNil) should be false")
	}
}

func TestNewTypedNil(t *testing.T) {
	{
		v := 42
		typedNil := NewTypedNil(reflect.TypeOf(&v))
		if u, n := IsNilValue(typedNil); u || !n {
			t.Errorf("IsNilValue(typedNil) = %v, %v", u, n)
		}

		if !IsValueTypedNil(typedNil) {
			t.Errorf("IsValueTypedNil(typedNil) = true")
		}

		if IsValueUntypedNil(typedNil) {
			t.Errorf("IsValueUntypedNil(typedNil) = false")
		}
	}

	{
		var i *int
		typedNil := NewTypedNil(reflect.TypeOf(i))
		if u, n := IsNilValue(typedNil); u || !n {
			t.Errorf("IsNilValue(typedNil) = %v, %v", u, n)
		}

		if !IsValueTypedNil(typedNil) {
			t.Error("IsValueTypedNil(typedNil) should be true")
		}

		if IsValueUntypedNil(typedNil) {
			t.Error("IsValueUntypedNil(typedNil) should be false")
		}
	}
}

func TestNewEmptyValueOfValue(t *testing.T) {
	{
		v := 42
		value := NewEmptyValueOfValue(reflect.ValueOf(v))
		if u, n := IsNilValue(value); u || n {
			t.Errorf("IsNilValue(int) = %v, %v", u, n)
		}

		if value.Type() != reflect.TypeOf(v) {
			t.Errorf("new value type is not %v, got %v", reflect.TypeOf(v), value.Type())
		}
	}

	{
		v := 42
		value := NewEmptyValueOfValue(reflect.ValueOf(&v))
		if u, n := IsNilValue(value); u || !n {
			t.Errorf("IsNilValue(*int) = %v, %v", u, n)
		}

		if value.Type() != reflect.TypeOf(&v) {
			t.Errorf("new value type is not %v, got %v", reflect.TypeOf(&v), value.Type())
		}
	}

	{
		var v *int
		value := NewEmptyValueOfValue(reflect.ValueOf(v))
		if u, n := IsNilValue(value); u || !n {
			t.Errorf("IsNilValue(*int) = %v, %v", u, n)
		}

		if value.Type() != reflect.TypeOf(v) {
			t.Errorf("new value type is not %v, got %v", reflect.TypeOf(v), value.Type())
		}
	}

	{
		value := NewEmptyValueOfValue(reflect.ValueOf(nil))
		if u, n := IsNilValue(value); !u || n {
			t.Errorf("IsNilValue(nil) = %v, %v", u, n)
		}

		if value.Kind() != reflect.Invalid {
			t.Errorf("new value type is not %v, got %v", reflect.Invalid, value.Kind())
		}
	}
}
