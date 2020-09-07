package MyLongestPalindromicSubstring

import "testing"

type question5 struct {
	param
	answer
}
type param struct {
	str string
}
type answer struct {
	substr string
}

func TestLongestPalindromicSubstring(t *testing.T) {
	qs := []question5{
		{
			param{"aaabaaaa"}, answer{"aaabaaa"},
		},
		{
			param{"aaa"}, answer{"aaa"},
		},
		{
			param{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}, answer{"aa"},
		}, {
			param{"bbaa"}, answer{"bb"},
		},
		{
			param{"bba"}, answer{"bb"},
		},
		{
			param{"cbb"}, answer{"bb"},
		},
		{
			param{"cbbd"}, answer{"bb"},
		},

		{
			param{"aaaa"}, answer{"aaaa"},
		},
		{
			param{"222020221"}, answer{"2202022"},
		},
		{
			param{"abcddcbf"}, answer{"bcddcb"},
		},
		{
			param{"aabaa"}, answer{"aabaa"},
		},
		{
			param{"aaabaaaa"}, answer{"aaabaaa"},
		},
		{
			param{"bananas"}, answer{"anana"},
		},
		{
			param{"babad"}, answer{"bab"},
		},

		{
			param{"bbbbbbbbb"}, answer{"bbbbbbbbb"},
		},
		{
			param{"abcbcbd"}, answer{"bcbcb"},
		},
		{
			param{""}, answer{""},
		},
		{
			param{"ab"}, answer{"ab"},
		},
		{
			param{"a"}, answer{"a"},
		},
		{
			param{"abcdcbahihba"}, answer{"abcdcba"},
		},
	}
	for _, v := range qs {
		t.Logf("input param %v, output answer : %s <=> %s ", v.param, v.answer.substr, longestPalindrome3(v.param.str))
	}

}
