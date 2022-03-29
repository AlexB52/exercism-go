package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	s = strings.ToLower(s)
	var groups []string

	for i := 0; i < len(s); {
		var group string

		for i < len(s) && len(group) < 5 {
			r := rune(s[i])
			if unicode.IsLetter(r) {
				group += string(rune('z' + 'a' - r))
			} else if unicode.IsDigit(r) {
				group += string(r)
			}
			i++
		}

		if len(group) > 0 {
			groups = append(groups, group)
		}
	}

	return strings.Join(groups, " ")
}
