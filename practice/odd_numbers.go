package practice

import "errors"

func AddOddNumbers(min, max int) (int, error) {
	sum := 0
	if min > max {
		return 0, errors.New("min must be less than max")
	}
	// if min input is even add one to make it odd
	if min%2 != 1 {
		min += 1
	}
	for i := min; i <= max; i += 2 {
		sum += i
	}
	return sum, nil
}
