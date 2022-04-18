package sieve

func Sieve(limit int) []int {
	if limit <= 1 {
		return nil
	}

	primes := []int{2}
	for i := 3; i <= limit; i++ {
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
	return primes
}
