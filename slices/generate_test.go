package slices_test

import (
	"fmt"
	"strconv"

	"github.com/synesissoftware/ANGoLS/slices"

	"testing"
)

func Test_GenerateSliceOfInt_2_SETTING_ALL_TO_CONSTANT_VALUE(t *testing.T) {

	req_len := 101
	ints, err := slices.GenerateSliceOfInt(req_len, func(index int) (int, error) {
		return -123, nil
	})
	if err != nil {

	} else {

		if len(ints) != req_len {

			t.Errorf("generated slice was required to be of length %d, but %d obtained", req_len, len(ints))
		}

		for i, v := range ints {

			if -123 != v {

				t.Errorf("element at index %d is expected to be %v, but is %v\n", i, 1, v)
			}
		}
	}
}

func Test_GenerateSliceOfInt_3_DOUBLING_index(t *testing.T) {

	req_len := 1001
	ints, err := slices.GenerateSliceOfInt(req_len, func(index int) (int, error) {
		return 2 * index, nil
	})
	if err != nil {

	} else {

		if len(ints) != req_len {

			t.Errorf("generated slice was required to be of length %d, but %d obtained", req_len, len(ints))
		}

		for i, v := range ints {

			if 2*i != v {

				t.Errorf("element at index %d is expected to be %v, but is %v\n", i, 2*i, v)
			}
		}
	}
}

func Test_GenerateSliceOfInt_3_SkipOneElement_FOR_ODD_INDEXES(t *testing.T) {

	req_len := 100
	ints, err := slices.GenerateSliceOfInt(req_len, func(index int) (int, error) {

		if 1 == (index % 2) {
			return -1, slices.SkipOneElement
		}

		return index, nil
	})
	if err != nil {

	} else {

		if len(ints) != 50 {

			t.Errorf("generated slice was required to be of length %d, but %d obtained", req_len, len(ints))
		}

		for i, v := range ints {

			if 2*i != v {

				t.Errorf("element at index %d is expected to be %v, but is %v\n", i, 2*i, v)
			}
		}
	}
}

func Test_GenerateSliceOfInt_4_SkipAllElements_AFTER_50(t *testing.T) {

	req_len := 100
	ints, err := slices.GenerateSliceOfInt(req_len, func(index int) (int, error) {

		if index == 50 {
			return -1, slices.SkipRemainingElements
		}

		return index, nil
	})
	if err != nil {

	} else {

		if len(ints) != 50 {

			t.Errorf("generated slice was required to be of length %d, but %d obtained", req_len, len(ints))
		}

		for i, v := range ints {

			if i != v {

				t.Errorf("element at index %d is expected to be %v, but is %v\n", i, 2*i, v)
			}
		}
	}
}

func Test_GenerateSliceOfString_1_itoa_FROM_INDEX_EXCEPT_22(t *testing.T) {

	req_len := 40
	ints, err := slices.GenerateSliceOfString(req_len, func(index int) (s string, err error) {
		if index == 22 {
			return "", slices.SkipOneElement
		}

		s = "[" + strconv.Itoa(index) + "]"

		return
	})
	if err != nil {

	} else {

		if len(ints) != req_len-1 {

			t.Errorf("generated slice was required to be of length %d, but %d obtained", req_len, len(ints))
		}

		for i, v := range ints {

			n := i
			if i >= 22 {
				n++
			}

			if fmt.Sprintf("[%d]", n) != v {

				t.Errorf("element at index %d is expected to be %v, but is %v\n", i, 2*i, v)
			}
		}
	}
}
