// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = "NaT" // not a triangle
	Equ Kind = "Equ" // equilateral
	Iso Kind = "Iso" // isosceles
	Sca Kind = "Sca" // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if !isValidSide(a) || !isValidSide(b) || !isValidSide(c) {
		return NaT
	}
	if lt(a+b, c) || lt(a+c, b) || lt(b+c, a) {
		return NaT
	}
	if eq(a, b) && eq(b, c) {
		return Equ
	}
	if eq(a, b) || eq(b, c) || eq(a, c) {
		return Iso
	}
	return Sca
}

func isValidSide(a float64) bool {
	return !math.IsNaN(a) && !math.IsInf(a, 0) && !lte(a, 0)
}

func eq(a, b float64) bool {
	if math.Abs(a-b) < 1e-9 {
		return true
	}
	return false
}

func lte(a, b float64) bool {
	return eq(a, b) || a < b
}

func lt(a, b float64) bool {
	return !eq(a, b) && a < b
}
