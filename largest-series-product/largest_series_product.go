package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span || span < 0 {
		return -1, errors.New("span is bigger than number supplied")
	}

	result := 0
	for i := 0; i <= len(digits)-span; i++ {
		product := 1
		for _, r := range digits[i : i+span] {
			n, err := strconv.Atoi(string(r))
			if err != nil {
				return -1, errors.New("only numbers are allowed")
			}

			product *= n
		}

		if product > result {
			result = product
		}
	}

	return int64(result), nil
}
