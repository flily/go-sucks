package math

import (
	"testing"

	"math"
)

func TestFloat32ToUint32(t *testing.T) {
	nums := []float32{
		0.0,
		1.0,
		-1.0,
		3.14,
		-3.14,
		math.MaxFloat32,
		-math.MaxFloat32,
	}
	for _, v := range nums {
		u := float32ToUint32(v)
		f := uint32ToFloat32(u)
		if f != v {
			t.Errorf("uint32ToFloat32(%v) = %v, want %v", u, f, v)
		}
	}
}

func TestFloat64ToUint64(t *testing.T) {
	nums := []float64{
		0.0,
		1.0,
		-1.0,
		3.14,
		-3.14,
		math.MaxFloat64,
		-math.MaxFloat64,
	}
	for _, v := range nums {
		u := float64ToUint64(v)
		f := uint64ToFloat64(u)
		if f != v {
			t.Errorf("uint64ToFloat64(%v) = %v, want %v", u, f, v)
		}
	}
}
