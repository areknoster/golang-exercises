// Package raindrops helps you find the correct answer to given number
// in the Raindrops game.
package raindrops

import (
	"strconv"
)

// Convert takes the n int number and returns string, which follows
// the game's rules.
func Convert(n int) (answer string) {
	answer = ""
	if n%3 == 0 {
		answer += "Pling"
	}
	if n%5 == 0 {
		answer += "Plang"
	}
	if n%7 == 0 {
		answer += "Plong"
	}
	if answer == "" {
		answer = strconv.Itoa(n)
	}
	return answer
}
