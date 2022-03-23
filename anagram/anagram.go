package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	anagrams := map[string][]string{}

	for _, candidate := range candidates {
		if strings.ToLower(candidate) == strings.ToLower(subject) {
			continue
		}
		anagrams[AnagramForm(candidate)] = append(anagrams[AnagramForm(candidate)], candidate)
	}

	return anagrams[AnagramForm(subject)]
}

func AnagramForm(s string) string {
	runes := []rune(strings.ToLower(s))
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}
