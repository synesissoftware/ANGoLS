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

func generateSliceOfBuiltIn[T bool | int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | uintptr | float32 | float64 | complex64 | complex128 | string](size int, generator func(index int) (result T, err error)) (result []T, err error) {

	result = make([]T, 0, size)

	var i int

	for {
		if size == 0 {
			break;
		} else {
			if size > 0 {
				size--
			}
		}

		value, err := generator(i)

		if err == nil {

			result = append(result, value)
		} else {

			if SkipOneElement == err {

			} else if SkipRemainingElements == err {

				break
			} else {

				return nil, err
			}
		}

		i++
	}

	return
}

// Creates a slice of []T up to a maximum size and populates its values with
// those created by the given generator.
//
// Parameters:
//   - `size` - maximum size if non-negative; otherwise size unlimited;
//   - `generator` - function that receives the given loop index, returning
//     a result to be added into the result-slice, or one of the specific
//     errors SkipOneElement (to skip a single element) or
//     SkipRemainingElements (to skip all remaining elements), or another
//     error to cause the operation to fail;
func GenerateSlice[T any](size int, generator func(index int) (result T, err error)) (result []T, err error) {

	result = make([]T, 0, size)

	var i int

	for {
		if size == 0 {
			break;
		} else {
			if size > 0 {
				size--
			}
		}

		value, err := generator(i)

		if err == nil {

			result = append(result, value)
		} else {

			if SkipOneElement == err {

			} else if SkipRemainingElements == err {

				break
			} else {

				return nil, err
			}
		}

		i++
	}

	return
}

// Creates a slice of []N up to a maximum size and populates its values with
// those created by the given generator.
//
// Parameters:
//   - `size` - maximum size if non-negative; otherwise size unlimited;
//   - `generator` - function that receives the given loop index, returning
//     a result to be added into the result-slice, or one of the specific
//     errors SkipOneElement (to skip a single element) or
//     SkipRemainingElements (to skip all remaining elements), or another
//     error to cause the operation to fail;
func GenerateSliceOfInteger[N int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | uintptr](size int, generator func(index int) (result N, err error)) (result []N, err error) {

	return generateSliceOfBuiltIn(size, generator)
}

// Creates a slice of []int up to a maximum size and populates its values
// with those created by the given generator.
//
// Parameters:
//   - `size` - maximum size if non-negative; otherwise size unlimited;
//   - `generator` - function that receives the given loop index, returning
//     a result to be added into the result-slice, or one of the specific
//     errors SkipOneElement (to skip a single element) or
//     SkipRemainingElements (to skip all remaining elements), or another
//     error to cause the operation to fail;
func GenerateSliceOfInt(size int, generator func(index int) (result int, err error)) (result []int, err error) {

	return GenerateSliceOfInteger(size, generator)
}

// Creates a slice of []uint up to a maximum size and populates its values
// with those created by the given generator.
//
// Parameters:
//   - `size` - maximum size if non-negative; otherwise size unlimited;
//   - `generator` - function that receives the given loop index, returning
//     a result to be added into the result-slice, or one of the specific
//     errors SkipOneElement (to skip a single element) or
//     SkipRemainingElements (to skip all remaining elements), or another
//     error to cause the operation to fail;
func GenerateSliceOfUint(size int, generator func(index int) (result uint, err error)) (result []uint, err error) {

	return GenerateSliceOfInteger(size, generator)
}

// Creates a slice of a given size and populates its values with the given
// generator (which may be nil).
func GenerateSliceOfString(size int, generator func(index int) (result string, err error)) (result []string, err error) {

	return generateSliceOfBuiltIn(size, generator)
}

/* ///////////////////////////// end of file //////////////////////////// */
