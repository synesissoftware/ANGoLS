// Copyright 2019-2025 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 30th March 2019
 * Updated: 24th February 2025
 */

package strings

// StringChomp() takes a single string and returns a chomped version of it,
// where chomping removes a single trailing '\n' character, a single
// trailing '\r' character, or a single trailing sequence of "\r\n"
func StringChomp(s string) string {

	switch l := len(s); l {

	case 0:

		break
	default:

		if '\r' == s[l-2] && '\n' == s[l-1] {

			s = s[0 : l-2]
			break
		}

		fallthrough
	case 1:

		switch s[l-1] {

		case '\r', '\n':

			s = s[0 : l-1]
		}
	}

	return s
}

// StringChompAll() takes a single string and returns a fully/repeatedly
// chomped version of, where full/repeated chomping removes all trailing
// '\r' and/or '\n' characters
func StringChompAll(s string) string {

out:
	for l := len(s); l > 0; {

		switch s[l-1] {

		case '\r', '\n':

			l--
			s = s[0:l]
		default:

			break out
		}
	}

	return s
}

/* ///////////////////////////// end of file //////////////////////////// */
