
package angols_test

import (

	"github.com/synesissoftware/ANGoLS"

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

