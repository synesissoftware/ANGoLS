// Copyright 2019-2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 1st March 2019
 * Updated: 7th February 2026
 */

package slices

// /////////////////////////////////////////////////////////////////////////
// Select*

// Selects elements of an input slice of []int to an output slice of []int,
// according to the given selector function.
func SelectSliceOfInt(input_slice []int, selector func(index int, input_item int) (bool, error)) ([]int, error) {

	return SelectSliceOfInteger(input_slice, selector)
}

// Selects elements of an input slice of []uint to an output slice of
// []uint, according to the given selector function.
func SelectSliceOfUint(input_slice []uint, selector func(index int, input_item uint) (bool, error)) ([]uint, error) {

	return SelectSliceOfInteger(input_slice, selector)
}

// Selects elements of an input slice of []N to an output slice of []N,
// according to the given selector function.
func SelectSliceOfInteger[N int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | uintptr](input_slice []N, selector func(index int, input_item N) (bool, error)) ([]N, error) {

	r := make([]N, 0, len(input_slice))

	for i, v := range input_slice {

		if selected, err := selector(i, v); err != nil {

			return nil, err
		} else {

			if selected {

				r = append(r, v)
			}
		}
	}

	return r, nil
}

// Selects elements of an input slice of []string to an output slice of
// []string, according to the given selector function.
func SelectSliceOfString(input_slice []string, selector func(index int, input_item string) (bool, error)) ([]string, error) {

	r := make([]string, 0, len(input_slice))

	for i, s := range input_slice {

		if selected, err := selector(i, s); err != nil {

			return nil, err
		} else {

			if selected {

				r = append(r, s)
			}
		}
	}

	return r, nil
}

/* ///////////////////////////// end of file //////////////////////////// */
