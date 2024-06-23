package divide_n_conquer

func Bigger(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}

	midle := len(numbers) / 2

	leftBigger := Bigger(numbers[:midle])
	rightBigger := Bigger(numbers[midle:])

	if leftBigger > rightBigger {
		return leftBigger
	}
	return rightBigger
}
