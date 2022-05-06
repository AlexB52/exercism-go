package prime

func Factors(n int64) (factors []int64) {
	if n <= 1 {
		return nil
	}

prime:
	for n >= 2 {
		for i := int64(2); i <= n; i++ {
			if n%i == 0 {
				factors = append(factors, i)
				n = n / i
				continue prime
			}
		}
	}

	return factors
}
