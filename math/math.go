package math

// OrderedNumber represents an ordered number type.
// Includes integers, floating-point numbers, and uintptr.
type OrderedNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~uintptr
}

// FloatNumber represents a floating-point number type (float32 or float64).
type FloatNumber interface {
	~float32 | ~float64
}
