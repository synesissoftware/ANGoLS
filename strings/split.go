// Copyright 2025 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 20th August 2025
 * Updated: 20th August 2025
 */

package strings

import (
	"slices"
	"strings"
	"unicode/utf8"
)

// SplitAfterByte slices a string into all substings after each instance of
// the byte sep and returns a slice of those substrings.
//
// If s does not contain sep, SplitAfterByte returns a slice of length 1
// whose only element is s.
//
// It is equivalent to [SplitAfterByteN] with a count of -1.
func SplitAfterByte(s string, sep byte) []string {

	r := make([]string, 0)

	b := 0

	for i := 0; i < len(s); i++ {
		if sep == s[i] {

			r = append(r, s[b:i+1])

			b = i + 1
		}
	}

	if b != len(s) {
		r = append(r, s[b:])
	}

	return r
}

// SplitAfterByteN slices s into substrings after each instance of the byte
// sep and returns a slice of those substrings.
//
// ix determines the number of substrings to return:
//   - n > 0: at most n substrings; the last substring being the unsplit
//     remainder;
//   - n == 0: the result is nil (zero substrings);
//   - n < 0: all substrings.
func SplitAfterByteN(s string, sep byte, ix int) []string {

	if ix < 0 {
		return SplitAfterByte(s, sep)
	} else {

		r := make([]string, 0)

		if 0 != ix {

			b := 0

			for i := 0; i < len(s); i++ {

				if 1 == ix {
					r = append(r, s[i:])

					return r
				} else {

					if sep == s[i] {

						r = append(r, s[b:i+1])

						b = i + 1

						ix -= 1
					}
				}
			}

			if b != len(s) {
				r = append(r, s[b:])
			}
		}

		return r
	}
}

// SplitAfterRune slices a string into all substings after each instance of
// the rune sep and returns a slice of those substrings.
//
// If s does not contain sep, SplitAfterRune returns a slice of length 1
// whose only element is s.
//
// It is equivalent to [SplitAfterRuneN] with a count of -1.
func SplitAfterRune(s string, sep rune) []string {

	r := make([]string, 0)

	b := 0
	p := 0

	for _, c := range s {
		n := utf8.RuneLen(c)

		if c == sep {
			r = append(r, s[b:p+n])

			b = p + n
		}

		p += n
	}

	if b != len(s) {
		r = append(r, s[b:])
	}

	return r
}

// SplitAfterRuneN slices s into substrings after each instance of the rune
// sep and returns a slice of those substrings.
//
// ix determines the number of substrings to return:
//   - n > 0: at most n substrings; the last substring being the unsplit
//     remainder;
//   - n == 0: the result is nil (zero substrings);
//   - n < 0: all substrings.
func SplitAfterRuneN(s string, sep rune, ix int) []string {

	if ix < 0 {
		return SplitAfterRune(s, sep)
	} else {

		r := make([]string, 0)

		if 0 != ix {

			b := 0
			p := 0

			for _, c := range s {
				n := utf8.RuneLen(c)

				if 1 == ix {
					r = append(r, s[b:])

					return r
				} else {

					if c == sep {

						r = append(r, s[b:p+n])

						b = p + n

						ix -= 1
					}

					p += n
				}
			}

			if b != len(s) {
				r = append(r, s[b:])
			}
		}

		return r
	}
}

// SplitAfterAny slices a string into all substings after each instance of
// any of the runes in chars and returns a slice of those substrings.
//
// If s does not contain any of the runes in chars and chars is not empty,
// SplitAfterAny returns a slice of length 1 whose only element is s.
//
// It is equivalent to [SplitAfterAnyN] with a count of -1.
func SplitAfterAny(s, chars string) []string {

	r := make([]string, 0)

	b := 0
	p := 0

	for _, c := range s {
		n := utf8.RuneLen(c)

		if strings.ContainsRune(chars, c) {
			r = append(r, s[b:p+n])

			b = p + n
		}

		p += n
	}

	if b != len(s) {
		r = append(r, s[b:])
	}

	return r
}

// SplitAfterAnyN slices s into substrings after each instance of any of the
// runes in chars and returns a slice of those substrings.
//
// ix determines the number of substrings to return:
//   - n > 0: at most n substrings; the last substring being the unsplit
//     remainder;
//   - n == 0: the result is nil (zero substrings);
//   - n < 0: all substrings.
func SplitAfterAnyN(s, chars string, ix int) []string {

	if ix < 0 {
		return SplitAfterAny(s, chars)
	} else {

		r := make([]string, 0)

		if 0 != ix {

			b := 0
			p := 0

			for _, c := range s {
				n := utf8.RuneLen(c)

				if 1 == ix {
					r = append(r, s[b:])

					return r
				} else {

					if strings.ContainsRune(chars, c) {

						r = append(r, s[b:p+n])

						b = p + n

						ix -= 1
					}

					p += n
				}
			}

			if b != len(s) {
				r = append(r, s[b:])
			}
		}

		return r
	}
}

// SplitAfterAnyBytes slices a string into all substings after each
// instance of any of the bytes in seps and returns a slice of those
// substrings.
//
// If s does not contain any of the bytes in seps and seps is not empty,
// SplitAfterAnyBytes returns a slice of length 1 whose only element is s.
//
// It is equivalent to [SplitAfterAnyBytesN] with a count of -1.
func SplitAfterAnyBytes(s string, seps []byte) []string {

	r := make([]string, 0)

	b := 0

	for i := 0; i < len(s); i++ {
		if slices.Contains(seps, s[i]) {

			r = append(r, s[b:i+1])

			b = i + 1
		}
	}

	if b != len(s) {
		r = append(r, s[b:])
	}

	return r
}

// SplitAfterAnyBytesN slices a string into all substings after each
// instance of any of the bytes in seps and returns a slice of those
// substrings.
//
// ix determines the number of substrings to return:
//   - n > 0: at most n substrings; the last substring being the unsplit
//     remainder;
//   - n == 0: the result is nil (zero substrings);
//   - n < 0: all substrings.
func SplitAfterAnyBytesN(s string, seps []byte, ix int) []string {

	if ix < 0 {
		return SplitAfterAnyBytes(s, seps)
	} else {

		r := make([]string, 0)

		if 0 != ix {

			b := 0

			for i := 0; i < len(s); i++ {
				if 1 == ix {
					r = append(r, s[b:])

					return r
				} else {

					if slices.Contains(seps, s[i]) {

						r = append(r, s[b:i+1])

						b = i + 1

						ix -= 1
					}
				}
			}

			if b != len(s) {
				r = append(r, s[b:])
			}
		}

		return r
	}
}

// SplitAfterAnyRunes slices a string into all substings after each
// instance of any of the runes in seps and returns a slice of those
// substrings.
//
// If s does not contain any of the runes in seps and seps is not empty,
// SplitAfterAnyRunes returns a slice of length 1 whose only element is s.
//
// It is equivalent to [SplitAfterAnyN] with a count of -1.
func SplitAfterAnyRunes(s string, seps []rune) []string {

	r := make([]string, 0)

	b := 0
	p := 0

	for _, c := range s {
		n := utf8.RuneLen(c)

		if slices.Contains(seps, c) {
			r = append(r, s[b:p+n])

			b = p + n
		}

		p += n
	}

	if b != len(s) {
		r = append(r, s[b:])
	}

	return r
}

// SplitAfterAnyRunesN slices a string into all substings after each
// instance of any of the runes in seps and returns a slice of those
// substrings.
//
// ix determines the number of substrings to return:
//   - n > 0: at most n substrings; the last substring being the unsplit
//     remainder;
//   - n == 0: the result is nil (zero substrings);
//   - n < 0: all substrings.
func SplitAfterAnyRunesN(s string, seps []rune, ix int) []string {

	if ix < 0 {
		return SplitAfterAnyRunes(s, seps)
	} else {

		r := make([]string, 0)

		if 0 != ix {

			b := 0
			p := 0

			for _, c := range s {
				n := utf8.RuneLen(c)

				if 1 == ix {
					r = append(r, s[b:])

					return r
				} else {

					if slices.Contains(seps, c) {

						r = append(r, s[b:p+n])

						b = p + n

						ix -= 1
					}

					p += n
				}
			}

			if b != len(s) {
				r = append(r, s[b:])
			}
		}

		return r
	}
}
