package hamming

import "errors"

func Distance(a, b string) (distance int, er error) {
	if len(a) != len(b) {
		return 0, errors.New("a & b of different length")
	}

	runeA := []rune(a)
	runeB := []rune(b)

	for index := range runeA {
		if runeA[index] != runeB[index] {
			distance++
		}
	}

	return
}
