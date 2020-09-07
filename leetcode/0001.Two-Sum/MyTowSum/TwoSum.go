package MyTowSum

func twoSum(nums []int, target int) []int {
	// 暴力循环
	index := []int{}
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			//从i的下一位开始判断，是否满足条件
			if v+nums[j] == target {
				index = append(index, i, j)
			}
		}
	}
	return index
}

func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[v] = i
	}
	return nil
}
