/* /////////////////////////////////////////////////////////////////////////
 * File:        select_slice.go
 *
 * Purpose:     SelectSlice*() functions
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

)

// /////////////////////////////////////////////////////////////////////////
// Select*

func SelectSliceOfInt(input_slice []int, selector func(index int, input_item int) (bool, error)) ([]int, error) {

	input_len	:=	len(input_slice)
	result		:=	make([]int, input_len)
	result_len	:=	0

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

	input_len	:=	len(input_slice)
	result		:=	make([]uint, input_len)
	result_len	:=	0

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

	input_len	:=	len(input_slice)
	result		:=	make([]string, input_len)
	result_len	:=	0

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


