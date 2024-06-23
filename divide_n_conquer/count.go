package divide_n_conquer

func Count(things []int) int {
	if len(things) == 0 {
		return 0
	}
	return 1 + Count(things[1:])
}
