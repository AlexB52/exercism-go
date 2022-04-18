package sieve

func Sieve(limit int) []int {
	var primes []int
	var sieve = map[int]bool{}

	for i := 2; i <= limit; i++ {
		_, ok := sieve[i]

		if !ok {
			primes = append(primes, i)
			for j := i; j*i <= limit; j++ {
				sieve[i*j] = false
			}
		}
	}

	return primes
}
