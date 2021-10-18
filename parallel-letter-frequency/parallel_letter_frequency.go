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

func ConcurrentFrequency(s []string) FreqMap {
	c := make(chan FreqMap)
	m := FreqMap{}

	for _, str := range s {
		go func(r string) { c <- Frequency(r) }(str)
	}

	for i := 0; i < len(s); i++ {
		for k, v := range <-c {
			m[k] += v
		}
	}

	return m
}
