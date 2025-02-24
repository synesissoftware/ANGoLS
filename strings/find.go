// Copyright 2019-2025 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 24th February 2025
 * Updated: 24th February 2025
 */

package strings

import (
	"github.com/synesissoftware/ANGoLS/_std_strings"
)

/*
 * C++ construct                            | Go construct                             | ANGoLS construct                                     |
 * ---------------------------------------- | ---------------------------------------- | ---------------------------------------------------- |
 * #find(string const& sf)                  | strings.Index(s, sf string)              | -                                                    |
 * #find(char c)                            | strings.IndexByte(s string, c byte),
 *                                            strings.IndexByte(s string, c rune)      | -                                                    |
 * #find_first_of(string const& chars)      | strings.IndexAny(s, chars string)        | -                                                    |
 * #rfind(string const& sf)                 | strings.LastIndex(s, sf string)          | -                                                    |
 * #rfind(char c)                           | strings.LastIndexByte(s string, c byte)  | -                                                    |
 * #find_last_of(string const& chars)       | strings.LastIndexAny(s, chars string)    | -                                                    |
 *
 *                                          |                                          | strings.IndexAfter(s, sf, string, ix int)            |
 *                                          |                                          | strings.IndexAnyAfter(s, chars, string, ix int)      |
 *                                          |                                          | strings.IndexByteAfter(s string, c byte, ix int)     |
 *                                          |                                          | strings.IndexFuncAfter(s string, fn, ix int)         |
 *                                          |                                          |                                                      |
 *                                          |                                          | strings.IndexNotAnyAfter(s, chars, string, ix int)   |
 */

// Finds the index of the given substring in the given string, starting from
// the position after the given index. -1 is returned if the find is not
// successful.
//
// To search from the start of the string, specify the value -1 for the
// index. Any index value less than -1 will be treated as if -1 specified.
// Any index value greater than the size of the string will result in a
// return value of -1.
//
// The returned value reflects the position of the found substring relative
// to the start of the string, not from the index.
func IndexAfter(s string, sf string, ix int) int {

	if ix < -1 {
		ix = -1
	}

	off := ix + 1

	if off > len(s) {
		return -1
	}

	r := _std_strings.PRIVATE_Index(s[off:], sf)

	if r < 0 {
		return r
	} else {
		return r + off
	}
}

func IndexAnyAfter(s string, chars string, ix int) int {
	if ix < -1 {
		ix = -1
	}

	off := ix + 1

	if off > len(s) {
		return -1
	}

	r := _std_strings.PRIVATE_IndexAny(s[off:], chars)

	if r < 0 {
		return r
	} else {
		return r + off
	}
}

// T.B.C.
func IndexByteAfter(s string, c byte, ix int) int {

	if ix < -1 {
		ix = -1
	}

	off := ix + 1

	if off > len(s) {
		return -1
	}

	r := _std_strings.PRIVATE_IndexByte(s[off:], c)

	if r < 0 {
		return r
	} else {
		return r + off
	}
}

// T.B.C.
func IndexFuncAfter(s string, f func(rune) bool, ix int) int {

	if ix < -1 {
		ix = -1
	}

	off := ix + 1

	if off > len(s) {
		return -1
	}

	r := _std_strings.PRIVATE_IndexFunc(s[off:], f)

	if r < 0 {
		return r
	} else {

		return r + off
	}
}

// T.B.C.
func IndexNotAnyAfter(s string, chars string, ix int) int {
	if ix < -1 {
		ix = -1
	}

	off := ix + 1

	if off > len(s) {
		return -1
	}

	if 0 == len(chars) {
		return 0
	}

	for ix2, c := range s[off:] {
		if -1 == _std_strings.PRIVATE_IndexRune(chars, c) {
			return ix2 + off
		}
	}

	return -1
}
