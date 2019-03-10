
package angols_test

import (

	"github.com/synesissoftware/ANGoLS"

	"strings"
	"testing"
)

const (

	int_equal_iterations	=	100000
	int_equal_size			=	100
)


// Tests

func Test_Collect_Array_1_ints_to_ints(t *testing.T) {

	fn := func(input interface{}) (interface{}, error) {

		if i, ok := input.(int); ok {

			if i < 0 {

				return -i, nil
			}
		}

		return input, nil
	}

	ints	:=	[]int { 1, 2, 3 }

	r, err	:=	angols.CollectSlice(ints, fn)
	if err != nil {

		t.Errorf("Collect() failed: %v", err)
	} else {

		if si, ok := r.([]interface{}); !ok {

			t.Errorf("Collect() failed: result is not of type []interface{}\n")
		} else {

			for ix, v := range si {

				if i, ok := v.(int); !ok {

					t.Errorf("Element (%T)%v of non-int type at index %d\n", v, v, ix)
				} else {

					if ints[ix] != i {

						t.Errorf("Element (%T)%v at index %d is different from expected value %d\n", v, v, ix, ints[ix])
					}
				}
			}
		}
	}
}

func Test_CollectSliceOfInt_1(t *testing.T) {

	fn := func(input int) int {

		if input < 0 {

			return -input
		}

		return input
	}

	input_slice	:=	[]int { -1, 0, -2, 3 }
	expected	:=	[]int {  1, 0,  2, 3 }
	actual		:=	angols.CollectSliceOfInt(input_slice, fn)

	if !angols.EqualSliceOfInt(expected, actual) {

		t.Errorf("actual value '%v' does not equal expected value '%v'", actual, expected)
	}

	if !angols.EqualSlice(expected, actual) {

		t.Errorf("actual value '%v' does not equal expected value '%v'", actual, expected)
	}
}

func Test_CollectSliceOfFloat64_1(t *testing.T) {

	fn := func(input float64) float64 {

		if input < 0 {

			return -input
		}

		return input
	}

	input_slice	:=	[]float64 { -1.1, 0, -2.1, 3.1 }
	expected	:=	[]float64 {  1.1, 0,  2.1, 3.1 }
	actual		:=	angols.CollectSliceOfFloat64(input_slice, fn)

	if !angols.EqualSliceOfFloat64(expected, actual) {

		t.Errorf("actual value '%v' does not equal expected value '%v'", actual, expected)
	}

	if !angols.EqualSlice(expected, actual) {

		t.Errorf("actual value '%v' does not equal expected value '%v'", actual, expected)
	}
}

func Test_CollectSliceOfString_1(t *testing.T) {

	input_slice	:=	[]string { "abc", "", "def", "G" }
	expected	:=	[]string { "ABC", "", "DEF", "G" }
	actual		:=	angols.CollectSliceOfString(input_slice, func(input string) string { return strings.ToUpper(input) })

	if !angols.EqualSliceOfString(expected, actual) {

		t.Errorf("actual value '%v' does not equal expected value '%v'", actual, expected)
	}

	if !angols.EqualSlice(expected, actual) {

		t.Errorf("actual value '%v' does not equal expected value '%v'", actual, expected)
	}
}


// Benchmarks

func Benchmark_equal_ints_by_EqualSliceOfInt(b *testing.B) {

	fn			:=	func(index int) (int, error) { return index, nil }
	ints_1, _	:=	angols.GenerateSliceOfInt(int_equal_size, fn)
	ints_2, _	:=	angols.GenerateSliceOfInt(int_equal_size, fn)

	for i := 0; i != int_equal_iterations; i++ {

		_ = angols.EqualSliceOfInt(ints_1, ints_2)
	}
}

func Benchmark_equal_ints_by_EqualSlice(b *testing.B) {

	fn			:=	func(index int) (int, error) { return index, nil }
	ints_1, _	:=	angols.GenerateSliceOfInt(int_equal_size, fn)
	ints_2, _	:=	angols.GenerateSliceOfInt(int_equal_size, fn)

	for i := 0; i != int_equal_iterations; i++ {

		_ = angols.EqualSlice(ints_1, ints_2)
	}
}

func Benchmark_equal_uints_by_EqualSlice(b *testing.B) {

	fn			:=	func(index int) (uint, error) { return uint(index), nil }
	ints_1, _	:=	angols.GenerateSliceOfUInt(int_equal_size, fn)
	ints_2, _	:=	angols.GenerateSliceOfUInt(int_equal_size, fn)

	for i := 0; i != int_equal_iterations; i++ {

		_ = angols.EqualSlice(ints_1, ints_2)
	}
}

