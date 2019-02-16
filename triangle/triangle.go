// Package triangle implements function, which checks the type of triangle
package triangle

import "math"

// Kind matches an int to the type of triangle
type Kind int

// Various types of triangles.
const (
	NaT = 0 //not a triangle
	Equ = 1 // equilateral triangle
	Iso = 2 // isosceles triangle
	Sca = 3 // scalene triangle
)

// KindFromSides takes float64 sides of a triangle and returns it's type in Kind type.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	switch {
	case !areValues(a, b, c) || a <= 0 || b <= 0 || c <= 0 || a+b < c || a+c < b || b+c < a:
		k = NaT
	case a == b && b == c:
		k = Equ
	case a == b || a == c || b == c:
		k = Iso
	default:
		k = Sca
	}
	return k
}

func areValues(values ...float64) bool {
	for _, v := range values {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	return true
}
