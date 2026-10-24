package math

import (
	"testing"
)

func TestSum(t *testing.T) {
	testFuncNumberArray(t, Sum, 4950)

	testFuncNumberArray[int8](t, Sum, 86) // overflowed
	testFuncNumberArray[int16](t, Sum, 4950)
	testFuncNumberArray[int32](t, Sum, 4950)
	testFuncNumberArray[int64](t, Sum, 4950)

	testFuncNumberArray[uint](t, Sum, 4950)
	testFuncNumberArray[uint8](t, Sum, 86) // overflowed
	testFuncNumberArray[uint16](t, Sum, 4950)
	testFuncNumberArray[uint32](t, Sum, 4950)
	testFuncNumberArray[uint64](t, Sum, 4950)
}

func TestSumWithNoArguments(t *testing.T) {
	testFuncCallArray(t, Sum, 0)

	testFuncCallArray[int8](t, Sum, 0)
	testFuncCallArray[int16](t, Sum, 0)
	testFuncCallArray[int32](t, Sum, 0)
	testFuncCallArray[int64](t, Sum, 0)

	testFuncCallArray[uint](t, Sum, 0)
	testFuncCallArray[uint8](t, Sum, 0)
	testFuncCallArray[uint16](t, Sum, 0)
	testFuncCallArray[uint32](t, Sum, 0)
	testFuncCallArray[uint64](t, Sum, 0)
}

func TestSumOf(t *testing.T) {
	testFuncNumberArgs(t, SumOf, 4950)

	testFuncNumberArgs[int8](t, SumOf, 86) // overflowed
	testFuncNumberArgs[int16](t, SumOf, 4950)
	testFuncNumberArgs[int32](t, SumOf, 4950)
	testFuncNumberArgs[int64](t, SumOf, 4950)

	testFuncNumberArgs[uint](t, SumOf, 4950)
	testFuncNumberArgs[uint8](t, SumOf, 86) // overflowed
	testFuncNumberArgs[uint16](t, SumOf, 4950)
	testFuncNumberArgs[uint32](t, SumOf, 4950)
	testFuncNumberArgs[uint64](t, SumOf, 4950)
}

func TestSumOfWithNoArguments(t *testing.T) {
	testFuncCallArgs(t, SumOf, 0)

	testFuncCallArgs[int8](t, SumOf, 0)
	testFuncCallArgs[int16](t, SumOf, 0)
	testFuncCallArgs[int32](t, SumOf, 0)
	testFuncCallArgs[int64](t, SumOf, 0)

	testFuncCallArgs[uint](t, SumOf, 0)
	testFuncCallArgs[uint8](t, SumOf, 0)
	testFuncCallArgs[uint16](t, SumOf, 0)
	testFuncCallArgs[uint32](t, SumOf, 0)
	testFuncCallArgs[uint64](t, SumOf, 0)
}
