package encode

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func RunLengthEncode(input string) (result string) {
	if len(input) == 0 {
		return ""
	}

	for i := 0; i < len(input); {
		j := i + 1
		for j < len(input) && input[i] == input[j] {
			j++
		}

		letter, count := string(input[i]), j-i

		if count > 1 {
			result += fmt.Sprintf("%d%s", count, letter)
		} else {
			result += letter
		}

		i = j
	}

	return result
}

func RunLengthDecode(input string) (result string) {
	re := regexp.MustCompile(`\d+[\w\s]|[\w\s]`)

	for _, match := range re.FindAll([]byte(input), -1) {
		letter := string(match[len(match)-1])
		count, err := strconv.Atoi(string(match[:len(match)-1]))
		if err != nil {
			count = 1
		}

		result += strings.Repeat(letter, count)
	}

	return result
}
