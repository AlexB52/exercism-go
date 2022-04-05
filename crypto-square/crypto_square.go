package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

func Encode(pt string) string {
	re := regexp.MustCompile(`\w`)
	matches := re.FindAllString(strings.ToLower(pt), -1)
	r, c := MatrixDetails(len(matches))

	var groups = make([]string, c)
	for i := 0; i < c*r; i++ {
		var s string
		if i < len(matches) {
			s = matches[i]
		} else {
			s = " "
		}

		groups[i%c] += s
	}

	return strings.Join(groups, " ")
}

func MatrixDetails(length int) (r int, c int) {
	r = int(math.Round(math.Sqrt(float64(length))))
	if r*r >= length {
		c = r
	} else {
		c = r + 1
	}
	return r, c
}
