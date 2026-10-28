package math

// Zero returns the zero value for the specified ordered number type.
// When T is a floating-point type, the positive parameter determines the sign of zero.
// For integer types, the positive parameter is ignored.
func Zero[T OrderedNumber](positive bool) T {
	var r T
	switch any(r).(type) {
	case float32:
		if positive {
			r = T(uint32ToFloat32(FloatBitsSingleZeroPos))
		} else {
			r = T(uint32ToFloat32(FloatBitsSingleZeroNeg))
		}

	case float64:
		if positive {
			r = T(uint64ToFloat64(FloatBitsDoubleZeroPos))
		} else {
			r = T(uint64ToFloat64(FloatBitsDoubleZeroNeg))
		}
	}

	return r
}
