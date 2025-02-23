// Copyright 2019-2025 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 1st March 2019
 * Updated: 23rd February 2025
 */

package angols

import (
	"errors"
)

var SkipOneElement = errors.New("Skip one element")
var SkipRemainingElements = errors.New("Skip remaining elements")

// GenerateSliceOfInt() creates a slice of a given size and populates its
// values with the given generator (which may be nil)
func GenerateSliceOfInt(size int, generator func(index int) (result int, err error)) (result []int, err error) {

	result = make([]int, size)

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

	result = make([]uint, size)

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

// GenerateSliceOfString() creates a slice of a given size and populates its
// values with the given generator (which may be nil)
func GenerateSliceOfString(size int, generator func(index int) (result string, err error)) (result []string, err error) {

	result = make([]string, size)

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
