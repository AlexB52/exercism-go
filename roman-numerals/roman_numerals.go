package romannumerals

import (
	"errors"
)

func ToRomanNumeral(input int) (string, error) {
	DecimalToRomanNumeral := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	if input <= 0 || input > 3000 {
		return "", errors.New("Invalid input")
	}

	numeralOrder := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var result string
	for _, i := range numeralOrder {
		for {
			if input-i < 0 {
				break
			}

			input -= i
			result += DecimalToRomanNumeral[i]
		}
	}

	return result, nil
}
