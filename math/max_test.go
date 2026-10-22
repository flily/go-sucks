package math

import (
	"math/rand/v2"
	"testing"
)

func TestMaxInt64(t *testing.T) {
	nums := []int64{1, 2, 3, 4, 5}

	for range 10 {
		rand.Shuffle(len(nums), func(i, j int) {
			nums[i], nums[j] = nums[j], nums[i]
		})

		got := Max(nums...)
		if got != 5 {
			t.Errorf("MaxInt64(%v) = %v, want %v", nums, got, 5)
		}
	}
}

func TestMinInt64(t *testing.T) {
	nums := []int64{1, 2, 3, 4, 5}

	for range 10 {
		rand.Shuffle(len(nums), func(i, j int) {
			nums[i], nums[j] = nums[j], nums[i]
		})

		got := Min(nums...)
		if got != 1 {
			t.Errorf("MinInt64(%v) = %v, want %v", nums, got, 1)
		}
	}
}

func TestMaxUint64(t *testing.T) {
	nums := []uint64{1, 2, 3, 4, 5}

	for range 10 {
		rand.Shuffle(len(nums), func(i, j int) {
			nums[i], nums[j] = nums[j], nums[i]
		})

		got := Max(nums...)
		if got != 5 {
			t.Errorf("MaxUint64(%v) = %v, want %v", nums, got, 5)
		}
	}
}

func TestMinUint64(t *testing.T) {
	nums := []uint64{1, 2, 3, 4, 5}

	for range 10 {
		rand.Shuffle(len(nums), func(i, j int) {
			nums[i], nums[j] = nums[j], nums[i]
		})

		got := Min(nums...)
		if got != 1 {
			t.Errorf("MinUint64(%v) = %v, want %v", nums, got, 1)
		}
	}
}

func TestMaxFloat64(t *testing.T) {
	nums := []float64{1, 2, 3, 4, 5}
	for range 10 {
		rand.Shuffle(len(nums), func(i, j int) {
			nums[i], nums[j] = nums[j], nums[i]
		})

		got := Max(nums...)
		if got != 5 {
			t.Errorf("MaxFloat64(%v) = %v, want %v", nums, got, 5)
		}
	}
}

func TestMinFloat64(t *testing.T) {
	nums := []float64{1, 2, 3, 4, 5}

	for range 10 {
		rand.Shuffle(len(nums), func(i, j int) {
			nums[i], nums[j] = nums[j], nums[i]
		})

		got := Min(nums...)
		if got != 1 {
			t.Errorf("MinFloat64(%v) = %v, want %v", nums, got, 1)
		}
	}
}

func TestMaxWithNoArguments(t *testing.T) {
	if got := Max[int](); got != 0 {
		t.Errorf("Max() = %v, want %v", got, 0)
	}

	if got := Max[uint](); got != 0 {
		t.Errorf("Max() = %v, want %v", got, 0)
	}

	if got := Max[int8](); got != 0 {
		t.Errorf("Max() = %v, want %v", got, 0)
	}

	if got := Max[int16](); got != 0 {
		t.Errorf("Max() = %v, want %v", got, 0)
	}

	if got := Max[int32](); got != 0 {
		t.Errorf("Max() = %v, want %v", got, 0)
	}

	if got := Max[int64](); got != 0 {
		t.Errorf("Max() = %v, want %v", got, 0)
	}

	if got := Max[uint8](); got != 0 {
		t.Errorf("Max() = %v, want %v", got, 0)
	}

	if got := Max[uint16](); got != 0 {
		t.Errorf("Max() = %v, want %v", got, 0)
	}

	if got := Max[uint32](); got != 0 {
		t.Errorf("Max() = %v, want %v", got, 0)
	}

	if got := Max[uint64](); got != 0 {
		t.Errorf("Max() = %v, want %v", got, 0)
	}

	if got := Max[float32](); got != 0 {
		t.Errorf("Max() = %v, want %v", got, 0)
	}

	if got := Max[float64](); got != 0 {
		t.Errorf("Max() = %v, want %v", got, 0)
	}
}

func TestMinWithNoArguments(t *testing.T) {
	if got := Min[int](); got != 0 {
		t.Errorf("Min() = %v, want %v", got, 0)
	}

	if got := Min[uint](); got != 0 {
		t.Errorf("Min() = %v, want %v", got, 0)
	}

	if got := Min[int8](); got != 0 {
		t.Errorf("Min() = %v, want %v", got, 0)
	}

	if got := Min[int16](); got != 0 {
		t.Errorf("Min() = %v, want %v", got, 0)
	}

	if got := Min[int32](); got != 0 {
		t.Errorf("Min() = %v, want %v", got, 0)
	}

	if got := Min[int64](); got != 0 {
		t.Errorf("Min() = %v, want %v", got, 0)
	}

	if got := Min[uint8](); got != 0 {
		t.Errorf("Min() = %v, want %v", got, 0)
	}

	if got := Min[uint16](); got != 0 {
		t.Errorf("Min() = %v, want %v", got, 0)
	}

	if got := Min[uint32](); got != 0 {
		t.Errorf("Min() = %v, want %v", got, 0)
	}

	if got := Min[uint64](); got != 0 {
		t.Errorf("Min() = %v, want %v", got, 0)
	}

	if got := Min[float32](); got != 0 {
		t.Errorf("Min() = %v, want %v", got, 0)
	}

	if got := Min[float64](); got != 0 {
		t.Errorf("Min() = %v, want %v", got, 0)
	}
}
