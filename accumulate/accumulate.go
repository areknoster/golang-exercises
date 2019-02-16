// Package accumulate implements function for function mapping on string slice.
package accumulate

// Accumulate performs an operation on each string in input slice and returns the slice with results
func Accumulate(input []string, f func(string) string) []string {
	result := make([]string, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}
