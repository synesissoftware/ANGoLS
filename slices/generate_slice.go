// Copyright 2019-2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 1st March 2019
 * Updated: 15th February 2026
 */

package slices

type skipOneElement struct {
	s string
}

func (e *skipOneElement) Error() string {
	return e.s
}

type skipRemainingElements struct {
	s string
}

func (e *skipRemainingElements) Error() string {
	return e.s
}

// Causes generation of a single element to be skipped.
var SkipOneElement = &skipOneElement{s: "Skip one element"}

// Causes generation of all remaining elements to be skipped.
var SkipRemainingElements = &skipRemainingElements{s: "Skip remaining elements"}

// Creates a slice of a given size and populates its values with the given
// generator (which may be nil).
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

// Creates a slice of a given size and populates its values with the given
// generator (which may be nil).
func GenerateSliceOfUint(size int, generator func(index int) (result uint, err error)) (result []uint, err error) {

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

// Creates a slice of a given size and populates its values with the given
// generator (which may be nil).
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
