package _std_strings

import "strings"

/* NOTE: currently waiting to find a good mechanism to use
 * standard packages under an alias.
 */

func PRIVATE_Index(s string, substr string) int {
	return strings.Index(s, substr)
}

func PRIVATE_IndexAny(s string, chars string) int {
	return strings.IndexAny(s, chars)
}

func PRIVATE_IndexByte(s string, c byte) int {
	return strings.IndexByte(s, c)
}

func PRIVATE_IndexFunc(s string, f func(rune) bool) int {
	return strings.IndexFunc(s, f)
}
