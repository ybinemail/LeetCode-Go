package MyMedianofTwoSortedArrays

import "testing"

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
	qs := []question4{
		{
			param{[]int{1, 3}, []int{3}},
			answer{2.1},
		},

		{
			param{[]int{1, 2}, []int{3, 4}},
			answer{2.5},
		},
	}
	t.Logf("【input】:%v", qs)
}
