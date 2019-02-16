// Package reverse reverses strings.
package reverse

import "unicode/utf8"

// String returns input string in reverse,
func String(input string) string {
	n := utf8.RuneCountInString(input)
	rev := make([]rune, n)
	for _, v := range input {
		n--
		rev[n] = v
	}
	return string(rev)
}
