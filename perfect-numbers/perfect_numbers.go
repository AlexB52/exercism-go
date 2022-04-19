package perfect

import (
	"errors"
)

// Define the Classification type here.
type Classification int

const (
	ClassificationPerfect Classification = iota
	ClassificationAbundant
	ClassificationDeficient
)

var ErrOnlyPositive = errors.New("can only be positive")

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}

	var sum int64
	for i := int64(1); i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	switch {
	case n > sum:
		return ClassificationDeficient, nil
	case n < sum:
		return ClassificationAbundant, nil
	default:
		return ClassificationPerfect, nil
	}
}
