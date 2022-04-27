package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	var sum int
	s := strconv.Itoa(n)
	for _, r := range s {
		sum += int(math.Pow(float64(r-'0'), float64(len(s))))
	}
	return sum == n
}
