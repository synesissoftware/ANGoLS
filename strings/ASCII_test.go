// Copyright 2019-2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	//strings "github.com/synesissoftware/ANGoLS/strings"
	stegol "github.com/synesissoftware/STEGoL"

	stdstrings "strings"
	"testing"
)

func TestAsciiToUpper(t *testing.T) {
	stegol.CheckStringEqual(t, "", asciiToUpper(""))
	stegol.CheckStringEqual(t, "ABC", asciiToUpper("abc"))
	stegol.CheckStringEqual(t, "ABC", asciiToUpper("ABC"))
	stegol.CheckStringEqual(t, "HELLO, WORLD!", asciiToUpper("Hello, World!"))
	stegol.CheckStringEqual(t, "12345", asciiToUpper("12345"))
	stegol.CheckStringEqual(t, "!@#$%", asciiToUpper("!@#$%"))
	stegol.CheckStringEqual(t, "AB1CD!", asciiToUpper("aB1cD!"))

	// non-ASCII bytes are left unchanged
	stegol.CheckStringEqual(t, "HéLLO世界", asciiToUpper("héllo世界"))
	stegol.CheckStringEqual(t, "ß", asciiToUpper("ß"))
}

func TestAsciiToLower(t *testing.T) {
	stegol.CheckStringEqual(t, "", asciiToLower(""))
	stegol.CheckStringEqual(t, "abc", asciiToLower("abc"))
	stegol.CheckStringEqual(t, "abc", asciiToLower("ABC"))
	stegol.CheckStringEqual(t, "hello, world!", asciiToLower("Hello, World!"))
	stegol.CheckStringEqual(t, "12345", asciiToLower("12345"))
	stegol.CheckStringEqual(t, "!@#$%", asciiToLower("!@#$%"))
	stegol.CheckStringEqual(t, "ab1cd!", asciiToLower("aB1cD!"))

	// non-ASCII bytes are left unchanged
	stegol.CheckStringEqual(t, "hÉllo世界", asciiToLower("HÉLLO世界"))
	stegol.CheckStringEqual(t, "Σ", asciiToLower("Σ"))
}

func BenchmarkAsciiToUpper_Short(b *testing.B) {
	const s = "Hello, World!"
	for i := 0; i < b.N; i++ {
		_ = asciiToUpper(s)
	}
}

func BenchmarkStringsToUpper(b *testing.B) {
	const s = "Hello, World!"
	for i := 0; i < b.N; i++ {
		_ = stdstrings.ToUpper(s)
	}
}

func BenchmarkAsciiToLower_Short(b *testing.B) {
	const s = "Hello, World!"
	for i := 0; i < b.N; i++ {
		_ = asciiToLower(s)
	}
}

func BenchmarkStringsToLower(b *testing.B) {
	const s = "Hello, World!"
	for i := 0; i < b.N; i++ {
		_ = stdstrings.ToLower(s)
	}
}
