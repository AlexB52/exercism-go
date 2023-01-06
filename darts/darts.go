package darts

func Score(x, y float64) int {
	distance := x*x + y*y

	switch {
	case distance <= 1:
		return 10
	case distance <= 25:
		return 5
	case distance <= 100:
		return 1
	default:
		return 0
	}
}
