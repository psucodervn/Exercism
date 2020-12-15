package hamming

import "errors"

// Distance calculates the Hamming distance between two strings
func Distance(sa, sb string) (int, error) {
	a, b := []rune(sa), []rune(sb)
	if len(a) != len(b) {
		return 0, errors.New("invalid input")
	}
	diff := 0
	for i := range a {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff, nil
}
