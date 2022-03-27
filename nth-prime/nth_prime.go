package prime

func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}

	primes := []int{}
	for i := 2; len(primes) < n; i++ {

		isPrime := true
		for _, p := range primes {
			if i%p == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes[len(primes)-1], true
}
