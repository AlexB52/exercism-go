package prime

func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}

	primes := []int{}

findPrimes:
	for i := 2; len(primes) < n; i++ {

		for _, p := range primes {
			if i%p == 0 {
				continue findPrimes
			}
		}

		primes = append(primes, i)
	}

	return primes[len(primes)-1], true
}
