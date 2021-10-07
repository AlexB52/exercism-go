package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	for index, letter := range word {
		if unicode.IsLetter(letter) && strings.ContainsRune(word[index+1:], letter) {
			return false
		}
	}

	return true
}
