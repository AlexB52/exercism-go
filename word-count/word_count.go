package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	regex := regexp.MustCompile("[a-z|'|0-9]+")
	words := regex.FindAll([]byte(strings.ToLower(phrase)), -1)

	result := Frequency{}
	for _, word := range words {
		result[strings.Trim(string(word), "'")]++
	}
	return result
}
