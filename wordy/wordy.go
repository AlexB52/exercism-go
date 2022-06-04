package wordy

import (
	"regexp"
	"strconv"
)

var (
	VALIDATION_REGEX = `What is (-?\d+)+( (plus|minus|multiplied by|divided by) (-?\d+))*\?`
	OPERATIONS_REGEX = `(-?\d+|plus|minus|multiplied by|divided by)`
)

func Answer(question string) (int, bool) {
	if regexp.MustCompile(VALIDATION_REGEX).FindString(question) == "" {
		return 0, false
	}

	words := regexp.MustCompile(OPERATIONS_REGEX).FindAllString(question, -1)

	result, _ := strconv.Atoi(words[0])

	for i := 1; i < len(words); i += 2 {
		n, _ := strconv.Atoi(words[i+1])
		switch words[i] {
		case "plus":
			result += n
		case "minus":
			result -= n
		case "multiplied by":
			result *= n
		case "divided by":
			result /= n
		default:
			return 0, false
		}
	}

	return result, true
}
