# **ANGoLS** Changes

## 0.10.0-beta2 - 15th February 2026

* refactored `SkipOneElement` and `SkipRemainingElements`;


## 0.10.0-beta1 - 7th February 2026

* simple refactoring of `ASCIIToLower()` and `ASCIIToUpper()`;
* changed signature of `CollectSlice()`, including making be a generic function (`CollectSlice[T, U]()`);
* refactored `CollectSliceIntoStringSlice()` in terms of `CollectSlice[T, U]()`;
* refactor `CollectSliceOfInt()` in terms of `CollectSliceOfInteger[T]()`;
* refactored `EqualSliceOfInt()` in terms of `EqualSliceOfInteger[T]()`;
* refactored `EqualSliceOfUint()` in terms of `EqualSliceOfInteger[T]()`;
* changed signature of `EqualSlice()`, including making be a generic function (`EqualSlice[T]()`);
* renamed `SelectSliceOfUinteger()` => `SelectSliceOfInteger()`;
* simplified implementation of `SelectSliceOfInteger[T]()`;
* added unit-tests for `SelectSliceOfInt()` and `SelectSliceOfString()`;
* refactored `SelectSliceOfInt()` in terms of `SelectSliceOfInteger[T]()`;
* refactored `SelectSliceOfUint()` in terms of `SelectSliceOfInteger[T]()`;
* simplified implementation of `SelectSliceOfString()`;


## 0.9.1 - 7th February 2026

* fixed `EqualSliceOfInteger[]()` return type;
* added unit-tests for equal and select;


## 0.9.0 - 3rd January 2026

* 0.9.0


## 0.9.0 (beta 2) - 2nd January 2026

* added `ASCIIToULower()` and `ASCIIToUpper()`;


## 0.9.0 (beta 1) - 27th November 2025

* added `CollectSliceIntoStringSlice[T any](input_slice []T, fn func(input_item *T) (string, error)) ([]string, error)`;
* changed return type of `CollectSlice()` from `(any, error)` to `([]any, error)`;


## 0.8.0 (alpha 1) - 27th August 2025

* added generic function `func CollectSliceOfInteger[N int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | uintptr](input_slice []N, fn func(input_item N) N) (result_slice []N)`;
* added generic function `func EqualSliceOfInteger[N int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | uintptr](lhs, rhs []int) bool`;
* added generic function `func SelectSliceOfUInteger[N int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | uintptr](lhs, rhs []int) bool`;


## 0.7.0 - 27th August 2025

* 0.7.0


## 0.7.0 (beta 1) - 27th August 2025

* moved private elements into **internal** (so as to hide properly);
* dependencies;
* GitHub Actions enhancements;
* boilerplate;


## 0.7.0 (alpha 1) - 20th August 2025

* added `SplitAfterByte()`, `SplitAfterByteN()` ,`SplitAfterRune()` ,`SplitAfterRuneN()`, `SplitAfterAny()`, `SplitAfterAnyN()`, `SplitAfterAnyBytes()`, `SplitAfterAnyBytesN()`, `SplitAfterAnyRunes()`, `SplitAfterAnyRunesN()`;


## 0.6.0 - 20th August 2025

* documentation;


## 0.6.0 (beta 1) - 18th August 2025

* GitHub Actions;
* `interface{}` => `any`;
* boilerplate;
* documentation;


## 0.6.0 (alpha 3) - 14th August 2025

* tidying;


## 0.6.0 (alpha 2) - 24th February 2025

* added **examples/libver.go**;
* boilerplate;


## 0.6.0 (alpha 1) - 24th February 2025

* refactored all slice functions into **slices** package;
* refactored all slice functions into **strings** package;
* added `strings.IndexAfter()`;
* added `strings.IndexAnyAfter()`;
* added `strings.IndexByteAfter()`;
* added `strings.IndexFuncAfter()`;
* added `strings.IndexNotAnyAfter()`;


## 0.5.0 - 30th March 2019

* + added ``StringChomp()`` and ``StringChompAll()`` functions
* + added unit-tests

## 0.4.0 - 11th March 2019

FIRST PUBLIC RELEASE

