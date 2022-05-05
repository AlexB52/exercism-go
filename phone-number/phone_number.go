package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func Number(phoneNumber string) (string, error) {
	re, _ := regexp.Compile("\\d")
	numbers := re.FindAllString(phoneNumber, -1)

	if len(numbers) != 10 && (len(numbers) != 11 || numbers[0] != "1") {
		return "", errors.New("invalid phone number")
	}

	if len(numbers) == 11 {
		numbers = numbers[1:]
	}

	if numbers[0] == "0" || numbers[0] == "1" {
		return "", errors.New("invalid phone number")
	}

	if numbers[3] == "0" || numbers[3] == "1" {
		return "", errors.New("invalid phone number")
	}

	return strings.Join(numbers, ""), nil
}

func AreaCode(phoneNumber string) (string, error) {
	numbers, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return numbers[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	numbers, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", numbers[0:3], numbers[3:6], numbers[6:10]), nil
}
