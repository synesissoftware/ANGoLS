// Copyright 2019-2025 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 1st March 2019
 * Updated: 27th August 2025
 */

package slices

import (
	"fmt"
	"reflect"
)

// /////////////////////////////////////////////////////////////////////////
// Collect*

// This function maps an input slice of arbitrary type to an output slice
// of type []any
func CollectSlice(input_slice any, fn func(input_item any) (any, error)) (any, error) {

	sl_t := reflect.TypeOf(input_slice)

	if reflect.Slice != sl_t.Kind() {

		msg := fmt.Sprintf("CollectSlice() called with input_slice of type %T; slice required", input_slice)

		panic(msg)
	}

	sl_v := reflect.ValueOf(input_slice)
	len := sl_v.Len()

	result := make([]any, len)

	for i := 0; len != i; i++ {

		p := sl_v.Index(i)
		v := p.Interface()

		r, e := fn(v)
		if e != nil {

			return nil, e
		}

		result[i] = r
	}

	return result, nil
}

// CollectSliceOfInt
func CollectSliceOfInt(input_slice []int, fn func(input_item int) int) (result_slice []int) {

	result_slice = make([]int, len(input_slice))

	for i, v := range input_slice {

		result := fn(v)

		result_slice[i] = result
	}

	return
}

// CollectSliceOfInteger
func CollectSliceOfInteger[N int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | uintptr](input_slice []N, fn func(input_item N) N) (result_slice []N) {

	result_slice = make([]N, len(input_slice))

	for i, v := range input_slice {

		result := fn(v)

		result_slice[i] = result
	}

	return
}

// CollectSliceOfFloat64
func CollectSliceOfFloat64(input_slice []float64, fn func(input_item float64) float64) (result_slice []float64) {

	result_slice = make([]float64, len(input_slice))

	for i, v := range input_slice {

		result := fn(v)

		result_slice[i] = result
	}

	return
}

// CollectSliceOfString
func CollectSliceOfString(input_slice []string, fn func(input_item string) string) (result_slice []string) {

	result_slice = make([]string, len(input_slice))

	for i, s := range input_slice {

		result_slice[i] = fn(s)
	}

	return
}

func CollectSliceIntoStringSlice[T any](input_slice []T, fn func(input_item *T) (string, error)) ([]string, error) {

	result_slice := make([]string, len(input_slice))

	for i, s := range input_slice {

		a, err := fn(&s)
		if err != nil {
			return []string{}, err
		}

		result_slice[i] = a
	}

	return result_slice, nil
}

/* ///////////////////////////// end of file //////////////////////////// */
