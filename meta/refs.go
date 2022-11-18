package meta

import (
	"reflect"
)

const (
	ElemInterface = 0
	ElemPointer   = 1
)

type ValueReferenceInfo struct {
	ElemType   int
	SourceType reflect.Type
}

// IsValueInstance returns true if value is an instance, which can not dereference.
func IsValueInstance(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Interface:
		return false

	case reflect.Ptr:
		return value.IsNil()

	default:
		return true
	}
}

// ValueToInstance returns the instance value and the reference information chain of the value.
func ValueToInstance(value reflect.Value) (reflect.Value, []ValueReferenceInfo) {
	refChain := make([]ValueReferenceInfo, 0, 4)

	for !IsValueInstance(value) {
		info := ValueReferenceInfo{
			SourceType: value.Type(),
		}

		switch value.Kind() {
		case reflect.Ptr:
			info.ElemType = ElemPointer

		case reflect.Interface:
			info.ElemType = ElemInterface
		}

		value = value.Elem()
		refChain = append(refChain, info)
	}

	return value, refChain
}

// ValueInstanceChainOf gets the instance of a value, dereference all levels.
// Return final value, and dereferenced chain.
func ValueInstanceChainOf(data interface{}) (reflect.Value, []ValueReferenceInfo) {
	value := reflect.ValueOf(data)
	return ValueToInstance(value)
}

// InstanceOf gets the actual instance of a value, dereference all levels.
// If a nil pointer is got, return the reflect.Value represents this pointer. If an untyped nil is
// got, return an invalid reflect.Value, which is returned by reflect.ValueOf().
func InstanceOf(data interface{}) reflect.Value {
	value := reflect.ValueOf(data)
	return ValueInstanceOf(value)
}

// ValueInstanceOf gets the instance of a value, dereference all levels.
func ValueInstanceOf(value reflect.Value) reflect.Value {
	final, _ := ValueToInstance(value)
	return final
}

// OriginOf makes original type value via reference chain.
func OriginOf(value reflect.Value, chain []ValueReferenceInfo) reflect.Value {
	if len(chain) <= 0 {
		return value
	}

	for i := len(chain) - 1; i >= 0; i-- {
		info := chain[i]

		switch info.ElemType {
		case ElemPointer:
			value = NewPointerTo(value)

		case ElemInterface:
			value = value.Convert(info.SourceType)
		}
	}

	return value
}
