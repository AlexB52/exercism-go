package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func longRunningTask(s string) <-chan FreqMap {
	r := make(chan FreqMap)

	go func() {
		defer close(r)
		r <- Frequency(s)
	}()

	return r
}

func ConcurrentFrequency(s []string) FreqMap {
	m := FreqMap{}

	for _, r := range s {
		for k, v := range <-longRunningTask(r) {
			m[k] += v
		}
	}

	return m
}
