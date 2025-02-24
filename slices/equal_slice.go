// Copyright 2019-2025 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 1st March 2019
 * Updated: 24th February 2025
 */

package slices

import (
	"reflect"
)

// /////////////////////////////////////////////////////////////////////////
// EqualSlice*()

func EqualSliceOfInt(lhs, rhs []int) bool {

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

func EqualSliceOfUInt(lhs, rhs []uint) bool {

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

func EqualSlice(lhs, rhs interface{}) bool {

	// Check for typed variants

	switch lhs_typed := lhs.(type) {

	case []int:

		if rhs_typed, ok := rhs.([]int); ok {

			return EqualSliceOfInt(lhs_typed, rhs_typed)
		} else {

			return false
		}
	case []float64:

		if rhs_typed, ok := rhs.([]float64); ok {

			return EqualSliceOfFloat64(lhs_typed, rhs_typed)
		} else {

			return false
		}
	case []string:

		if rhs_typed, ok := rhs.([]string); ok {

			return EqualSliceOfString(lhs_typed, rhs_typed)
		} else {

			return false
		}
	}

	// Generic comparison

	lhs_t := reflect.TypeOf(lhs)
	rhs_t := reflect.TypeOf(rhs)

	// check both are slices

	if reflect.Slice != lhs_t.Kind() {

		return false
	}

	if reflect.Slice != rhs_t.Kind() {

		return false
	}

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
