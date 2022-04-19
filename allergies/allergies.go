package allergies

var ALLERGIES = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(allergies uint) []string {
	var out []string
	for allergen, _ := range ALLERGIES {
		if AllergicTo(allergies, allergen) {
			out = append(out, allergen)
		}
	}
	return out
}

func AllergicTo(allergies uint, allergen string) bool {
	return allergies&ALLERGIES[allergen] > 0
}
