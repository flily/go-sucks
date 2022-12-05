package reflectutil

import (
	"reflect"
	"testing"
)

func TestNilType(t *testing.T) {
	{
		v := 42
		if u, n := NilType(reflect.ValueOf(v)); u || n {
			t.Errorf("IsNilValue(int) = %v, %v", u, n)
		}
	}

	{
		var v *int
		if u, n := NilType(reflect.ValueOf(v)); u || !n {
			t.Errorf("IsNilValue(*int) = %v, %v", u, n)
		}
	}

	{
		var v []int
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
}
