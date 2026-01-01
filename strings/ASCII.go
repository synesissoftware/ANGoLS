// Copyright 2019-2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 1st January 2026
 * Updated: 1st January 2026
 */

package strings

// asciiToUpper returns a copy of s with all lowercase ASCII letters converted
// to their uppercase equivalents. Non-ASCII bytes and non-lowercase letters
// are left unchanged.
func asciiToUpper(s string) string {
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if b[i] >= 'a' && b[i] <= 'z' {
			b[i] -= 32
		}
	}
	return string(b)
}

// asciiToUpper returns a copy of s with all lowercase ASCII letters converted
// to their lowercase equivalents. Non-ASCII bytes and non-lowercase letters
// are left unchanged.
func asciiToLower(s string) string {
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] += 32
		}
	}
	return string(b)
}

/* ///////////////////////////// end of file //////////////////////////// */
