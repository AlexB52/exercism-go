package pangram

import (
	"strings"
)

func IsPangram(input string) bool {
	result := map[rune]bool{}

	for _, r := range strings.ToLower(input) {
		result[r] = true
	}

	for l := 'a'; l < 'z'; l++ {
		if !result[l] {
			return false
		}
	}

	return true
}
