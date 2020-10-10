package mergeSort

func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	var sort []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			sort = append(sort, left[0])
			left = left[1:]
		} else {
			sort = append(sort, right[0])
			right = right[1:]
		}
	}
	for _, v := range left {
		sort = append(sort, v)
	}
	for _, v := range right {
		sort = append(sort, v)
	}

	return sort
}
