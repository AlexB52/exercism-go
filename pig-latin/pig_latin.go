package piglatin

import (
	"regexp"
	"strings"
)

// From aolshev solution. It was too good not to use it.
// https://exercism.org/tracks/go/exercises/pig-latin/solutions/aolshev

func Sentence(text string) string {
	var result []string
	for _, w := range strings.Split(text, " ") {
		result = append(result, word(w))
	}
	return strings.Join(result, " ")
}

var (
	beginWithVowels     = regexp.MustCompile(`^([aeiou].|xr|yt).+$`)
	beginWithConsonants = regexp.MustCompile(`^([^aeiou]+)(.+)$`)
	beginWithQuVariants = regexp.MustCompile(`^([^aeiou]*qu)(.+)$`)
	beginWithYVariants  = regexp.MustCompile(`^([^aeiou]+|.)(y.+)$`)
)

func word(word string) string {
	// Rule 1
	if beginWithVowels.MatchString(word) {
		return word + "ay"
	}
	// Rule 3
	if match := beginWithQuVariants.FindStringSubmatch(word); match != nil {
		return match[2] + match[1] + "ay"
	}
	// Rule 4
	if match := beginWithYVariants.FindStringSubmatch(word); match != nil {
		return match[2] + match[1] + "ay"
	}
	// Rule 2
	if match := beginWithConsonants.FindStringSubmatch(word); match != nil {
		return match[2] + match[1] + "ay"
	}
	return word
}
