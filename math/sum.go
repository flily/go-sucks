package math

// Sum calculates the sum of a slice of numbers.
func Sum[T OrderedNumber](nums []T) T {
	var sum T
	for _, n := range nums {
		sum += n
	}
	return sum
}

// SumOf calculates the sum of the given numbers.
func SumOf[T OrderedNumber](nums ...T) T {
	return Sum(nums)
}
