package insertionSort

func insertionSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		for j := i-1; j>=0 ; j-- {
			if tmp < arr[j] {
				arr[j+1] = arr[j]
				arr[j] = tmp
			}
		}
	}

	return arr
}


func insertionSort2(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i - 1
		for j >= 0 && tmp < arr[j] {
			arr[j+1] = arr[j]
			//arr[j] = tmp
			j--
		}
		arr[j+1] = tmp
	}

	return arr
}