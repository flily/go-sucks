package reflectutil

import (
	"testing"

	"reflect"
)

func TestNilType(t *testing.T) {
	{
		v := 42
		if u, n := NilType(reflect.ValueOf(v)); u || n {
			t.Errorf("IsNilValue(int) = %v, %v", u, n)
		}
	}

	{
		var v *int // typed nil
		if u, n := NilType(reflect.ValueOf(v)); u || !n {
			t.Errorf("IsNilValue(*int) = %v, %v", u, n)
		}
	}

	{
		var v []int // typed nil
		if u, n := NilType(reflect.ValueOf(v)); u || !n {
			t.Errorf("IsNilValue(nil) = %v, %v", u, n)
		}
	}

	{
		var v interface{} // untyped nil, go sucks
		if u, n := NilType(reflect.ValueOf(v)); !u || n {
			t.Errorf("IsNilValue(interface{}(nil)) = %v, %v", u, n)
		}
	}

	{
		var v interface{} = (*int)(nil) // typed nil, go sucks
		if u, n := NilType(reflect.ValueOf(v)); u || !n {
			t.Errorf("IsNilValue(interface{}(nil)) = %v, %v", u, n)
		}
	}

	{
		f := func(t *testing.T, v interface{}) {
			if u, n := NilType(reflect.ValueOf(v)); !u || n {
				t.Errorf("IsNilValue(%T(nil)) = %v, %v", v, u, n)
			}
		}

		f(t, nil)
	}

	{
		v := []interface{}{
			(*int)(nil),         // typed nil
			interface{}(nil),    // untyped nil, go sucks
			(*interface{})(nil), // typed nil, go sucks
		}
		if u, n := NilType(reflect.ValueOf(v[0])); u || !n {
			t.Errorf("IsNilValue([]interface{}[0](nil)) = %v, %v", u, n)
		}

		if u, n := NilType(reflect.ValueOf(v[1])); !u || n {
			t.Errorf("IsNilValue([]interface{}[1](nil)) = %v, %v", u, n)
		}

		if u, n := NilType(reflect.ValueOf(v[2])); u || !n {
			t.Errorf("IsNilValue([]interface{}[2](nil)) = %v, %v", u, n)
		}
	}

	{
		var v chan int // typed nil
		if u, n := NilType(reflect.ValueOf(v)); u || !n {
			t.Errorf("IsNilValue(nil) = %v, %v", u, n)
		}
	}
}

func TestIsNil(t *testing.T) {
	if !IsNil(reflect.ValueOf(nil)) {
		t.Error("IsNil(nil) should be true")
	}

	if IsNil(reflect.ValueOf(0)) {
		t.Error("IsNil(0) should be false")
	}

	if IsNil(reflect.ValueOf(42)) {
		t.Error("IsNil(42) should be false")
	}

	var pointer *int
	if !IsNil(reflect.ValueOf(pointer)) {
		t.Error("IsNil(pointer) should be true")
	}

	var ifaceUntyped interface{}
	if !IsNil(reflect.ValueOf(ifaceUntyped)) {
		t.Error("IsNil(interface{}) should be true")
	}

	var ifaceTyped interface{} = (*int)(nil)
	if !IsNil(reflect.ValueOf(ifaceTyped)) {
		t.Error("IsNil(interface{}(nil)) should be true")
	}
}

type testEmptyInterface interface{}

type testEmptyStruct struct{}

func TestNilInArray(t *testing.T) {
	array := []interface{}{
		nil,                       // typed nil in struct, untyped nil when copy to local variable, go sucks
		(*int)(nil),               // typed nil in struct, go sucks
		(testEmptyInterface)(nil), // typed nil in struct, untyped nil when copy to local variable, go sucks
		(*testEmptyStruct)(nil),   // typed nil in struct, go sucks
	}

	valueArray := reflect.ValueOf(array)
	{
		valueD := array[0]            // untyped nil after copy to local variable, go sucks
		valueR := valueArray.Index(0) // typed nil, go sucks

		if u, n := NilType(reflect.ValueOf(valueD)); !u || n {
			t.Errorf("IsNilValue(array[0]) = %v, %v", u, n)
		}

		if u, n := NilType(valueR); u || !n {
			t.Errorf("IsNilValue(array.Index(0)) = %v, %v", u, n)
		}
	}

	{
		valueD := array[1]            // typed nil, go sucks
		valueR := valueArray.Index(1) // typed nil, go sucks

		if u, n := NilType(reflect.ValueOf(valueD)); u || !n {
			t.Errorf("IsNilValue(array[1]) = %v, %v", u, n)
		}

		if u, n := NilType(valueR); u || !n {
			t.Errorf("IsNilValue(array.Index(1)) = %v, %v", u, n)
		}
	}

	{
		valueD := array[2]            // untyped nil after copy to local variable, go sucks
		valueR := valueArray.Index(2) // typed nil, go sucks

		if u, n := NilType(reflect.ValueOf(valueD)); !u || n {
			t.Errorf("IsNilValue(array[2]) = %v, %v", u, n)
		}

		if u, n := NilType(valueR); u || !n {
			t.Errorf("IsNilValue(array.Index(2)) = %v, %v", u, n)
		}
	}

	{
		valueD := array[3]            // typed nil, go sucks
		valueR := valueArray.Index(3) // typed nil, go sucks

		if u, n := NilType(reflect.ValueOf(valueD)); u || !n {
			t.Errorf("IsNilValue(array[3]) = %v, %v", u, n)
		}

		if u, n := NilType(valueR); u || !n {
			t.Errorf("IsNilValue(array.Index(3)) = %v, %v", u, n)
		}
	}
}
