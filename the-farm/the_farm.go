package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	number int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.number)
}

var nonScaleError = errors.New("non-scale error")
var negativeFodder = errors.New("negative fodder")
var divisionByZero = errors.New("division by zero")

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()

	if err != nil {
		if err == ErrScaleMalfunction {
			fodder *= 2.0
		} else {
			return 0, err
		}
	}

	if fodder < 0 {
		return 0, negativeFodder
	}

	if cows == 0 {
		return 0, divisionByZero
	}

	if cows < 0 {
		return 0, &SillyNephewError{number: cows}
	}

	return fodder / float64(cows), nil
}
