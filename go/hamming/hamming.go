package hamming

import (
	"fmt"
)

// Distance calculates the Hamming distance between two strings
func Distance(sa, sb string) (int, error) {
	a, b := []rune(sa), []rune(sb)
	if len(a) != len(b) {
		return 0, fmt.Errorf(`invalid input: length of "%s" (%d) and "%s" (%d) are not equal`, sa, len(a), sb, len(b))
	}
	diff := 0
	for i := range a {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff, nil
}
