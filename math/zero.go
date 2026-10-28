package math

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
