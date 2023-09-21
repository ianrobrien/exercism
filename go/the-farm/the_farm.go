package thefarm

import (
	"errors"
	"fmt"
)

type SillyNephewError struct {
	amountCows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.amountCows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	amount, err := weightFodder.FodderAmount()

	if cows == 0 {
		return 0, errors.New("division by zero")
	}
	if err != nil && err != ErrScaleMalfunction {
		return 0, err
	}
	if err == ErrScaleMalfunction {
		if amount >= 0 {
			return (amount * 2) / float64(cows), nil
		}
		return 0, errors.New("negative fodder")
	}
	if (err == ErrScaleMalfunction || err == nil) && amount < 0 {
		return 0, errors.New("negative fodder")
	}
	if cows < 0 {
		return 0, &SillyNephewError{
			amountCows: cows,
		}
	}

	return amount / float64(cows), nil
}
