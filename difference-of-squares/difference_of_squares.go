// Package diffsquares implements functions for computing the values on functions on natural sequences.
package diffsquares

// SquareOfSum takes n integer and returns the square of sum of all elements of sequence an = n from 1 to n.
func SquareOfSum(n int) int {
	return (n * n * (n + 1) * (n + 1)) / 4
}

// SumOfSquares takes n integer and returns the sum of all elements of sequence an = n * n from 1 to n.
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference takes n integer. Let sequence a(n) = n for all natural i: i <= n.
// The function returns the difference of square of sum of a(i) and the sum of squares of a(i).
func Difference(n int) int {
	return n * (3*n + 2) * (n*n - 1) / 12
}
