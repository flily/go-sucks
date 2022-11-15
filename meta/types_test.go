package meta

import (
	"testing"
)

func TestIsNil(t *testing.T) {
	if !IsNil(nil) {
		t.Error("untyped nil is not nil")
	}

	if IsNil(0) {
		t.Error("int is nil")
	}

	var pointer *int
	if !IsNil(pointer) {
		t.Error("nil pointer is not nil")
	}

	var function func()
	if !IsNil(function) {
		t.Error("nil function is not nil")
	}

	var channel chan int
	if !IsNil(channel) {
		t.Error("nil channel is not nil")
	}

	var m map[int]int
	if !IsNil(m) {
		t.Error("nil map is not nil")
	}

	var nilSlice []int
	if !IsNil(nilSlice) {
		t.Error("nil slice is not nil")
	}

	emptySlice := []int{}
	if IsNil(emptySlice) {
		t.Error("empty slice is nil")
	}

	var array [0]int
	if IsNil(array) {
		t.Error("zero-size array is nil")
	}
}
