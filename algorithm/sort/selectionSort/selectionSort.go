package selectionSort

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		index := i
		min := arr[i]
		for j := i + 1; j < len(arr); j++ {
			if min > arr[j] {
				min = arr[j]
				index = j
			}
		}
		arr[i], arr[index] = arr[index], arr[i]
	}
	return arr
}
