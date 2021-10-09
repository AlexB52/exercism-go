package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(card string) bool {
	card = strings.ReplaceAll(card, " ", "")

	if InvalidFormat(card) {
		return false
	}

	sum := 0
	for i := len(card) - 1; i >= 0; i-- {
		number, _ := strconv.Atoi(string(card[i]))
		position := len(card) - i
		sum += MappingLuhnNumber(position, number)
	}

	return sum%10 == 0
}

func InvalidFormat(card string) bool {
	if len(card) < 2 {
		return true
	}

	for _, char := range card {
		if !unicode.IsNumber(char) {
			return true
		}
	}

	return false
}

func MappingLuhnNumber(position int, number int) int {
	if position%2 == 0 {
		number *= 2
		if number > 9 {
			number -= 9
		}
	}

	return number
}
