package isbn

import "unicode"

func IsValidISBN(input string) bool {
	var ar []int
	for _, c := range input {
		if unicode.IsDigit(c) {
			ar = append(ar, int(c)-int('0'))
		} else if c == 'X' {
			ar = append(ar, 10)
		}
	}
	if len(ar) != 10 {
		return false
	}

	s := 0
	for i := 0; i < 10; i++ {
		if ar[i] == 10 && i < 9 {
			return false
		}
		s += ar[i] * (10 - i)
	}
	return s%11 == 0
}
