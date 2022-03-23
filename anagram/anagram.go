package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) (result []string) {
	for _, candidate := range candidates {
		c, s := strings.ToLower(candidate), strings.ToLower(subject)

		if c != s && AnagramForm(c) == AnagramForm(s) {
			result = append(result, candidate)
		}
	}
	return result
}

func AnagramForm(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}
