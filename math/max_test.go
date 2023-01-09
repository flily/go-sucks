package math

import (
	"testing"
)

func TestMaxInt64(t *testing.T) {
	nums := []int64{1, 2, 3, 4, 5}
	got := MaxInt64(nums...)
	if got != 5 {
		t.Errorf("MaxInt64(%v) = %v, want %v", nums, got, 5)
	}
}

func TestMinInt64(t *testing.T) {
	nums := []int64{1, 2, 3, 4, 5}
	got := MinInt64(nums...)
	if got != 1 {
		t.Errorf("MinInt64(%v) = %v, want %v", nums, got, 1)
	}
}

func TestMaxUint64(t *testing.T) {
	nums := []uint64{1, 2, 3, 4, 5}
	got := MaxUint64(nums...)
	if got != 5 {
		t.Errorf("MaxUint64(%v) = %v, want %v", nums, got, 5)
	}
}

func TestMinUint64(t *testing.T) {
	nums := []uint64{1, 2, 3, 4, 5}
	got := MinUint64(nums...)
	if got != 1 {
		t.Errorf("MinUint64(%v) = %v, want %v", nums, got, 1)
	}
}

func TestMaxFloat64(t *testing.T) {
	nums := []float64{1, 2, 3, 4, 5}
	got := MaxFloat64(nums...)
	if got != 5 {
		t.Errorf("MaxFloat64(%v) = %v, want %v", nums, got, 5)
	}
}

func TestMinFloat64(t *testing.T) {
	nums := []float64{1, 2, 3, 4, 5}
	got := MinFloat64(nums...)
	if got != 1 {
		t.Errorf("MinFloat64(%v) = %v, want %v", nums, got, 1)
	}
}
