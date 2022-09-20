package grains

import "errors"

func Square(number int) (uint64, error) {
	if number < 1 {
		return 0, errors.New("number must be greater than 0")
	}
	if number > 64 {
		return 0, errors.New("number cannot be greater than 64")
	}

	return square(number)(), nil
}

func square(number int) func() uint64 {
	squares := make([]uint64, 65)
	squares[1] = 1

	return func() uint64 {
		if squares[number] != 0 {
			return squares[number]
		}
		for i := 1; i <= number; i++ {
			if squares[i] != 0 {
				continue
			}
			squares[i] = squares[i-1] * 2
		}

		return squares[number]
	}
}

func Total() uint64 {
	var sum uint64

	for i := 1; i <= 64; i++ {
		sum += square(i)()
	}

	return sum
}
