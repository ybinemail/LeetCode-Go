package MyContainerWithMostWater

import "testing"

/*给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

type question struct {
	param
	answer
}
type param struct {
	heigth []int
}
type answer struct {
	res int
}

func TestMaxArea(t *testing.T) {
	qs := []question{
		{
			param{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, answer{49},
		},
		{
			param{[]int{1, 1}}, answer{1},
		},
		{
			param{[]int{0, 0, 0, 0}}, answer{0},
		},
		{
			param{[]int{1, 0, 0, 1}}, answer{3},
		},
		{
			param{[]int{1, 0, 0, 0}}, answer{0},
		},
	}
	for _, v := range qs {
		t.Logf("input param %+v, output answer : %+v <=> %+v ", v.param, v.res, maxArea(v.param.heigth))
	}

}
