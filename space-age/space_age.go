package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case "Earth":
		return 31.69
	case "Mercury":
		return 280.88
	case "Venus":
		return 9.78
	case "Mars":
		return 35.88
	case "Jupiter":
		return 2.41
	case "Saturn":
		return 2.15
	case "Uranus":
		return 0.46
	case "Neptune":
		return 0.35
	default:
		return 0
	}
}
