// Package luhn is package for checking numbers with Luhn formula.
package luhn

// Valid uses Luhn formula to validate the input number.
// Returns false, if the number is invalid or it contains signs other than whitespace and digits.
func Valid(input string) bool {
	var index, sum int
	for i := len(input) - 1; i >= 0; i-- {
		v := input[i]
		if v == ' ' {
			continue
		}
		if v < '0' || '9' < v {
			return false
		}
		index++
		val := int(v) - '0'
		if index%2 == 1 {
			sum += val
			continue
		}
		val *= 2
		if val > 9 {
			val -= 9
		}
		sum += val
	}
	if index <= 1 {
		return false
	}
	return sum%10 == 0
}
