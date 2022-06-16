package say

import (
	"fmt"
	"strings"
)

func Say(n int64) (string, bool) {
	if n < 0 || n > 999_999_999_999 {
		return "", false
	}

	var result []string
	switch {
	case n < 100:
		result = append(result, Hundreds(n))
	case n < 1000:
		if n/1000 > 0 {
			result = append(result, fmt.Sprintf("%s thousand", Hundreds(n/1000)))
			n = n % 1000
		}

		if n/100 > 0 {
			result = append(result, fmt.Sprintf("%s hundred", Hundreds(n/100)))
			n = n % 100
		}

		if n/1 > 0 {
			result = append(result, Hundreds(n/1))
		}
	case n < 1_000_000:
		if n/1_000_000 > 0 {
			result = append(result, fmt.Sprintf("%s million", Hundreds(n/1_000_000)))
			n = n % 1_000_000
		}

		if n/1000 > 0 {
			result = append(result, fmt.Sprintf("%s thousand", Hundreds(n/1000)))
			n = n % 1000
		}

		if n/100 > 0 {
			result = append(result, fmt.Sprintf("%s hundred", Hundreds(n/100)))
			n = n % 100
		}

		if n/1 > 0 {
			result = append(result, Hundreds(n/1))
		}
	case n < 1_000_000_000:
		if n/1_000_000 > 0 {
			result = append(result, fmt.Sprintf("%s million", Hundreds(n/1_000_000)))
			n = n % 1_000_000
		}

		if n/1000 > 0 {
			result = append(result, fmt.Sprintf("%s thousand", Hundreds(n/1000)))
			n = n % 1000
		}

		if n/100 > 0 {
			result = append(result, fmt.Sprintf("%s hundred", Hundreds(n/100)))
			n = n % 100
		}

		if n/1 > 0 {
			result = append(result, Hundreds(n/1))
		}
	case n < 1_000_000_000_000:
		if n/1_000_000_000 > 0 {
			result = append(result, fmt.Sprintf("%s billion", Hundreds(n/1_000_000_000)))
			n = n % 1_000_000_000
		}

		if n/1_000_000 > 0 {
			result = append(result, fmt.Sprintf("%s million", Hundreds(n/1_000_000)))
			n = n % 1_000_000
		}

		if n/1000 > 0 {
			result = append(result, fmt.Sprintf("%s thousand", Hundreds(n/1000)))
			n = n % 1000
		}

		if n/100 > 0 {
			result = append(result, fmt.Sprintf("%s hundred", Hundreds(n/100)))
			n = n % 100
		}

		if n/1 > 0 {
			result = append(result, Hundreds(n/1))
		}
	}
	return strings.Join(result, " "), true
}

func Hundreds(n int64) string {
	if n == 0 {
		return NUMBERS[n]
	}

	var result []string
	if n/100 > 0 {
		result = append(result, fmt.Sprintf("%s hundred", NUMBERS[n/100]))
		n = n % 100
	}

	if n/1 > 0 {
		result = append(result, NUMBERS[n/1])
	}

	return strings.Join(result, " ")
}
