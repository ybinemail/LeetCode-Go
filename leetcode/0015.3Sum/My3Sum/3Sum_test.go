package My3Sum

import (
	"reflect"
	"testing"
)

/*给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func Test3Sum(t *testing.T) {

	type test struct {
		input []int
		want  [][]int
	}
	tests := map[string]test{
		"case_1": {input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, want: [][]int{}},
	}

	for name, ts := range tests {
		got := threeSum(ts.input)
		if !reflect.DeepEqual(got, ts.want) {
			t.Errorf("name:%s excepted:%#v, got:%#v", name, ts.want, got)
		}
	}
}
