package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	var result string
	re := regexp.MustCompile(`\b[^a-zA-Z]*([a-zA-Z])[\w']*\b`)
	for _, match := range re.FindAllStringSubmatch(s, -1) {
		result += match[1]
	}
	return strings.ToUpper(result)
}
