package angols_test

import (
	"github.com/stretchr/testify/require"
	strings "github.com/synesissoftware/ANGoLS/strings"

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
		require.LessOrEqual(t, -1, strings.IndexAfter("abcdefghijklmnopqrstuvwxyz", "nopqrs", 100))
	}
}
