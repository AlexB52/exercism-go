package pangram

import (
	"strings"
)

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	for l := 'a'; l < 'z'; l++ {
		if !strings.ContainsRune(input, l) {
			return false
		}
	}

	return true
}
