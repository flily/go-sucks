package reflectutil

import (
	"reflect"
)

const (
	ElemInterface = 0
	ElemPointer   = 1
)

type ReferenceInfo struct {
	ElemType   int
	SourceType reflect.Type
}

// IsInstance returns true if value is an instance, which can not dereference.
func IsInstance(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Interface:
		return false

	case reflect.Ptr:
		return value.IsNil()

	default:
		return true
	}
}

// DereferenceToInstance returns the instance value and the reference information chain of the value.
func DereferenceToInstance(value reflect.Value) (reflect.Value, []ReferenceInfo) {
	refChain := make([]ReferenceInfo, 0, 4)

	for !IsInstance(value) {
		info := ReferenceInfo{
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

// ReferenceChainOf gets the instance of a value, dereference all levels.
// Return final value, and dereferenced chain.
func ReferenceChainOf(data interface{}) (reflect.Value, []ReferenceInfo) {
	value := reflect.ValueOf(data)
	return DereferenceToInstance(value)
}

// InstanceOf gets the actual instance of a value, dereference all levels.
// If a nil pointer is got, return the reflect.Value represents this pointer. If an untyped nil is
// got, return an invalid reflect.Value, which is returned by reflect.ValueOf().
func InstanceOf(data interface{}) reflect.Value {
	value := reflect.ValueOf(data)
	instance, _ := DereferenceToInstance(value)
	return instance
}

// OriginOf makes original type value via reference chain.
func OriginOf(value reflect.Value, chain []ReferenceInfo) reflect.Value {
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
