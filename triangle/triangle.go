package triangle

import "sort"

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

func KindFromSides(a, b, c float64) Kind {
	n := []float64{a, b, c}
	sort.Float64s(n)

	a, b, c = n[0], n[1], n[2]

	var k Kind

	if (a + b) <= c {
		k = NaT
	} else if a == b && a == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}

	return k
}
