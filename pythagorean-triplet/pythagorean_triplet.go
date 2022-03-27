package pythagorean

import (
	"math"
)

type Triplet [3]int

func hypotenuse(a, b int) float64 {
	return math.Sqrt(float64(a*a + b*b))
}

func Range(min, max int) (result []Triplet) {
	for a := min; a <= max; a++ {
		for b := a + 1; hypotenuse(a, b) <= float64(max); b++ {
			if c := hypotenuse(a, b); c == float64(int(c)) {
				result = append(result, Triplet{a, b, int(c)})
			}
		}
	}
	return result
}

func Sum(p int) (result []Triplet) {
	for _, t := range Range(1, p/2) {
		if t[0]+t[1]+t[2] == p {
			result = append(result, t)
		}
	}
	return result
}
