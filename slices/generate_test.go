package slices_test

import (
	"github.com/synesissoftware/ANGoLS/slices"

	"testing"
)

func Test_GenerateSliceOfInt_1_nil_generator(t *testing.T) {

	req_len := 1001
	ints, err := slices.GenerateSliceOfInt(req_len, nil)
	if err != nil {

	} else {

		if len(ints) != req_len {

			t.Errorf("generated slice was required to be of length %d, but %d obtained", req_len, len(ints))
		}

		for i, v := range ints {

			if 0 != v {

				t.Errorf("element at index %d is expected to be %v, but is %v\n", i, 0, v)
			}
		}
	}
}

func Test_GenerateSliceOfInt_2_setting_all_to_const(t *testing.T) {

	req_len := 1001
	ints, err := slices.GenerateSliceOfInt(req_len, func(index int) (int, error) { return 1, nil })
	if err != nil {

	} else {

		if len(ints) != req_len {

			t.Errorf("generated slice was required to be of length %d, but %d obtained", req_len, len(ints))
		}

		for i, v := range ints {

			if 1 != v {

				t.Errorf("element at index %d is expected to be %v, but is %v\n", i, 1, v)
			}
		}
	}
}

func Test_GenerateSliceOfInt_3_doubling_index(t *testing.T) {

	req_len := 1001
	ints, err := slices.GenerateSliceOfInt(req_len, func(index int) (int, error) { return 2 * index, nil })
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
