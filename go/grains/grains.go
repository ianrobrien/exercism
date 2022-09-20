package grains

import "errors"

func Square(number int) (uint64, error) {
	if number < 1 {
		return 0, errors.New("number must be greater than 0")
	}
	if number > 64 {
		return 0, errors.New("number cannot be greater than 64")
	}
	sum := uint64(1)
	for i := 1; i < number; i++ {
		sum *= 2
	}
	return sum, nil
}

func Total() uint64 {
	var sum uint64
	for i := 1; i <= 64; i++ {
		square, _ := Square(i)
		sum += square
	}
	return sum
}
