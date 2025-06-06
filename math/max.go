package math

// MaxInt64 returns the largest integer in a list of integers.
func MaxInt64(nums ...int64) int64 {
	result, found := int64(0), false

	for _, n := range nums {
		if !found || n > result {
			result = n
			found = true
		}
	}

	return result
}

func MinInt64(nums ...int64) int64 {
	result, found := int64(0), false

	for _, n := range nums {
		if !found || n < result {
			result = n
			found = true
		}
	}

	return result
}

func MaxUint64(nums ...uint64) uint64 {
	result, found := uint64(0), false

	for _, n := range nums {
		if !found || n > result {
			result = n
			found = true
		}
	}

	return result
}

func MinUint64(nums ...uint64) uint64 {
	result, found := uint64(0), false

	for _, n := range nums {
		if !found || n < result {
			result = n
			found = true
		}
	}

	return result
}

func MaxFloat64(nums ...float64) float64 {
	result, found := float64(0), false

	for _, n := range nums {
		if !found || n > result {
			result = n
			found = true
		}
	}

	return result
}

func MinFloat64(nums ...float64) float64 {
	result, found := float64(0), false

	for _, n := range nums {
		if !found || n < result {
			result = n
			found = true
		}
	}

	return result
}
