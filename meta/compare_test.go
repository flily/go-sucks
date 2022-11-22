package meta

import (
	"reflect"
	"testing"
)

func TestValueEqualOnSimpleValue(t *testing.T) {
	var a, b interface{}

	a = 1
	b = 1
	if !Equal(a, b) {
		t.Errorf("unexpected result: %v", a)
	}

	a = 1
	b = 2
	if Equal(a, b) {
		t.Errorf("unexpected result: %v", a)
	}

	a = nil
	b = nil
	if !Equal(a, b) {
		t.Errorf("unexpected result: %v", a)
	}

	a = nil
	b = 1
	if Equal(a, b) {
		t.Errorf("unexpected result: %v", a)
	}

	a = 1
	b = nil
	if Equal(a, b) {
		t.Errorf("unexpected result: %v", a)
	}

	a = 1
	b = "a"
	if Equal(a, b) {
		t.Errorf("unexpected result: %v", a)
	}

	a = "a"
	b = "a"
	if !Equal(a, b) {
		t.Errorf("unexpected result: %v", a)
	}
}

func TestValueEqualOnNil(t *testing.T) {
	if !Equal(nil, nil) {
		t.Errorf("unexpected result: nil != nil")
	}

	var a *int
	var b *int
	if !Equal(a, b) {
		t.Errorf("unexpected result: nil != nil")
	}
}

func TestValueEqualOnUnknownKind(t *testing.T) {
	a := reflect.Value{}
	b := reflect.Value{}

	// use for coverage test, untyped nil will never pass into equalForSameKindValue
	if equalForSameKindValue(a, b) {
		t.Errorf("unexpected result: %+v == %+v", a, b)
	}
}

func TestValueEqualOnBool(t *testing.T) {
	if !Equal(true, true) {
		t.Errorf("unexpected result")
	}

	if Equal(true, false) {
		t.Errorf("unexpected result")
	}

	if !Equal(false, false) {
		t.Errorf("unexpected result")
	}

	if Equal(false, true) {
		t.Errorf("unexpected result")
	}
}

func TestValueEqualOnInt(t *testing.T) {
	var a int = 42

	if !Equal(a, 42) {
		t.Errorf("unexpected result, %v != %v", a, 42)
	}

	if Equal(a, 43) {
		t.Errorf("unexpected result, %v == %v", a, 43)
	}

	if Equal(a, int8(42)) {
		t.Errorf("unexpected result, %v == int8(%v)", a, int8(42))
	}
}

func TestValueEqualOnUint(t *testing.T) {
	var a uint = 42

	if !Equal(a, uint(42)) {
		t.Errorf("unexpected result, %v != %v", a, 42)
	}

	if Equal(a, uint(43)) {
		t.Errorf("unexpected result, %v == %v", a, 43)
	}

	if Equal(a, 42) {
		t.Errorf("unexpected result, %v == int(%v)", a, 42)
	}

	if Equal(a, uint8(42)) {
		t.Errorf("unexpected result, %v == uint8(%v)", a, uint8(42))
	}
}

func TestValueEqualOnComplex(t *testing.T) {
	var a complex64 = 1 + 2i
	var b complex64 = 1 + 2i
	var c complex64 = 2 + 2i

	if !Equal(a, b) {
		t.Errorf("unexpected result: %v != %v", a, b)
	}

	if Equal(a, c) {
		t.Errorf("unexpected result: %v == %v", a, c)
	}
}

func TestValueEqualOnSimpleArray(t *testing.T) {
	aa := []int{1, 3, 5, 7, 9}
	ab := []int{1, 3, 5, 7, 9}

	if !Equal(aa, ab) {
		t.Errorf("unexpected result: %v != %v", aa, ab)
	}

	ac := []int{1, 3, 5, 7, 9, 11}
	if Equal(aa, ac) {
		t.Errorf("unexpected result: %v == %v", aa, ac)
	}

	ad := []int32{1, 3, 5, 7, 9}
	if Equal(aa, ad) {
		t.Errorf("unexpected result: %v == %v", aa, ad)
	}

	ae := []int{1, 3, 6, 7, 9}
	if Equal(aa, ae) {
		t.Errorf("unexpected result: %v == %v", aa, ae)
	}
}

func TestValueEqualOnStruct(t *testing.T) {
	type test struct {
		Name string
		Age  int
	}

	a := test{"a", 1}
	b := test{"b", 2}
	c := test{"a", 1}

	if Equal(a, b) {
		t.Errorf("unexpected result: %+v == %+v", a, b)
	}

	if !Equal(a, c) {
		t.Errorf("unexpected result: %+v != %+v", a, c)
	}
}

func TestValueEqualOnStructWithPrivateField(t *testing.T) {
	type test struct {
		Name string
		age  int
	}

	a := test{"a", 1}
	b := test{"b", 2}
	c := test{"a", 1}

	if Equal(a, b) {
		t.Errorf("unexpected result: %+v == %+v", a, b)
	}

	if !Equal(a, c) {
		t.Errorf("unexpected result: %+v != %+v", a, c)
	}
}

func TestValueEqualOnChannel(t *testing.T) {
	var a chan int
	var b chan int

	if !Equal(a, b) {
		t.Errorf("unexpected result: %v == %v", a, b)
	}

	a = make(chan int)
	b = make(chan int)
	if Equal(a, b) {
		t.Errorf("unexpected result: %v == %v", a, b)
	}

	a = b
	if !Equal(a, a) {
		t.Errorf("unexpected result: %v != %v", a, a)
	}
}
