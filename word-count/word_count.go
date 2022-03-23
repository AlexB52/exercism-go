package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	regex := regexp.MustCompile(`\b[\w']+\b`)
	words := regex.FindAll([]byte(strings.ToLower(phrase)), -1)

	result := Frequency{}
	for _, word := range words {
		result[string(word)]++
	}
	return result
}
