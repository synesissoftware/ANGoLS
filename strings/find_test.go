package strings_test

import (
	strings "github.com/synesissoftware/ANGoLS/strings"

	"github.com/stretchr/testify/require"

	"testing"
)

func Test_IndexAfter(t *testing.T) {

	{
		require.Equal(t, 0, strings.IndexAfter("", "", -1))
	}

	{
		require.Equal(t, 0, strings.IndexAfter("abc", "", -3))
		require.Equal(t, 0, strings.IndexAfter("abc", "", -2))
		require.Equal(t, 0, strings.IndexAfter("abc", "", -1))
		require.Equal(t, +1, strings.IndexAfter("abc", "", 0))
		require.Equal(t, +2, strings.IndexAfter("abc", "", 1))
		require.Equal(t, +3, strings.IndexAfter("abc", "", 2))
		require.Equal(t, -1, strings.IndexAfter("abc", "", 3))
		require.Equal(t, -1, strings.IndexAfter("abc", "", 4))

		require.Equal(t, 0, strings.IndexAfter("abc", "a", -3))
		require.Equal(t, 0, strings.IndexAfter("abc", "a", -2))
		require.Equal(t, 0, strings.IndexAfter("abc", "a", -1))
		require.Equal(t, -1, strings.IndexAfter("abc", "a", 0))
		require.Equal(t, -1, strings.IndexAfter("abc", "a", +1))
		require.Equal(t, -1, strings.IndexAfter("abc", "a", +2))
		require.Equal(t, -1, strings.IndexAfter("abc", "a", +3))

		require.Equal(t, +1, strings.IndexAfter("abc", "b", -3))
		require.Equal(t, +1, strings.IndexAfter("abc", "b", -2))
		require.Equal(t, +1, strings.IndexAfter("abc", "b", -1))
		require.Equal(t, +1, strings.IndexAfter("abc", "b", 0))
		require.Equal(t, -1, strings.IndexAfter("abc", "b", +1))
		require.Equal(t, -1, strings.IndexAfter("abc", "b", +2))
		require.Equal(t, -1, strings.IndexAfter("abc", "b", +3))

		require.Equal(t, +2, strings.IndexAfter("abc", "c", -3))
		require.Equal(t, +2, strings.IndexAfter("abc", "c", -2))
		require.Equal(t, +2, strings.IndexAfter("abc", "c", -1))
		require.Equal(t, +2, strings.IndexAfter("abc", "c", 0))
		require.Equal(t, +2, strings.IndexAfter("abc", "c", +1))
		require.Equal(t, -1, strings.IndexAfter("abc", "c", +2))
		require.Equal(t, -1, strings.IndexAfter("abc", "c", +3))
	}

	{
		require.LessOrEqual(t, 0, strings.IndexAfter("abcdefghijklmnopqrstuvwxyz", "abc", -3))
		require.LessOrEqual(t, 0, strings.IndexAfter("abcdefghijklmnopqrstuvwxyz", "abc", -2))
		require.LessOrEqual(t, 0, strings.IndexAfter("abcdefghijklmnopqrstuvwxyz", "abc", -1))
		require.LessOrEqual(t, -1, strings.IndexAfter("abcdefghijklmnopqrstuvwxyz", "abc", 0))
		require.LessOrEqual(t, -1, strings.IndexAfter("abcdefghijklmnopqrstuvwxyz", "abc", 10))
		require.LessOrEqual(t, -1, strings.IndexAfter("abcdefghijklmnopqrstuvwxyz", "abc", 20))

		require.LessOrEqual(t, 13, strings.IndexAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", -3))
		require.LessOrEqual(t, 13, strings.IndexAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", -2))
		require.LessOrEqual(t, 13, strings.IndexAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", -1))
		require.LessOrEqual(t, 13, strings.IndexAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 0))
		require.LessOrEqual(t, 13, strings.IndexAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 1))
		require.LessOrEqual(t, 13, strings.IndexAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 10))
		require.LessOrEqual(t, 13, strings.IndexAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 11))
		require.LessOrEqual(t, 13, strings.IndexAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 12))
		require.LessOrEqual(t, -1, strings.IndexAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 13))
		require.LessOrEqual(t, -1, strings.IndexAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 14))
		require.LessOrEqual(t, -1, strings.IndexAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 15))
		require.LessOrEqual(t, -1, strings.IndexAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 100))
	}
}

func Test_IndexAnyAfter(t *testing.T) {

	{
		require.Equal(t, -1, strings.IndexAnyAfter("", "", -1))
	}

	{
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "", -3))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "", -2))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "", -1))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "", 0))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "", 1))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "", 2))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "", 3))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "", 4))

		require.Equal(t, 0, strings.IndexAnyAfter("abc", "a", -3))
		require.Equal(t, 0, strings.IndexAnyAfter("abc", "a", -2))
		require.Equal(t, 0, strings.IndexAnyAfter("abc", "a", -1))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "a", 0))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "a", +1))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "a", +2))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "a", +3))

		require.Equal(t, +1, strings.IndexAnyAfter("abc", "b", -3))
		require.Equal(t, +1, strings.IndexAnyAfter("abc", "b", -2))
		require.Equal(t, +1, strings.IndexAnyAfter("abc", "b", -1))
		require.Equal(t, +1, strings.IndexAnyAfter("abc", "b", 0))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "b", +1))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "b", +2))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "b", +3))

		require.Equal(t, +2, strings.IndexAnyAfter("abc", "c", -3))
		require.Equal(t, +2, strings.IndexAnyAfter("abc", "c", -2))
		require.Equal(t, +2, strings.IndexAnyAfter("abc", "c", -1))
		require.Equal(t, +2, strings.IndexAnyAfter("abc", "c", 0))
		require.Equal(t, +2, strings.IndexAnyAfter("abc", "c", +1))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "c", +2))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "c", +3))

		require.Equal(t, -1, strings.IndexAnyAfter("abc", "d", -3))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "d", -2))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "d", -1))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "d", 0))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "d", +1))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "d", +2))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "d", +3))

		require.Equal(t, 0, strings.IndexAnyAfter("abc", "abcd", -3))
		require.Equal(t, 0, strings.IndexAnyAfter("abc", "abcd", -2))
		require.Equal(t, 0, strings.IndexAnyAfter("abc", "abcd", -1))
		require.Equal(t, 1, strings.IndexAnyAfter("abc", "abcd", 0))
		require.Equal(t, 2, strings.IndexAnyAfter("abc", "abcd", +1))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "abcd", +2))
		require.Equal(t, -1, strings.IndexAnyAfter("abc", "abcd", +3))
	}

	{
		require.LessOrEqual(t, 0, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "abc", -3))
		require.LessOrEqual(t, 0, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "abc", -2))
		require.LessOrEqual(t, 0, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "abc", -1))
		require.LessOrEqual(t, -1, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "abc", 0))
		require.LessOrEqual(t, -1, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "abc", 10))
		require.LessOrEqual(t, -1, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "abc", 20))

		require.LessOrEqual(t, 13, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", -3))
		require.LessOrEqual(t, 13, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", -2))
		require.LessOrEqual(t, 13, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", -1))
		require.LessOrEqual(t, 13, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 0))
		require.LessOrEqual(t, 13, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 1))
		require.LessOrEqual(t, 13, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 10))
		require.LessOrEqual(t, 13, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 11))
		require.LessOrEqual(t, 13, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 12))
		require.LessOrEqual(t, 14, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 13))
		require.LessOrEqual(t, 15, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 14))
		require.LessOrEqual(t, 16, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 15))
		require.LessOrEqual(t, 17, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 16))
		require.LessOrEqual(t, 18, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 17))
		require.LessOrEqual(t, -1, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 18))
		require.LessOrEqual(t, -1, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 19))
		require.LessOrEqual(t, -1, strings.IndexAnyAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 100))
	}
}

func Test_IndexByteAfter(t *testing.T) {

	{
		require.Equal(t, -1, strings.IndexByteAfter("", '~', -1))
	}

	{
		require.Equal(t, -1, strings.IndexByteAfter("abc", '~', -3))
		require.Equal(t, -1, strings.IndexByteAfter("abc", '~', -2))
		require.Equal(t, -1, strings.IndexByteAfter("abc", '~', -1))
		require.Equal(t, -1, strings.IndexByteAfter("abc", '~', 0))
		require.Equal(t, -1, strings.IndexByteAfter("abc", '~', 1))
		require.Equal(t, -1, strings.IndexByteAfter("abc", '~', 2))
		require.Equal(t, -1, strings.IndexByteAfter("abc", '~', 3))
		require.Equal(t, -1, strings.IndexByteAfter("abc", '~', 4))

		require.Equal(t, 0, strings.IndexByteAfter("abc", 'a', -3))
		require.Equal(t, 0, strings.IndexByteAfter("abc", 'a', -2))
		require.Equal(t, 0, strings.IndexByteAfter("abc", 'a', -1))
		require.Equal(t, -1, strings.IndexByteAfter("abc", 'a', 0))
		require.Equal(t, -1, strings.IndexByteAfter("abc", 'a', +1))
		require.Equal(t, -1, strings.IndexByteAfter("abc", 'a', +2))
		require.Equal(t, -1, strings.IndexByteAfter("abc", 'a', +3))

		require.Equal(t, +1, strings.IndexByteAfter("abc", 'b', -3))
		require.Equal(t, +1, strings.IndexByteAfter("abc", 'b', -2))
		require.Equal(t, +1, strings.IndexByteAfter("abc", 'b', -1))
		require.Equal(t, +1, strings.IndexByteAfter("abc", 'b', 0))
		require.Equal(t, -1, strings.IndexByteAfter("abc", 'b', +1))
		require.Equal(t, -1, strings.IndexByteAfter("abc", 'b', +2))
		require.Equal(t, -1, strings.IndexByteAfter("abc", 'b', +3))

		require.Equal(t, +2, strings.IndexByteAfter("abc", 'c', -3))
		require.Equal(t, +2, strings.IndexByteAfter("abc", 'c', -2))
		require.Equal(t, +2, strings.IndexByteAfter("abc", 'c', -1))
		require.Equal(t, +2, strings.IndexByteAfter("abc", 'c', 0))
		require.Equal(t, +2, strings.IndexByteAfter("abc", 'c', +1))
		require.Equal(t, -1, strings.IndexByteAfter("abc", 'c', +2))
		require.Equal(t, -1, strings.IndexByteAfter("abc", 'c', +3))
	}
}

func Test_IndexFuncAfter(t *testing.T) {

	is_ASCII_vowel := func(c rune) bool {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			return true
		default:
			return false
		}
	}

	is_not_ASCII_vowel := func(c rune) bool {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			return false
		default:
			return true
		}
	}

	{
		require.Equal(t, -1, strings.IndexFuncAfter("", is_ASCII_vowel, -1))
	}

	{
		require.Equal(t, 0, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, -3))
		require.Equal(t, 0, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, -2))
		require.Equal(t, 0, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, -1))
		require.Equal(t, 4, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 0))
		require.Equal(t, 4, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 1))
		require.Equal(t, 4, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 2))
		require.Equal(t, 4, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 3))
		require.Equal(t, 8, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 4))
		require.Equal(t, 8, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 5))
		require.Equal(t, 8, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 6))
		require.Equal(t, 8, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 7))
		require.Equal(t, 14, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 8))
		require.Equal(t, 14, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 9))
		require.Equal(t, 14, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 10))
		require.Equal(t, 14, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 11))
		require.Equal(t, 14, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 12))
		require.Equal(t, 14, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 13))
		require.Equal(t, 20, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 14))
		require.Equal(t, 20, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 15))
		require.Equal(t, 20, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 16))
		require.Equal(t, 20, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 17))
		require.Equal(t, 20, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 18))
		require.Equal(t, 20, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 19))
		require.Equal(t, -1, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 20))
		require.Equal(t, -1, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 21))
		require.Equal(t, -1, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 22))
		require.Equal(t, -1, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 23))
		require.Equal(t, -1, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 24))
		require.Equal(t, -1, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 25))
		require.Equal(t, -1, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 26))
		require.Equal(t, -1, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_ASCII_vowel, 27))
	}

	{
		require.Equal(t, 1, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, -3))
		require.Equal(t, 1, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, -2))
		require.Equal(t, 1, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, -1))
		require.Equal(t, 1, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 0))
		require.Equal(t, 2, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 1))
		require.Equal(t, 3, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 2))
		require.Equal(t, 5, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 3))
		require.Equal(t, 5, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 4))
		require.Equal(t, 6, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 5))
		require.Equal(t, 7, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 6))
		require.Equal(t, 9, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 7))
		require.Equal(t, 9, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 8))
		require.Equal(t, 10, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 9))
		require.Equal(t, 11, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 10))
		require.Equal(t, 12, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 11))
		require.Equal(t, 13, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 12))
		require.Equal(t, 15, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 13))
		require.Equal(t, 15, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 14))
		require.Equal(t, 16, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 15))
		require.Equal(t, 17, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 16))
		require.Equal(t, 18, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 17))
		require.Equal(t, 19, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 18))
		require.Equal(t, 21, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 19))
		require.Equal(t, 21, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 20))
		require.Equal(t, 22, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 21))
		require.Equal(t, 23, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 22))
		require.Equal(t, 24, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 23))
		require.Equal(t, 25, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 24))
		require.Equal(t, -1, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 25))
		require.Equal(t, -1, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 26))
		require.Equal(t, -1, strings.IndexFuncAfter("abcdefghijklmnopqrstuvwxyz", is_not_ASCII_vowel, 27))
	}
}

func Test_IndexNotAnyAfter(t *testing.T) {

	{
		require.Equal(t, 0, strings.IndexNotAnyAfter("", "", -1))
		require.Equal(t, -1, strings.IndexNotAnyAfter("", "", 0))
	}

	{
		require.Equal(t, -1, strings.IndexNotAnyAfter("abc", "abc", -1))
		require.Equal(t, -1, strings.IndexNotAnyAfter("abc", "abc", 0))

		require.Equal(t, 2, strings.IndexNotAnyAfter("abc", "ab", -1))
		require.Equal(t, 2, strings.IndexNotAnyAfter("abc", "ab", 0))
		require.Equal(t, 2, strings.IndexNotAnyAfter("abc", "ab", +1))
		require.Equal(t, -1, strings.IndexNotAnyAfter("abc", "ab", +2))

		require.Equal(t, 0, strings.IndexNotAnyAfter("abc", "bc", -1))
		require.Equal(t, -1, strings.IndexNotAnyAfter("abc", "bc", 0))
		require.Equal(t, -1, strings.IndexNotAnyAfter("abc", "bc", +1))
		require.Equal(t, -1, strings.IndexNotAnyAfter("abc", "bc", +2))
	}

	{
		require.Equal(t, -1, strings.IndexNotAnyAfter(".", "./", -1))
		require.Equal(t, -1, strings.IndexNotAnyAfter("./", "./", -1))
		require.Equal(t, -1, strings.IndexNotAnyAfter("..", "./", -1))
		require.Equal(t, -1, strings.IndexNotAnyAfter("../", "./", -1))

		require.Equal(t, 1, strings.IndexNotAnyAfter(".a", "./", -1))
		require.Equal(t, 2, strings.IndexNotAnyAfter("./a", "./", -1))
	}
}
