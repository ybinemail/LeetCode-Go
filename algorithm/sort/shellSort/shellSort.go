package shellSort

func shellSort(arr []int) []int {
	length := len(arr)

	gas := 1
	for gas < length/3 {
		gas = gas*3 + 1
	}

	for gas := gas; gas > 0; gas = gas / 3 {

		for i := gas; i < len(arr); i++ {
			tmp := arr[i]
			j := i - gas
			for j >= 0 && tmp < arr[j] {
				arr[j+gas] = arr[j]
				//arr[j] = tmp
				j -= gas
			}
			arr[j+gas] = tmp
		}
	}
	return arr
}
