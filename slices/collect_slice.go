// Copyright 2019-2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 1st March 2019
 * Updated: 7th February 2026
 */

package slices

// /////////////////////////////////////////////////////////////////////////
// Collect*

// This function maps an input slice of T[] to an output slice of []T.
func CollectSlice[T any](input_slice []T, collector func(index int, input_item *T) (T, error)) ([]T, error) {

	result := make([]T, len(input_slice))

	for i := 0; i != len(input_slice); i++ {

		if r, err := collector(i, &input_slice[i]); err != nil {
			return nil, err
		} else {

			result[i] = r
		}
	}

	return result, nil
}

// This function maps an input slice of []int to an output slice of []int.
func CollectSliceOfInt(input_slice []int, collector func(input_item int) int) ([]int) {

	return CollectSliceOfInteger(input_slice, collector)
}

// This function maps an input slice of []N to an output slice of []N, where
// N is any integer type.
func CollectSliceOfInteger[N int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | uintptr](input_slice []N, collector func(input_item N) N) (result_slice []N) {

	result_slice = make([]N, len(input_slice))

	for i, v := range input_slice {

		result := collector(v)

		result_slice[i] = result
	}

	return
}

// This function maps an input slice of []float64 to an output slice of
// []float64.
func CollectSliceOfFloat64(input_slice []float64, collector func(input_item float64) float64) (result_slice []float64) {

	result_slice = make([]float64, len(input_slice))

	for i, v := range input_slice {

		result := collector(v)

		result_slice[i] = result
	}

	return
}

// This function maps an input slice of []string to an output slice of
// []string.
func CollectSliceOfString(input_slice []string, collector func(input_item string) string) (result_slice []string) {

	result_slice = make([]string, len(input_slice))

	for i, s := range input_slice {

		result_slice[i] = collector(s)
	}

	return
}

// This function maps an input slice of []T to an output slice of []string.
func CollectSliceIntoStringSlice[T any](input_slice []T, collector func(input_item *T) (string, error)) ([]string, error) {

	result_slice := make([]string, len(input_slice))

	for i, t := range input_slice {

		s, err := collector(&t)
		if err != nil {
			return []string{}, err
		}

		result_slice[i] = s
	}

	return result_slice, nil
}

/* ///////////////////////////// end of file //////////////////////////// */
