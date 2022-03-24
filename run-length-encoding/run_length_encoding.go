package encode

import (
	"fmt"
	"regexp"
	"strconv"
)

type Fragment struct {
	count  int
	letter byte
}

func (f *Fragment) ToString() (result string) {
	for i := 1; i <= f.count; i++ {
		result += string(f.letter)
	}
	return result
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

	var fragments = []*Fragment{&Fragment{1, input[0]}}
	for i := 1; i < len(input); i++ {
		f := fragments[len(fragments)-1]

		if input[i] == f.letter {
			f.count++
		} else {
			fragments = append(fragments, &Fragment{1, input[i]})
		}
	}

	for _, f := range fragments {
		result += f.ToCode()
	}
	return result
}

func RunLengthDecode(input string) (result string) {
	re := regexp.MustCompile(`\d+[\w\s]|[\w\s]`)
	matches := re.FindAll([]byte(input), -1)

	for _, match := range matches {
		var f Fragment
		count, err := strconv.Atoi(string(match[:len(match)-1]))

		if err == nil {
			f = Fragment{count, match[len(match)-1]}
		} else {
			f = Fragment{1, match[0]}
		}

		result += f.ToString()
	}

	return result
}
