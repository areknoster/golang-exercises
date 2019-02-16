// Package hamming helps you calculate the Hamming difference between two DNA strands.
package hamming

import "errors"

// Distance takes two DNA strands encoded as strings. It counts how many of the nucleotides
// are different from their equivalent in the other string. Returns error in case strings are not the same length.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("error: strings are not the same length")
	}
	total := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			total++
		}
	}
	return total, nil
}
