package allergies

import (
	"math"
)

var ALLERGIES = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}

func Allergies(allergies uint) []string {
	allergies %= 256

	var out []string
	for i := 7; i >= 0; i-- {
		d := uint(math.Pow(2, float64(i)))
		if allergies/d == 1 {
			out = append(out, ALLERGIES[i])
		}
		allergies %= d
	}
	return out
}

func AllergicTo(allergies uint, allergen string) bool {
	for _, a := range Allergies(allergies) {
		if a == allergen {
			return true
		}
	}
	return false
}
