package summultiples

func SumMultiples(limit int, divisors ...int) (result int) {
	var numbers = map[int]bool{}

	for _, divisor := range divisors {
		for i := 0; i < limit; i++ {
			if divisor != 0 && i%divisor == 0 {
				numbers[i] = true
			}
		}
	}

	for n, _ := range numbers {
		result += n
	}
	return result
}
