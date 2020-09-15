package MyMedianofTwoSortedArrays

import (
	"reflect"
	"testing"
)

/*给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。

请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。



示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

type question4 struct {
	param
	answer
}

type param struct {
	num1 []int
	num2 []int
}

type answer struct {
	one float64
}

func TestMedianofTwoSortedArrays(t *testing.T) {

	type test struct {
		nums1 []int
		nums2 []int
		want  float64
	}
	tests := map[string]test{
		"case_1": {nums1: []int{1, 3}, nums2: []int{3}, want: 2.1},
		"case_2": {nums1: []int{1, 2}, nums2: []int{3, 4}, want: 2.5},
	}

	for name, ts := range tests {
		got := findMedianSortedArrays(ts.nums1, ts.nums2)
		if !reflect.DeepEqual(got, ts.want) {
			t.Errorf("name:%s excepted:%#v, got:%#v", name, ts.want, got)
		}
	}
}
