package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("number is not within a valid range of [1, 64]")
	}

	return 1 << (number - 1), nil
}

func Total() uint64 {
	// total number of grains on a board of 64 squares
	// 2**0 + 2**1 + ... + 2**63
	return math.MaxUint64
}
