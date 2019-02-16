// Package bob implements function which help you predict his answer on remarks.
package bob

import "strings"

// Hey takes the remark on Bob string. It returns his answer.
func Hey(remark string) string {
	remark = removeSpaces(remark)
	if remark == "" {
		return "Fine. Be that way!"
	}
	question, capital := isQuestion(remark), isCapital(remark)
	switch {
	case question && capital:
		return "Calm down, I know what I'm doing!"
	case question:
		return "Sure."
	case capital:
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

func removeSpaces(remark string) string {
	return strings.Join(strings.Fields(remark), "")
}

func isQuestion(remark string) bool {
	return remark[len(remark)-1] == '?'
}

func isCapital(remark string) bool {
	return strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark
}
