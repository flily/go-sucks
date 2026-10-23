package math

import (
	"cmp"
)

// Max returns the maximum value among the given numbers.
// If the slice is empty, returns the zero value of the type.
func Max[S ~[]T, T cmp.Ordered](nums S) T {
	if len(nums) <= 0 {
		var zero T
		return zero
	}

	result := nums[0]
	for _, n := range nums {
		if n > result {
			result = n
		}
	}

	return result
}

// MaxOf returns the maximum value among the given numbers.
// If no arguments, returns the zero value of the type.
func MaxOf[T cmp.Ordered](nums ...T) T {
	if len(nums) <= 0 {
		var zero T
		return zero
	}

	result := nums[0]

	for _, n := range nums {
		if n > result {
			result = n
		}
	}

	return result
}

// Min returns the minimum value among the given numbers.
// If the slice is empty, returns the zero value of the type.
func Min[S ~[]T, T cmp.Ordered](nums S) T {
	if len(nums) <= 0 {
		var zero T
		return zero
	}

	result := nums[0]
	for _, n := range nums {
		if n < result {
			result = n
		}
	}

	return result
}

// MinOf returns the minimum value among the given numbers.
// If no arguments, returns the zero value of the type.
func MinOf[T cmp.Ordered](nums ...T) T {
	if len(nums) <= 0 {
		var zero T
		return zero
	}

	result := nums[0]

	for _, n := range nums {
		if n < result {
			result = n
		}
	}

	return result
}
