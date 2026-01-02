// Copyright 2019-2026 Harold Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 1st January 2026
 * Updated: 2nd January 2026
 */

package strings

// Returns a copy of s with all lowercase ASCII letters converted to their
// uppercase equivalents. Non-ASCII bytes and non-lowercase letters are left
// unchanged.
func ASCIIToUpper(s string) string {
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if b[i] >= 'a' && b[i] <= 'z' {
			b[i] -= 32
		}
	}
	return string(b)
}

// Returns a copy of s with all uppercase ASCII letters converted to their
// lowercase equivalents. Non-ASCII bytes and non-uppercase letters are left
// unchanged.
func ASCIIToLower(s string) string {
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] += 32
		}
	}
	return string(b)
}

/* ///////////////////////////// end of file //////////////////////////// */
