package math

import (
	"cmp"
)

func Max[T cmp.Ordered](nums ...T) T {
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

func Min[T cmp.Ordered](nums ...T) T {
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
