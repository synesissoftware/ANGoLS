package angols_test

import (
	"github.com/synesissoftware/ANGoLS/slices"

	"strconv"
	"strings"
	"testing"
)

func Test_SelectSliceOfInt_1(t *testing.T) {

	input, err := slices.GenerateSliceOfInt(10, func(index int) (int, error) { return index, nil })
	if err != nil {

		t.Errorf("GenerateSliceOfInt() failed: %v\n", err)
	} else {

		actual, err := slices.SelectSliceOfInt(input, func(index int, value int) (bool, error) { return 0 == (value % 2), nil })
		if err != nil {

			t.Errorf("SelectSliceOfInt() failed: %v\n", err)
		} else {

			expected := []int{0, 2, 4, 6, 8}

			if !slices.EqualSliceOfInt(expected, actual) {

				t.Errorf("actual value '%v' does not equal expected value '%v'", actual, expected)
			}

			if !slices.EqualSlice(expected, actual) {

				t.Errorf("actual value '%v' does not equal expected value '%v'", actual, expected)
			}
		}
	}
}

func Test_SelectSliceOfUint_1(t *testing.T) {

	input, err := slices.GenerateSliceOfUint(10, func(index int) (uint, error) { return uint(index), nil })
	if err != nil {

		t.Errorf("GenerateSliceOfUint() failed: %v\n", err)
	} else {

		actual, err := slices.SelectSliceOfUint(input, func(index int, value uint) (bool, error) { return 0 == (value % 2), nil })
		if err != nil {

			t.Errorf("SelectSliceOfUint() failed: %v\n", err)
		} else {

			expected := []uint{0, 2, 4, 6, 8}

			if !slices.EqualSliceOfUint(expected, actual) {

				t.Errorf("actual value '%v' does not equal expected value '%v'", actual, expected)
			}

			if !slices.EqualSlice(expected, actual) {

				t.Errorf("actual value '%v' does not equal expected value '%v'", actual, expected)
			}
		}
	}
}

func Test_SelectSliceOfString_1(t *testing.T) {

	input, err := slices.GenerateSliceOfString(10, func(index int) (string, error) { return strconv.Itoa(index), nil })
	if err != nil {

		t.Errorf("GenerateSliceOfString() failed: %v\n", err)
	} else {

		actual, err := slices.SelectSliceOfString(input, func(index int, value string) (bool, error) { return strings.IndexAny(value[:1], "2457") >= 0, nil })
		if err != nil {

			t.Errorf("SelectSliceOfString() failed: %v\n", err)
		} else {

			expected := []string{"2", "4", "5", "7"}

			if !slices.EqualSliceOfString(expected, actual) {

				t.Errorf("actual value '%v' does not equal expected value '%v'", actual, expected)
			}

			if !slices.EqualSlice(expected, actual) {

				t.Errorf("actual value '%v' does not equal expected value '%v'", actual, expected)
			}
		}
	}
}
