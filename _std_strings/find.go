package _std_strings

import "strings"

/* NOTE: currently waiting to find a good mechanism to use
 * standard packages under an alias.
 */

func PRIVATE_Index(s string, substr string) int {
	return strings.Index(s, substr)
}
