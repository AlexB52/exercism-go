package encode

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Fragment struct {
	count  int
	letter byte
}

func (f *Fragment) ToCode() (result string) {
	if f.count == 1 {
		return string(f.letter)
	} else {
		return fmt.Sprintf("%d%s", f.count, string(f.letter))
	}
}

func RunLengthEncode(input string) (result string) {
	if len(input) == 0 {
		return ""
	}

	fragments := []*Fragment{&Fragment{1, input[0]}}

	for i := 1; i < len(input); i++ {
		current := fragments[len(fragments)-1]

		if input[i] == current.letter {
			current.count++
			continue
		}

		fragments = append(fragments, &Fragment{1, input[i]})
	}

	for _, f := range fragments {
		result += f.ToCode()
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
