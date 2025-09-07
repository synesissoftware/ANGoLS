// Copyright 2019-2025 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 1st March 2019
 * Updated: 24th February 2025
 */

package slices

// /////////////////////////////////////////////////////////////////////////
// Select*

func SelectSliceOfInt(input_slice []int, selector func(index int, input_item int) (bool, error)) ([]int, error) {

	input_len := len(input_slice)
	result := make([]int, input_len)
	result_len := 0

	for i, v := range input_slice {

		selected, err := selector(i, v)
		if err != nil {

			return nil, err
		}

		if selected {

			result[result_len] = v
			result_len++
		}
	}

	return result[0:result_len], nil
}

func SelectSliceOfUInt(input_slice []uint, selector func(index int, input_item uint) (bool, error)) ([]uint, error) {

	input_len := len(input_slice)
	result := make([]uint, input_len)
	result_len := 0

	for i, v := range input_slice {

		selected, err := selector(i, v)
		if err != nil {

			return nil, err
		}

		if selected {

			result[result_len] = v
			result_len++
		}
	}

	return result[0:result_len], nil
}

func SelectSliceOfUInteger[N int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | uintptr](input_slice []N, selector func(index int, input_item N) (bool, error)) ([]N, error) {

	input_len := len(input_slice)
	result := make([]N, input_len)
	result_len := 0

	for i, v := range input_slice {

		selected, err := selector(i, v)
		if err != nil {

			return nil, err
		}

		if selected {

			result[result_len] = v
			result_len++
		}
	}

	return result[0:result_len], nil
}

func SelectSliceOfString(input_slice []string, selector func(index int, input_item string) (bool, error)) ([]string, error) {

	input_len := len(input_slice)
	result := make([]string, input_len)
	result_len := 0

	for i, v := range input_slice {

		selected, err := selector(i, v)
		if err != nil {

			return nil, err
		}

		if selected {

			result[result_len] = v
			result_len++
		}
	}

	return result[0:result_len], nil
}

/* ///////////////////////////// end of file //////////////////////////// */
