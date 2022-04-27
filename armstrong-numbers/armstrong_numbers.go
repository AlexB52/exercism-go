package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	var result int
	s := strconv.Itoa(n)
	for _, r := range s {
		a, _ := strconv.Atoi(string(r))
		result += int(math.Pow(float64(a), float64(len(s))))
	}
	return result == n
}
