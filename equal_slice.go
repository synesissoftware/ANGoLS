/* /////////////////////////////////////////////////////////////////////////
 * File:        equal_slice.go
 *
 * Purpose:     EqualSlice*() functions
 *
 * Created:     1st March 2019
 * Updated:     11th March 2019
 *
 * Home:        http://github.com/synesissoftware/ANGOLS
 *
 * Copyright (c) 2019, Matthew Wilson and Synesis Software
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 * - Redistributions of source code must retain the above copyright notice,
 *   this list of conditions and the following disclaimer.
 * - Redistributions in binary form must reproduce the above copyright
 *   notice, this list of conditions and the following disclaimer in the
 *   documentation and/or other materials provided with the distribution.
 * - Neither the names of Matthew Wilson, Synesis Software nor
 *   the names of any contributors may be used to endorse or promote products
 *   derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS
 * IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO,
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
 * PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR
 * CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
 * PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
 * LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
 * NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 * ////////////////////////////////////////////////////////////////////// */

package angols

import (

	"reflect"
)

// /////////////////////////////////////////////////////////////////////////
// EqualSlice*()

func EqualSliceOfInt(lhs, rhs []int) bool {

	len_l	:=	len(lhs)
	len_r	:=	len(rhs)

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

	len_l	:=	len(lhs)
	len_r	:=	len(rhs)

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

	len_l	:=	len(lhs)
	len_r	:=	len(rhs)

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

	len_l	:=	len(lhs)
	len_r	:=	len(rhs)

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

	lhs_t	:=	reflect.TypeOf(lhs)
	rhs_t	:=	reflect.TypeOf(rhs)

	// check both are slices

	if reflect.Slice != lhs_t.Kind() {

		return false
	}

	if reflect.Slice != rhs_t.Kind() {

		return false
	}

	// check element type

	lhs_k	:=	lhs_t.Elem()
	rhs_k	:=	rhs_t.Elem()

	if lhs_k != rhs_k {

		return false
	} else {

		lhs_v	:=	reflect.ValueOf(lhs)
		rhs_v	:=	reflect.ValueOf(rhs)

		lhs_n	:=	lhs_v.Len()
		rhs_n	:=	rhs_v.Len()

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


