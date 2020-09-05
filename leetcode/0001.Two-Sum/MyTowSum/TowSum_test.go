package MyTowSum

import "testing"

func TestTowSum(t *testing.T) {

	t.Log("start testing")

	nums := []int{2, 7, 11, 15}
	target := 9
	index := twoSum(nums, target)
	t.Log(index)
}
