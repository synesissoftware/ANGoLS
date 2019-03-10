
package angols_test

import (

	"github.com/synesissoftware/ANGoLS"

	"strconv"
	"strings"
	"testing"
)

func Test_SelectSliceOfInt_1(t *testing.T) {

	input, err	:=	angols.GenerateSliceOfInt(10, func(index int) (int, error) { return index, nil })
	if err != nil {

		t.Errorf("GenerateSliceOfInt() failed: %v\n", err)
	} else {

		actual, err	:=	angols.SelectSliceOfInt(input, func(index int, value int) (bool, error) { return 0 == (value % 2), nil })
		if err != nil {

			t.Errorf("SelectSliceOfInt() failed: %v\n", err)
		} else {

			expected	:=	[]int{ 0, 2, 4, 6, 8 }

			if !angols.EqualSliceOfInt(expected, actual) {

				t.Errorf("actual value '%v' does not equal expected value '%v'", actual, expected)
			}

			if !angols.EqualSlice(expected, actual) {

				t.Errorf("actual value '%v' does not equal expected value '%v'", actual, expected)
			}
		}
	}
}

func Test_SelectSliceOfUInt_1(t *testing.T) {

	input, err	:=	angols.GenerateSliceOfUInt(10, func(index int) (uint, error) { return uint(index), nil })
	if err != nil {

		t.Errorf("GenerateSliceOfUInt() failed: %v\n", err)
	} else {

		actual, err	:=	angols.SelectSliceOfUInt(input, func(index int, value uint) (bool, error) { return 0 == (value % 2), nil })
		if err != nil {

			t.Errorf("SelectSliceOfUInt() failed: %v\n", err)
		} else {

			expected	:=	[]uint{ 0, 2, 4, 6, 8 }

			if !angols.EqualSliceOfUInt(expected, actual) {

				t.Errorf("actual value '%v' does not equal expected value '%v'", actual, expected)
			}

			if !angols.EqualSlice(expected, actual) {

				t.Errorf("actual value '%v' does not equal expected value '%v'", actual, expected)
			}
		}
	}
}

func Test_SelectSliceOfString_1(t *testing.T) {

	input, err	:=	angols.GenerateSliceOfString(10, func(index int) (string, error) { return strconv.Itoa(index), nil })
	if err != nil {

		t.Errorf("GenerateSliceOfString() failed: %v\n", err)
	} else {

		actual, err	:=	angols.SelectSliceOfString(input, func(index int, value string) (bool, error) { return strings.IndexAny(value[:1], "2457") >= 0, nil })
		if err != nil {

			t.Errorf("SelectSliceOfString() failed: %v\n", err)
		} else {

			expected	:=	[]string{ "2", "4", "5", "7" }

			if !angols.EqualSliceOfString(expected, actual) {

				t.Errorf("actual value '%v' does not equal expected value '%v'", actual, expected)
			}

			if !angols.EqualSlice(expected, actual) {

				t.Errorf("actual value '%v' does not equal expected value '%v'", actual, expected)
			}
		}
	}
}

