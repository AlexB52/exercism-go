package hamming

import "errors"

func Distance(a, b string) (distance int, er error) {
	if len(a) != len(b) {
		return 0, errors.New("a & b of different length")
	}

	for index := range a {
		if a[index] != b[index] {
			distance++
		}
	}

	return
}
