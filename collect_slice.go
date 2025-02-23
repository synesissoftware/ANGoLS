/* /////////////////////////////////////////////////////////////////////////
 * File:        collect_slice.go
 *
 * Purpose:     CollectSlice*() functions
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
	"fmt"
	"reflect"
)

// /////////////////////////////////////////////////////////////////////////
// Collect*

// This function maps an input slice of arbitrary type to an output slice
// of type []interface{}
func CollectSlice(input_slice interface{}, fn func(input_item interface{}) (interface{}, error)) (interface{}, error) {

	sl_t := reflect.TypeOf(input_slice)

	if reflect.Slice != sl_t.Kind() {

		msg := fmt.Sprintf("CollectSlice() called with input_slice of type %T; slice required", input_slice)

		panic(msg)
	}

	sl_v := reflect.ValueOf(input_slice)
	len := sl_v.Len()

	result := make([]interface{}, len)

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

/* ///////////////////////////// end of file //////////////////////////// */
