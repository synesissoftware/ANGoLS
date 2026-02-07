// Copyright 2019-2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 1st March 2019
 * Updated: 7th February 2026
 */

package slices

import (
	"reflect"
)

// /////////////////////////////////////////////////////////////////////////
// EqualSlice*()

// Indicates whether two []int slices have the same size, contents, and
// order.
func EqualSliceOfInt(lhs, rhs []int) bool {

	return EqualSliceOfInteger(lhs, rhs)
}

// Indicates whether two []uint slices have the same size, contents, and
// order.
func EqualSliceOfUint(lhs, rhs []uint) bool {

	return EqualSliceOfInteger(lhs, rhs)
}

// Indicates whether two []N slices have the same size, contents, and
// order.
func EqualSliceOfInteger[N int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | uintptr](lhs, rhs []N) bool {

	len_l := len(lhs)
	len_r := len(rhs)

	if len_l != len_r {

		return false
	} else {

		for i := 0; len_l != i; i++ {

			if lhs[i] != rhs[i] {

				return false
			}
		}

		return true
	}
}

// Indicates whether two []float64 slices have the same size, contents, and
// order.
func EqualSliceOfFloat64(lhs, rhs []float64) bool {

	len_l := len(lhs)
	len_r := len(rhs)

	if len_l != len_r {

		return false
	} else {

		for i := 0; len_l != i; i++ {

			if lhs[i] != rhs[i] {

				return false
			}
		}

		return true
	}
}

// Indicates whether two []string slices have the same size, contents, and
// order.
func EqualSliceOfString(lhs, rhs []string) bool {

	len_l := len(lhs)
	len_r := len(rhs)

	if len_l != len_r {

		return false
	} else {

		for i := 0; len_l != i; i++ {

			if lhs[i] != rhs[i] {

				return false
			}
		}

		return true
	}
}

func EqualSlice[T any](lhs, rhs []T) bool {

	// Check for typed variants

	switch lhs_typed := any(lhs).(type) {

	case []int8:

		rhs_typed, _ := any(rhs).([]int8)

		return EqualSliceOfInteger(lhs_typed, rhs_typed)
	case []int16:

		rhs_typed, _ := any(rhs).([]int16)

		return EqualSliceOfInteger(lhs_typed, rhs_typed)
	case []int32:

		rhs_typed, _ := any(rhs).([]int32)

		return EqualSliceOfInteger(lhs_typed, rhs_typed)
	case []int64:

		rhs_typed, _ := any(rhs).([]int64)

		return EqualSliceOfInteger(lhs_typed, rhs_typed)
	case []int:

		rhs_typed, _ := any(rhs).([]int)

		return EqualSliceOfInteger(lhs_typed, rhs_typed)
	case []uint8:

		rhs_typed, _ := any(rhs).([]uint8)

		return EqualSliceOfInteger(lhs_typed, rhs_typed)
	case []uint16:

		rhs_typed, _ := any(rhs).([]uint16)

		return EqualSliceOfInteger(lhs_typed, rhs_typed)
	case []uint32:

		rhs_typed, _ := any(rhs).([]uint32)

		return EqualSliceOfInteger(lhs_typed, rhs_typed)
	case []uint64:

		rhs_typed, _ := any(rhs).([]uint64)

		return EqualSliceOfInteger(lhs_typed, rhs_typed)
	case []uint:

		rhs_typed, _ := any(rhs).([]uint)

		return EqualSliceOfInteger(lhs_typed, rhs_typed)
	case []uintptr:

		rhs_typed, _ := any(rhs).([]uintptr)

		return EqualSliceOfInteger(lhs_typed, rhs_typed)
	case []float64:

		rhs_typed, _ := any(rhs).([]float64)

		return EqualSliceOfFloat64(lhs_typed, rhs_typed)
	case []string:

		rhs_typed, _ := any(rhs).([]string)

		return EqualSliceOfString(lhs_typed, rhs_typed)
	}

	// Generic comparison

	lhs_t := reflect.TypeOf(lhs)
	rhs_t := reflect.TypeOf(rhs)

	// check element type

	lhs_k := lhs_t.Elem()
	rhs_k := rhs_t.Elem()

	if lhs_k != rhs_k {

		return false
	} else {

		lhs_v := reflect.ValueOf(lhs)
		rhs_v := reflect.ValueOf(rhs)

		lhs_n := lhs_v.Len()
		rhs_n := rhs_v.Len()

		if lhs_n != rhs_n {

			return false
		}

		for i := 0; lhs_n != i; i++ {

			if lhs_v.Index(i).Interface() != rhs_v.Index(i).Interface() {

				return false
			}
		}

		return true
	}
}

/* ///////////////////////////// end of file //////////////////////////// */
