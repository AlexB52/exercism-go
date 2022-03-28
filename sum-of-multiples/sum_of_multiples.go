package summultiples

func SumMultiples(limit int, divisors ...int) (result int) {
OUTER:
	for i := 0; i < limit; i++ {
		for _, divisor := range divisors {
			if divisor != 0 && i%divisor == 0 {
				result += i
				continue OUTER
			}
		}
	}
	return result
}
