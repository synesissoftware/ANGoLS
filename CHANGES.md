# **ANGoLS** Changes

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

