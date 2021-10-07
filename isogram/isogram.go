package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	letterCount := make(map[rune]int)
	trimedWord := strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(word), "-", ""), " ", "")

	for _, letter := range trimedWord {
		letterCount[letter]++

		if letterCount[letter] > 1 {
			return false
		}
	}

	return true
}
