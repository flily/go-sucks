package meta

import (
	"testing"

	"reflect"
)

func TestValueIsNil(t *testing.T) {
	{
		value := ValueOf(nil)
		if !value.IsNil() {
			t.Errorf("value.IsNil() should be true")
		}

		if u, n := value.NilType(); !u || n {
			t.Errorf("value.NilType() should be (true, false)")
		}

		if !value.IsUntypedNil() {
			t.Errorf("value.IsUntypedNil() should be true")
		}

		if value.IsTypedNil() {
			t.Errorf("value.IsTypedNil() should be false")
		}

		innerValue := value.Value()
		if innerValue.Kind() != reflect.Invalid {
			t.Errorf("value.Value() should be reflect.Invalid")
		}
	}
}

func TestValueKind(t *testing.T) {
	{
		value := ValueOf(1)
		if value.Kind() != reflect.Int {
			t.Errorf("value.Kind() should be reflect.Int")
		}
	}
}
