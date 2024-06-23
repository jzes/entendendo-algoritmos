package divide_n_conquer

func Sum(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}
	return numbers[0] + Sum(numbers[1:])
}
