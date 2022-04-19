package allergies

var ALLERGIES = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func Allergies(allergies uint) []string {
	var out []string
	for i, a := range ALLERGIES {
		if allergies&(1<<i) != 0 {
			out = append(out, a)
		}
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
