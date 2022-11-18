package meta

import (
	"testing"

	"reflect"
)

func TestIsNil(t *testing.T) {
	if u, n := IsNil(nil); !u || n {
		t.Errorf("IsNil(nil) = %v, %v", u, n)
	}

	if u, n := IsNil(0); u || n {
		t.Errorf("IsNil(0) = %v, %v", u, n)
	}

	var pointer *int
	if u, n := IsNil(pointer); u || !n {
		t.Errorf("IsNil(pointer) = %v, %v", u, n)
	}

	var channel chan int
	if u, n := IsNil(channel); u || !n {
		t.Errorf("IsNil(channel) = %v, %v", u, n)
	}

	var function func()
	if u, n := IsNil(function); u || !n {
		t.Errorf("IsNil(function) = %v, %v", u, n)
	}

	var mapValue map[int]int
	if u, n := IsNil(mapValue); u || !n {
		t.Errorf("IsNil(mapValue) = %v, %v", u, n)
	}

	var slice []int
	if u, n := IsNil(slice); u || !n {
		t.Errorf("IsNil(slice) = %v, %v", u, n)
	}

	// GO SUCKS
	var interfaceValue error
	if u, n := IsNil(interfaceValue); !u || n {
		t.Errorf("IsNil(interfaceValue) = %v, %v", u, n)
	}
}

func TestIsPointer(t *testing.T) {
	v := 42
	if IsPointer(v) {
		t.Error("IsPointer(v) should be false")
	}

	if !IsPointer(&v) {
		t.Error("IsPointer(&v) should be true")
	}
}

func TestIsPointerOfNil(t *testing.T) {
	var v *int
	if !IsPointer(v) {
		t.Error("IsPointer(v) should be true")
	}

	if IsPointer(nil) {
		t.Error("IsPointer(nil) should be false")
	}
}

type ttIsStruct1 struct {
}

func (ttIsStruct1) One()    {}
func (ttIsStruct1) Two()    {}
func (*ttIsStruct1) Three() {}
func (*ttIsStruct1) Four()  {}

type ttIsStruct1Interface1 interface {
	One()
	Two()
}

type ttIsStruct1Interface2 interface {
	Three()
	Four()
}

func TestStructBasis(t *testing.T) {
	data := ttIsStruct1{}

	{
		// An interface type variable is struct.
		var i ttIsStruct1Interface1 = data
		v := reflect.ValueOf(i)

		if v.Kind() != reflect.Struct {
			t.Errorf("%s is not struct: %s", v, v.Kind())
		}
	}

	{
		// An interface type variable is pointer.
		var i ttIsStruct1Interface2 = &data

		v := reflect.ValueOf(i)
		if v.Kind() != reflect.Ptr {
			t.Errorf("%s is not ptr: %s", v, v.Kind())
		}
	}

	{
		// An interface type variable, which is a pointer, element type is interface.
		var i ttIsStruct1Interface2 = &data

		v := reflect.ValueOf(&i).Type().Elem()
		if v.Kind() != reflect.Interface {
			t.Errorf("%s is not interface: %s", v, v.Kind())
		}
	}

	{
		// A pointer to interface type is pointer.
		var i ttIsStruct1Interface2 = &data
		x := &i

		v := reflect.ValueOf(*x)
		if v.Kind() != reflect.Ptr {
			t.Errorf("%s is not ptr: %s", v, v.Kind())
		}
	}
}

func TestIsStruct(t *testing.T) {
	{
		v := ttIsStruct1{}
		if !IsStruct(v) {
			t.Error("IsStruct(v) should be true")
		}

		if !IsStruct(&v) {
			t.Error("IsStruct(&v) should be true")
		}
	}

	{
		var v ttIsStruct1Interface1 = ttIsStruct1{}
		if !IsStruct(v) {
			t.Error("IsStruct(v) should be true")
		}

		if !IsStruct(&v) {
			t.Error("IsStruct(&v) should be true")
		}
	}
}
