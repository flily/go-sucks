package meta

import (
	"reflect"

	"github.com/flily/go-sucks/meta/reflectutil"
)

func equalForArray(a reflect.Value, b reflect.Value) bool {
	if a.Len() != b.Len() {
		return false
	}

	for i := 0; i < a.Len(); i++ {
		ai := a.Index(i)
		bi := b.Index(i)
		if !equalForValue(ai, bi) {
			return false
		}
	}

	return true
}

func equalForMap(a reflect.Value, b reflect.Value) bool {
	if a.Len() != b.Len() {
		return false
	}

	keys := map[interface{}]bool{}
	for _, key := range a.MapKeys() {
		copyKey, err := duplicateValueInstance(key)
		if err != nil {
			return false
		}

		keys[copyKey.Interface()] = true
		ai := a.MapIndex(copyKey)
		bi := b.MapIndex(copyKey)

		if !ai.IsValid() || !bi.IsValid() {
			return false
		}

		if !equalForValue(ai, bi) {
			return false
		}
	}

	return true
}

func equalForStruct(a reflect.Value, b reflect.Value) bool {
	for i := 0; i < a.NumField(); i++ {
		af := a.Field(i)
		bf := b.Field(i)

		if !equalForValue(af, bf) {
			return false
		}
	}

	return true
}

func equalForBool(a reflect.Value, b reflect.Value) bool {
	return a.Bool() == b.Bool()
}

func equalForInt(a reflect.Value, b reflect.Value) bool {
	return a.Int() == b.Int()
}

func equalForUint(a reflect.Value, b reflect.Value) bool {
	return a.Uint() == b.Uint()
}

func equalForFloat(a reflect.Value, b reflect.Value) bool {
	return a.Float() == b.Float()
}

func equalForComplex(a reflect.Value, b reflect.Value) bool {
	return a.Complex() == b.Complex()
}

func equalForString(a reflect.Value, b reflect.Value) bool {
	return a.String() == b.String()
}

func equalForUntypedNil(a reflect.Value, b reflect.Value) bool {
	isAUntypedNil, isATypedNil := reflectutil.NilType(a)
	isBUntypedNil, isBTypedNil := reflectutil.NilType(b)
	return (isAUntypedNil || isATypedNil) && (isBUntypedNil || isBTypedNil)
}

func equalForSameKindValue(a reflect.Value, b reflect.Value) bool {
	switch a.Kind() {
	case reflect.Bool:
		return equalForBool(a, b)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return equalForInt(a, b)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return equalForUint(a, b)

	case reflect.Float32, reflect.Float64:
		return equalForFloat(a, b)

	case reflect.Complex64, reflect.Complex128:
		return equalForComplex(a, b)

	case reflect.Array, reflect.Slice:
		return equalForArray(a, b)

	case reflect.Chan, reflect.Func, reflect.UnsafePointer:
		return a.Pointer() == b.Pointer()

	case reflect.Interface, reflect.Ptr:
		ea := a.Elem()
		eb := b.Elem()
		return equalForValue(ea, eb)

	case reflect.Map:
		return equalForMap(a, b)

	case reflect.Struct:
		return equalForStruct(a, b)

	case reflect.String:
		return equalForString(a, b)

	default:
		return false
	}
}

func equalForValue(a reflect.Value, b reflect.Value) bool {
	if !a.IsValid() || !b.IsValid() {
		return equalForUntypedNil(a, b)
	}

	if a.Type() != b.Type() {
		return false
	}

	return equalForSameKindValue(a, b)
}

func ValueEqual(a reflect.Value, b reflect.Value) bool {
	return equalForValue(a, b)
}

func Equal(a interface{}, b interface{}) bool {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)
	return ValueEqual(va, vb)
}
