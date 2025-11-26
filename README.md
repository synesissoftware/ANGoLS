# ANGoLS <!-- omit in toc -->

[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![GitHub release](https://img.shields.io/github/v/release/synesissoftware/ANGoLS.svg)](https://github.com/synesissoftware/ANGoLS/releases/latest)
[![Last Commit](https://img.shields.io/github/last-commit/synesissoftware/ANGoLS)](https://github.com/synesissoftware/ANGoLS/commits/master)
[![Go](https://github.com/synesissoftware/ANGoLS/actions/workflows/go.yml/badge.svg)](https://github.com/synesissoftware/ANGoLS/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/synesissoftware/ANGoLS)](https://goreportcard.com/report/github.com/synesissoftware/ANGoLS)
[![Go Reference](https://pkg.go.dev/badge/github.com/synesissoftware/ANGoLS.svg)](https://pkg.go.dev/github.com/synesissoftware/ANGoLS)

**A**gorithms **N**ot in **Go** **L**anguage **S**tandard library


## Table of Contents <!-- omit in toc -->

- [Introduction](#introduction)
- [Installation](#installation)
- [Components](#components)
	- [Slices](#slices)
	- [Strings](#strings)
- [Examples](#examples)
- [Project Information](#project-information)
	- [Where to get help](#where-to-get-help)
	- [Contribution guidelines](#contribution-guidelines)
	- [Dependencies](#dependencies)
		- [Development/Testing Dependencies](#developmenttesting-dependencies)
	- [Related projects](#related-projects)
	- [License](#license)


## Introduction

As the name implies, there are (or were at the time of writing) commonly required operations/algorithms that are not provided in the language standard library. This library is a slowly growing collection of such things. (As of time of writing it does not contain any generic, but the plan is to sort that in the coming weeks.)

As an example, the standard library's **strings** package provides the following functions to split a `string` into `[]string`:

```Go
func Split(s, sep string) []string
func SplitN(s, sep string, n int) []string

func SplitAfter(s, sep string) []string
func SplitAfterN(s, sep string, n int) []string
```

**ANGoLS** provides a bunch of additional split functions that provide more expressiveness and/or functionality not possible with the standard library functions, including:

```Go
func SplitAfterByte(s string, sep byte) []string
func SplitAfterByteN(s string, sep byte, ix int) []string

func SplitAfterRune(s string, sep rune) []string
func SplitAfterRuneN(s string, sep rune, ix int) []string

func SplitAfterAny(s, chars string) []string
func SplitAfterAnyN(s, chars string, ix int) []string

func SplitAfterAnyBytes(s string, seps []byte) []string
func SplitAfterAnyBytesN(s string, seps []byte, ix int) []string

func SplitAfterAnyRunes(s string, seps []rune) []string
func SplitAfterAnyRunesN(s string, seps []rune, ix int) []string
```


## Installation

```bash
go get "github.com/synesissoftware/ANGoLS"
```


```Go
import (
  angols_slices "github.com/synesissoftware/ANGoLS/slices"
  angols_strings "github.com/synesissoftware/ANGoLS/strings"
)
```


## Components


### Slices

```Go
// in "github.com/synesissoftware/ANGoLS/slices"

func CollectSlice(input_slice any, fn func(input_item any) (any, error)) (any, error)

func CollectSliceOfInt(input_slice []int, fn func(input_item int) int) (result_slice []int)

func CollectSliceOfInteger[N int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | uintptr](input_slice []N, fn func(input_item N) N) (result_slice []N)

func CollectSliceOfFloat64(input_slice []float64, fn func(input_item float64) float64) (result_slice []float64)

func CollectSliceOfString(input_slice []string, fn func(input_item string) string) (result_slice []string)

func CollectSliceIntoStringSlice[T any](input_slice []T, fn func(input_item *T) (string, error)) ([]string, error)
```

```Go
// in "github.com/synesissoftware/ANGoLS/slices"

func EqualSliceOfInt(lhs, rhs []int) bool

func EqualSliceOfUInt(lhs, rhs []uint) bool

func EqualSliceOfInteger[N int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | uintptr](lhs, rhs []int) bool

func EqualSliceOfFloat64(lhs, rhs []float64) bool

func EqualSliceOfString(lhs, rhs []string) bool

func EqualSlice(lhs, rhs any) bool
```

```Go
// in "github.com/synesissoftware/ANGoLS/slices"

// GenerateSliceOfInt() creates a slice of a given size and populates its
// values with the given generator (which may be nil)
func GenerateSliceOfInt(size int, generator func(index int) (result int, err error)) (result []int, err error)

// GenerateSliceOfUInt() creates a slice of a given size and populates its
// values with the given generator (which may be nil)
func GenerateSliceOfUInt(size int, generator func(index int) (result uint, err error)) (result []uint, err error)

// GenerateSliceOfString() creates a slice of a given size and populates its
// values with the given generator (which may be nil)
func GenerateSliceOfString(size int, generator func(index int) (result string, err error)) (result []string, err error)
```

```Go
// in "github.com/synesissoftware/ANGoLS/slices"

func SelectSliceOfInt(input_slice []int, selector func(index int, input_item int) (bool, error)) ([]int, error)

func SelectSliceOfUInt(input_slice []uint, selector func(index int, input_item uint) (bool, error)) ([]uint, error)

func SelectSliceOfUInteger[N int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | uintptr](input_slice []N, selector func(index int, input_item N) (bool, error)) ([]N, error)

func SelectSliceOfString(input_slice []string, selector func(index int, input_item string) (bool, error)) ([]string, error)
```


### Strings

```Go
// in "github.com/synesissoftware/ANGoLS/strings"

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
func IndexAfter(s string, sf string, ix int) int

// Finds the index of the first instance of any character in chars, starting
// from the position after the given index. -1 is returned if the find is
// not successful.
//
// To search from the start of the string, specify the value -1 for the
// index. Any index value less than -1 will be treated as if -1 specified.
// Any index value greater than the size of the string will result in a
// return value of -1.
//
// The returned value reflects the position of the found character relative
// to the start of the string, not from the index.
func IndexAnyAfter(s string, chars string, ix int) int

// Finds the index of the given byte in the given string, starting from the
// position after the given index. -1 is returned if the find is not
// successful.
//
// To search from the start of the string, specifying the value -1 for the
// index. Any index value less than -1 will be treated as if -1 specified.
// Any index value greater than the size of the string will result in a
// return value of -1.
//
// The returned value reflects the position of the found byte relative to
// the start of the string, not from the index.
func IndexByteAfter(s string, c byte, ix int) int

// Finds the index of a character identified by the given function, starting
// from the position after the given index. -1 is returned if the find is
// not successful.
//
// To search from the start of the string, specifying the value -1 for the
// index. Any index value less than -1 will be treated as if -1 specified.
// Any index value greater than the size of the string will result in a
// return value of -1.
//
// The returned value reflects the position of the identified character
// relative to the start of the string, not from the index.
func IndexFuncAfter(s string, f func(rune) bool, ix int) int

// Finds the index of the first instance of any character not in chars,
// starting from the position after the given index. -1 is returned if the
// find is not successful.
//
// To search from the start of the string, specify the value -1 for the
// index. Any index value less than -1 will be treated as if -1 specified.
// Any index value greater than the size of the string will result in a
// return value of -1.
//
// The returned value reflects the position of the found character relative
// to the start of the string, not from the index.
func IndexNotAnyAfter(s string, chars string, ix int) int
```

```Go
// in "github.com/synesissoftware/ANGoLS/strings"

// SplitAfterByte slices a string into all substings after each instance of
// the byte sep and returns a slice of those substrings.
//
// If s does not contain sep, SplitAfterByte returns a slice of length 1
// whose only element is s.
//
// It is equivalent to [SplitAfterByteN] with a count of -1.
func SplitAfterByte(s string, sep byte) []string

// SplitAfterByteN slices s into substrings after each instance of the byte
// sep and returns a slice of those substrings.
//
// ix determines the number of substrings to return:
//   - n > 0: at most n substrings; the last substring being the unsplit
//     remainder;
//   - n == 0: the result is nil (zero substrings);
//   - n < 0: all substrings.
func SplitAfterByteN(s string, sep byte, ix int) []string

// SplitAfterRune slices a string into all substings after each instance of
// the rune sep and returns a slice of those substrings.
//
// If s does not contain sep, SplitAfterRune returns a slice of length 1
// whose only element is s.
//
// It is equivalent to [SplitAfterRuneN] with a count of -1.
func SplitAfterRune(s string, sep rune) []string

// SplitAfterRuneN slices s into substrings after each instance of the rune
// sep and returns a slice of those substrings.
//
// ix determines the number of substrings to return:
//   - n > 0: at most n substrings; the last substring being the unsplit
//     remainder;
//   - n == 0: the result is nil (zero substrings);
//   - n < 0: all substrings.
func SplitAfterRuneN(s string, sep rune, ix int) []string

// SplitAfterAny slices a string into all substings after each instance of
// any of the runes in chars and returns a slice of those substrings.
//
// If s does not contain any of the runes in chars and chars is not empty,
// SplitAfterAny returns a slice of length 1 whose only element is s.
//
// It is equivalent to [SplitAfterAnyN] with a count of -1.
func SplitAfterAny(s, chars string) []string

// SplitAfterAnyN slices s into substrings after each instance of any of the
// runes in chars and returns a slice of those substrings.
//
// ix determines the number of substrings to return:
//   - n > 0: at most n substrings; the last substring being the unsplit
//     remainder;
//   - n == 0: the result is nil (zero substrings);
//   - n < 0: all substrings.
func SplitAfterAnyN(s, chars string, ix int) []string

// SplitAfterAnyBytes slices a string into all substings after each
// instance of any of the bytes in seps and returns a slice of those
// substrings.
//
// If s does not contain any of the bytes in seps and seps is not empty,
// SplitAfterAnyBytes returns a slice of length 1 whose only element is s.
//
// It is equivalent to [SplitAfterAnyBytesN] with a count of -1.
func SplitAfterAnyBytes(s string, seps []byte) []string

// SplitAfterAnyBytesN slices a string into all substings after each
// instance of any of the bytes in seps and returns a slice of those
// substrings.
//
// ix determines the number of substrings to return:
//   - n > 0: at most n substrings; the last substring being the unsplit
//     remainder;
//   - n == 0: the result is nil (zero substrings);
//   - n < 0: all substrings.
func SplitAfterAnyBytesN(s string, seps []byte, ix int) []string

// SplitAfterAnyRunes slices a string into all substings after each
// instance of any of the runes in seps and returns a slice of those
// substrings.
//
// If s does not contain any of the runes in seps and seps is not empty,
// SplitAfterAnyRunes returns a slice of length 1 whose only element is s.
//
// It is equivalent to [SplitAfterAnyN] with a count of -1.
func SplitAfterAnyRunes(s string, seps []rune) []string

// SplitAfterAnyRunesN slices a string into all substings after each
// instance of any of the runes in seps and returns a slice of those
// substrings.
//
// ix determines the number of substrings to return:
//   - n > 0: at most n substrings; the last substring being the unsplit
//     remainder;
//   - n == 0: the result is nil (zero substrings);
//   - n < 0: all substrings.
func SplitAfterAnyRunesN(s string, seps []rune, ix int) []string
```

```Go
// in "github.com/synesissoftware/ANGoLS/strings"

// StringChomp() takes a single string and returns a chomped version of it,
// where chomping removes a single trailing '\n' character, a single
// trailing '\r' character, or a single trailing sequence of "\r\n"
func StringChomp(s string) string

// StringChompAll() takes a single string and returns a fully/repeatedly
// chomped version of, where full/repeated chomping removes all trailing
// '\r' and/or '\n' characters
func StringChompAll(s string) string
```


## Examples

Examples are provided in the `examples` directory, along with a markdown description for each. A detailed list TOC of them is provided in [EXAMPLES.md](./EXAMPLES.md).


## Project Information


### Where to get help

[GitHub Page](https://github.com/synesissoftware/ANGoLS "GitHub Page")


### Contribution guidelines

Defect reports, feature requests, and pull requests are welcome on https://github.com/synesissoftware/ANGoLS.


### Dependencies

* [**ver2go**](https://github.com/synesissoftware/ver2go/);


#### Development/Testing Dependencies

* [**require**](https://github.com/stretchr/testify/);
* [**STEGoL**](https://github.com/synesissoftware/STEGoL/);


### Related projects

T.B.C.


### License

**ANGoLS** is released under the 3-clause BSD license. See [LICENSE](./LICENSE) for details.


<!-- ########################### end of file ########################### -->

