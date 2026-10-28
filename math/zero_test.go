package math

import (
	"testing"
)

func TestZeroFloat(t *testing.T) {
	zeroPosStd32 := float32(0.0)

	if zeroPosStd32 != Zero[float32](true) {
		t.Errorf("float32 +0.0 != +0.0")
	}

	if zeroPosStd32 != Zero[float32](false) {
		t.Errorf("float32 +0.0 != -0.0")
	}

	zeroPosStd64 := float64(0.0)

	if zeroPosStd64 != Zero[float64](true) {
		t.Errorf("float64 +0.0 != +0.0")
	}

	if zeroPosStd64 != Zero[float64](false) {
		t.Errorf("float64 +0.0 != -0.0")
	}
}
