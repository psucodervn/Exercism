// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", "")
	ar := strings.Split(s, " ")
	res := ""
	for _, w := range ar {
		if len(w) == 0 {
			continue
		}
		res = res + strings.ToUpper(string(w[0]))
	}
	return res
}
