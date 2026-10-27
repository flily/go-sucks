package math

import (
	"unsafe"
)

func float32ToUint32(v float32) uint32 {
	p := unsafe.Pointer(&v)
	return *(*uint32)(p)
}

func uint32ToFloat32(v uint32) float32 {
	p := unsafe.Pointer(&v)
	return *(*float32)(p)
}

func float64ToUint64(v float64) uint64 {
	p := unsafe.Pointer(&v)
	return *(*uint64)(p)
}

func uint64ToFloat64(v uint64) float64 {
	p := unsafe.Pointer(&v)
	return *(*float64)(p)
}
