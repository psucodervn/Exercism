package hamming

import "errors"

var (
	ErrInvalidInput = errors.New("invalid input")
)

func Distance(a, b string) (int, error) {
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
