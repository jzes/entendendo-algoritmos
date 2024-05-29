package main

func lower(arr []int) int {
	li := 0
	lv := arr[li]

	for i := 1; i < len(arr); i++ {
		if arr[i] < lv {
			li = i
			lv = arr[i]
		}
	}

	return li
}

func selectSort(arr []int) []int {
	orderedArr := []int{}
	li := 0
	for len(arr) > 0 {
		li = lower(arr)
		orderedArr = append(orderedArr, arr[li])
		arr = append(arr[:li], arr[li+1:]...)
	}
	return orderedArr
}
