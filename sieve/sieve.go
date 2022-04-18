package sieve

func Sieve(limit int) []int {
	var primes []int

findPrimes:
	for i := 2; i <= limit; i++ {
		for _, p := range primes {
			if i%p == 0 {
				continue findPrimes
			}
		}
		primes = append(primes, i)
	}

	return primes
}
