package leetcode

import (
	"reflect"
	"testing"
)

type question1 struct {
	para1
	ans1
}

// para 是参数
// one 代表第一个参数
type para1 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans1 struct {
	one []int
}

func Test_Problem1(t *testing.T) {

	type test struct {
		nums   []int
		target int
		want   []int
	}
	tests := map[string]test{
		"case_1": {nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		"case_2": {nums: []int{3, 2, 4}, target: 5, want: []int{1, 2}},
		"case_3": {nums: []int{0, 8, 7, 3, 3, 4, 2}, target: 11, want: []int{1, 3}},
	}

	for name, ts := range tests {
		got := twoSum(ts.nums, ts.target)
		if !reflect.DeepEqual(got, ts.want) {
			t.Errorf("name:%s excepted:%#v, got:%#v", name, ts.want, got)
		}
	}
}
