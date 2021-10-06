package practice

func AddOddNumbers(min, max int) int {
	sum := 0
	for i := min; i <= max; i++ {
		if i%2 == 1 {
			sum = sum + i
		}
	}
	return sum
}
