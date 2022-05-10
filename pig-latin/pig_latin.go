package piglatin

import (
	"fmt"
	"regexp"
	"strings"
)

var PREFIXES = []string{"ch", "qu", "squ", "thr", "sch", "th", "rh", "f", "m", "r", "p", "k", "x", "q", "c", "h", "y"}
var EXCEPTIONS = []string{"xr", "yt"}

func Sentence(sentence string) string {
	var result []string
	re := regexp.MustCompile("\\w+")

	for _, w := range re.FindAllString(sentence, -1) {
		result = append(result, PigLatin(w))
	}

	return strings.Join(result, " ")
}

func PigLatin(word string) string {
	swapPrefix := true
	for _, e := range EXCEPTIONS {
		if strings.HasPrefix(word, e) {
			swapPrefix = false
			break
		}
	}

	if swapPrefix {
		for _, c := range PREFIXES {
			if strings.HasPrefix(word, c) {
				word = fmt.Sprintf("%s%s", strings.TrimPrefix(word, c), c)
				break
			}
		}
	}

	return fmt.Sprintf("%say", word)
}
