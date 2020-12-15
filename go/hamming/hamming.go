package hamming

import "errors"

var (
	// ErrInvalidInput is invalid input error
	ErrInvalidInput = errors.New("invalid input")
)

// Distance calculates the Hamming distance between two strings
func Distance(sa, sb string) (int, error) {
	a, b := []rune(sa), []rune(sb)
	if len(a) != len(b) {
		return 0, ErrInvalidInput
	}
	diff := 0
	for i := range a {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff, nil
}
