# ANGoLS <!-- omit in toc -->

[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![GitHub release](https://img.shields.io/github/v/release/synesissoftware/ANGoLS.svg)](https://github.com/synesissoftware/ANGoLS/releases/latest)
[![Last Commit](https://img.shields.io/github/last-commit/synesissoftware/ANGoLS)](https://github.com/synesissoftware/ANGoLS/commits/master)
[![Go](https://github.com/synesissoftware/ANGoLS/actions/workflows/go.yml/badge.svg)](https://github.com/synesissoftware/ANGoLS/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/synesissoftware/ANGoLS)](https://goreportcard.com/report/github.com/synesissoftware/ANGoLS)
[![Go Reference](https://pkg.go.dev/badge/github.com/synesissoftware/ANGoLS.svg)](https://pkg.go.dev/github.com/synesissoftware/ANGoLS)

**A**gorithms **N**ot in **Go** **L**anguage **S**tandard library


## Introduction

T.B.C.


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


## Installation

```Go
import angols "github.com/synesissoftware/ANGoLS"
```


## Components


### Slices

```Go
// in "github.com/synesissoftware/ANGoLS/slices"

func CollectSlice(input_slice any, fn func(input_item any) (any, error)) (any, error)

func CollectSliceOfInt(input_slice []int, fn func(input_item int) int) (result_slice []int)

func CollectSliceOfFloat64(input_slice []float64, fn func(input_item float64) float64) (result_slice []float64)

func CollectSliceOfString(input_slice []string, fn func(input_item string) string) (result_slice []string)
```

```Go
// in "github.com/synesissoftware/ANGoLS/slices"

func EqualSliceOfInt(lhs, rhs []int) bool

func EqualSliceOfUInt(lhs, rhs []uint) bool

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

func SelectSliceOfString(input_slice []string, selector func(index int, input_item string) (bool, error)) ([]string, error)
```


### Strings

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

func IndexAnyAfter(s string, chars string, ix int) int

func IndexByteAfter(s string, c byte, ix int) int

func IndexFuncAfter(s string, f func(rune) bool, ix int) int

func IndexNotAnyAfter(s string, chars string, ix int) int
```

T.B.C.


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

