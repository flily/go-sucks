package math

import (
	"testing"

	"cmp"
	"fmt"
	"math/rand/v2"
)

func testFuncNumberArray[T OrderedNumber](t *testing.T, f func([]T) T, exp T) {
	t.Helper()

	array := make([]T, 100)
	for i := range array {
		array[i] = T(i)
	}

	for range 100 {
		rand.Shuffle(len(array), func(i, j int) {
			array[i], array[j] = array[j], array[i]
		})

		got := f(array)
		if got != exp {
			t.Errorf("testFuncNumber[%T]() = %v, want %v", array[0], got, exp)
		}
	}
}

func testFuncNumberArgs[T OrderedNumber](t *testing.T, f func(...T) T, exp T) {
	t.Helper()

	array := make([]T, 100)
	for i := range array {
		array[i] = T(i)
	}

	for range 100 {
		rand.Shuffle(len(array), func(i, j int) {
			array[i], array[j] = array[j], array[i]
		})

		got := f(array...)
		if got != exp {
			t.Errorf("testFuncNumber() = %v, want %v", got, exp)
		}
	}
}

func testFuncStringArray[T ~string](t *testing.T, f func([]T) T, exp T) {
	t.Helper()

	array := make([]T, 100)
	for i := range array {
		array[i] = T(fmt.Sprintf("a%d", i))
	}

	for range 100 {
		rand.Shuffle(len(array), func(i, j int) {
			array[i], array[j] = array[j], array[i]
		})

		got := f(array)
		if got != exp {
			t.Errorf("testFuncStringArray() = %v, want %v", got, exp)
		}
	}
}

func testFuncCallArray[T cmp.Ordered](t *testing.T, f func([]T) T, exp T, args ...T) {
	t.Helper()

	got := f(args)
	if got != exp {
		t.Errorf("testFuncCall() = %v, want %v", got, exp)
	}
}

func testFuncCallArgs[T cmp.Ordered](t *testing.T, f func(...T) T, exp T, args ...T) {
	t.Helper()

	got := f(args...)
	if got != exp {
		t.Errorf("testFuncCall() = %v, want %v", got, exp)
	}
}

func TestMax(t *testing.T) {
	testFuncNumberArray(t, Max, 99) // test for int

	testFuncNumberArray[int8](t, Max, 99)
	testFuncNumberArray[int16](t, Max, 99)
	testFuncNumberArray[int32](t, Max, 99)
	testFuncNumberArray[int64](t, Max, 99)

	testFuncNumberArray[uint](t, Max, 99)
	testFuncNumberArray[uint8](t, Max, 99)
	testFuncNumberArray[uint16](t, Max, 99)
	testFuncNumberArray[uint32](t, Max, 99)
	testFuncNumberArray[uint64](t, Max, 99)

	testFuncNumberArray[float32](t, Max, 99)
	testFuncNumberArray[float64](t, Max, 99)

	testFuncNumberArray[uintptr](t, Max, 99)
	testFuncStringArray(t, Max, "a99")
}

func TestMaxWithNoArguments(t *testing.T) {
	testFuncCallArray(t, Max, 0) // test for int

	testFuncCallArray[int8](t, Max, 0)
	testFuncCallArray[int16](t, Max, 0)
	testFuncCallArray[int32](t, Max, 0)
	testFuncCallArray[int64](t, Max, 0)

	testFuncCallArray[uint](t, Max, 0)
	testFuncCallArray[uint8](t, Max, 0)
	testFuncCallArray[uint16](t, Max, 0)
	testFuncCallArray[uint32](t, Max, 0)
	testFuncCallArray[uint64](t, Max, 0)

	testFuncCallArray[float32](t, Max, 0)
	testFuncCallArray[float64](t, Max, 0)

	testFuncCallArray[uintptr](t, Max, 0)
	testFuncCallArray(t, Max, "")
}

func TestMin(t *testing.T) {
	testFuncNumberArray(t, Min, 0) // test for int

	testFuncNumberArray[int8](t, Min, 0)
	testFuncNumberArray[int16](t, Min, 0)
	testFuncNumberArray[int32](t, Min, 0)
	testFuncNumberArray[int64](t, Min, 0)

	testFuncNumberArray[uint](t, Min, 0)
	testFuncNumberArray[uint8](t, Min, 0)
	testFuncNumberArray[uint16](t, Min, 0)
	testFuncNumberArray[uint32](t, Min, 0)
	testFuncNumberArray[uint64](t, Min, 0)

	testFuncNumberArray[float32](t, Min, 0)
	testFuncNumberArray[float64](t, Min, 0)

	testFuncNumberArray[uintptr](t, Min, 0)
	testFuncStringArray(t, Min, "a0")
}

func TestMinWithNoArguments(t *testing.T) {
	testFuncCallArray(t, Min, 0) // test for int

	testFuncCallArray[int8](t, Min, 0)
	testFuncCallArray[int16](t, Min, 0)
	testFuncCallArray[int32](t, Min, 0)
	testFuncCallArray[int64](t, Min, 0)

	testFuncCallArray[uint](t, Min, 0)
	testFuncCallArray[uint8](t, Min, 0)
	testFuncCallArray[uint16](t, Min, 0)
	testFuncCallArray[uint32](t, Min, 0)
	testFuncCallArray[uint64](t, Min, 0)

	testFuncCallArray[float32](t, Min, 0)
	testFuncCallArray[float64](t, Min, 0)

	testFuncCallArray[uintptr](t, Min, 0)
	testFuncCallArray(t, Min, "")
}

func TestMaxOf(t *testing.T) {
	testFuncNumberArgs(t, MaxOf, 99) // test for int

	testFuncNumberArgs[int8](t, MaxOf, 99)
	testFuncNumberArgs[int16](t, MaxOf, 99)
	testFuncNumberArgs[int32](t, MaxOf, 99)
	testFuncNumberArgs[int64](t, MaxOf, 99)

	testFuncNumberArgs[uint](t, MaxOf, 99)
	testFuncNumberArgs[uint8](t, MaxOf, 99)
	testFuncNumberArgs[uint16](t, MaxOf, 99)
	testFuncNumberArgs[uint32](t, MaxOf, 99)
	testFuncNumberArgs[uint64](t, MaxOf, 99)

	testFuncNumberArgs[float32](t, MaxOf, 99)
	testFuncNumberArgs[float64](t, MaxOf, 99)

	testFuncNumberArgs[uintptr](t, MaxOf, 99)
}

func TestMaxOfWithNoArguments(t *testing.T) {
	testFuncCallArgs(t, MaxOf, 0) // test for int

	testFuncCallArgs[int8](t, MaxOf, 0)
	testFuncCallArgs[int16](t, MaxOf, 0)
	testFuncCallArgs[int32](t, MaxOf, 0)
	testFuncCallArgs[int64](t, MaxOf, 0)

	testFuncCallArgs[uint](t, MaxOf, 0)
	testFuncCallArgs[uint8](t, MaxOf, 0)
	testFuncCallArgs[uint16](t, MaxOf, 0)
	testFuncCallArgs[uint32](t, MaxOf, 0)
	testFuncCallArgs[uint64](t, MaxOf, 0)

	testFuncCallArgs[float32](t, MaxOf, 0)
	testFuncCallArgs[float64](t, MaxOf, 0)

	testFuncCallArgs[uintptr](t, MaxOf, 0)
	testFuncCallArgs(t, MaxOf, "")
}

func TestMinOf(t *testing.T) {
	testFuncNumberArgs(t, MinOf, 0) // test for int

	testFuncNumberArgs[int8](t, MinOf, 0)
	testFuncNumberArgs[int16](t, MinOf, 0)
	testFuncNumberArgs[int32](t, MinOf, 0)
	testFuncNumberArgs[int64](t, MinOf, 0)

	testFuncNumberArgs[uint](t, MinOf, 0)
	testFuncNumberArgs[uint8](t, MinOf, 0)
	testFuncNumberArgs[uint16](t, MinOf, 0)
	testFuncNumberArgs[uint32](t, MinOf, 0)
	testFuncNumberArgs[uint64](t, MinOf, 0)

	testFuncNumberArgs[float32](t, MinOf, 0)
	testFuncNumberArgs[float64](t, MinOf, 0)

	testFuncNumberArgs[uintptr](t, MinOf, 0)
}

func TestMinOfWithNoArguments(t *testing.T) {
	testFuncCallArgs(t, MinOf, 0)
	testFuncCallArgs[int8](t, MinOf, 0)
	testFuncCallArgs[int16](t, MinOf, 0)
	testFuncCallArgs[int32](t, MinOf, 0)
	testFuncCallArgs[int64](t, MinOf, 0)

	testFuncCallArgs[uint](t, MinOf, 0)
	testFuncCallArgs[uint8](t, MinOf, 0)
	testFuncCallArgs[uint16](t, MinOf, 0)
	testFuncCallArgs[uint32](t, MinOf, 0)
	testFuncCallArgs[uint64](t, MinOf, 0)

	testFuncCallArgs[float32](t, MinOf, 0)
	testFuncCallArgs[float64](t, MinOf, 0)

	testFuncCallArgs[uintptr](t, MinOf, 0)
	testFuncCallArgs(t, MinOf, "")
}
