package allyourbase

import (
	"errors"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase <= 1 {
		return nil, errors.New("input base must be >= 2")
	}

	if outputBase <= 1 {
		return nil, errors.New("output base must be >= 2")
	}

	for _, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}

	var number int
	for i := 0; i < len(inputDigits); i++ {
		d := float64(inputDigits[i])
		e := float64(len(inputDigits) - 1 - i)
		b := float64(inputBase)
		number += int(d * math.Pow(b, e))
	}

	if number == 0 {
		return []int{0}, nil
	}

	var result []int
	for number > 0 {
		result = append([]int{number % outputBase}, result...)
		number /= outputBase
	}
	return result, nil
}
