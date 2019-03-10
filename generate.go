/* /////////////////////////////////////////////////////////////////////////
 * File:        generate.go
 *
 * Purpose:     Generator functions
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
 * - Neither the names of Matthew Wilson, Sean Kelly, Synesis Software nor
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

	"errors"
)

var SkipOneElement			=	errors.New("Skip one element")
var SkipRemainingElements	=	errors.New("Skip remaining elements")

// GenerateSliceOfInt() creates a slice of a given size and populates its
// values with the given generator (which may be nil)
func GenerateSliceOfInt(size int, generator func(index int) (result int, err error)) (result []int, err error) {

	result	=	make([]int, size)

	if generator != nil {

		for i := 0; size != i; i++ {

			value, err := generator(i)

			if err == nil {

				result[i] = value
			} else {

				if SkipOneElement == err {

				} else if SkipRemainingElements == err {

					break
				} else {

					return nil, err
				}
			}
		}
	}

	return
}

// GenerateSliceOfUInt() creates a slice of a given size and populates its
// values with the given generator (which may be nil)
func GenerateSliceOfUInt(size int, generator func(index int) (result uint, err error)) (result []uint, err error) {

	result	=	make([]uint, size)

	if generator != nil {

		for i := 0; size != i; i++ {

			value, err := generator(i)

			if err == nil {

				result[i] = value
			} else {

				if SkipOneElement == err {

				} else if SkipRemainingElements == err {

					break
				} else {

					return nil, err
				}
			}
		}
	}

	return
}

/* ///////////////////////////// end of file //////////////////////////// */


