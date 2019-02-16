// Package isogram implements function which checks whether the phrase is an isogram.
package isogram

import "unicode"

// IsIsogram takes string and returns true if the string is a word or phrase without a repeating letter.
func IsIsogram(phrase string) bool {

	charset := make(map[rune]bool)
	for _, v := range phrase {
		if !unicode.IsLetter(v) {
			continue
		}
		v = unicode.ToLower(v)
		if charset[v] {
			return false
		}
		charset[v] = true
	}
	return true
}
