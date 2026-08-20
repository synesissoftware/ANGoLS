// Copyright 2019-2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	"errors"
	"strings"
)

type ParseKeyValuePairsListOption int64

// NOTE: do not ever insert into this list, only ever append to it

const (
	ParseKeyValuePairsListOption_None                 ParseKeyValuePairsListOption = 0
	ParseKeyValuePairsListOption_IgnoreAnonymousValue ParseKeyValuePairsListOption = 1 << iota // causes an anonymous value - e.g. ",=val1" - to be omitted from the results; otherwise this causes parse failure
	ParseKeyValuePairsListOption_PermitAnonymousValue                                          // causes an anonymous value - e.g. ",=val1" - to be included in the results, with the key `""`; otherwise this causes parse failure
	ParseKeyValuePairsListOption_IgnoreValuelessKey                                            // causes a valueless key - e.g. ",key1," or ",key1=," - to be omitted from the results; otherwise this causes parse failure
	ParseKeyValuePairsListOption_PermitValuelessKey                                            // causes a valueless key - e.g. ",key1," or ",key1=," - to be included in the results, with the value `""`; otherwise this causes parse failure
	ParseKeyValuePairsListOption_PermitRepeatedKeys                                            // causes repeated key(s) to be included in the results; otherwise this will cause parse failure
	ParseKeyValuePairsListOption_TakeFirstRepeatedKey                                          // causes the first repeated key only to be obtained in the results; otherwise this will cause parse failure
	ParseKeyValuePairsListOption_TakeLastRepeatedKey                                           // causes the last repeated key only to be obtained in the results; otherwise this will cause parse failure
	ParseKeyValuePairsListOption_PreserveOrder                                                 // causes key-value pairs order to be preserved in the results, which may cause more processing cost; otherwise arbitrary order will obtain
)

// anonymous value,
// valueless key ",key1," or ",key1=,"
// repeated key

var (
	errAnonymousValueNotPermitted = errors.New("missing key")
	errInvalidSeparators          = errors.New("invalid separator(s): neither `pairSeparator` nor `keyValueSeparator` may be `nil`, nor may they be equal")
	errRepeatedKeys               = errors.New("repeated keys")
	errValuelessKeyNotPermitted   = errors.New("missing value")
)

// Result type for `ParseKeyValuePairsList`
type KeyValuePair struct {
	Key   string // The key, or an empty string for an anonymous value
	Value string // The value, or an empty string for a valueless key
}

// Parses a string that contains a list of key-value pairs
//
// Parameters:
//   - input - The input string;
//   - pairSeparator - The pair separator, which may not be `nil` and may
//     not equal `keyValueSeparator`;
//   - keyValueSeparator - The key/value separator, which may not be `nil`
//     and may not equal `pairSeparator`;
//   - parseOptions - Options that control the parsing;
//
// Preconditions:
//   - `pairSeparator != ""`;
//   - `keyValueSeparator != ""`;
//   - `pairSeparator != keyValueSeparator;
//
// Note:
// Whitespace is elided from around key-value pairs, and from around keys,
// but is preserved from values. Thus, if the only/last pair is to contain
// trailing whitespace, it must be followed by a pairSeparator, as in the
// following example " key1 =  a  really  spacey  value  ," that would
// obtain a key of "key1" and a value of "  a  really  spacey  value  ".
//
// Note:
// If input contains either/both of the separate sequences in the keys or
// values they are recognised as separators nonetheless - no escaping is
// supported in the current implementation. The exception to this is that a
// value may be obtained with a prefix of `keyValueSeparator` by specifying
// it twice (or more) in the pair, as in "key1==value1==," would obtain the
// value "=value1==".
func ParseKeyValuePairsList(
	input string,
	pairSeparator string, // e.g. ","
	keyValueSeparator string, // e.g. "="
	parseOptions ParseKeyValuePairsListOption,
) (
	pairs []KeyValuePair,
	err error,
) {
	// fmt.Fprintf(os.Stderr, "%s(input=%s, pairSeparator=%s, keyValueSeparator=%s, parseOptions=%x)\n", d.FileLineFunction(), input, pairSeparator, keyValueSeparator, parseOptions)

	// precondition enforcement

	if pairSeparator == "" || keyValueSeparator == "" || pairSeparator == keyValueSeparator {
		panic(errInvalidSeparators)
	}

	pairs = make([]KeyValuePair, 0, len(input)/10) // 10 be a guess ...

	input = strings.TrimSpace(input)

	splits0 := strings.Split(input, pairSeparator)

	for i, s0 := range splits0 {

		splits1 := strings.SplitN(s0, keyValueSeparator, 2)

		var k string
		var v string

		switch len(splits1) {
		case 0: // e.g. 2nd from "key1=val1,,key3=val3"

			// ignore
			continue
		case 1: // e.g. 2nd from "key1=val1,  ,key3=val3"

			k = strings.TrimSpace(splits1[0])

			if k == "" {

				// ignore
				continue
			} else {

				if _hasFlag(parseOptions, ParseKeyValuePairsListOption_IgnoreValuelessKey) {

					// ignore
					continue
				}

				if !_hasFlag(parseOptions, ParseKeyValuePairsListOption_PermitValuelessKey) {

					err = errValuelessKeyNotPermitted

					return
				}

				pairs = append(pairs, KeyValuePair{
					Key: k,
				})
			}
		case 2:

			k = strings.TrimSpace(splits1[0])
			v = splits1[1]

			if k == "" {

				if _hasFlag(parseOptions, ParseKeyValuePairsListOption_IgnoreAnonymousValue) {

					// ignore
					continue
				}

				if !_hasFlag(parseOptions, ParseKeyValuePairsListOption_PermitAnonymousValue) {

					err = errAnonymousValueNotPermitted

					return
				}
			}

			if v == "" {

				if _hasFlag(parseOptions, ParseKeyValuePairsListOption_IgnoreValuelessKey) {

					// ignore
					continue
				}

				if !_hasFlag(parseOptions, ParseKeyValuePairsListOption_PermitValuelessKey) {

					err = errValuelessKeyNotPermitted

					return
				}

				_ = i // TODO: report error, including pair index
			}

			pairs = append(pairs, KeyValuePair{
				Key:   k,
				Value: v,
			})
		}
	}

	// now deal with repeated keys

	if !_hasFlag(parseOptions, ParseKeyValuePairsListOption_PermitRepeatedKeys) {

		// need to

		switch {
		case _hasFlag(parseOptions, ParseKeyValuePairsListOption_TakeLastRepeatedKey):

			pairs2 := make([]KeyValuePair, 0, len(pairs))

			// map of string => int, where value is "index of first (i.e. only)"

			m := make(map[string]int, len(pairs))

			for _, kv := range pairs {

				if ix, exists := m[kv.Key]; exists {

					if _hasFlag(parseOptions, ParseKeyValuePairsListOption_PreserveOrder) {

					} else {

						pairs2[ix] = kv
					}
				} else {

					pairs2 = append(pairs2, kv)

					m[kv.Key] = len(pairs2) - 1
				}
			}

			pairs = pairs2
		default:

			pairs2 := make([]KeyValuePair, 0, len(pairs))

			// map of string => bool, where true is "seen before"

			m := make(map[string]bool, len(pairs))

			for _, kv := range pairs {

				if m[kv.Key] {

					if !_hasFlag(parseOptions, ParseKeyValuePairsListOption_TakeFirstRepeatedKey) {

						err = errRepeatedKeys

						return
					}
				} else {

					m[kv.Key] = true

					pairs2 = append(pairs2, kv)
				}
			}

			pairs = pairs2
		}
	}

	return
}

func _hasFlag(
	parseOptions ParseKeyValuePairsListOption,
	flag ParseKeyValuePairsListOption,
) bool {

	return (parseOptions & flag) == flag
}
