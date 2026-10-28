package math

const (
	SingleFloatExponentBias           = 127  // 0x7f
	DoubleFloatExponentBias           = 1023 // 0x3ff
	FloatBitsSingleZeroPos     uint32 = 0x00000000
	FloatBitsSingleZeroNeg     uint32 = 0x80000000
	FloatBitsSingleInfinityPos uint32 = 0x7f800000
	FloatBitsSingleInfinityNeg uint32 = 0xff800000
	FloatBitsSingleNaN         uint32 = 0x7fc00001
	FloatBitsDoubleZeroPos     uint64 = 0x0000000000000000
	FloatBitsDoubleZeroNeg     uint64 = 0x8000000000000000
	FloatBitsDoubleInfinityPos uint64 = 0x7ff0000000000000
	FloatBitsDoubleInfinityNeg uint64 = 0xfff0000000000000
	FloatBitsDoubleNaN         uint64 = 0x7ff8000000000001
)

// IEEE 754 floating-point number
//
// Single precision (32-bit)
// +------+----------+----------+
// | sign | exponent | fraction |
// +------+----------+----------+
// |   1  |     8    |    23    |
// +------+----------+----------+
//
// Double precision (64-bit)
// +------+----------+----------+
// | sign | exponent | fraction |
// +------+----------+----------+
// |   1  |    11    |    52    |
// +------+----------+----------+

func getSingleFloatBits(v float32) (bool, int16, uint64) {
	u32 := float32ToUint32(v)

	sign := (u32 >> 31) != 0
	exp := int16((u32 >> 23) & 0xff)
	fraction := uint64(u32 & 0x7f_ffff)

	exponent := exp - SingleFloatExponentBias
	return sign, exponent, fraction
}

func makeSingleFloatBits(sign bool, exponent int16, fraction uint64) float32 {
	signBit := uint32(0)
	if sign {
		signBit = 1
	}

	signBit <<= 31
	exp := uint32(exponent+SingleFloatExponentBias) << 23
	return uint32ToFloat32(signBit | exp | uint32(fraction))
}

func getDoubleFloatBits(v float64) (bool, int16, uint64) {
	u64 := float64ToUint64(v)

	sign := (u64 >> 63) != 0
	exp := int16((u64 >> 52) & 0x7ff)
	fraction := uint64(u64 & 0xf_ffff_ffff_ffff)

	exponent := exp - DoubleFloatExponentBias
	return sign, exponent, fraction
}

func makeDoubleFloatBits(sign bool, exponent int16, fraction uint64) float64 {
	signBit := uint64(0)
	if sign {
		signBit = 1
	}

	signBit <<= 63
	exp := uint64(exponent+DoubleFloatExponentBias) << 52
	return uint64ToFloat64(signBit | exp | fraction)
}

// GetFloatBits returns the sign, exponent, and fraction parts of a floating-point number.
// The sign is true for positive numbers, false for negative numbers.
// The exponent is the unbiased exponent.
// The fraction is the raw fraction bits.
func GetFloatBits[T FloatNumber](v T) (bool, int16, uint64) {
	sign := false
	exponent, fraction := int16(0), uint64(0)

	switch any(v).(type) {
	case float32:
		sign, exponent, fraction = getSingleFloatBits(float32(v))

	case float64:
		sign, exponent, fraction = getDoubleFloatBits(float64(v))
	}

	return sign, exponent, fraction
}

// MakeFloat constructs a floating-point number from its sign, exponent, and fraction parts.
// The sign should be true for positive numbers, false for negative numbers.
// The exponent should be the unbiased exponent.
// The fraction should be the raw fraction bits.
func MakeFloat[T FloatNumber](sign bool, exponent int16, fraction uint64) T {
	var r T
	switch any(r).(type) {
	case float32:
		v := makeSingleFloatBits(sign, exponent, fraction)
		r = T(v)

	case float64:
		v := makeDoubleFloatBits(sign, exponent, fraction)
		r = T(v)
	}

	return r
}

func floatSingleInfinity(positive bool) float32 {
	if positive {
		return uint32ToFloat32(FloatBitsSingleInfinityPos)
	}
	return uint32ToFloat32(FloatBitsSingleInfinityNeg)
}

func floatDoubleInfinity(positive bool) float64 {
	if positive {
		return uint64ToFloat64(FloatBitsDoubleInfinityPos)
	}
	return uint64ToFloat64(FloatBitsDoubleInfinityNeg)
}

// Inf returns positive or negative infinity for the specified floating-point type.
func Inf[T FloatNumber](positive bool) T {
	var r T
	switch any(r).(type) {
	case float32:
		v := floatSingleInfinity(positive)
		r = T(v)

	case float64:
		v := floatDoubleInfinity(positive)
		r = T(v)
	}

	return r
}

// InfPositive returns positive infinity for the specified floating-point type.
func InfPositive[T FloatNumber]() T {
	return Inf[T](true)
}

// InfNegative returns negative infinity for the specified floating-point type.
func InfNegative[T FloatNumber]() T {
	return Inf[T](false)
}
