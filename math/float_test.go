package math

import (
	"testing"

	"math"
)

func TestFloatParts(t *testing.T) {
	{
		v := float32(0.0)
		sign, exponent, fraction := GetFloatBits(v)
		if sign != false || exponent != -127 || fraction != 0 {
			t.Errorf("float32(0.0) parts incorrect: sign=%v, exponent=%v, fraction=%v", sign, exponent, fraction)
		}
	}

	{
		v := float64(0.0)
		sign, exponent, fraction := GetFloatBits(v)
		if sign != false || exponent != -1023 || fraction != 0 {
			t.Errorf("float64(0.0) parts incorrect: sign=%v, exponent=%v, fraction=%v", sign, exponent, fraction)
		}
	}
}

func TestFloatMakeBits(t *testing.T) {
	{
		v1 := float32(3.1415926)
		sign, exponent, fraction := GetFloatBits(v1)

		v2 := MakeFloat[float32](sign, exponent, fraction)
		if v2 != v1 {
			t.Errorf("MakeFloat[float32] failed: got %v, want %v", v2, v1)
		}
	}

	{
		v1 := float32(-3.1415926)
		sign, exponent, fraction := GetFloatBits(v1)

		v2 := MakeFloat[float32](sign, exponent, fraction)
		if v2 != v1 {
			t.Errorf("MakeFloat[float32] failed: got %v, want %v", v2, v1)
		}
	}

	{
		v1 := float64(3.1415926)
		sign, exponent, fraction := GetFloatBits(v1)

		v2 := MakeFloat[float64](sign, exponent, fraction)
		if v2 != v1 {
			t.Errorf("MakeFloat[float64] failed: got %v, want %v", v2, v1)
		}
	}

	{
		v1 := float64(-3.1415926)
		sign, exponent, fraction := GetFloatBits(v1)

		v2 := MakeFloat[float64](sign, exponent, fraction)
		if v2 != v1 {
			t.Errorf("MakeFloat[float64] failed: got %v, want %v", v2, v1)
		}
	}
}

func TestInf(t *testing.T) {
	max32, max64 := float32(math.MaxFloat32), float64(math.MaxFloat64)

	if Inf[float32](true) <= max32 {
		t.Errorf("Inf[float32](true) is not greater than max32")
	}

	if Inf[float32](false) >= -max32 {
		t.Errorf("Inf[float32](false) is not less than -max32")
	}

	if Inf[float64](true) <= max64 {
		t.Errorf("Inf[float64](true) is not greater than max64")
	}

	if Inf[float64](false) >= -max64 {
		t.Errorf("Inf[float64](false) is not less than -max64")
	}

	if InfPositive[float32]() <= max32 {
		t.Errorf("InfPositive[float32]() is not greater than max32")
	}

	if InfNegative[float32]() >= -max32 {
		t.Errorf("InfNegative[float32]() is not less than -max32")
	}

	if InfPositive[float64]() <= max64 {
		t.Errorf("InfPositive[float64]() is not greater than max64")
	}

	if InfNegative[float64]() >= -max64 {
		t.Errorf("InfNegative[float64]() is not less than -max64")
	}
}
