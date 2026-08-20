# ANGoLS - TODO <!-- omit in toc -->


## Table of Contents <!-- omit in toc -->

- [Functional improvements](#functional-improvements)
- [Performance improvements](#performance-improvements)
- [Packaging improvements](#packaging-improvements)


## Functional improvements

* ~~simplify `CollectSlice()`~~ - ✅;
* ~~add unit-testing for `slices` functions~~ - ✅;
* ~~enhance code coverage in unit-tests~~ - ✅;
* ~~change semantics of `CollectSlice()`, `EqualSlice()` to not take `any`, but (at least) `[]any`~~ - ✅;
* ~~correct all `slices` functions to proper naming, including `Uint` rather than `Uint`~~ - ✅;
* ~~`GenerateSliceOfInteger[T]()`~~ - ✅;
* [ ] `SelectSlice[T]()`;
* ~~expand coverage of specific types for `EqualSlice[T]()`~~ - ✅;
* ~~enhance code coverage in unit-tests~~ - ✅;
* [ ] simplify names, e.g. `CollectSlice()` => `Collect()`, and so forth;
* [ ] apply `~` on generics;


## Performance improvements

* \<none>


## Packaging improvements

* ~~Before the next official release: confirm **`go.mod`** (`go 1.21`) and the CI Go-version matrix, bump Synesis `require`s to newly published tags, then run **`go mod tidy`**~~ - ✅;


<!-- ########################### end of file ########################### -->
