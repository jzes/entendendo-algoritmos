package toolBox

func MakeRange(start, end int) []int {
	r := make([]int, end-start+1)

	for i := range r {
		r[i] = start + i
	}
	return r
}
