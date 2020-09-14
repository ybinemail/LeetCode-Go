package MyRegularExpressionMatching

import "testing"

type question struct {
	param
	answer
}
type param struct {
	str string
	pre string
}
type answer struct {
	res bool
}

func TestRegularExpressionMatching(t *testing.T) {

	qs := []question{
		{
			param{"aa", "a"}, answer{false},
		},
		{
			param{"aa", "a*"}, answer{true},
		},
		{
			param{"ab", ".*"}, answer{true},
		},
		{
			param{"aab", "c*a*b"}, answer{true},
		}, {
			param{"mississippi", "mis*is*p*."}, answer{false},
		},
	}

	for _, v := range qs {
		t.Logf("input param %+v, output answer : %+v <=> %+v ", v.param, v.res, isMatch(v.str, v.pre))
	}
}
