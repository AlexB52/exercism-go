package sieve

func Sieve(limit int) []int {
	var primes []int
	var sieve = map[int]bool{}

	for i := 2; i <= limit; i++ {
		s, ok := sieve[i]

		if !ok {
			primes = append(primes, i)
			s = true
		}

		if !s {
			continue
		}

		for j := 2; j*i <= limit; j++ {
			sieve[i*j] = false
		}
	}

	return primes
}
