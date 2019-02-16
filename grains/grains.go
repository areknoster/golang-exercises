// Package grains finds the number of grains the king had to pay.
package grains

import (
	"fmt"
)

// Square finds the number of grains for given given square on a chessboard.
func Square(n int) (val uint64, err error) {
	if n < 1 || 64 < n {
		return 0, fmt.Errorf("error: there is no square number %v on the board", n)
	}
	return uint64(1 << (uint64(n) - 1)), err
}

// Total returns the total number of grains which the king should pay for all squares on chessboard.
func Total() uint64 {
	return ^uint64(0)
}
