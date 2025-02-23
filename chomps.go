/* /////////////////////////////////////////////////////////////////////////
 * File:        chomps.go
 *
 * Purpose:     string chomp functions
 *
 * Created:     30th March 2019
 * Updated:     30th March 2019
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
