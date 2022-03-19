package romannumerals

import (
	"errors"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {
	ArabicToRoman := map[int]string{
		1000: "M",
		500:  "D",
		100:  "C",
		50:   "L",
		10:   "X",
		5:    "V",
		1:    "I",
	}

	if input <= 0 || input > 3000 {
		return "", errors.New("Invalid input")
	}

	var result string
	for _, i := range [7]int{1000, 500, 100, 50, 10, 5, 1} {
		for input >= i {
			input -= i
			result += ArabicToRoman[i]
		}
	}

	// It appears that the long version of roman numerals was a valid one.
	// The short version replaced these patterns with their short equivalent.
	// Source: https://sandimetz.com/blog/2016/6/9/make-everything-the-same

	result = strings.Replace(result, "DCCCC", "CM", 1)
	result = strings.Replace(result, "CCCC", "CD", 1)
	result = strings.Replace(result, "LXXXX", "XC", 1)
	result = strings.Replace(result, "XXXX", "XL", 1)
	result = strings.Replace(result, "VIIII", "IX", 1)
	result = strings.Replace(result, "IIII", "IV", 1)

	return result, nil
}
