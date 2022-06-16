package say

import (
	"fmt"
	"strings"
)

func Say(n int64) (string, bool) {
	if n < 0 || n > 999_999_999_999 {
		return "", false
	}

	if n == 0 {
		return "zero", true
	}

	var result []string
	var i int64
	for i = 1_000_000_000; i > 100; i /= 1000 {
		if n/i > 0 {
			result = append(result, fmt.Sprintf("%s %s", Hundreds(n/i), NUMBERS[i]))
			n = n % i
		}
	}

	if n/1 > 0 {
		result = append(result, fmt.Sprintf("%s", Hundreds(n/1)))
	}

	return strings.Join(result, " "), true
}

func Hundreds(n int64) string {
	var result []string
	if n/100 > 0 {
		result = append(result, fmt.Sprintf("%s %s", NUMBERS[n/100], NUMBERS[100]))
		n = n % 100
	}

	if n/1 > 0 {
		result = append(result, fmt.Sprintf("%s", NUMBERS[n/1]))
	}

	return strings.Join(result, " ")
}
