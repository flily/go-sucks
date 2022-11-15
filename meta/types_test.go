package meta

import (
	"io"
	"testing"
)

func TestIsNil(t *testing.T) {
	if u, n := IsNil(nil); !u || n {
		t.Errorf("IsNil(nil) = %v, %v", u, n)
	}

	if u, n := IsNil(0); u || n {
		t.Errorf("IsNil(0) = %v, %v", u, n)
	}

	var pointer *int
	if u, n := IsNil(pointer); u || !n {
		t.Errorf("IsNil(pointer) = %v, %v", u, n)
	}

	var channel chan int
	if u, n := IsNil(channel); u || !n {
		t.Errorf("IsNil(channel) = %v, %v", u, n)
	}

	var function func()
	if u, n := IsNil(function); u || !n {
		t.Errorf("IsNil(function) = %v, %v", u, n)
	}

	var mapValue map[int]int
	if u, n := IsNil(mapValue); u || !n {
		t.Errorf("IsNil(mapValue) = %v, %v", u, n)
	}

	var slice []int
	if u, n := IsNil(slice); u || !n {
		t.Errorf("IsNil(slice) = %v, %v", u, n)
	}

	// GO SUCKS
	var interfaceValue io.Writer
	if u, n := IsNil(interfaceValue); !u || n {
		t.Errorf("IsNil(interfaceValue) = %v, %v", u, n)
	}
}
