// Copyright 2019-2026 Harold Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings_test

import (
	strings "github.com/synesissoftware/ANGoLS/strings"
	stegol "github.com/synesissoftware/STEGoL"

	stdstrings "strings"
	"testing"
)

func Test_ASCIIToUpper(t *testing.T) {
	stegol.CheckStringEqual(t, "", strings.ASCIIToUpper(""))
	stegol.CheckStringEqual(t, "ABC", strings.ASCIIToUpper("abc"))
	stegol.CheckStringEqual(t, "ABC", strings.ASCIIToUpper("ABC"))
	stegol.CheckStringEqual(t, "HELLO, WORLD!", strings.ASCIIToUpper("Hello, World!"))
	stegol.CheckStringEqual(t, "12345", strings.ASCIIToUpper("12345"))
	stegol.CheckStringEqual(t, "!@#$%", strings.ASCIIToUpper("!@#$%"))
	stegol.CheckStringEqual(t, "AB1CD!", strings.ASCIIToUpper("aB1cD!"))

	// non-ASCII bytes are left unchanged
	stegol.CheckStringEqual(t, "HéLLO世界", strings.ASCIIToUpper("héllo世界"))
	stegol.CheckStringEqual(t, "ß", strings.ASCIIToUpper("ß"))
}

func Test_ASCIIToLower(t *testing.T) {
	stegol.CheckStringEqual(t, "", strings.ASCIIToLower(""))
	stegol.CheckStringEqual(t, "abc", strings.ASCIIToLower("abc"))
	stegol.CheckStringEqual(t, "abc", strings.ASCIIToLower("ABC"))
	stegol.CheckStringEqual(t, "hello, world!", strings.ASCIIToLower("Hello, World!"))
	stegol.CheckStringEqual(t, "12345", strings.ASCIIToLower("12345"))
	stegol.CheckStringEqual(t, "!@#$%", strings.ASCIIToLower("!@#$%"))
	stegol.CheckStringEqual(t, "ab1cd!", strings.ASCIIToLower("aB1cD!"))

	// non-ASCII bytes are left unchanged
	stegol.CheckStringEqual(t, "hÉllo世界", strings.ASCIIToLower("HÉLLO世界"))
	stegol.CheckStringEqual(t, "Σ", strings.ASCIIToLower("Σ"))
}

func Benchmark_ASCIIToUpper_Short(b *testing.B) {
	const s = "Hello, World!"
	for i := 0; i < b.N; i++ {
		_ = strings.ASCIIToUpper(s)
	}
}

func Benchmark_StringsToUpper(b *testing.B) {
	const s = "Hello, World!"
	for i := 0; i < b.N; i++ {
		_ = stdstrings.ToUpper(s)
	}
}

func Benchmark_ASCIIToLower_Short(b *testing.B) {
	const s = "Hello, World!"
	for i := 0; i < b.N; i++ {
		_ = strings.ASCIIToLower(s)
	}
}

func Benchmark_StringsToLower(b *testing.B) {
	const s = "Hello, World!"
	for i := 0; i < b.N; i++ {
		_ = stdstrings.ToLower(s)
	}
}
