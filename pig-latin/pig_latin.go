package piglatin

import (
	"fmt"
	"regexp"
	"strings"
)

func Sentence(sentence string) string {
	var result []string
	re := regexp.MustCompile("\\w+")

	for _, w := range re.FindAllString(sentence, -1) {
		result = append(result, Convert(w))
	}

	return strings.Join(result, " ")
}

func Convert(word string) string {
	consonants := []string{"ch", "qu", "squ", "thr", "sch", "th", "rh", "f", "m", "r", "p", "k", "x", "q", "c", "h", "y"}
	exceptions := []string{"xr", "yt"}

swap:
	for _, c := range consonants {
		for _, e := range exceptions {
			if strings.HasPrefix(word, e) {
				break swap
			}
		}

		if strings.HasPrefix(word, c) {
			word = fmt.Sprintf("%s%s", strings.TrimPrefix(word, c), c)
			break
		}
	}

	return fmt.Sprintf("%say", word)
}
