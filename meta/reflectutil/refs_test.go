package reflectutil

import (
	"testing"

	"reflect"
)

type ttTypeCastData struct {
	Name string
	Age  int
}

func (t *ttTypeCastData) Say(word string) string {
	return t.Name + " says " + word
}

func (t ttTypeCastData) Shout(word string) string {
	return t.Name + " shouts " + word + "!"
}

type ttTypeSayingInterface interface {
	Say(string) string
}

type ttTypeShoutingInterface interface {
	Shout(string) string
}

func TestInstanceOf(t *testing.T) {
	data := 3

	dataPtr := &data
	value := reflect.ValueOf(data)
	{
		got := InstanceOf(dataPtr)
		if got.Interface() != value.Interface() {
			t.Errorf("got error value: %v <=> %v", got, value)
		}
	}

	{
		got := reflect.Indirect(reflect.ValueOf(dataPtr))
		if got.Interface() != value.Interface() {
			t.Errorf("got error value: %v <=> %v", got, value)
		}
	}

	dataPtrPtr := &dataPtr
	{
		got := InstanceOf(dataPtrPtr)
		if got.Interface() != value.Interface() {
			t.Errorf("got error value: %v <=> %v", got, value)
		}
	}

	{
		// NOTES: reflect.Indirect returns the first data that points to.
		got := reflect.Indirect(reflect.ValueOf(dataPtrPtr))
		if got.Interface() == value.Interface() {
			t.Errorf("got error value: %v <=> %v", got, value)
		}
	}

	{
		var nilPtr *int = nil
		got := InstanceOf(nilPtr)
		if !got.IsValid() {
			t.Errorf("got error value: %v, is valid %v", got, got.IsValid())
		}

		if !got.IsNil() {
			t.Errorf("got error value: %v <=> %v", got, value)
		}
	}

	{
		// Untyped nil
		got := InstanceOf(nil)
		if got.IsValid() {
			t.Errorf("got error value: %v, is valid %v", got, got.IsValid())
		}
	}
}

func TestOrigineOf(t *testing.T) {
	data := ttTypeCastData{
		Name: "Lily",
		Age:  33,
	}

	ptr := &data

	{
		var inf ttTypeSayingInterface = ptr

		v := reflect.ValueOf(inf)
		if v.Kind() != reflect.Ptr {
			t.Errorf("%s is not pointer", v.Kind())
		}

		ins, chain := ReferenceChainOf(inf)
		ori := OriginOf(ins, chain)

		if *(ori.Interface().(*ttTypeCastData)) != *ptr {
			t.Errorf("origin of %v is %v, expect %v", ins, ori, ptr)
		}
	}

	{
		var inf ttTypeShoutingInterface = data

		v := reflect.ValueOf(inf)
		if v.Kind() != reflect.Struct {
			t.Errorf("%s is not Struct", v.Kind())
		}

		ins, chain := ReferenceChainOf(inf)
		ori := OriginOf(ins, chain)

		if ori.Interface() != data {
			t.Errorf("origin of %v is %v, expect %v", ins, ori, data)
		}
	}

	{
		var s ttTypeSayingInterface = ptr
		inf := &s

		v := reflect.ValueOf(inf)
		if v.Kind() != reflect.Ptr {
			t.Errorf("%s is not pointer", v.Kind())
		}

		ins, chain := ReferenceChainOf(inf)
		ori := OriginOf(ins, chain)

		p, ok := ori.Interface().(*ttTypeSayingInterface)
		if !ok {
			t.Errorf("origin of %v is %v, expect %v", ins, ori, ptr)
		}

		if *(*p).(*ttTypeCastData) != *ptr {
			t.Errorf("origin of %v is %v, expect %v", ins, ori, ptr)
		}
	}
}
